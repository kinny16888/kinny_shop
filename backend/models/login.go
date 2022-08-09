package models

import (
	"kinnyShop_api/entity"
	"kinnyShop_api/tools"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func FindAccount(u entity.User) bool {
	//fmt.Println(u)
	err := tools.Ormconnect().First(&u).Error
	if err == nil {
		log.Println("查詢到 user 為 ", u)
		return true
	} else {
		log.Println("查詢 user 失敗，原因為 " + err.Error())
		return false
	}
}
func FindPhone(u entity.User) bool {
	row := tools.Sqlconnect().QueryRow("SELECT phone FROM users WHERE phone = ?", u.Phone)
	err := row.Scan(&u.Phone)
	if err == nil {
		log.Println("查詢到手機號碼為 ", u)
		return true
	} else {
		log.Println("查詢手機號碼失敗，原因為 " + err.Error())
		return false
	}
}
func AddAccount(u entity.User) (bool, error) {
	u.Permission = 0
	err := tools.Ormconnect().Create(u).Error
	if err != nil {
		log.Println("資料庫 Migrate 失敗，原因為 " + err.Error())
		return false, err
	}
	return true, err
}
func CheckAccount(u entity.User) (bool, int) {
	row := tools.Sqlconnect().QueryRow("select * from users where account=? and password = ?", u.Account, u.Password)
	err := row.Scan(&u.Account, &u.Password, &u.Name, &u.Phone, &u.Permission)
	if err != nil {
		log.Println("查詢該帳戶失敗，原因為：" + err.Error())
		return false, -1
	}
	return true, u.Permission
}
