package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var err error


//SCHEMA
type User struct {
	ID 				int 			`gorm:"primaryKey" json:"id"`
	Username 		string 			`json:"username"`
	Email 			string 			`json:"email"`
	Password 		string 			`json:"password"`
	User_Balances 	[]User_Balance 	`gorm:"ForeignKey:User_ID" json:"user_balances"`
}

type User_Balance struct {
	ID 						int					 	`gorm:"primaryKey" json:"id"`
	User_ID 				int 					`json:"user_id"`
	Balance 				int 					`json:"balance"`
	Balance_Achieve 		int 					`json:"balance_achieve"`
	User_Balance_History	[]User_Balance_History	`gorm:"ForeignKey:ID" json:"user_balance_history"`
}

type User_Balance_History struct {
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
	ID 						int 					`gorm:"primaryKey" json:"id"`
	Balance 				int 					`json:"balance"`
	Balance_Achieve 		int 					`json:"balance_achieve"`
	Code 					string 					`json:"code"`
	Enable 					bool 					`json:"enable"`
	Blance_Bank_History 	[]Blance_Bank_History 	`gorm:"ForeignKey:Balance_Bank_ID" json:"blance_bank_history"`
}

type Blance_Bank_History struct {
	ID 					int 	`gorm:"primaryKey" json:"id"`
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

func main()  {
	db, err := gorm.Open("mysql", "root:@/e-wallet?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failde", err)
	}else{
		log.Println("Connection established")
	}
	//Migrate the Schema
	db.AutoMigrate(&User{}, &User_Balance{}, &User_Balance_History{}, &Blance_Bank{}, &Blance_Bank_History{})
	db.Model(&User_Balance{}).AddForeignKey("user_id","users(id)","cascade","cascade")
	db.Model(&User_Balance_History{}).AddForeignKey("id","user_balances(id)","cascade","cascade")
	db.Model(&Blance_Bank_History{}).AddForeignKey("balance_bank_id","blance_banks(id)","cascade","cascade")
}
