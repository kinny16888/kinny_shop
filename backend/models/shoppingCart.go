package models

import (
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/tools"
	"log"
)

func CreateShoppingCart(s entity.ShoppingCart) bool {
	row := tools.Sqlconnect().QueryRow("SELECT 1 FROM shopping_cart WHERE pid = ? AND uid = ?", s.Pid, s.Uid)
	var isHaveInfo bool
	err := row.Scan(&isHaveInfo)
	var db = tools.Sqlconnect()
	if err != nil {
		fmt.Printf("讀取購物車失敗，原因為：%v\n", err)
		_, err := db.Exec(
			"INSERT INTO shopping_cart(uid, pid, quantity) VALUES(?,?,?)",
			s.Uid, s.Pid, s.Quantity)
		if err != nil {
			log.Println("讀取購物車失敗，原因為: " + err.Error())
			return false
		}
	} else {
		var tmp int
		row := tools.Sqlconnect().QueryRow("SELECT quantity FROM shopping_cart WHERE uid = ? AND pid = ?", s.Uid, s.Pid)
		err2 := row.Scan(&tmp)
		if err2 != nil {
			log.Println("讀取庫存失敗，原因為: " + err2.Error())
			return false
		}
		tmp = tmp + s.Quantity
		_, err := db.Exec("UPDATE shopping_cart SET quantity = ? WHERE uid = ? AND pid = ?",
			tmp, s.Uid, s.Pid)
		if err != nil {
			log.Println("讀取購物車失敗，原因為: " + err.Error())
			return false
		}
	}
	defer db.Close()
	return true
}
func ReadShoppingCart(a string) []entity.ShoppingCart {
	var ss []entity.ShoppingCart
	rows, _ := tools.Sqlconnect().Query("SELECT s.*, p.name, p.price, p.stock FROM shopping_cart s LEFT JOIN products p ON s.pid = p.pkno WHERE uid = ?", a)
	defer rows.Close()
	for rows.Next() {
		var s entity.ShoppingCart
		if err := rows.Scan(&s.Pkno, &s.Uid, &s.Pid, &s.Quantity, &s.Pname, &s.Price, &s.Stock); err != nil {
			log.Println("查詢購物車失敗，原因為: " + err.Error())
		}
		ss = append(ss, s)
	}
	return ss
}
func UpdateShoppingCart(t string, a string, s entity.ShoppingCart) bool {
	var err error
	var db = tools.Sqlconnect()
	var tmp int
	row := tools.Sqlconnect().QueryRow("SELECT quantity FROM shopping_cart WHERE uid = ? AND pid = ?", a, s.Pid)
	err2 := row.Scan(&tmp)
	if err2 != nil {
		log.Println("讀取庫存失敗，原因為: " + err2.Error())
		return false
	}
	if t == "+" {
		tmp = tmp + 1
		_, err = db.Exec("UPDATE shopping_cart SET quantity = ? WHERE uid = ? AND pid = ?", tmp, a, s.Pid)
	} else if t == "-" {
		if tmp >= 1 {
			tmp = tmp - 1
		} else {
			tmp = 0
		}
		_, err = db.Exec("UPDATE shopping_cart SET quantity = ? WHERE uid = ? AND pid = ?", tmp, a, s.Pid)
	}
	if err != nil {
		log.Println("更新購物車失敗，原因為: " + err.Error())
		return false
	}
	defer db.Close()
	return true
}
func DeleteShoppingCart(a string, s entity.ShoppingCart) bool {
	var db = tools.Sqlconnect()
	_, err := db.Exec("DELETE FROM shopping_cart WHERE uid = ? AND pid = ?", a, s.Pid)
	if err != nil {
		log.Println("刪除購物車失敗，原因為: " + err.Error())
		return false
	}
	defer db.Close()
	return true
}
