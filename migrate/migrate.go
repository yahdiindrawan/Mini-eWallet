package migrate

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
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

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func Migrate()  {
	db, err = gorm.Open("mysql","root:@/e-wallet?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed", err)
	} else {
		log.Println("Connection established ")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserBalance{})
	db.Model(&UserBalance{}).AddForeignKey("user_id","users(id)","cascade","cascade")
}
