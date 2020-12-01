//package main
//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"github.com/jinzhu/gorm"
//	"log"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"net/http"
//)
//
//var db *gorm.DB
//var err error
//
//
//
////SCHEMA
//type User struct {
//	ID 				int 			`json:"id"`
//	Username 		string 			`json:"username"`
//	Email 			string 			`json:"email"`
//	Password 		string 			`json:"password"`
//}
//
//type User_Balance struct {
//	ID 						int					 	`json:"id"`
//	User_ID 				int 					`json:"user_id"`
//	Balance 				int 					`json:"balance"`
//	Balance_Achieve 		int 					`json:"balance_achieve"`
//}
//
//type User_Balance_History struct {
//	ID 					int 	`json:"id"`
//	User_Balance_ID 	int 	`json:"user_balance_id"`
//	Balance_Before 		int 	`json:"balance_before"`
//	Balance_After 		int 	`json:"balance_after"`
//	Activity		 	string 	`json:"activity"`
//	Type 				string 	`gorm:"type:enum('credit','debit')" json:"type"`
//	IP 					string 	`json:"ip"`
//	Location 			string 	`json:"location"`
//	User_Agent 			string 	`json:"user_agent"`
//	Author 				string 	`json:"author"`
//}
//
//type Blance_Bank struct {
//	ID 						int 					`json:"id"`
//	Balance 				int 					`json:"balance"`
//	Balance_Achieve 		int 					`json:"balance_achieve"`
//	Code 					string 					`json:"code"`
//	Enable 					bool 					`json:"enable"`
//}
//
//type Blance_Bank_History struct {
//	ID 					int 	`json:"id"`
//	Balance_Bank_ID 	int 	`json:"balance_bank_id"`
//	Balance_Before 		int 	`json:"balance_before"`
//	Balance_After 		int 	`json:"balance_after"`
//	Activity 			string	`json:"activity"`
//	Type 				string 	`gorm:"type:enum('credit','debit')" json:"type"`
//	IP 					string 	`json:"ip"`
//	Location 			string 	`json:"location"`
//	User_Agent 			string 	`json:"user_agent"`
//	Author 				string 	`json:"author"`
//}
//
//type Result struct {
//	Code	int 		`json:"code"`
//	Data	interface{} `json:"data"`
//	Message string 		`json:"message"`
//}
//
//func main()  {
//	db, err := gorm.Open("mysql", "root:@/e-wallet?charset=utf8&parseTime=True")
//
//	if err != nil {
//		log.Println("Connection Failed", err)
//	}else{
//		log.Println("Connection established")
//	}
//	//migrate the Schema
//	db.AutoMigrate(&User{}, &User_Balance{}, &User_Balance_History{}, &Blance_Bank{}, &Blance_Bank_History{})
//	db.Model(&User_Balance{}).AddForeignKey("user_id","users(id)","cascade","cascade")
//	db.Model(&User_Balance_History{}).AddForeignKey("id","user_balances(id)","cascade","cascade")
//	db.Model(&Blance_Bank_History{}).AddForeignKey("balance_bank_id","blance_banks(id)","cascade","cascade")
//
//	handleRequests()
//}
//
//func handleRequests()  {
//	log.Println("Start the development server at http://127.0.0.1:9999")
//
//	myRouter := mux.NewRouter().StrictSlash(true)
//
//	myRouter.HandleFunc("/", homePage)
//
//
//	log.Fatal(http.ListenAndServe(":9999", myRouter))
//}
//
//func homePage(w http.ResponseWriter, r *http.Request)  {
//	fmt.Fprintf(w, "Welcome")
//}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	ID 	int	`json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserBalance struct {
	ID             int `json:"id"`
	UserId         int `json:"user_id"`
	Balance        int `json:"balance"`
	BalanceAchieve int `json:"balance_achieve"`
}

type UserBalanceHistory struct {
	ID 					int 	`json:"id"`
	User_Balance_ID 	int 	`json:"user_balance_id"`
	Balance_Before 		int 	`json:"balance_before"`
	Balance_After 		int 	`json:"balance_after"`
	Activity		 	string 	`json:"activity"`
	Type 				string 	`gorm:"type:enum('credit','debit')" json:"type"`
	IP 					string 	`json:"ip"`
	Location 			string 	`json:"location"`
	User_Agent 			string 	`json:"user_agent"`
	Author 				string 	`json:"author"`
}

type Blance_Bank struct {
	ID 					int 	`json:"id"`
	Balance 			int 	`json:"balance"`
	Balance_Achieve 	int 	`json:"balance_achieve"`
	Code 				string 	`json:"code"`
	Enable 				bool 	`json:"enable"`
}

