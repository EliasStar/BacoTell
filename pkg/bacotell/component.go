package bacotell

type Component interface {
	CustomID() (string, error)
	Handle(HandleProxy) error
}

type HandleProxy interface {
	InteractionProxy
}
