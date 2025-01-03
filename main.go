package main

import (
	"github.com/realtemirov/go-sqlc-grpc-http/apps"
)

func main() {
	// nums := []int{1, 2, 3, 4}
	// for _, v := range nums[1:1] {
	// 	fmt.Println(v)
	// }
	app := apps.App{}
	app.Run()
}
