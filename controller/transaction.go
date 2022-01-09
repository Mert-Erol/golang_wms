package controller

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/mert-erol/golang_wms/database"
	"github.com/mert-erol/golang_wms/models"
	"github.com/mert-erol/golang_wms/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if utils.IsloggedIn(r) == "0" {	// Auth control. If there is no cookie send the user to the login page
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

// Reading csv data
func AddReceipt(w http.ResponseWriter, r *http.Request) {
	if utils.IsloggedIn(r) == "0" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/receipt.gohtml")
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("receiptFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("uploads", "upload-*.csv")	// Allow only .csv files and create temp file
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	ReadCsvFile(tempFile.Name())
	// return that we have successfully uploaded our file!

}

func ReadCsvFile(FileName string) {

	db = database.ConnectDB()

	records, err := ReadData(FileName)	// Getting data
	if err != nil {
		log.Fatal()
	}
	for _, record := range records {	// parse data row to row
		receipt := models.ReceiptDetail{
			StockCode:       record[0],
			Quantity:        record[1],
			TransactionType: 0,
		}

		// insert data one by one
		_, err = db.Exec("INSERT INTO transactions(stock_code, quantity, type) VALUES(?, ?, ?)", &receipt.StockCode, &receipt.Quantity, &receipt.TransactionType)
		if err != nil {
			fmt.Println(err)
			return
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
