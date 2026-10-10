package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	cnf := config.Getconfig()
	rest.Start(cnf)

}
