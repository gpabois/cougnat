package mongo

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/reporting/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportRepository struct {
	coll *mongo.Collection
}

func (repo ReportRepository) Create(report models.Report) result.Result[models.ReportID] {
	res, err := repo.coll.InsertOne(context.TODO(), report)

	// The operations failed
	if err != nil {
		return result.Failed[models.ReportID](err)
	}

	objID := res.InsertedID.(primitive.ObjectID).String()

	return result.Success(objID)
}
