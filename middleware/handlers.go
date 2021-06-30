package middleware

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/Akhenaten-Sama/go-gres/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
)

func LoadDb() *sql.DB{
	err := godotenv.Load(".env")
	if err != nil {
	log.Fatalf("error loading %s", err)
	} else {
	log.Println("env loaded")
	}	

	username := os.Getenv("APP_DB_USERNAME")
pass := os.Getenv("APP_DB_PASSWORD")
host := os.Getenv("APP_DB_HOST")
dbName := os.Getenv("APP_DB_NAME")
port := os.Getenv("APP_DB_PORT")
connects := fmt.Sprintf("host=%s port=%s user=%s "+
"password=%s dbname=%s sslmode=disable", host, port, username, pass, dbName)
db, err := sql.Open("postgres", connects)
if err != nil{
panic(err)
}
err = db.Ping()
if err != nil{
panic(err)
}
fmt.Println("Successfully connected!")

return db
}


func GetBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books := []models.Book{}
	db := LoadDb()
	statement := `select * from books`
	rows, err := db.Query(statement)
	if err != nil {
	log.Println(fmt.Sprintf("error occurred doing this: %s", err))
	}
	for rows.Next() {
	err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year)
	if err != nil {
	log.Println(fmt.Sprintf("error occurred doing this: %s", err))
	}
	books = append(books, book)
	
	}
	json.NewEncoder(w).Encode(books)
	
	defer db.Close()
	defer rows.Close()
	
	}


func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

		statement := `select * from books where id=$1`
		db := LoadDb()
		defer db.Close()
		params := mux.Vars(r)
		rows := db.QueryRow(statement, params["id"])
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year)
		if err != nil {
		log.Println(fmt.Sprintf("error occurred doing this: %s", err))
		}
		json.NewEncoder(w).Encode(book)
		}



func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	
			db := LoadDb()
			json.NewDecoder(r.Body).Decode(&book)
			defer db.Close()
			statement := `insert into books(title, author, year) values($1,$2,$3) returning id, author, title, year`
			err := db.QueryRow(statement, book.Title, book.Author, book.Year).Scan(&book.ID, &book.Author, &book.Title, &book.Year)
			if err != nil {
			log.Println(fmt.Sprintf("error occurred while doing this: %s", err))
			}
			json.NewEncoder(w).Encode(book)
			}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
				params := mux.Vars(r)
				var book models.Book

				db := LoadDb()
				defer db.Close()
				statement := `update books set author=$2, title=$3, year=$4 where id=$1`
				_, err := db.Exec(statement, params["id"], book.Author, book.Title, book.Year)
				if err != nil {
				log.Println(fmt.Sprintf("error occurred doing this: %s", err))
				}
				json.NewEncoder(w).Encode("row updated")
				}


				func DeleteBook(w http.ResponseWriter, r *http.Request) {
					db := LoadDb()
					defer db.Close()
					params := mux.Vars(r)
					statement := `delete from books where id=$1 `
					_, err := db.Exec(statement, params["id"])
					if err != nil {
					log.Println(fmt.Sprintf("error occurred doing this: %s", err))
					} else {
					json.NewEncoder(w).Encode("row deleted")
					}
					
					}

