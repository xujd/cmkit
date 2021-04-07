package fileupload

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Service 服务接口
type Service interface {
	Upload(r *http.Request) (*FileData, error)
}

// FileUploadService 文件上传服务
type FileUploadService struct {
	WebDir string
}

// 文件上传
func (s FileUploadService) Upload(r *http.Request) (*FileData, error) {
	r.ParseMultipartForm(32 << 20) // max memory is set to 32MB
	clientfd, _, err := r.FormFile("uploadfile")
	if err != nil {
		return nil, err
	}
	defer clientfd.Close()
	destLocalPath := "./files/"
	fileName := "temp.png"
	if r.Form.Get("fileName") != "" {
		fileName = r.Form.Get("fileName")
	}
	localpath := fmt.Sprintf("%s%s", destLocalPath, fileName)
	localfd, err := os.OpenFile(localpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer localfd.Close()

	// 利用io.TeeReader在读取文件内容时计算hash值
	fhash := sha1.New()
	io.Copy(localfd, io.TeeReader(clientfd, fhash))
	hstr := hex.EncodeToString(fhash.Sum(nil))

	// 拷贝到web目录
	copy(localpath, fmt.Sprintf("%s%s", s.WebDir, fileName))

	return &FileData{FileName: fileName, Hash: hstr}, nil
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
