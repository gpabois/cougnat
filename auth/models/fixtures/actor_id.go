package fixtures

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/option"
)

func RandomAnonymousID() models.ActorID {
	return models.AnonymousID(option.Some("test"))
}
