package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST
func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	//decoder
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "INvalied request", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid Cradentials", http.StatusBadRequest)
		return
	}
	cnf := config.Getconfig()

	aceessToken,err:=util.CreateJwt(cnf.JwtSecreteKey,util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,

	})

	if err!=nil{
		http.Error(w,"Internal sever error",http.StatusInternalServerError)
	}
	util.SendData(w, aceessToken, http.StatusCreated)
}