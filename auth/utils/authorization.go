package utils

import (
	"context"

	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/option"
)

func SetCurrentActorID(ctx context.Context, actorID models.ActorID) context.Context {
	return context.WithValue(ctx, "CurrentActor.ID", actorID)
}

func GetCurrentActorID(ctx context.Context) option.Option[models.ActorID] {
	return models.ActorID_TryFromAny(ctx.Value("CurrentActor.ID"))
}
