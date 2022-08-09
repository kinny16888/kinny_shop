package models

import (
	"database/sql"
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/tools"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ReadProduct(name string) []entity.Product {

	var pp []entity.Product
	var rows *sql.Rows
	if name == "recommend" {
		rows, _ = tools.Sqlconnect().Query("SELECT p.pkno, p.name, p.price, p.introduce, p.stock, f.path FROM products p LEFT JOIN( SELECT pid, path FROM file_path WHERE SUBSTRING(path, 1, 4) = 'home' ) f ON p.pkno = f.pid LEFT JOIN detailed_order d ON p.pkno = d.pid WHERE 1 GROUP BY p.pkno ORDER BY SUM(d.quantity) DESC LIMIT 10")
	} else {
		sql := "SELECT products.*,f.path FROM products LEFT JOIN  (SELECT pid,path FROM file_path WHERE SUBSTRING(path,1,4) = 'home') AS f ON products.pkno = f.pid WHERE products.name LIKE '%" + name + "%' "
		rows, _ = tools.Sqlconnect().Query(sql)
	}

	for rows.Next() {
		var p entity.Product
		var tmp string
		if err := rows.Scan(&p.Pkno, &p.Name, &p.Price, &p.Introduce, &p.Stock, &tmp); err != nil {
			fmt.Printf("查詢產品失敗，原因為：%v\n", err)
		}
		p.FilePath = append(p.FilePath, tmp)
		pp = append(pp, p)
	}
	defer rows.Close()
	return pp
}

func ReadProductDetailed(pkno int) entity.Product {
	var p entity.Product
	row := tools.Sqlconnect().QueryRow("SELECT * FROM products WHERE pkno = ?", pkno)
	if err := row.Scan(&p.Pkno, &p.Name, &p.Price, &p.Introduce, &p.Stock); err != nil {
		fmt.Printf("查詢產品失敗，原因為：%v\n", err)
	}

	rows, _ := tools.Sqlconnect().Query("SELECT path FROM file_path WHERE pid = ?", p.Pkno)
	for rows.Next() {
		var tmp string
		if err2 := rows.Scan(&tmp); err2 != nil {
			fmt.Printf("查詢產品細項失敗，原因為：%v\n", err2)
		}
		p.FilePath = append(p.FilePath, tmp)
	}
	defer rows.Close()
	return p
}
func ReadAllProduct() []entity.Product {
	var pp []entity.Product
	rows, err := tools.Sqlconnect().Query("SELECT products.*, f.path FROM products LEFT JOIN  (SELECT pid,path FROM file_path WHERE SUBSTRING(path,1,4) = 'home') AS f ON products.pkno = f.pid")
	if err != nil {
		fmt.Printf("查詢所有產品失敗，原因為：%v\n", err)
	}
	for rows.Next() {
		var p entity.Product
		var tmp string
		if err := rows.Scan(&p.Pkno, &p.Name, &p.Price, &p.Introduce, &p.Stock, &tmp); err != nil {
			log.Fatal(err)
		}
		p.FilePath = append(p.FilePath, tmp)
		pp = append(pp, p)
	}
	defer rows.Close()
	return pp
}
func CreateProduct(p entity.Product) (bool, entity.Product) {
	fmt.Println(p)
	var db = tools.Sqlconnect()
	_, err := db.Exec(
		"INSERT INTO products(name, price, introduce, stock) VALUES (?,?,?,?)",
		p.Name, p.Price, p.Introduce, p.Stock)
	row := tools.Sqlconnect().QueryRow("SELECT MAX(pkno) FROM products")
	if err2 := row.Scan(&p.Pkno); err2 != nil {
		log.Println("主鍵讀取失敗 原因為: ", err2.Error())
	}
	defer db.Close()
	if err == nil {
		log.Println("新增 product成功", p)
		return true, p
	} else {
		log.Println("新增 product失敗，原因為: " + err.Error())
		return false, p
	}
}
func DeleteProduct(pkno int) bool {
	var p entity.Product
	p.Pkno = pkno
	var db = tools.Sqlconnect()
	_, err := db.Exec("DELETE FROM products WHERE pkno = ?", pkno)
	if err != nil {
		log.Println("刪除購物車失敗，原因為: " + err.Error())
		return false
	}
	defer db.Close()
	return true
}
func UpdateProduct(pkno int, p entity.Product) bool {
	p.Pkno = pkno
	var db = tools.Sqlconnect()
	_, err := db.Exec("UPDATE products SET name=?,price=?,introduce=?,stock=? WHERE pkno = ?",
		p.Name, p.Price, p.Introduce, p.Stock, p.Pkno)
	if err != nil {
		log.Println("更新庫存失敗，原因為: " + err.Error())
		return false
	}
	defer db.Close()
	return true
}
func InsertFile(p entity.Product) bool {
	var db = tools.Sqlconnect()
	for i := 0; i < len(p.FilePath); i++ {
		_, err := db.Exec(
			"INSERT INTO `file_path`(pid, path) VALUES (?,?)",
			p.Pkno, p.FilePath[i])
		if err != nil {
			log.Println("新增檔案路徑至資料庫時失敗, 原因為: ", err.Error())
			return false
		}
	}
	defer db.Close()
	return true
}
