package define;

type Handler interface {
	Log(ie IEntity) error
}
