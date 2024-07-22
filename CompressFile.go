package GoCommon_File

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ZipFolder 递归地将文件夹压缩成ZIP文件
func ZipFolder(srcDir, destZip string) error {
	// 将目标ZIP文件路径按"/"分割成列表
	des_path_list := strings.Split(destZip, "/")
	// 获取ZIP文件所在的目录路径
	des_path := strings.Join(des_path_list[:len(des_path_list)-1], "/")
	// 如果目录不存在，则创建目录
	if !FileIsExisted(des_path) {
		MakeDir(des_path)
	}

	// 创建ZIP文件
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建ZIP写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 递归遍历文件夹并添加到ZIP中
	err = filepath.Walk(srcDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建ZIP条目的相对路径
		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			return err
		}

		// 如果是文件夹，则在ZIP中添加一个目录条目
		if fileInfo.IsDir() {
			// 创建目录的ZIP文件头
			zipHeader, err := zip.FileInfoHeader(fileInfo)
			if err != nil {
				return err
			}
			// 设置目录条目的名称
			zipHeader.Name = relPath + "/"
			// 在ZIP中添加目录条目
			_, err = zipWriter.CreateHeader(zipHeader)
			if err != nil {
				return err
			}
			return nil
		}

		// 打开文件并创建ZIP文件条目
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// 创建文件的ZIP文件头
		// 在ZIP中添加文件条目
		zipHeader, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		// 设置文件条目的名称
		zipHeader.Name = relPath
		// 在ZIP中添加文件条目
		writer, err := zipWriter.CreateHeader(zipHeader)
		if err != nil {
			return err
		}

		// 将文件内容复制到ZIP条目中
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
func FindDir(dir string, num int) {
	fileinfo, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {
		// 重复输出制表符，模拟层级结构
		print(strings.Repeat("\t", num))

		// 判断是不是目录
		if fi.IsDir() {
			// 如果是目录，则递归调用FindDir函数，并增加层级数
			println(`目录：`, fi.Name())
			FindDir(dir+`/`+fi.Name(), num+1)
		} else {
			// 如果是文件，则直接输出文件名
			println(`文件：`, fi.Name())
		}
	}
}
func FindCurrentDir(dir string) ([]string, []string) {
	// 存储目录列表
	dir_list := []string{}
	// 存储文档列表
	document_list := []string{}

	// 读取目录下的所有文件和文件夹
	fileinfo, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 判断是不是目录
		if fi.IsDir() {
			// 如果是目录，则添加到目录列表中
			dir_list = append(dir_list, fi.Name())
		} else {
			// 如果不是目录，则添加到文档列表中
			document_list = append(document_list, fi.Name())
		}
	}

	// 返回目录列表和文档列表
	return dir_list, document_list
}
