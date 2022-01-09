package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mert-erol/golang_wms/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var db *sql.DB

// Auth operation
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {	// Checking if request is POST
		http.ServeFile(w, r, "views/login.gohtml")	//IF not redirect to login page
		return
	}

	db = database.ConnectDB()	// Connect to DB

	// Form fields from the login page
	username := r.FormValue("username")
	password := r.FormValue("password")

	var UserName string
	var UserPassword string
	var UserRole string

	// Checiking if user exist
	err := db.QueryRow("SELECT username, password,role FROM users WHERE username=?", username).Scan(&UserName, &UserPassword, &UserRole)

	// IF not refresh the page
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	// Checking the password
	err = bcrypt.CompareHashAndPassword([]byte(UserPassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	// If everything is okay create the cookie
	c := &http.Cookie{
		Name:  "UserRole",
		Value: UserRole,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", 301)

}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // Checking if request is POST
		http.ServeFile(w, r, "views/signup.gohtml")
		return
	}

	db = database.ConnectDB() // Connect to DB

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user string

	// Checking the same username has been taken
	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)	// Encode the password value
		if err != nil {
			http.Error(w, "Server hatası, kullanıcı oluşturulamadı.", 500)
			return
		}

		// Creating user
		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Server hatası, kullanıcı oluşturulamadı.", 500)
			return
		}

		w.Write([]byte("Kullanıcı Oluşturuldu!"))
		return
	case err != nil:
		http.Error(w, "Server hatası, kullanıcı oluşturulamadı.", 500)
		return
	default:
		http.Redirect(w, r, "/login", 301)
	}
}

// Logout operation
func Logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{	// Destroying tho cookie
		Name:   "UserRole",
		Value:  "0",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", 301)
}
