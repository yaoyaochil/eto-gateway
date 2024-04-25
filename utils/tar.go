package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// TarGzDir 将指定目录打包成tar.gz格式的文件
func TarGzDir(sourceDir, targetDir string) (fileName string, err error) {
	// 检查目的文件夹是否存在
	fmt.Println("sourceDir:", sourceDir)
	fmt.Println("targetDir:", targetDir)
	_, err = os.Stat(targetDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(targetDir, os.ModePerm)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	// 创建目标文件
	now := time.Now()
	fileName = fmt.Sprintf("%s_%s.tar.gz", now.Format("20060102150405"), filepath.Base(sourceDir))
	targetFile := filepath.Join(targetDir, fileName)
	file, err := os.Create(targetFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建gzip写入器
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// 创建tar写入器
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// 遍历源文件夹下的所有文件和子文件夹
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构造文件头信息
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name, _ = filepath.Rel(sourceDir, filePath)

		// 写入文件头信息
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件则写入文件数据
		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// UnTarGzDir 解压tar.gz文件
func UnTarGzDir(sourceFile, targetDir string) error {
	// 打开源文件
	tarGzFile, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("err1:", err)
		return err
	}
	defer tarGzFile.Close()

	// 创建gzip读取器
	gzipReader, err := gzip.NewReader(tarGzFile)
	if err != nil {
		fmt.Println("err2:", err)
		return err
	}
	defer gzipReader.Close()

	// 创建tar读取器
	tarReader := tar.NewReader(gzipReader)

	// 遍历tar文件中的所有文件
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			fmt.Println("err3:", err)
			break
		}
		if err != nil {
			fmt.Println("err4:", err)
			return err
		}

		// 构造目标文件路径
		targetFile := filepath.Join(targetDir, header.Name)

		if header.FileInfo().IsDir() {
			// 如果是文件夹则创建文件夹
			err = os.MkdirAll(targetFile, os.ModePerm)
			if err != nil {
				fmt.Println("err5:", err)
				return err
			}
		} else {
			// 如果是文件则创建文件并写入文件数据
			file, err := os.Create(targetFile)
			if err != nil {
				fmt.Println("err6:", err)
				return err
			}
			_, err = io.Copy(file, tarReader)
			if err != nil {
				fmt.Println("err7:", err)
				return err
			}
			file.Close()
		}
	}
	return nil
}

// GetFileSize 获取文件大小 (单位: kb)
func GetFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size() / 1024
	return fileSize, nil
}

// Removefile 删除文件
func Removefile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
