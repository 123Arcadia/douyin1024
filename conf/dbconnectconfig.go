package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func DbInit() {
	v := viper.New()
	// 1.
	//v.SetConfigName("conf")
	//v.SetConfigType("yaml")
	// 2.
	v.SetConfigFile("conf/config.yaml")
	v.AddConfigPath("./conf")
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("打开配置文件\"conf.yaml\" 出错: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("未发现 \"conf.yaml\" 文件.")
			return
		}
		return
	}

	// 尝试将配置文件的内容解析到 conf.Config 变量中。
	//println(v.AllKeys())
	ip := v.GetString("mysql.ip")
	port := v.GetString("mysql.port")
	user := v.GetString("mysql.username")
	password := v.GetString("mysql.password")
	database := v.GetString("mysql.db")

	v.WatchConfig() // WriteConfig 将当前的 viper 配置写入到 (最后一次读取到的配置文件名 + ConfiogType 为后缀)，覆盖写入。
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变化:", e.Name)
		fmt.Println("操作类型:", e.Op.String())
	})

	// 注意不能有空格!!!
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	fmt.Println(dsn)
	//root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印 sql
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "dy_",
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		SkipDefaultTransaction: true,                                //禁用事务
		Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句

	})
	if err != nil {
		fmt.Println(err, "mysql 链接错误!", err)
		os.Exit(1)
	}
	// // 第二种打印sql：设置gorm的LogMode为true
	//DB = DB.Debug()
	fmt.Println("mysql 链接成功!")
}
