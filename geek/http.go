package geek

import (
	"bytes"
	"encoding/json"
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
		"Referer":      columnBaseUrl,
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	}
	return
}

func Post(url string, jsonStr []byte, resp interface{}) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	header := getHeaders()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("err ", err, url)

	}

	//fmt.Println(string(body))

	return
}
