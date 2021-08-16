package main

import (
	"fmt"
	"geektime-ebook/geek"
	"github.com/mattn/godown"
	"log"
	"os"
	"strings"
	"time"
)

func generateMd(title string, content string) {
	file, err := os.OpenFile("./ebook/"+title+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	godown.Convert(file, strings.NewReader(content), &godown.Option{})

}
func main() {
	//var articleId  = "68319"
	//article := geek.GetArticle(articleId)
	//title := strings.Replace(article.Title,"|","-",1)
	//generateMd(title,article.Content)

	var courseId = 100020801
	columns := geek.ListColumn(courseId)
	for _, v := range columns {

		article := geek.GetArticle(geek.Int2String(v.Id))
		fmt.Println(article.Title)
		title := strings.Replace(article.Title, "|", "-", 1)

		generateMd(title, article.Content)
		fmt.Println(article.Title)
		time.Sleep(30 * time.Second)
	}
}
