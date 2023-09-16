package main

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initModelsExamples"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/messgage"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	go messgage.RunMessageServer()

	conf.DbInit()
	//建表
	if conf.DB != nil {
		if err := initModelsExamples.CreateTable(conf.DB); err != nil {
			log.Fatalf("无法创建或更新表: %v", err)
		}
		log.Println("数据库表初始化成功!")
	} else {
		log.Fatalf("数据库连接为空，无法创建或更新表，请检查...")
	}
	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
