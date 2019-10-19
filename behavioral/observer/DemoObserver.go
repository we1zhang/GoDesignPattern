package observer

import "fmt"

type DemoObserver struct {
}

func (do *DemoObserver) OnNotify(event Event) {
	fmt.Println("i am demoObserver,i receive event=", event)
}
