package guards

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/auth/services"
	"github.com/gpabois/cougnat/core/result"
)

// Check if the user has a specific permission
func CheckAccessControl(authz services.IAuthorizationService) func(acl models.AccessControl) result.Result[bool] {
	return func(acl models.AccessControl) result.Result[bool] {
		return authz.HasPermission(acl.Actor, acl.Permission, acl.Object)
	}
}
