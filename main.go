package main

import (
	"github.com/ray-yd/gin-blog/model"
	"github.com/ray-yd/gin-blog/routes"
)

func main() {
	// 調用資料庫
	model.InitDb()

	routes.InitRouter()
}
