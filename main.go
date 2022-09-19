package main

import (
	"fmt"
	"time"
	"myproject/components"
)

type userLoginInfo struct {
	id string
	username string ""
	action string ""
	action_timestamp string ""
	token string
}

func main(){
	components.TestPrint()
    fmt.Printf("Test current time %v", time.Now().Format("2006-01-02 15:04:05"))
}