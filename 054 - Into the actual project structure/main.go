package main

// import (
// 	"ecommerce/util"
// 	"fmt"
// )

import "ecommerce/cmd"

func main() {
	cmd.Serve()



	//JWT


	// // var s string

	// // s = "a,b,c"

	// // byteArr := []byte(s)
	// // fmt.Println(s)
	// // fmt.Println(byteArr)

	// // enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	// // b64Str:= enc.EncodeToString(byteArr)
	// // fmt.Println(b64Str)

	// secret := []byte("my-secret")
	// message := []byte("Hello world")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)

	// text := h.Sum(nil)
	// fmt.Println(text)
	// jwt, err := util.CreateJwt("my-secret", util.Payload{
	// 	Sub:         "45",
	// 	FirstName:   "Mohaiminul",
	// 	LastName:    "Ayon",
	// 	Email:       "mohaiminul5@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)

	
}
