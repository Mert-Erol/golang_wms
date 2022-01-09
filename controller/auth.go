package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mert-erol/golang_wms/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var db *sql.DB

type User struct {
	UserID   int
	UserName string
	Roles    int
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.gohtml")
		return
	}

	db = database.ConnectDB()

	username := r.FormValue("username")
	password := r.FormValue("password")

	var UserName string
	var UserPassword string
	var UserRole string

	err := db.QueryRow("SELECT username, password,role FROM users WHERE username=?", username).Scan(&UserName, &UserPassword, &UserRole)

	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(UserPassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	c := &http.Cookie{
		Name:  "UserRole",
		Value: UserRole,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", 301)

}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/signup.gohtml")
		return
	}

	db = database.ConnectDB()

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server hatası, kullanıcı oluşturulamadı1.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Server hatası, kullanıcı oluşturulamadı2.", 500)
			return
		}

		w.Write([]byte("Kullanıcı Oluşturuldu!"))
		return
	case err != nil:
		http.Error(w, "Server hatası, kullanıcı oluşturulamadı3.", 500)
		return
	default:
		http.Redirect(w, r, "/login", 301)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:   "UserRole",
		Value:  "0",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", 301)
}
