//看错了
//x.(T)断言
//T(x)类型转换
//qwq

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty:pair<type:*os.File,value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	/*	func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		testlog.Open(name)
		f, err := openFileNolog(name, flag, perm)
		if err != nil {
			return nil, err
		}
		f.appendMode = flag&O_APPEND != 0

		return f, nil
		}
	*/

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//r:pair<type: ,value: >
	var r io.Reader
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	//r:pair<type:*os.File,value:"/dev/tty"文件描述符>
	r = tty

	//w:pair<type: ,value: >
	var w io.Writer
	//w:<type:*os.File,value:"/dev/tty"文件描述符>
	w = r.(io.Writer) //将r强制转换成w

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))
}
