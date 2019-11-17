package main

import (
	"github.com/ysbayram/TestRepo/server"
	"github.com/ysbayram/TestRepo/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	db := server.CreateDBCon()
	r := server.SetupRouter(db)

	defer func() {
		if r := recover(); r != nil {
			logrus.Panicf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(6, 6))
		}
		db.Close()
	}()

	r.Run(":4141")
}
