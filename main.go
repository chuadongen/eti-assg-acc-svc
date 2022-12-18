package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "strings"
)

type Passenger struct {
	PassengerId  int    `json:"passengerId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	EmailAddr    string `json:"emailAddr"`
	Password     string `json:"password"`
}

type Driver struct {
	DriverId      int    `json:"driverId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MobileNumber  string `json:"mobileNumber"`
	EmailAddr     string `json:"emailAddr"`
	Password      string `json:"password"`
	LicenseNumber string `json:"licenseNumber"`
	IdNumber      string `json:"idNumber"`
	DriverStatus  string `json:"driverStatus"`
}

func main() {
	//TODO:: Check that DB connections are properly closed
	//TODO:: http return status codes
	//TODO:: error handling
	router := mux.NewRouter()
	router.HandleFunc("/passenger", passengerEndpoint).Methods("POST", "PATCH")
	router.HandleFunc("/driver", driverEndpoint).Methods("POST", "PATCH")
	router.HandleFunc("/driver/logout", driverLogoutEndpoint).Methods("PATCH")
	router.HandleFunc("/auth/passenger", authPassengerEndpoint).Methods("GET")
	router.HandleFunc("/auth/driver", authDriverEndpoint).Methods("GET")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func passengerEndpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST": //Tested
		//Digest Passenger object from Body
		var p Passenger
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Insert Passenger into DB
		query := fmt.Sprintf("INSERT INTO passenger (firstName, lastName, mobileNumber, emailAddr, password) VALUES ('%s', '%s', '%s', '%s', '%s');", p.FirstName, p.LastName, p.MobileNumber, p.EmailAddr, p.Password)
		insert, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insert.Close()
		w.WriteHeader(http.StatusAccepted)
		defer db.Close()
	case "PATCH": //Tested
		//Digest Passenger object from Body
		var p Passenger
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if p.PassengerId == 0 { //0 is default value, which means it did not recieve PassengerId
			http.Error(w, "PassengerId missing", http.StatusBadRequest)
			return
		}
		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			panic(err.Error())
		}
		//Update Passenger in DB
		query := fmt.Sprintf("UPDATE passenger SET firstName='%s',lastName='%s',mobileNumber='%s',emailAddr='%s',password='%s'WHERE passengerId=%d;", p.FirstName, p.LastName, p.MobileNumber, p.EmailAddr, p.Password, p.PassengerId)
		update, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		update.Close()
		w.WriteHeader(http.StatusAccepted)
		defer db.Close()
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func driverEndpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST": //Tested
		//Digest Driver object from Body
		var d Driver
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Insert Driver into DB
		query := fmt.Sprintf("INSERT INTO driver (firstName, lastName, mobileNumber, emailAddr, password, idNumber, licenseNumber,driverStatus) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s','Offline');", d.FirstName, d.LastName, d.MobileNumber, d.EmailAddr, d.Password, d.IdNumber, d.LicenseNumber)
		insert, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insert.Close()
		w.WriteHeader(http.StatusAccepted)
		defer db.Close()
	case "PATCH": //Tested
		//Digest Driver object from Body
		var d Driver
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if d.DriverId == 0 { //0 is default value, which means it did not recieve DriverId
			http.Error(w, "DriverId missing", http.StatusBadRequest)
			return
		}

		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Update Driver in DB
		query := fmt.Sprintf("UPDATE driver SET firstName='%s',lastName='%s',mobileNumber='%s',emailAddr='%s',password='%s', idNumber='%s',licenseNumber='%s'WHERE driverId=%d;", d.FirstName, d.LastName, d.MobileNumber, d.EmailAddr, d.Password, d.IdNumber, d.LicenseNumber, d.DriverId)
		update, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		update.Close()
		w.WriteHeader(http.StatusAccepted)
		defer db.Close()
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func driverLogoutEndpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "PATCH": //Tested
		//Digest Driver object from Body
		var d Driver
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if d.DriverId == 0 { //0 is default value, which means it did not recieve DriverId
			http.Error(w, "DriverId missing", http.StatusBadRequest)
			return
		}

		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Update Driver in DB
		query := fmt.Sprintf("UPDATE driver SET driverStatus='Offline' WHERE driverId=%d;", d.DriverId)
		update, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		update.Close()
		w.WriteHeader(http.StatusAccepted)
		defer db.Close()
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func authPassengerEndpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method { //Tested
	case "GET":
		querystringmap := r.URL.Query()
		emailAddr := querystringmap.Get("emailAddr")
		password := querystringmap.Get("password")
		//Digest PassengerAuth info from query params
		if emailAddr == "" || password == "" {
			http.Error(w, "Data missing", http.StatusBadRequest)
			return
		}

		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			panic(err.Error())
		}
		//Select Passenger from DB
		query := fmt.Sprintf("SELECT * FROM  passenger WHERE password='%s' AND emailAddr ='%s';", password, emailAddr)
		var result Passenger
		err = db.QueryRow(query).Scan(&result.PassengerId, &result.FirstName, &result.LastName, &result.MobileNumber, &result.EmailAddr, &result.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(output))
		defer db.Close()
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func authDriverEndpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET": //Tested
		querystringmap := r.URL.Query()
		emailAddr := querystringmap.Get("emailAddr")
		password := querystringmap.Get("password")
		//Digest PassengerAuth info from query params
		if emailAddr == "" || password == "" {
			http.Error(w, "Data missing", http.StatusBadRequest)
			return
		}

		//Init DB Connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/etiassignone")
		if err != nil {
			panic(err.Error())
		}
		//Select Driver from DB
		query := fmt.Sprintf("SELECT * FROM  driver WHERE password='%s' AND emailAddr ='%s';", password, emailAddr)
		var result Driver
		err = db.QueryRow(query).Scan(&result.DriverId, &result.FirstName, &result.LastName, &result.MobileNumber, &result.EmailAddr, &result.Password, &result.IdNumber, &result.LicenseNumber, &result.DriverStatus)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Update Driver status to Available in DB
		statusUpdateQuery := fmt.Sprintf("UPDATE driver SET driverStatus='Available' WHERE driverId =%d;",result.DriverId)
		update, err := db.Query(statusUpdateQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		update.Close()
		output, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(output))
		defer db.Close()
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
