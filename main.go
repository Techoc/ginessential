package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/techoc/ginessential/common"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
}

func InitConfig() {
	//获取项目工作路径
	workDir, err := os.Getwd()
	if err != nil {
		return
	}
	//设置配置文件名称
	viper.SetConfigName("application")
	//设置配置文件类型
	viper.SetConfigType("yaml")
	//添加设置文件
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
}
