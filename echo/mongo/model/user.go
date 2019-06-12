package model

type User struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}
