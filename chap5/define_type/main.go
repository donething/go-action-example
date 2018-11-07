package main

import "log"

type User struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type Admin struct {
	User
	level string
}

// notify 使用值接收者实现了一个方法
func (u User) notify() {
	log.Printf("Send User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *User) changeEmail(email string) {
	u.email = email
}

type Duration int64

func main() {
	var now Duration = 12
	var time int64
	time = int64(now)
	log.Println(time)

	var bill = User{name: "Bill", email: "bill@email.com", ext: 124, privileged: true}
	bill.notify()
	var lisa = &User{name: "Lisa", email: "lisa@email.com", ext: 123, privileged: true}
	lisa.notify()

	bill.changeEmail("billnew@email.com")
	lisa.changeEmail("lisanew@email.com")
	bill.notify()
	lisa.notify()

	var admin = Admin{
		User: User{
			name:       "lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "",
	}
	admin.notify()
}
