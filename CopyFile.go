package GoCommon_File

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyDir 函数用于复制一个目录及其子目录和文件到另一个目录
// 参数srcPath为源目录路径，desPath为目标目录路径
// 返回值为error类型，如果复制过程中发生错误则返回错误信息，否则返回nil
func CopyDir(srcPath, desPath string) error {
	// 检查源路径是否存在，并且是一个目录
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	// 检查目标路径是否存在，并且是一个目录
	if desInfo, err := os.Stat(desPath); err != nil {
		MakeDir(desPath)
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	// 源路径与目标路径不能相同
	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	// 遍历源目录，并复制文件或创建目录到目标目录
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		// 跳过源目录本身
		if path == srcPath {
			return nil
		}

		// 替换路径中的源目录为目标目录
		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		// 如果是文件，则复制文件
		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			// 如果是目录，则检查目标路径是否存在，不存在则创建目录
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

// CopyFile 用于复制文件内容
//
// src: 源文件路径
// des: 目标文件路径
//
// written: 实际写入的字节数
// err: 错误信息，如果复制成功则为nil
func CopyFile(src, des string) (written int64, err error) {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// 获取源文件信息
	fi, _ := srcFile.Stat()
	// 获取源文件权限
	perm := fi.Mode()

	des_path_list := strings.Split(des, "/")
	des_path := strings.Join(des_path_list[:len(des_path_list)-1], "/")
	if !FileIsExisted(des_path) {
		MakeDir(des_path)
	}
	// 创建目标文件
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	// 复制文件内容
	return io.Copy(desFile, srcFile)
}

// FileIsExisted 函数检查指定文件是否存在
// 参数：
// filename string - 文件名
// 返回值：
// bool - 文件是否存在，存在返回true，不存在返回false
func FileIsExisted(filename string) bool {
	// 初始化变量existed为true，表示文件存在
	existed := true

	// 使用os.Stat函数检查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 如果文件不存在，将existed设置为false
		existed = false
	}

	// 返回文件是否存在的结果
	return existed
}

// MakeDir 函数用于创建指定路径的目录，若目录已存在则不进行任何操作。
//
// 参数：
//
//	dir: 待创建的目录路径
//
// 返回值：
//
//	error: 返回错误信息，如果目录创建成功则返回nil
func MakeDir(dir string) error {
	// 判断目录是否存在
	if !FileIsExisted(dir) {
		// 如果目录不存在，则创建目录
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}
