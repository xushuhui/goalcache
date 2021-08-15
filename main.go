package main

import (
	"encoding/json"
	"fmt"
)

type ArticleReq struct {
	ID               string `json:"id"`
	IncludeNeighbors bool   `json:"include_neighbors"`
	IsFreelyread     bool   `json:"is_freelyread"`
}
type columnReq struct {
	ID int `json:"cid"`
}

var articleUrl = "https://time.geekbang.org/serv/v1/article"
var commentUrl = "https://time.geekbang.org/serv/v1/comments"
var columnBaseUrl = "https://time.geekbang.org/column/article/"
var columnUrl = "https://time.geekbang.org/serv/v1/column/articles"

func main() {
	//var articleId string = "67888"
	//GetArticle(articleId)
	var chapterId = 100020801
	ListChapter(chapterId)
}

func ListChapter(chapterId int) {
	jsonStr, _ := json.Marshal(ChapterReq{
		ID: chapterId,
	})

	var resp ChapterResp
	Post(articleUrl, jsonStr, resp)
	fmt.Println(resp.Data)
}
func GetArticle(articleId string) {
	jsonStr, _ := json.Marshal(ArticleReq{
		ID:               articleId,
		IncludeNeighbors: true,
		IsFreelyread:     true,
	})
	var resp ArticleResp
	Post(articleUrl, jsonStr, resp)
	fmt.Println(resp.Data.ArticleContent, resp.Data.ArticleTitle)

}
