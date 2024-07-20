package types

type Subscriber struct {
	Channel chan interface{}
	Unsubscribe chan bool
}