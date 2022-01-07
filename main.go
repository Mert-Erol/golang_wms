package main

import (
	"database/sql"
	"encoding/csv"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
var err error

var tpl *template.Template

type User struct {
	UserName string
	Password string
}

type Dispatch struct {
	DispatchNote string
	Type         int
	StockCode    string
	Quantity     int
}

var databaseUsers = map[string]User{}      //user ID, user object
var databaseSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server hatası, kullanıcı oluşturulamadı.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server hatası, kullanıcı oluşturulamadı.", 500)
			return
		}

		res.Write([]byte("Kullanıcı Oluşturuldu!"))
		return
	case err != nil:
		http.Error(res, "Server hatası, kullanıcı oluşturulamadı.", 500)
		return
	default:
		http.Redirect(res, req, "/login", 301)
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.gohtml")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	//sID := uuid.NewV4()
	c := &http.Cookie{
		Name:  "user",
		Value: username,
	}
	http.SetCookie(w, c)
	databaseSessions[c.Value] = username

	http.Redirect(w, r, "/dashboard", 301)
	//res.Write([]byte("Hello" + databaseUsername))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
func dashboard(w http.ResponseWriter, r *http.Request) {
	if isloggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	user, _ := r.Cookie("user")

	tpl.ExecuteTemplate(w, "dashboard.gohtml", user)
}
func dispatch(w http.ResponseWriter, r *http.Request)  {

	file, _, _ := r.FormFile("file")

	records, err := ReadData(file)
	if err != nil {
		log.Fatal()
	}

	for _, record := range records {
		dispatch := Dispatch{
			DispatchNote: record[0],
			Type:         record[1],
			StockCode:    record[2],
			Quantity:     record[3],
		}

		_, err = db.Exec("INSERT INTO transactions(dispatch_note, type, stock_code, quantity) VALUES(?, ?, ?, ?)", dispatch.DispatchNote, dispatch.Type, dispatch.StockCode, dispatch.Quantity)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func ReadData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func main() {
	db, err = sql.Open("mysql", "root:@/wms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/dispatch", dispatch)
	http.ListenAndServe(":4447", nil)
}
