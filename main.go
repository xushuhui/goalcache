package main

import (
	"geektime-ebook/geek"
	"github.com/mattn/godown"
	"log"
	"os"
	"strings"
)

func generateMd(dir string, title string, content string) {
	file, err := os.OpenFile(dir+"/"+title+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	title = `<h1>` + title + `</h1>`
	godown.Convert(file, strings.NewReader(title+content), &godown.Option{
		GuessLang: func(s string) (string, error) { return "sql", nil },
	})

}
func hasDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 777)
	}
}
func main() {
	//var err error
	//err = geek.Login("","")
	//if err != nil {
	//	log.Println(err)
	//}
	geek.IsLogin()

	//var sku = 100020801
	//cid := 139
	//course := geek.GetCourse(cid)
	//hasDir(course.Title)
	//
	//columns := geek.ListColumn(course.Sku)
	//for _, v := range columns {
	//
	//	article := geek.GetArticle(geek.Int2String(v.Id))
	//	//过滤 "|"
	//	title := strings.Replace(article.Title, "|", "-", 1)
	//
	//	generateMd(course.Title, title, article.Content)
	//	fmt.Println(article.Title)
	//	time.Sleep(10 * time.Second)
	//}
}
