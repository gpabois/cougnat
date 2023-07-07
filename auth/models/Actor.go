package auth_models

type Actor interface {
	IsSuper() bool
}
