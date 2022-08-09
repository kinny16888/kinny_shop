package models

import (
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/tools"

	_ "github.com/go-sql-driver/mysql"
)

func Ranking(month string) []entity.Statistics {
	var ss []entity.Statistics
	rows, _ := tools.Sqlconnect().Query("SELECT p.name, SUM(d.quantity*p.price) AS num FROM products p LEFT JOIN detailed_order d ON p.pkno = d.pid LEFT JOIN orders o ON d.opkno = o.pkno WHERE SUBSTRING(o.date,6,2) = ? GROUP BY p.pkno ORDER BY SUM(d.quantity) DESC LIMIT 5", month)
	for rows.Next() {
		var s entity.Statistics
		if err := rows.Scan(&s.Key, &s.Vlaue); err != nil {
			fmt.Printf("查詢統計排名失敗，原因為：%v\n", err)
		}
		ss = append(ss, s)
	}
	defer rows.Close()
	return ss
}
func BarChart() []entity.Statistics {
	var ss []entity.Statistics
	var month = [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	var db = tools.Sqlconnect()
	for i := 0; i < len(month); i++ {
		var s entity.Statistics
		s.Key = month[i]
		row := db.QueryRow("SELECT IFNULL(SUM(d.quantity * p.price),0) AS VALUE FROM orders o LEFT JOIN detailed_order d ON o.pkno = d.opkno LEFT JOIN products p ON d.pid = p.pkno WHERE SUBSTRING(o.date, 6, 2) = ?", month[i])
		if err := row.Scan(&s.Vlaue); err != nil {
			fmt.Printf("查詢長條圖失敗，原因為：%v\n", err)
		}
		ss = append(ss, s)
	}
	defer db.Close()
	return ss
}
func PieChart(month string) []entity.Statistics {
	var ss []entity.Statistics
	rows, _ := tools.Sqlconnect().Query("SELECT p.name, SUM(d.quantity) AS num FROM products p LEFT JOIN detailed_order d ON p.pkno = d.pid LEFT JOIN orders o ON d.opkno = o.pkno WHERE SUBSTRING(o.date,6,2) = ? GROUP BY p.pkno ORDER BY SUM(d.quantity) DESC LIMIT 5", month)
	for rows.Next() {
		var s entity.Statistics
		if err := rows.Scan(&s.Key, &s.Vlaue); err != nil {
			fmt.Printf("查詢圓餅圖失敗，原因為：%v\n", err)
		}
		ss = append(ss, s)
	}
	defer rows.Close()
	return ss
}
