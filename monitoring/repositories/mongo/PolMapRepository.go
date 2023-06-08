package mongo

import (
	"context"
	"fmt"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PolMapRepository struct {
	coll *mongo.Collection
}

func (repo *PolMapRepository) IncPollutionTileMany(commands []repositories.IncPollutionCommand) result.Result[bool] {

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
					bson.E{"cluster_id", serde.Normalise(cmd.ClusterIndex)},
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

func window(ul slippy_map.TileIndex, dr slippy_map.TileIndex, begin int, end int) bson.D {
	return bson.D{
		bson.E{"id.x", bson.D{
			bson.E{"$gt", ul.X},
			bson.E{"$lt", dr.X},
		}},
		bson.E{"id.y", bson.D{
			bson.E{"$gt", ul.Y},
			bson.E{"$lt", dr.Y},
		}},
		bson.E{"id.z", bson.D{
			bson.E{"$gt", ul.Z},
			bson.E{"$lt", dr.Z},
		}},
		bson.E{"id.t", bson.D{
			bson.E{"$gt", begin},
			bson.E{"$lt", end},
		}},
	}
}

func (repo *PolMapRepository) GetPollutionTiles(ul slippy_map.TileIndex, dr slippy_map.TileIndex, begin int, end int) result.Result[models.PolTimeSlice] {
	pipeline := bson.A{
		// Select the polmap subset
		bson.D{
			bson.E{"$match", window(ul, dr, begin, end)},
		},
		// Group by time
		bson.D{
			bson.E{"$group", bson.D{
				bson.E{"_id", "id.t"},
				bson.E{"matrix", bson.D{bson.E{"$push", "$$ROOT"}}},
			}},
		},
		// copy _id value to t
		bson.D{
			bson.E{"$set", bson.D{
				bson.E{"t", "$_id"},
			}},
		},
		// sort chronologically
		bson.D{
			bson.E{"$sort", bson.D{
				bson.E{"t", 1},
			}},
		},
	}
	cursor, err := repo.coll.Aggregate(context.TODO(), pipeline)

	if err != nil {
		return result.Result[models.PolTimeSlice]{}.Failed(err)
	}

	timeSlice := models.PolTimeSlice{}

	for cursor.Next(context.TODO()) {
		decodeResult := serde.UnMarshalBson[models.PolTimeSerieEntry](cursor.Current)
		if decodeResult.HasFailed() {
			return result.Result[models.PolTimeSlice]{}.Failed(decodeResult.UnwrapError())
		}

		timeSlice = append(timeSlice, decodeResult.Expect())
	}

	return result.Success(timeSlice)
}
