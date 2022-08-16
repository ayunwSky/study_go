/*
 * @ -*- Author: ayunwSky
 * @ -*- Date  : 2022/8/16 13:02
 * @ -*- Desc  : parse yaml config
 */

package viperYamlConf

import (
	"fmt"

	"github.com/spf13/viper"
	// "os"
)

var config *viper.Viper
var AllConfig TotalConfig

/*
SetConfigName 指定配置文件名，不需要加后缀
SetConfigType 指定配置文件类型，可以省略，viper会自动识别
AddConfigPath 添加配置文件所在目录，可以多个，告诉viper区哪里寻找配置文件
ReadInConfig  从前面告知的配置中加载配置文件
MergeInConfig 从前面告知的配置中加载配置文件并合并到之前
SetConfigFile 指定配置文件的全路径
*/

// 从IO流中读取配置文件 方法一 (不推荐)
// Init 初始化配置文件解析
// func Init() {
// 	config = viper.New()
// 	c, _ := os.Open("./configs/config/config.yaml")
// 	defer c.Close()
// 	config.SetConfigType("yaml")
// 	config.ReadConfig(c)

// 	s, _ := os.Open("./configs/secret/secret.yaml")
// 	defer s.Close()
// 	config.SetConfigType("yaml")
// 	config.MergeConfig(s)
// }

// 读取配置文件方法二 （不是很推荐）
// func Init() {
// 	config = viper.New()

// 	// 1. 解析 config 配置文件

// 	// 设置 public 的配置文件名字
// 	config.SetConfigName("config")
// 	// 设置 public 配置文件类型
// 	config.SetConfigType("yaml")
// 	// 设置配置文件的所在目录
// 	config.AddConfigPath("./configs/config")
// 	// 读取该配置文件
// 	config.ReadInConfig()

// 	// 2. 解析 secret 配置文件
// 	config.SetConfigName("secret")
// 	config.SetConfigType("yaml")
// 	config.AddConfigPath("./configs/secret")
// 	config.MergeInConfig()
// }

// 读取配置文件方法三 (推荐)
func Init() {
	config = viper.New()
	// 设置 public 配置文件名
	config.SetConfigFile("./configs/config/config.yaml")
	// 读取该配置文件
	config.ReadInConfig()
	// 解析 secret 配置文件
	config.SetConfigFile("./configs/secret/secret.yaml")
	config.MergeInConfig()

	config.Unmarshal(&AllConfig)
}

type basicinfo struct {
	RealName string   `yaml:"realname"`
	Age      int      `yaml:"age"`
	Language []string `yaml:"language"`
	Married  bool     `yaml:"married"`
}

type hobby struct {
	Sport       []string `yaml:"sport:`
	Music       []string `yaml:"music"`
	LuckyNumber int      `yaml:"luckNumber"`
}

type admin struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type redis struct {
	Admin admin `yaml:"admin"`
}

type mysql struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}

type TotalConfig struct {
	TimeStamp string    `yaml:"timestamp"`
	Username  string    `yaml:"username"`
	BasicInfo basicinfo `yaml:"basicinfo"`
	Hobby     hobby     `yaml:"hobby"`
	Redis     redis     `yaml:"redis"`
	MySQL     mysql     `yaml:"mysql"`
	Jwt       jwt       `yaml:"jwt"`
}

func ParseYamlConf() {
	fmt.Println("Go parse yaml config.")
	fmt.Println()

	Init()

	fmt.Println(AllConfig)
	fmt.Println()
	fmt.Println(AllConfig.TimeStamp)
	fmt.Println(AllConfig.BasicInfo.Language)
	fmt.Println(AllConfig.BasicInfo.Language[1])
	fmt.Println(AllConfig.Hobby.Sport)
	fmt.Println(AllConfig.Hobby.Sport[0])
	fmt.Println(AllConfig.Redis.Admin.Password)
	fmt.Println(AllConfig.Jwt.Secret)
}
