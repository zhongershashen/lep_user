package main

import (
	lep_user "bupt/lep_user/kitex_gen/lep_user/lepuser"
	"log"

	"github.com/zhongershashen/lep_lib/dal"
)

func main() {
	svr := lep_user.NewServer(new(LepUserImpl))
	dal.Init()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
