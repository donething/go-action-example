package main

import (
	"log"
)

// notifier是一个定义了通知行为的接口
type notifier interface {
	notify()
}

// user在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify是使用指针接收者实现的方法
func (u *user) notify() {
	log.Printf("Send User Email To %s<%s>\n", u.name, u.email)
}

// admin 定义了程序里的管理员
type admin struct {
	name  string
	email string
}

type admin2 struct {
	user
	level string
}

// notify 使用指针接收者实现了 notifier 接口
func (a *admin) notify() {
	log.Printf("Send Admin Email To %s<%s>\n", a.name, a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	lisa := admin{"Lisa", "lisa@email.com"}
	gate := admin2{user{"Gate", "gate@email.com"}, "1"}

	sendNotification(&gate)
	// Cannot use 'bill' (type user) as type notifier,
	// Type does not implement 'notifier' as 'notify' method has a pointer receiver l
	// sendNotification(bill)
	sendNotification(&bill)
	sendNotification(&lisa)

	david := admin2{level: "1"}
	david.name = "David"
	david.email = "david@email.com"
	sendNotification(&david)
}

// sendNotification 接受一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
