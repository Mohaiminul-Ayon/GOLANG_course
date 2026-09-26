package cmd

import (
	"ecommerce/rest"
	"ecommerce/config"
)

func Serve() {
	cnf := config.Getconfig()
	rest.Start(cnf)

}