package auth_pg

import (
	auth_models "github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/pg"
)

// Scan the row to return an ActorID
// Expected order: ID, Nature
func ScanActorID(actorID *auth_models.ActorID) pg.ScanCommands {
	return pg.ScanStringOption(&actorID.ID).Append(pg.ScanPrimaryType(&actorID.Nature))

}
