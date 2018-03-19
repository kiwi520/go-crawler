package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte , error)  {
	resp , err :=http.Get(url)
	if err != nil{
		return  nil,err
	}
	defer  resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code",resp.StatusCode)
		return nil,fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	encode:= deteRmineEnconding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader,encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

/**
自动识别html字符编码
 */
func deteRmineEnconding(r *bufio.Reader)  encoding.Encoding {
	bytes ,err  := r.Peek(1024)
	if err !=nil {
		return unicode.UTF8
	}
	encode,_,_ := charset.DetermineEncoding(bytes,"")

	if encode !=nil {

	}
	return encode
}
