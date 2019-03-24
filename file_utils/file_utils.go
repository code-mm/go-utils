package file_utils

import (
	"os"
	"path"
	_ "path"
	"path/filepath"
)

// 创建文件
func CreateFile(path string) {
	os.Create(path)
}

// 创建文件夹
func Mkdir(path string) {
	os.Mkdir(path, 777)
}

// 强制创建文件夹
func Mkdirs(path string) {
	os.MkdirAll(path, 777)
}

// 删除文件
func DeleteFile(path string) {
	os.Remove(path)
}

// 重命名文件
func ReName(oldPath string, newPath string) {
	os.Rename(oldPath, newPath)

}

// 获取文件路径
func GetPath(path string) string {
	files := path
	paths, _ := filepath.Split(files)
	return paths
}

// 获取文件名称
func GetFileName(path string) string {
	files := path
	return filepath.Base(files)
}

// 获取文件后缀
func GetFileSuffix(paths string) string {
	files := paths
	return path.Ext(files)
}
