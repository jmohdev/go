package main

import (
	"Decorator_pattern/cipher"
	"Decorator_pattern/lzw"
	"fmt"
)

var sentData string
var recvData string

type Component interface {
	Operator(string)
}

type SendComponent struct{}

// Operator
func (self *SendComponent) Operator(data string) {
	// send data
	sentData = data
}

// Decorator
type ZipComponent struct {
	com Component
}

// Operator
func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

// Decorator
type EncryptComponent struct {
	key string
	com Component
}

// Operator
func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)

	if err != nil {
		panic(err) // exit
	}
	self.com.Operator(string(encryptData))
}

// Decorator
type DecryptComponent struct {
	key string
	com Component
}

// Operator
func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err) // exit
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

// Operator
func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))

	if err != nil {
		panic(err) // exit
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	// receive data
	recvData = data
}

// func main() {
// 	sender := &EncryptComponent{key: "abcde",
// 		com: &ZipComponent{
// 			com: &SendComponent{}}}

// 	sender.Operator("Hello World")
// 	fmt.Println(sentData)

// 	receiver := &UnzipComponent{
// 		com: &DecryptComponent{key: "abcde",
// 			com: &ReadComponent{}}}

// 	receiver.Operator(sentData)
// 	fmt.Println(recvData)

// }

func main() {
	sender := &ZipComponent{
		com: &SendComponent{}}

	sender.Operator("Hello World")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &ReadComponent{}}

	receiver.Operator(sentData)
	fmt.Println(recvData)

}
