package mongo

import (
	"context"
	"fmt"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PollutionRepository struct {
	coll *mongo.Collection
}

func (repo *PollutionRepository) IncPollutionTileMany(commands []repositories.IncPollutionCommand) result.Result[bool] {

	session, err := repo.coll.Database().Client().StartSession()
	if err != nil {
		return result.Result[bool]{}.Failed(err)
	}
	defer session.EndSession(context.TODO())

	opts := options.Update().SetUpsert(true)

	res, err := session.WithTransaction(context.TODO(), func(sessCtx mongo.SessionContext) (any, error) {
		for _, cmd := range commands {
			filter := serde.MarshalBson(map[string]any{
				"id": serde.Normalise(cmd.TileIndex),
			}).Expect()

			update := bson.D{
				bson.E{"$setOnInsert", bson.D{
					bson.E{"id", serde.Normalise(cmd.TileIndex)},
					bson.E{"types", bson.M{}},
				}},
				bson.E{"$add", bson.D{
					bson.E{fmt.Sprintf("types.%s.count", cmd.Report.Type.Name), 1},
					bson.E{fmt.Sprintf("types.%s.weight", cmd.Report.Type.Name), cmd.Report.Rate},
				}},
			}

			_, err := repo.coll.UpdateOne(context.TODO(), filter, update, opts)

			if err != nil {
				sessCtx.AbortTransaction(context.TODO())
				return false, err
			}
		}
		return true, nil
	})

	if err != nil {
		return result.Result[bool]{}.Failed(err)
	}

	return result.Success(res.(bool))
}

func bounds(tileBounds slippy_map.TileBounds, timeBounds time_serie.TimeInterval) bson.D {
	return bson.D{
		bson.E{"id.x", bson.D{
			bson.E{"$gte", tileBounds.MinX()},
			bson.E{"$lte", tileBounds.MaxX()},
		}},
		bson.E{"id.y", bson.D{
			bson.E{"$gte", tileBounds.MinY()},
			bson.E{"$lte", tileBounds.MaxY()},
		}},
		bson.E{"id.z", bson.D{
			bson.E{"$gte", tileBounds.MinZ()},
			bson.E{"$lte", tileBounds.MaxZ()},
		}},
		bson.E{"id.t", bson.D{
			bson.E{"$gte", timeBounds.Min()},
			bson.E{"$lte", timeBounds.Max()},
		}},
	}
}

func (repo *PollutionRepository) GetPollutionTiles(tileBounds slippy_map.TileBounds, timeBounds time_serie.TimeInterval) result.Result[models.PollutionTileCollection] {
	cursor, err := repo.coll.Find(context.TODO(), bounds(tileBounds, timeBounds))

	if err != nil {
		return result.Result[models.PollutionTileCollection]{}.Failed(err)
	}

	collection := models.PollutionTileCollection{}

	for cursor.Next(context.TODO()) {
		decodeResult := serde.UnMarshalBson[models.PollutionTile](cursor.Current)
		if decodeResult.HasFailed() {
			return result.Result[models.PollutionTileCollection]{}.Failed(decodeResult.UnwrapError())
		}

		collection = append(collection, decodeResult.Expect())
	}

	return result.Success(collection)
}
