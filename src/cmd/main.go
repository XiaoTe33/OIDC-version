package main

import (
	"winter-examination/src/app"
	"winter-examination/src/dao"
)

func main() {
	dao.InitDb()
	app.InitRouters()
}
