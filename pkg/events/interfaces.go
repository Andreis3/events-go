package events

type EventInterface interface {
	GetName() string
	GetDateTime() string
	GetPayload() any
}

type EventHandlerInterface interface {
	HandleEvent(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
