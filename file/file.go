package file

import (
	"encoding/json"
	"io"
	"os"
)

// 写入字符串
func WriteString(filename string, content string) bool {
	f, err := os.Create(filename)
	if err != nil {
		println("failed to create file :" + filename + "\nerror:" + err.Error())
		return false
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		println("failed to write file :" + filename + "\nerror:" + err.Error())
		return false
	}
	return true
}

// 写入字节
func WriteBytes(filename string, content []byte) bool {
	f, err := os.Create(filename)
	if err != nil {
		println("failed to create file :" + filename + "\nerror:" + err.Error())
		return false
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		println("failed to write file :" + filename + "\nerror:" + err.Error())
		return false
	}
	return true
}

// 检测文件存在
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 读取文件字符串
func ReadString(filename string) string {
	if !Exists(filename) {
		println("file not exist :" + filename)
		return ""
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		println("failed to read file :" + filename + "\nerror:" + err.Error())
		return ""
	}
	return string(data)
}

// 读取文件字节
func ReadBytes(filename string) []byte {
	if !Exists(filename) {
		println("file not exist :" + filename)
		return nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		println("failed to read file :" + filename + "\nerror:" + err.Error())
		return nil
	}
	return data
}

// 写入Json数据
func WriteJson(filename string, data any) bool {
	bdata, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		println("failed to marshal json :" + filename + "\nerror:" + err.Error())
		return false
	}
	return WriteBytes(filename, bdata)
}

// 读取Json数据
func ReadJson(filename string, data any) bool {
	if !Exists(filename) {
		println("file not exist :" + filename)
		return false
	}
	bdata := ReadBytes(filename)
	if bdata == nil {
		println("failed to read file :" + filename)
		return false
	}
	err := json.Unmarshal(bdata, data)
	if err != nil {
		println("failed to unmarshal json :" + filename + "\nerror:" + err.Error())
		return false
	}
	return true
}

// 复制文件
func Copy(filename_s string, filename_t string) bool {
	if Exists(filename_t) {
		println("file already exist :" + filename_t)
		return false
	}
	if !Exists(filename_s) {
		println("file not exist :" + filename_s)
		return false
	}
	sourceFile, err := os.Open(filename_s)
	if err != nil {
		println("failed to open file :" + filename_s + "\nerror:" + err.Error())
		return false
	}
	defer sourceFile.Close()

	destFile, err := os.Create(filename_t)
	if err != nil {
		println("failed to create file :" + filename_t + "\nerror:" + err.Error())
		return false
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		println("failed to copy file :" + filename_s + " to " + filename_t + "\nerror:" + err.Error())
		return false
	}
	return true
}

// 删除文件
func Remove(filename string) bool {
	if !Exists(filename) {
		println("file not exist :" + filename)
		return false
	}
	err := os.Remove(filename)
	if err != nil {
		println("failed to delete file :" + filename + "\nerror:" + err.Error())
		return false
	}
	return true
}
