# VSCode with Go

## Downloads Page : https://go.dev/dl/  

(Windows 기준) .msi 파일을 실행하거나 .zip 파일을 압축 해제해서 Go 폴더 생성  

![Go_downloads_page_head](./img/../../img/go_001.png)
![Go_downloads_page_body](./img/../../img/go_002.png)

- .msi File  
    > .msi 파일의 기본 설치 경로는 "C:\Program Files\Go" 이지만 다른 경로로 변경해서 설치해도 된다.  

    ![.msi](./img/../../img/go_003.png)

- .zip File  
   
    ![package_file](./img/../../img/go_004.png)
    ![unzip_package_file](./img/../../img/go_005.png)

- Go Folder  

    ![GO_Folder](./img/../../img/go_006.png)

## env 
Command 창을 통해 go 설치 확인
```
go version  
```
![go_version](./img/../../img/go_007.png)

go 환경변수에서 GOPATH 확인
```
go env  
```
> set GOPATH=C:\Users\username\go
![go_env](./img/../../img/go_008.png)

GOPATH 경로 안에 bin, pkg, src 폴더 생성

```
mkdir C:\Users\username\go\bin  
mkdir C:\Users\username\go\pkg  
mkdir C:\Users\username\go\src  
```
![gopath_folder](./img/../../img/go_009.png)

## VSCode with Go
src 폴더 안에 작업 폴더(project) 생성
```
mkdir C:\Users\username\go\src\project
```
<p></p>
<p>VSCode 를 실행하여 Go 폴더 열기  </p>
<img src="./img/../../img/go_010.png" width="50%" height="50%" alt="exec_VSCode">
<img src="./img/../../img/go_011.png" width="50%" height="50%" alt="open_Gofolder_compact_on">
<p></p>
<details>
<summary>explorer 창 디렉터리 트리 설정 변경</summary>
<p></p>
<p>Settings (Ctrl + ,)</p>
<img src="./img/../../img/go_012.png" width="50%" height="50%" alt="VSCode_Settings">
<p></p>
<p>Search 창에서 explorer.compactFolders 검색</p>
<img src="./img/../../img/go_013.png" width="50%" height="50%" alt="Search_Settings">
<p></p>
<img src="./img/../../img/go_014.png" width="50%" height="50%" alt="Search_explorer.compactFolders">
<p></p>
<p>Explorer.Compact Folders 설정 체크박스 해제</p>
<img src="./img/../../img/go_015.png" width="100%" height="100%" alt="Explorer.Compact_Folders_on">
<img src="./img/../../img/go_016.png" width="100%" height="100%" alt="Explorer.Compact_Folders_off">
<p></p>
<p>explorer 창 디렉터리 트리 확인</p>
<img src="./img/../../img/go_017.png" width="50%" height="50%" alt="open_Gofolder_compact_off">
</details>  
<p></p>
<p>main.go 파일 생성</p>
<p></p>
<img src="./img/../../img/go_018.png" width="50%" height="50%" alt="make_main.go">
<p></p>
<p>Go Extension 설치</p>
<p></p>
<img src="./img/../../img/go_019.png" width="50%" height="50%" alt="message_go_extension">
<p></p>
<p>Go Extension 설치 확인</p>
<img src="./img/../../img/go_020.png" width="100%" height="100%" alt="go_extension_market">
<p></p>
<p>main.go 작성</p>
<p></p>
<img src="./img/../../img/go_021.png" width="50%" height="50%" alt="write_main.go">
<p></p>
<p>Go command Extension 설치</p>
<p></p>
<img src="./img/../../img/go_022.png" width="50%" height="50%" alt="message_go_command_extension">
<p></p>
<p>Go command Extension 설치 확인</p>
<p></p>
<img src="./img/../../img/go_023.png" width="50%" height="50%" alt="install_go_command_extension">
<p></p>
<p>현재 디렉터리를 Module Root 로 설정 후 main.go 실행</p>
<p></p>

```
go mod init
```
<img src="./img/../../img/go_024.png" width="75%" height="75%" alt="go_mod_init">
