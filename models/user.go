package models

type Address struct {
	Province string `json:"province" bson:"province"`
	District string `json:"district" bson:"district"`
	Ward     string `json:"ward" bson:"ward"`
}

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     int     `json:"age" bson:"user-age"`
	Address Address `json:"address" bson:"user_address"`
}
