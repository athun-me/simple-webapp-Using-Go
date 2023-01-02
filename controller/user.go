package controller

import (
	"fmt"
	"github.com/athun/config"
	"github.com/athun/models"
	"html/template"
	"net/http"
)

type UserInput struct {
	FName    string
	UName    string
	Lame     string
	Password string
	Email    string
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		tmp, err := template.ParseFiles("view/signup.html")
		if err != nil {
			panic(err)
		}
		tmp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		RegisterInput := &UserInput{
			FName:    r.Form.Get("fname"),
			UName:    r.Form.Get("uname"),
			Lame:     r.Form.Get("lname"),
			Password: r.Form.Get("password"),
			Email:    r.Form.Get("email"),
		}
		DB := config.Dbconnect()
		var tempUser models.User
		result := DB.First(&tempUser, "uname LIKE ?", RegisterInput.UName)
		if result.Error != nil {
			DB.Create(&models.User{
				FName:    RegisterInput.FName,
				Lame:     RegisterInput.Lame,
				UName:    RegisterInput.UName,
				Email:    RegisterInput.Email,
				Password: RegisterInput.Password,
			})
			fmt.Println(RegisterInput.Password)
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			http.Redirect(w, r, "/signup", http.StatusFound)

		}

	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tem, err := template.ParseFiles("view/login.html")
		if err != nil {
			panic(err)
		}
		tem.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			UName:    r.Form.Get("uname"),
			Password: r.Form.Get("password"),
		}

		DB := config.Dbconnect()
		var tempUser models.User
		result := DB.First(&tempUser, "u_name LIKE ? AND password LIKE ?", UserInput.UName, UserInput.Password)
		if result.Error != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			http.Redirect(w, r, "/home", http.StatusFound)
		}
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
	tmp.Execute(w, nil)
}

func AdminPanel(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	DB := config.Dbconnect()
	result := DB.Find(&users)

	if result.Error != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	tem, err := template.ParseFiles("view/admin.html")
	if err != nil {
		panic(err)
	}
	tem.Execute(w, users)
}
