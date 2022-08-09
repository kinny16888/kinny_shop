package entity

type Order struct {
	Pkno      string          `json:"pkno" gorm:"primary_key"`
	Uid       string          `json:"uid"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	Date      string          `json:"date"`
	Time      string          `json:"time"`
	Amount    int             `json:"amount"`
	SendMark  bool            `json:"sendMark"`
	Commodity []DetailedOrder `json:"commodity"`
}

type DetailedOrder struct {
	Opkno    string `json:"opkno"`
	Pid      int    `json:"pid"`
	Pname    string `json:"pname"`
	Quantity int    `json:"quantity"`
}
