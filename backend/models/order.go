package models

import (
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/tools"
	"log"
	"strconv"
	"time"
)

func CreateOrder(o entity.Order) (bool, []entity.ReturnProducts) {
	falseSign := true
	var returnItem []entity.ReturnProducts
	for i := 0; i < len(o.Commodity); i++ {
		var stock int
		var tmpReturnProducts entity.ReturnProducts
		row := tools.Sqlconnect().QueryRow("SELECT stock FROM products WHERE pkno = ?", o.Commodity[i].Pid)
		if err := row.Scan(&stock); err != nil {
			log.Println("獲取商品庫存失敗 原因為: ", err.Error())
		}
		if o.Commodity[i].Quantity > stock {
			falseSign = false
			row2 := tools.Sqlconnect().QueryRow("SELECT name, stock FROM products WHERE pkno = ?", o.Commodity[i].Pid)
			if err := row2.Scan(&tmpReturnProducts.Name, &tmpReturnProducts.Stock); err != nil {
				log.Println("獲取商品名稱失敗 原因為: ", err.Error())
			}
			returnItem = append(returnItem, tmpReturnProducts)
		}
	}
	if !falseSign {
		return false, returnItem
	}
	var date string = strconv.Itoa(time.Now().Year()) + "-" +
		strconv.Itoa(int(time.Now().Month())) + "-" +
		strconv.Itoa(time.Now().Day())
	var time string = strconv.Itoa(time.Now().Hour()) +
		strconv.Itoa(time.Now().Minute()) +
		strconv.Itoa(int(time.Now().Second()))
	log.Println("日期" + date)
	log.Println("時間" + time)
	_, err := tools.Sqlconnect().Exec(
		"insert INTO orders(pkno,uid,address,date,time,amount,sendMark) values(?,?,?,?,?,?,?)",
		o.Pkno, o.Uid, o.Address, date, time, o.Amount, o.SendMark)
	if err != nil {
		log.Println("建立訂單失敗，原因為: " + err.Error())
		return false, returnItem
	}
	fmt.Println("建立訂單成功！")
	for i := 0; i < len(o.Commodity); i++ {
		_, err = tools.Sqlconnect().Exec(
			"insert INTO detailed_order(opkno,pid,quantity) values(?,?,?)",
			o.Pkno, o.Commodity[i].Pid, o.Commodity[i].Quantity)
		if err != nil {
			log.Println("第" + (strconv.Itoa(i)) + "筆訂單細項新增失敗 原因為: " + err.Error())
			return false, returnItem
		}
	}
	fmt.Println("建立訂單細項成功！")
	for i := 0; i < len(o.Commodity); i++ {
		var tmp int
		row := tools.Sqlconnect().QueryRow("SELECT stock FROM products WHERE pkno = ?", o.Commodity[i].Pid)
		err2 := row.Scan(&tmp)
		if err2 != nil {
			log.Println("讀取庫存失敗，原因為: " + err2.Error())
			return false, returnItem
		}
		var db = tools.Sqlconnect()
		_, err := db.Exec("UPDATE products SET stock=? WHERE pkno = ?", tmp-o.Commodity[i].Quantity, o.Commodity[i].Pid)
		//UPDATE products SET stock=(SELECT stock FROM products WHERE pkno = ?)-? WHERE pkno = ?
		if err != nil {
			log.Println("更新庫存失敗，原因為: " + err.Error())
			return false, returnItem
		}
		db.Close()
	}
	_, err2 := tools.Sqlconnect().Exec("DELETE FROM shopping_cart WHERE uid = ?", o.Uid)
	if err2 != nil {
		log.Println("刪除用戶購物車失敗，原因為: " + err.Error())
		return false, returnItem
	}
	return true, returnItem
}
func ReadOrder(pkno string) []entity.Order {
	var oo []entity.Order
	rows, _ := tools.Sqlconnect().Query("SELECT orders.*,users.name FROM `orders` LEFT JOIN users ON orders.uid = users.account WHERE uid = ?", pkno)
	defer rows.Close()
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(&o.Pkno, &o.Uid, &o.Address, &o.Date, &o.Time, &o.Amount, &o.SendMark, &o.Name); err != nil {
			log.Println("讀取訂單失敗，原因為: " + err.Error())
		}
		oo = append(oo, o)
	}
	return oo
}
func ReadAllOrder(a string, p string) []entity.Order {

	var oo []entity.Order
	if p == "1" {
		rows, _ := tools.Sqlconnect().Query("SELECT orders.*,users.name FROM `orders` LEFT JOIN users ON orders.uid = users.account WHERE 1=1")
		defer rows.Close()
		for rows.Next() {
			var o entity.Order
			if err := rows.Scan(&o.Pkno, &o.Uid, &o.Address, &o.Date, &o.Time, &o.Amount, &o.SendMark, &o.Name); err != nil {
				log.Println("讀取訂單失敗，原因為: " + err.Error())
			}
			oo = append(oo, o)
		}
	} else {
		rows, _ := tools.Sqlconnect().Query("SELECT orders.*,users.name FROM `orders` LEFT JOIN users ON orders.uid = users.account WHERE uid = ?", a)
		defer rows.Close()
		for rows.Next() {
			var o entity.Order
			if err := rows.Scan(&o.Pkno, &o.Uid, &o.Address, &o.Date, &o.Time, &o.Amount, &o.SendMark, &o.Name); err != nil {
				log.Println("讀取訂單失敗，原因為: " + err.Error())
			}
			oo = append(oo, o)
		}
	}
	return oo
}
func ReadDetailedOrder(pkno string) entity.Order {
	var o entity.Order
	rows, _ := tools.Sqlconnect().Query("SELECT o.pkno, p.name,d.quantity FROM detailed_order d LEFT JOIN orders o ON d.opkno = o.pkno LEFT JOIN products p ON d.pid = p.pkno WHERE d.opkno = ?", pkno)
	defer rows.Close()
	for rows.Next() {

		var tmp entity.DetailedOrder
		if err := rows.Scan(&o.Pkno, &tmp.Pname, &tmp.Quantity); err != nil {
			log.Println("讀取訂單細項失敗，原因為: " + err.Error())
		}
		o.Commodity = append(o.Commodity, tmp)
	}
	return o
}
func UpdateSendMrk(pkno string) bool {
	var tmp bool
	row := tools.Sqlconnect().QueryRow("SELECT sendMark FROM orders WHERE pkno = ?", pkno)
	err2 := row.Scan(&tmp)
	if err2 != nil {
		log.Println("讀取出貨註記失敗，原因為: " + err2.Error())
		return false
	}
	tmp = !tmp
	_, err := tools.Sqlconnect().Exec("UPDATE orders SET sendMark=? WHERE pkno = ?", tmp, pkno)
	//UPDATE orders SET sendMark=!(SELECT sendMark FROM orders WHERE pkno = ?) WHERE pkno = ?
	if err != nil {
		log.Println("更新出貨註記失敗，原因為: " + err.Error())
		return false
	}
	return true
}