type BlanceBankHistory struct {
	ID 					int 	`json:"id"`
	Balance_Bank_ID 	int 	`json:"balance_bank_id"`
	Balance_Before 		int 	`json:"balance_before"`
	Balance_After 		int 	`json:"balance_after"`
	Activity 			string	`json:"activity"`
	Type 				string 	`gorm:"type:enum('credit','debit')" json:"type"`
	IP 					string 	`json:"ip"`
	Location 			string 	`json:"location"`
	User_Agent 			string 	`json:"user_agent"`
	Author 				string 	`json:"author"`
}

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql","root:@/e-wallet?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed", err)
	} else {
		log.Println("Connection established ")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserBalance{})
	db.AutoMigrate(&UserBalanceHistory{})
	db.AutoMigrate(&Blance_Bank{})
	db.AutoMigrate(&BlanceBankHistory{})
	db.Model(&UserBalance{}).AddForeignKey("user_id","users(id)","cascade","cascade")
	db.Model(&UserBalanceHistory{}).AddForeignKey("id","user_balances(id)","cascade","cascade")
	db.Model(&BlanceBankHistory{}).AddForeignKey("balance_bank_id","blance_banks(id)","cascade","cascade")
	HandleRequests()

}

func HandleRequests()  {
	log.Println("Start the development server at http://127.0.0.1:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)

	myRouter.HandleFunc("/api/users", CreateUser).Methods("POST")
	myRouter.HandleFunc("/api/users", GetUsers).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	myRouter.HandleFunc("/api/userbalances", CreateUserBalance).Methods("POST")
	myRouter.HandleFunc("/api/userbalances", GetUsersBalances).Methods("GET")
	myRouter.HandleFunc("/api/userbalances/{id}", GetUserBalance).Methods("GET")
	myRouter.HandleFunc("/api/userbalances/{id}", UpdateUserBalance).Methods("PUT")
	myRouter.HandleFunc("/api/userbalances/{id}", DeleteUserBalance).Methods("DELETE")

	myRouter.HandleFunc("/api/userbalancehistories", CreateUserBalanceHistory).Methods("POST")
	myRouter.HandleFunc("/api/userbalancehistories", GetUserBalanceHistories).Methods("GET")
	myRouter.HandleFunc("/api/userbalancehistories/{id}", GetUserBalanceHistory).Methods("GET")
	myRouter.HandleFunc("/api/userbalancehistories/{id}", UpdateUserBalanceHistory).Methods("PUT")
	myRouter.HandleFunc("/api/userbalancehistories/{id}", DeleteUserBalanceHistory).Methods("DELETE")

	myRouter.HandleFunc("/api/blancebanks", CreateBlanceBank).Methods("POST")
	myRouter.HandleFunc("/api/blancebanks", GetBlanceBanks).Methods("GET")
	myRouter.HandleFunc("/api/blancebanks/{id}", GetBlanceBank).Methods("GET")
	myRouter.HandleFunc("/api/blancebanks/{id}", UpdateBlanceBank).Methods("PUT")
	myRouter.HandleFunc("/api/blancebanks/{id}", DeleteBlanceBank).Methods("DELETE")

	myRouter.HandleFunc("/api/blancebankhistories", CreateBlanceBankHistory).Methods("POST")
	myRouter.HandleFunc("/api/blancebankhistories", GetBlanceBankHistories).Methods("GET")
	myRouter.HandleFunc("/api/blancebankhistories/{id}", GetBlanceBankHistory).Methods("GET")
	myRouter.HandleFunc("/api/blancebankhistories/{id}", UpdateBlanceBankHistory).Methods("PUT")
	myRouter.HandleFunc("/api/blancebankhistories/{id}", DeleteBlanceBankHistory).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func HomePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(payloads, &user)
	db.Create(&user)
	res := Result{Code: 200, Data: user, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	users := []User{}

	db.Find(&users)

	res := Result{Code: 200, Data: users, Message: "Success get users"}
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

	var user User
	db.First(&user, userID)

	res := Result{Code: 200, Data: user, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	UserID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var userUpdate User
	json.Unmarshal(payloads, &userUpdate)

	var user User
	db.First(&user, UserID)
	db.Model(&user).Updates(userUpdate)

	res := Result{Code: 200, Data: user, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	UserID := vars["id"]

	var user User
	db.First(&user, UserID)
	db.Delete(&user)

	res := Result{Code:200, Message: "Success delete user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}



func CreateUserBalance(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var userbalance UserBalance
	json.Unmarshal(payloads, &userbalance)
	db.Create(&userbalance)
	res := Result{Code: 200, Data: userbalance, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetUsersBalances(w http.ResponseWriter, r *http.Request)  {
	userbalances := []UserBalance{}

	db.Find(&userbalances)

	res := Result{Code: 200, Data: userbalances, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetUserBalance(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalanceID := vars["id"]

	var userbalance UserBalance
	db.First(&userbalance, userbalanceID)

	res := Result{Code: 200, Data: userbalance, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUserBalance(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalanceID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var userbalanceUpdate UserBalance
	json.Unmarshal(payloads, &userbalanceUpdate)

	var userbalance UserBalance
	db.First(&userbalance, userbalanceID)
	db.Model(&userbalance).Updates(userbalanceUpdate)

	res := Result{Code: 200, Data: userbalance, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUserBalance(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalanceID := vars["id"]

	var userbalance UserBalance
	db.First(&userbalance, userbalanceID)
	db.Delete(&userbalance)

	res := Result{Code:200, Message: "Success delete user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}



func CreateUserBalanceHistory(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var userbalancehistory UserBalanceHistory
	json.Unmarshal(payloads, &userbalancehistory)
	db.Create(&userbalancehistory)
	res := Result{Code: 200, Data: userbalancehistory, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetUserBalanceHistories(w http.ResponseWriter, r *http.Request)  {
	userbalancehistories := []UserBalanceHistory{}

	db.Find(&userbalancehistories)

	res := Result{Code: 200, Data: userbalancehistories, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetUserBalanceHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalancehistoryID := vars["id"]

	var userbalancehistory UserBalanceHistory
	db.First(&userbalancehistory, userbalancehistoryID)

	res := Result{Code: 200, Data: userbalancehistory, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUserBalanceHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalancehistoryID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var userbalancehistoryUpdate UserBalanceHistory
	json.Unmarshal(payloads, &userbalancehistoryUpdate)

	var userbalancehistory UserBalanceHistory
	db.First(&userbalancehistory, userbalancehistoryID)
	db.Model(&userbalancehistory).Updates(userbalancehistoryUpdate)

	res := Result{Code: 200, Data: userbalancehistory, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUserBalanceHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userbalancehistoryID := vars["id"]

	var userbalancehistory UserBalanceHistory
	db.First(&userbalancehistory, userbalancehistoryID)
	db.Delete(&userbalancehistory)

	res := Result{Code:200, Message: "Success delete user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


func CreateBlanceBank(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var blancebank Blance_Bank
	json.Unmarshal(payloads, &blancebank)
	db.Create(&blancebank)
	res := Result{Code: 200, Data: blancebank, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetBlanceBanks(w http.ResponseWriter, r *http.Request)  {
	blancebanks := []Blance_Bank{}

	db.Find(&blancebanks)

	res := Result{Code: 200, Data: blancebanks, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetBlanceBank(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankID := vars["id"]

	var blancebank Blance_Bank
	db.First(&blancebank, blancebankID)

	res := Result{Code: 200, Data: blancebank, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateBlanceBank(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var blancebankUpdate Blance_Bank
	json.Unmarshal(payloads, &blancebankUpdate)

	var blancebank Blance_Bank
	db.First(&blancebank, blancebankID)
	db.Model(&blancebank).Updates(blancebankUpdate)

	res := Result{Code: 200, Data: blancebank, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteBlanceBank(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankID := vars["id"]

	var blancebank Blance_Bank
	db.First(&blancebank, blancebankID)
	db.Delete(&blancebank)

	res := Result{Code:200, Message: "Success delete user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}



func CreateBlanceBankHistory(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var blancebankhistory BlanceBankHistory
	json.Unmarshal(payloads, &blancebankhistory)
	db.Create(&blancebankhistory)
	res := Result{Code: 200, Data: blancebankhistory, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetBlanceBankHistories(w http.ResponseWriter, r *http.Request)  {
	blancebankhistories := []BlanceBankHistory{}

	db.Find(&blancebankhistories)

	res := Result{Code: 200, Data: blancebankhistories, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetBlanceBankHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankhistoryID := vars["id"]

	var blancebankhistory BlanceBankHistory
	db.First(&blancebankhistory, blancebankhistoryID)

	res := Result{Code: 200, Data: blancebankhistory, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateBlanceBankHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankhistoryID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var blancebankhistoryUpdate BlanceBankHistory
	json.Unmarshal(payloads, &blancebankhistoryUpdate)

	var blancebankhistory BlanceBankHistory
	db.First(&blancebankhistory, blancebankhistoryID)
	db.Model(&blancebankhistory).Updates(blancebankhistoryUpdate)

	res := Result{Code: 200, Data: blancebankhistory, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteBlanceBankHistory(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	blancebankhistoryID := vars["id"]

	var blancebankhistory BlanceBankHistory
	db.First(&blancebankhistory, blancebankhistoryID)
	db.Delete(&blancebankhistory)

	res := Result{Code:200, Message: "Success delete user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

