package initModelsExamples

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"gorm.io/gorm"
	"log"
)

var err error

func CreateTable(db *gorm.DB) error {
	err = db.AutoMigrate(models.User{}, models.Comment{}, models.Favorite{}, models.Relation{}, models.Message{}, models.Video{}, models.Relation{})
	if err != nil {
		log.Fatal("建表异常:", err)
		return err
	}
	// 输入初始样例
	tableInit(db)
	return nil
}

// 表格-样例数据初始化
func tableInit(db *gorm.DB) {
	InitUsers(db)
	InitVideos(db)
	InitFavorites(db)
	InitComments(db)
	InitRelations(db)
	InitMessages(db)
}
