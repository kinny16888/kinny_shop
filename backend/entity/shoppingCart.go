package entity

type ShoppingCart struct {
	Pkno     int    `json:"pkno"`
	Uid      string `json:"uid"`
	Pid      int    `json:"pid"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Pname    string `json:"pname"`
	Stock    int    `json:"stock"`
}
