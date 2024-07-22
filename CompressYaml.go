package GoCommon_File

import (
	"fmt"

	"github.com/spf13/viper"
)

func YamlConfig(ConfigPath, ConfigName string) *viper.Viper {
	// 设置配置文件的名字
	viper.SetConfigName(ConfigName)
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(ConfigPath)
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.GetViper()
}
