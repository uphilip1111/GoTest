package dborm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	AccountID int    `gorm:"primary_key"`
	UserAcc   string `gorm:"column:user_acc"`
	UserPwd   string `gorm:"column:user_pwd"`
	UserName  string `gorm:"column:user_name"`
	UserPhone string `gorm:"column:phone"`
	UserAddr  string `gorm:"column:address"`
}

func (User) TableName() string {
	return "user_account"
}

func QueryFunc() {
	db, err := gorm.Open("mysql", "root:admin@/project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	var user User
	//db.First(&user, 1)
	db.First(&user)
	//fmt.Println("accid = ", user.UserAddr)
}
