package GoCommon_File

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_File"
)

func TestCopyFile(t *testing.T) {
	// 调用 ControlFile 的 CopyFile 方法，复制文件，将文件 "/root/cn_swarm_restart.sh" 复制到 "/root/gitlab/cn_swarm_restart.sh"
	_, error := GoCommon_File.CopyFile("/root/cn_swarm_restart.sh", "/root/gitlab/cn_swarm_restart.sh")
	// 调用 ControlFile 的 CopyDir 方法,复制文件夹
	// error := ControlFile.CopyDir("/root/shell_test", "/root/gitlab/shell_test")
	fmt.Println(error)
}
