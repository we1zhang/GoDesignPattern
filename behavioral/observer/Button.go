package observer

type Button struct {
	Text      string
	observers []Observer
}

func (b *Button) Register(observer Observer) {
	if b.observers == nil {
		b.observers = make([]Observer, 0)
	}
	b.observers = append(b.observers, observer)
}

func (b *Button) UnRegister(observer Observer) {
	for index, o := range b.observers {
		if o == observer {
			b.observers = append(b.observers[:index], b.observers[index+1:]...)
			break
		}
	}
}

func (b *Button) Notify(event Event) {
	for _, o := range b.observers {
		if o != nil {
			o.OnNotify(event)
		}
	}
}

//
//func (b *Button) SetOnClickListener(listener ClickListener)  {
//
//}
