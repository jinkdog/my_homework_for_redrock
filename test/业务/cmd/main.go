package main

import (
	"modulename/basic_homework/homework_for_redrock/test/业务/api"
	"modulename/basic_homework/homework_for_redrock/test/业务/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
