package main

// accessできる情報を絞る

import (
	gohttp "accessLevel/goHttp"
	"fmt"
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
