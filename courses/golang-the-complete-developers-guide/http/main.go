package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func writeResponseToCustomWriter(resp *http.Response) {
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}


func writeResponseToStdout(resp *http.Response)	 {
	io.Copy(os.Stdout, resp.Body)
}

func getStringFromResponse(resp *http.Response) string {
	bs := make([]byte, 32*1024) // 32KB
	resp.Body.Read(bs)
	return string(bs)
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 1. write directly to stdout
	writeResponseToStdout(resp)

	// 2. get string from resp
	str := getStringFromResponse(resp)
	fmt.Println(str)

	// 3. write to custom writer
	writeResponseToCustomWriter(resp)
}
