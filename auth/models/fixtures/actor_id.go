package fixtures

import "github.com/gpabois/cougnat/auth/models"

func RandomAnonymousID() models.ActorID {
	return models.AnonymousID("random")
}
