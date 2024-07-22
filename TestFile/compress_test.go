package GoCommon_File

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_File"
)

func TestCompressFile(t *testing.T) {
	// 定义本地文件夹路径
	localPath := "/root/gitlab"
	// 定义压缩文件输出路径
	outputFile := localPath + "_zip/gitlab.zip"

	// 调用 ControlFile 对象的 ZipFolder 方法进行文件夹压缩
	err := GoCommon_File.ZipFolder(localPath, outputFile)
	if err != nil {
		// 如果压缩过程中出现错误，打印错误信息并返回
		fmt.Println("Error compressing folder:", err)
		return
	}

}
func TestFindDir(t *testing.T) {
	// 调用 ControlFile 结构体的 FindDir 方法，传入根目录 "/root/" 和递归深度 0
	GoCommon_File.FindDir("/root/", 0)
}

func TestFindCurrentDir(t *testing.T) {
	// 调用 ControlFile 包的 FindCurrentDir 函数，传入参数 "/root/"，返回两个结果：dir_list 和 document_list
	dir_list, document_list := GoCommon_File.FindCurrentDir("/root/")
	fmt.Println(dir_list)
	fmt.Println(document_list)
}
