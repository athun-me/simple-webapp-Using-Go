package main

import (
	"fmt"
	"github.com/athun/controller"
	"net/http"
)

func main() {

	fmt.Println("server started localhost:7000")
	http.HandleFunc("/home", controller.HomePage)
	http.HandleFunc("/login", controller.LoginPage)
	http.HandleFunc("/signup", controller.SignUpPage)
	http.HandleFunc("/admin", controller.AdminPanel)

	http.ListenAndServe(":7000", nil)
}
