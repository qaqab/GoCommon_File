package GoCommon_File

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_File"
)

func TestReadYaml(t *testing.T) {
	viper := GoCommon_File.YamlConfig("/root/go-project/EasyCommon/YamlFile/", "test")
	fmt.Println(viper.Get("es"))
	fmt.Println(viper.Get("es.Addresses"))
}
