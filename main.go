package main

import (
	"fmt"
	"time"
	"os"
	"myproject/db"
)

type userLoginInfo struct {
	id string
	username string ""
	action string ""
	action_timestamp string ""
	token string
}

func main(){
	os.Setenv("RUNTIME_ENV", "sit")
	db.InitConnection()
	fmt.Printf("execute complete, current time %v", time.Now().Format("2006-01-02 15:04:05"))
}
