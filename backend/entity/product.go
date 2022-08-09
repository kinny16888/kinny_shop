package entity

type Product struct {
	Pkno      int      `json:"pkno" gorm:"primary_key"`
	Name      string   `json:"name"`
	Price     int      `json:"price"`
	Introduce string   `json:"introduce"`
	Stock     int      `json:"stock"`
	FilePath  []string `json:"filePath"`
}
type ReturnProducts struct {
	Name  string
	Stock int
}
