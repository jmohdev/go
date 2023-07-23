package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "jmoh", Email: "jmoh.developer@gmail.com", Age: 23}
	user2 := User{Name: "ex", Email: "ex@gmail.com", Age: 32}
	users := []User{user, user2}
	tmpl, err := template.New("Templ1").ParseFiles("templates/tmpl_01", "templates/tmpl_02")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl_02", users)
}
