package main

import (
	"fmt"
	"github.com/yayayapluto/api_lelang_online_golang_revisi_2/cmd/config"
	"github.com/yayayapluto/api_lelang_online_golang_revisi_2/internal/utils"
)

func main() {
	env, err := utils.LoadEnv()
	fmt.Println(env)

	if err != nil {
		panic(err)
	}

	db, err := config.ConnectDB(env.DBHOST, env.DBUSER, env.DBPASSWORD, env.DBNAME, env.DBPORT)
	if err != nil {
		panic(err)
	}

	_ = db
	fmt.Println("Success connected to database!")
}
