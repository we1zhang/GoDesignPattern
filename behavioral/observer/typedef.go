package observer

type Event struct {
	Data int64
}

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Register(Observer)
	UnRegister(Observer)
	Notify(Event)
}

//type ClickListener interface {
//	OnClick(event Event)
//}


