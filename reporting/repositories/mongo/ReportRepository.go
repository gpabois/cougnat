package mongo

import (
	"context"

	"github.com/gpabois/cougnat/core/mongo"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/reporting/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportRepository struct {
	coll mongo.Collection[models.Report]
}

func (repo ReportRepository) Create(report models.Report) result.Result[models.ReportID] {
	return result.Map(
		repo.coll.InsertOne(context.TODO(), report),
		func(obj primitive.ObjectID) string {
			return obj.String()
		},
	)
}

func (repo ReportRepository) GetById(id models.ReportID) result.Result[option.Option[models.Report]] {
	objId, err := primitive.ObjectIDFromHex(id)

	if err == nil {
		return result.Failed[option.Option[models.Report]](err)
	}

	return repo.coll.FindOne(context.TODO(), bson.D{{"_id", objId}})
}
