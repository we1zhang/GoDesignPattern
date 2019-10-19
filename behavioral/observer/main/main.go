package main

import (
	"project/GoDesignPattern/behavioral/observer"
)

func main() {

	button := &observer.Button{
		Text: "Hello",
	}

	demoObserver := &observer.DemoObserver{}

	button.Register(demoObserver)
	button.Notify(observer.Event{Data: 123})

	button.UnRegister(demoObserver)
	button.Notify(observer.Event{Data: 456})

	//button.SetOnClickListener(&DemoClickListener{})
}

//
//type DemoClickListener struct {
//}
//
//func (dcl *DemoClickListener) OnClick(event observer.Event) {
//	fmt.Println("i am demoClickListener,event=", event)
//}
