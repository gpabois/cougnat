package models

import "github.com/gpabois/cougnat/core/option"

type User struct {
	ID        option.Option[string]
	UserName  string
	FirstName option.Option[string]
	LastName  option.Option[string]
	Email     string
	Password  string
	// Has all the permissions
	IsSuper bool
}
