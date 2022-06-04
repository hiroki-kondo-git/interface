package main

import "fmt"

type BookDb struct {
	id    string
	title string
}

func NewBookDb() *BookDb {
	return new(BookDb)
}

func (bd *BookDb) Remove() {
	fmt.Printf("aaa")
}

func (bd *BookDb) GetById() {
	fmt.Println("get from id")
}

func (bd *BookDb) GetByTitle() {
	fmt.Printf("get from title")
}
