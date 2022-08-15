/*
 * @ -*- Author: ayunwSky
 * @ -*- Date  : 2022/8/14 21:01
 * @ -*- Desc  :
 */

package viperTomlAndJson

import (
	"fmt"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
}

type RedisConfig struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Password string `toml:"password" json:"password"`
	Database int    `toml:"database" json:"database"`
}

type Config struct {
	MySQL MySQLConfig `toml:"mysql" json:"mysql"`
	Redis RedisConfig `toml:"redis" json:"redis"`
	Title string       `toml:"title" json:"title"`
}

func ParseTomlAndJson() {
	//fmt.Println("# ---------- viperTomlAndJson start ----------")
	// 定义一个 Config 结构体类型的变量 configToml

	// 读取 toml 格式配置文件使用
	var configToml Config

	vToml := viper.New()
	vToml.SetConfigName("config")
	vToml.SetConfigType("toml")
	vToml.AddConfigPath("./configs")

	if err := vToml.ReadInConfig(); err != nil {
		fmt.Printf("read toml config failed, err: %v\n", err)
		return
	}

	vToml.Unmarshal(&configToml)

	fmt.Println("---------- viperTomlJson.go ----------")
	fmt.Println()
	fmt.Println("Config All:", configToml)
	fmt.Println("Redis:", configToml.Redis)
	fmt.Println("MySQL:", configToml.MySQL)
	fmt.Printf("Redis Host: %v\n", configToml.Redis.Host)
	fmt.Printf("MySQL Host: %v\n", configToml.MySQL.Host)
	fmt.Printf("title: %v\n", configToml.Title)

	fmt.Println("# ----------------------------------------")

	// 定义一个 Config 结构体类型的变量 configJson

	// 读取 json 格式配置文件
	var configJson Config

	vJson := viper.New()
	vJson.SetConfigName("jsconfig")
	vJson.SetConfigType("json")
	vJson.AddConfigPath("./configs")

	if err := vJson.ReadInConfig(); err != nil {
		fmt.Printf("read json config failed, err: %v\n", err)
		return
	}

	vJson.Unmarshal(&configJson)
	fmt.Println("Config All:", configJson)
	fmt.Println("Redis:", configJson.Redis)
	fmt.Println("MySQL:", configJson.MySQL)
	fmt.Printf("Redis Host: %v\n", configJson.Redis.Host)
	fmt.Printf("MySQL Host: %v\n", configJson.MySQL.Host)

	fmt.Println()
}
