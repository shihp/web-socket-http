package main

import (
"fmt"
"net/http"
"io/ioutil"
)

func main() {

	url := "http://www.sobot.com/call-open/getAccessToken/4?companyId=&email=hzlizhen@sobot.com"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "7798f348-41bb-4902-9f68-60d3e773336d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}