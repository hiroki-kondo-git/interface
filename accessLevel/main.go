package main

// accessできる情報を絞る

import (
	"fmt"

	gohttp "github.com/hiroki-kondo-git/interface/accessLevel/goHttp"
)

func main() {
	client := gohttp.HttpClient{}
	client.Get()
	client.Post()
	//関係ないところaccessできてしまう
	fmt.Println(client.Password)

	client2 := gohttp.New()
	client2.Get()
	client2.Post()
}
