package main

import (
	lep_user "bupt/lep_user/kitex_gen/lep_user/lepuser"
	"log"
)

func main() {
	svr := lep_user.NewServer(new(LepUserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
