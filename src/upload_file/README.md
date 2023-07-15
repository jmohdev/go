# upload file

```bash
├── public
│   └── index.html
├── go.mod
├── go.sum
├── main_test.go
├── main.go
└── README.md
```
*** 
## /main.go

### func main
Go의 표준 패키지인 net/http 패키지에서 제공하는 Server, Client 기능 사용
```go
func main() {
    http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/uploads", uploadsHandler)
    http.ListenAndServe(":3000", nil)
}
```
#### 자주 사용하는 net/http 패키지 메소드
- ListenAndServe : 정의된 TCP network address로 Client의 Request를 받아서 ServeMux를 통해 작업 할당
 	```go
	func http.ListenAndServe(addr string, handler http.Handler) error
	```
  - addr : 개발시에는 일반적으로 3000번 포트로 설정(e.g. ":3000")
  - handler : Request를 Handle하기 위한 handler call로 ServeMux 정의하거나 nil 사용 (nil일 경우 DefaultServeMux 으로 동작)
  - error : ListenAndServe는 항상 non-nil error 반환
  #
- Handle : Request에서 요청된 Request Path에서 동작을 처리하는 Handler로 Routing
- HandleFunc : Request에서 요청된 Request Path에서 동작을 처리하는 Handler Function으로 Routing
	```go
	func http.Handle(pattern string, handler http.Handler)
	```
	```go
	func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	```
  - pattern : pattern 값 설정
  - handler : Handler 또는 Handler function로 pattern 값 처리 (ServeMux에 정의된 pattern 형식에 맞게 처리해야 함)
  - http.FileServer
  : 파일 시스템의 컨텐츠를 HTTP requests 로 Handler 를 통해 전달 (fs.FS 에서는 http.FS 를 사용)
	```go
	func http.FileServer(root http.FileSystem) http.Handler
	```
  - http.Dir
	```go
	func (http.Dir).Open(name string) (http.File, error)
	```
  - http.FS
	```go
	func http.FS(fsys fs.FS) http.FileSystem
	```
***
### func uploadsHandler
/public/index.html 의 Form 으로 "upload_file" 파일을 받고 ./uploads 폴더 안에 파일 생성

```go
func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// read
	uploadfile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadfile.Close()

	// make
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	
	// io.Copy(Dst, Src)
	io.Copy(file, uploadfile)
	// response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}
```
***
## /public/index.html
Form 으로 "upload_file" 파일을 ./uploads 폴더에 등록하는 POST Method 생성
```html
<html>
<head>
    <title> File Transfer </title>
</head>
<body>
    <p><h1> File Uploads!!!</h1></p>
    <p> </p>
    <form action="/uploads" method="POST" accept-charset="utf-8" enctype="multipart/form-data">
        <p><input type="file" id="upload_file" name="upload_file" /></p>
        <p><input type="submit" name="upload" /></p>
    </form>
</body>
</html>
```
***
## /main_test.go
```go
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
```