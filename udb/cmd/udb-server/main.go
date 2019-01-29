package main

import (
	"fmt"
	"github.com/hoenn/mcrosvc/proto"
)

func main() {
	u := &proto.User{
		Name: "n",
	}
	fmt.Println(u)
}
