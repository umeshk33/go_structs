package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Printf("%+v", resp)

	//interface can be defined as a field inside a struct
	//as long as it satisfies the requirements of that interface

	//we can embed one interface inside the other
	/*
		type ReadCloser interface {
			Reader
			Closer
		}
		type Reader interface {
			Read([]byte) (int, error)
		}
	*/
	//saves us rewriting a lot of code

	//declare a byte slice that will contain the response body
	//reader interface
	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	//writer interface copy(dst Writer, src Reader)
	//dst has to implement writer interface, here *File implements writer
	//src has to have a implementation of reader interface, Body->read
	//	io.Copy(os.Stdout, resp.Body)

	//implement our own writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes: ", len(bs))
	return len(bs), nil
}
