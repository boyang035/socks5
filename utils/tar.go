package utils

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// ZipTarDir 先打包后压缩
func ZipTarDir(srcDir, outName string) error{
	if err := TarDir(srcDir,outName); err != nil {
		return err
	}

	// 创建一个压缩对象
	f,err := os.Create(outName + ".zip")
	if err != nil {
		return err
	}
	defer f.Close()
	zp := zip.NewWriter(f)
	defer zp.Close()

	// 写入头部和内容
	fileInfo,err := os.Stat(outName + ".tar")
	if err != nil {
		return err
	}
	fileContent,err := os.Open(outName + ".tar")
	if err != nil {
		return err
	}
	hdr,err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	//hdr.Name =
	hdr.Method = zip.Deflate
	wt,err := zp.CreateHeader(hdr)
	if err != nil {
		return err
	}
	if _,err = io.Copy(wt,fileContent); err != nil {
		return err
	}
	fmt.Println("压缩成功！", outName + ".zip")
	return nil
}

// TarDir 打包目录
func TarDir(srcDir, outName string) error{
	// 创建一个tar文件
	outName += ".tar"
	baseFolder := filepath.Base(srcDir)		//sublime
	f,err := os.Create(outName)
	if err != nil{
		return err
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()

	// 遍历目录下的所有文件和文件夹
	err = filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		// 写入文件或目录的头部信息
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = filepath.Join(baseFolder, strings.TrimPrefix(path, srcDir))
		tw.WriteHeader(hdr)

		// 如果是文件，则写入内容
		if !info.IsDir() {
			fileContent,err := os.ReadFile(path)
			if err != nil{
				return err
			}
			if _,err = tw.Write(fileContent); err != nil {
				return err
			}
		}
		fmt.Println("打包:", path)
		return nil
	})
	if err != nil{
		return err
	}
	fmt.Println("打包成功！", outName)
	return nil
}