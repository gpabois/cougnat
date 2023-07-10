package auth_pg

import (
	auth_models "github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/pg"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/jackc/pgx/v5"
)

// Scan the row to return an ActorID
// Expected order: ID, Nature
func ScanActorID(actorID *auth_models.ActorID) pg.ScanCommands {
	return pg.Scan(&actorID.ID).Append(pg.Scan(&actorID.Nature))
}

func ScanOptionalActorID(optActorID *option.Option[auth_models.ActorID]) pg.ScanCommands {
	return []pg.ScanCommand{{
		Scan: func(rows pgx.Rows) result.Result[bool] {
			actorID := auth_models.ActorID{}
			res := ScanActorID(&actorID).Exec(rows)
			if res.HasFailed() {
				return res
			}
			if actorID.ID.IsNone() && actorID.Nature == "" {
				*optActorID = option.None[auth_models.ActorID]()
			}
			*optActorID = option.Some(actorID)
			return result.Success(true)
		},
	}}
}
