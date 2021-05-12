package main

import "net/http"

func main(){
	var(
		resp *http.Response
		err  error
		url string
	)
	url = "http://tracking.miui.com/test"
	if resp,err = http.Get(url); err != nil{

	}
}
