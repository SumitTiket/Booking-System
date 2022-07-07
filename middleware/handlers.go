package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SumitTiket/booking-system/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//create connection with postgresql :/

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("USER"), os.Getenv("NAME"), os.Getenv("PASSWORD"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the db!")
	return db
}

func GetAllHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := getAllHotels()

	if err != nil {
		log.Fatalf("Unable to get all users.%v", err)
	}
	json.NewEncoder(w).Encode(hotels)
}

//handlers ----------------->>>>>>>>>>>>>>>

func getAllHotels() ([]models.Hotels, error) {
	db := createConnection()

	defer db.Close()

	var hotels []models.Hotels

	sqlStatement := `SELECT * FROM hotels`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query.%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var hotel models.Hotels
		err = rows.Scan(&hotel.ID, &hotel.Name, &hotel.Rooms, &hotel.Location)
		if err != nil {
			log.Fatalf("unable to scan the row.%v", err)
		}

		hotels = append(hotels, hotel)
	}
	return hotels, err
}
