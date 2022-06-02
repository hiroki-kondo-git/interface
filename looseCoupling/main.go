package main

import (
	"log"
)

type DBInterface interface {
	Get()
	Delete()
}

type DB3rdParty struct {
	user     string
	password int
}

type DBHandler struct {
	DB3rdParty
}

// interfaceを引数にした関数で処理
func DBUseInterface(d DBInterface) {
	d.Get()
	d.Delete()
}

func main() {
	db := &DBHandler{
		DB3rdParty{
			user:     "あああ",
			password: 22,
		},
	}
	DBUseInterface(db)
}

func (d *DBHandler) Get() {
	log.Println(d.user, d.password)
}

func (d *DBHandler) Delete() {
	log.Println("delete", d.user)
}
