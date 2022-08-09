package entity

type User struct {
	Account    string `json:"account" gorm:"primary_key"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Permission int    `json:"permission"`
}
