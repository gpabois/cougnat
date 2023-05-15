package guards

import (
	"context"
	"errors"

	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/auth/utils"
	"github.com/gpabois/cougnat/core/result"
)

// Check if the user is authenticated
func IsAuthenticated(ctx context.Context) result.Result[models.ActorID] {
	return utils.GetCurrentActorID(ctx).IntoResult(errors.New("not authenticated"))
}
