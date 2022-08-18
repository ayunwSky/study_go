/*
 * @ -*- Author: ayunwSky
 * @ -*- Date: 2022/8/14 21:32
 * @ -*- Desc:
 */

package viperGetItem

import (
	"fmt"

	"github.com/spf13/viper"
)

// ParseItems 读取配置项的值
func ParseItems() {
	// 设置要解析的配置文件名
	viper.SetConfigName("item")
	// 设置要解析的配置文件格式
	viper.SetConfigType("json")
	// 设置配置文件所在路径，如果是当前路径，就写成 "."
	viper.AddConfigPath("./configs")

	// 根据上面的配置加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("load item config file failed, err:%v\n", err)
		return
	}

	mysqlHost := viper.Get("mysql.host")
	mysqlUsername := viper.Get("mysql.username")
	mysqlPort := viper.Get("mysql.port")
	mysqlPortsSlice := viper.GetIntSlice("mysql.ports")

	mysqlMetricsPort := viper.GetInt("mysql.metrics.port")
	redisConf := viper.Get("redis")

	mysqlStringMap := viper.GetStringMapString("mysql")

	fmt.Println("---------- viperGetItem.go ----------")

	if viper.IsSet("mysql.host") {
		fmt.Println("[IsSet()] mysql.host is set")
	} else {
		fmt.Println("[IsSet()] mysql.host is not set")
	}

	fmt.Println()
	fmt.Println("mysqlHost:", mysqlHost)
	fmt.Println("mysqlUsername:", mysqlUsername)
	fmt.Println("mysqlPort:", mysqlPort)
	fmt.Println("mysqlPortsSlice:", mysqlPortsSlice)
	fmt.Println("mysqlPortsSlice[0]:", mysqlPortsSlice[0])
	fmt.Println("mysqlPortsSlice[1]:", mysqlPortsSlice[1])
	fmt.Println("mysqlMetricsPort:", mysqlMetricsPort)
	fmt.Println("redisConf:", redisConf)

	fmt.Println()

	fmt.Println("mysqlStringMap:", mysqlStringMap)
	fmt.Println("mysqlStringMapUsername:", mysqlStringMap["username"])

	fmt.Println("mysqlStringMapHost:", mysqlStringMap["mysql.metrics.host"])
	fmt.Println("mysqlStringMapMetricsPort:", mysqlStringMap["mysql.metrics.port"])

	fmt.Println()
}
