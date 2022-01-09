package controller

import (
	"database/sql"
	"github.com/mert-erol/golang_wms/database"
	"github.com/mert-erol/golang_wms/utils"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if utils.IsloggedIn(r) == "0" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		http.ServeFile(w, r, "views/dashboard.gohtml")
		return
	}
}

func Shelfs(w http.ResponseWriter, r *http.Request) {
	if utils.IsloggedIn(r) == "0" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/shelfs.gohtml")
		return
	}

	db = database.ConnectDB()

	shelfName := r.FormValue("shelf_name")
	shelfCapacity := r.FormValue("shelf_capacity")

	var shelf string

	err := db.QueryRow("SELECT name FROM shelfs WHERE name=?", shelfName).Scan(&shelf)

	switch {
	case err == sql.ErrNoRows:
		_, err = db.Exec("INSERT INTO shelfs(name, capacity) VALUES(?, ?)", shelfName, shelfCapacity)
		if err != nil {
			http.Error(w, "Server hatası, raf oluşturulamadı.", 500)
			return
		}

		w.Write([]byte("Raf Oluşturuldu!"))
		return
	case err != nil:
		http.Error(w, "Server hatası, raf oluşturulamadı.", 500)
		return
	default:
		http.Redirect(w, r, "/shelfs", 301)
	}
}
