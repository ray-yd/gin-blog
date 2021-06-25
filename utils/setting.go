package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket    string
	VultrUrl  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("ini 讀取錯誤")
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadVultr(file)
}

// LoadServer 讀取 ini 中的 Server 配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("hxyMtz6PkhqNuzwvaDF5b8vF46Rtv6zK6PcdZW9CZRZu4vzepVdKYhHbUpCW5bfKCBp7A3RrbKwmFEYXnu4FZCPR8ZqdaGbcApE5eesGea8fdAysV3NAhuPsxqc5fAmnMuHhEPTmeWWSQRsSqFX5CS73fuhFPnEH5GyPuSGKcFfU5KBRgwbkdBgZZh2BGuDUQ6X9yTeWKf6W2nXYFS9dugFtSAnmdKgCYCkE63423qv7eGWDmYC2RMhrggp49VVH")

}

// LoadDatabase 讀取 ini 中的 database 配置
func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("198.13.53.73")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("gin-blog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("tNBSND4C3JxcPkxd")
	DbName = file.Section("database").Key("DbName").MustString("gin-blog")
}

func LoadVultr(file *ini.File) {
	AccessKey = file.Section("Vultr").Key("AccessKey").String()
	SecretKey = file.Section("Vultr").Key("SecretKey").String()
	Bucket = file.Section("Vultr").Key("Bucket").String()
	VultrUrl = file.Section("Vultr").Key("VultrUrl").String()
}
