package TestFile

import (
	"fmt"
	"testing"
)

func TestReadYaml(t *testing.T) {
	viper := YamlConfig("test", "/root/go-project/EasyCommon/YamlFile/")
	fmt.Println(viper.Get("es"))
	fmt.Println(viper.Get("es.Addresses"))
}
