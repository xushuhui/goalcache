package geek

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var articleUrl = "https://time.geekbang.org/serv/v1/article"
var commentUrl = "https://time.geekbang.org/serv/v1/comments"
var columnBaseUrl = "https://time.geekbang.org/column/article/"
var columnUrl = "https://time.geekbang.org/serv/v1/column/articles"
var courseUrl = "https://time.geekbang.org/serv/v1/column/intro"
var loginUrl = "https://account.geekbang.org/account/ticket/login"
var allCourseUrl = "https://time.geekbang.org/serv/v1/column/all"
var checkUserUrl = "https://u.geekbang.org/serv/v1/user/check_user"
var cookie string

func init() {
	file, err := os.Open("./cookie.txt")
	if err != nil {
		log.Println(err)

	}
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	cookie = strings.TrimSpace(string(b))

}

func getHeaders() (header map[string]string) {
	header = map[string]string{
		"Content-Type": "application/json",
		"Cookie":       cookie,
		"Referer":      "https://time.geekbang.org/dashboard/course",
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	}
	return
}

func Post(url string, jsonStr []byte) (res *http.Response, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	header := getHeaders()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err = client.Do(req)
	if err != nil {
		log.Println("err", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}

	//fmt.Println(string(body))

	return
}
func Send(url string, jsonStr []byte, resp interface{}) (cookies []*http.Cookie, err error) {
	res, _ := Post(url, jsonStr)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	cookies = res.Cookies()

	fmt.Println(string(body))
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("err ", err, url)
	}
	return
}
func Get(url string, resp interface{}) (err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	header := getHeaders()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println("err", err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("json err ", err, url)
	}

	return
}
