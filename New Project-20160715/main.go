package main

import (
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
)

func main() {

    message := "Hi hello ho & #42 *&23 w are you...."

    
    urlToPost := "http://api.msg91.com/api/sendhttp.php"


    form := url.Values{}
    form.Add("authkey", "myAuthKey")
    form.Add("mobiles", "9981534313")
    form.Add("message", message)
    form.Add("sender", "MSGIND")
    form.Add("route", "4")
    
    fmt.Println(form.Encode())
    
	req, _ := http.NewRequest("POST", urlToPost, strings.NewReader(form.Encode()))

     req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}