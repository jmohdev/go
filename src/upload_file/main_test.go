package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/usernamne/Downloads/a.txt"
	// file open
	file, _ := os.Open(path)
	defer file.Close()

	// reset
	os.RemoveAll("./uploads")

	// buf: Data Storag
	buf := &bytes.Buffer{}
	// MIME
	// Data Store
	writer := multipart.NewWriter(buf)
	// filepath.Base(path): Extract only file name from path
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	// io.Copy(Dst, Src)
	io.Copy(multi, file)
	writer.Close()

	// Data Transfer
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// File check
	uploadFilePath := "./uploads/" + filepath.Base(path)
	// Load File Info.
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	// Downloads file vs Uploads file
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	// setting Byte Array
	uploadData := []byte{}
	originData := []byte{}
	// Read Downloads file vs Uploads file
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	// Check Downloads file vs Uploads file
	assert.Equal(originData, uploadData)
}
