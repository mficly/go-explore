package main

import (
	"fmt"
	"time"
	"main01/mytest"
)

type userLoginInfo struct {
	id string
	username string ""
	action string ""
	action_timestamp string ""
	token string
}

func main(){
	test.TestPrint()
    fmt.Printf("Test current time %v", time.Now().Format("2006-01-02 15:04:05"))
}