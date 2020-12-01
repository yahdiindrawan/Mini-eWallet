package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yahdiindrawan/Mini-eWallet/migrate"
	"io/ioutil"
	"log"
	"net/http"
)

var db *gorm.DB

func HandleRequests()  {
	log.Println("Start the development server at http://127.0.0.1:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)

	myRouter.HandleFunc("/api/users", CreateUser).Methods("POST")
	myRouter.HandleFunc("/api/users", GetUsers).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	//myRouter.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	//myRouter.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func HomePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user migrate.User
	json.Unmarshal(payloads, &user)
	db.Create(&user)
	res := migrate.Result{Code: 200, Data: user, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	users := []migrate.User{}

	db.Find(&users)

	res := migrate.Result{Code: 200, Data: users, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user migrate.User
	db.First(&user, userID)

	res := migrate.Result{Code: 200, Data: user, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

