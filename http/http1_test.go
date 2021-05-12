package http

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHttp(t *testing.T) {
	var(
		resp *http.Response
		err error
	)
	url := "http://tracking.miui.com/test"
	if resp, err = http.Get(url); err != nil{
		fmt.Println(err)
	}
	fmt.Println(resp)
}