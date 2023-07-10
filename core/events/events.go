package events

import "github.com/gpabois/gostd/result"

//go:generate mockery
type IEventService interface {
	Notify(eventName string, data any) result.Result[bool]
	Listen(eventName string) result.Result[chan any]
}
