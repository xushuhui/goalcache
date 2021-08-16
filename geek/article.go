package geek

import (
	"encoding/json"
	"fmt"
)

type ArticleData struct {
	VideoID             string        `json:"video_id"`
	Sku                 interface{}   `json:"sku"`
	VideoCover          string        `json:"video_cover"`
	AuthorName          string        `json:"author_name"`
	TextReadVersion     int           `json:"text_read_version"`
	AudioSize           int           `json:"audio_size"`
	ArticleCover        string        `json:"article_cover"`
	Subtitles           []interface{} `json:"subtitles"`
	HadFreelyread       bool          `json:"had_freelyread"`
	AudioURL            string        `json:"audio_url"`
	ChapterID           string        `json:"chapter_id"`
	ColumnHadSub        bool          `json:"column_had_sub"`
	AudioDubber         string        `json:"audio_dubber"`
	ColumnCover         string        `json:"column_cover"`
	Like                Like          `json:"like"`
	Neighbors           interface{}   `json:"neighbors"`
	AudioTime           string        `json:"audio_time"`
	VideoHeight         int           `json:"video_height"`
	FreeGet             bool          `json:"free_get"`
	ArticleContent      string        `json:"article_content"`
	FooterCoverData     interface{}   `json:"footer_cover_data"`
	FreelyreadCount     int           `json:"freelyread_count"`
	HlsVideos           []interface{} `json:"hls_videos"`
	ArticleCoverHidden  bool          `json:"article_cover_hidden"`
	ColumnIsExperience  bool          `json:"column_is_experience"`
	IsRequired          bool          `json:"is_required"`
	Share               Share         `json:"share"`
	ProductID           int           `json:"product_id"`
	FreelyreadTotal     int           `json:"freelyread_total"`
	LikeCount           int           `json:"like_count"`
	HadLiked            bool          `json:"had_liked"`
	ArticleSubtitle     string        `json:"article_subtitle"`
	AudioDownloadURL    string        `json:"audio_download_url"`
	ID                  int           `json:"id"`
	AudioTimeArr        AudioTimeArr  `json:"audio_time_arr"`
	ArticleTitle        string        `json:"article_title"`
	ColumnBgcolor       string        `json:"column_bgcolor"`
	ArticlePosterWxlite string        `json:"article_poster_wxlite"`
	ArticleFeatures     int           `json:"article_features"`
	IsVideoPreview      bool          `json:"is_video_preview"`
	ArticleSummary      string        `json:"article_summary"`
	Score               string        `json:"score"`
	ColumnSaleType      int           `json:"column_sale_type"`
	OfflinePackage      string        `json:"offline_package"`
	ArticleCouldPreview bool          `json:"article_could_preview"`
	ColumnID            int           `json:"column_id"`
	VideoTime           string        `json:"video_time"`
	ArticleSharetitle   string        `json:"article_sharetitle"`
	AudioTitle          string        `json:"audio_title"`
	AudioMd5            string        `json:"audio_md5"`
	VideoSize           int           `json:"video_size"`
	TextReadPercent     int           `json:"text_read_percent"`
	CommentCount        int           `json:"comment_count"`
	Cid                 int           `json:"cid"`
	Offline             Offline       `json:"offline"`
	ColumnIsOnboard     bool          `json:"column_is_onboard"`
	ArticleCshort       string        `json:"article_cshort"`
	VideoWidth          int           `json:"video_width"`
	ColumnCouldSub      bool          `json:"column_could_sub"`
	ArticleCtime        int           `json:"article_ctime"`
	ProductType         string        `json:"product_type"`
}
type Like struct {
	HadDone bool `json:"had_done"`
	Count   int  `json:"count"`
}
type Neighbors struct {
	Left  []interface{} `json:"left"`
	Right []interface{} `json:"right"`
}
type FooterCoverData struct {
	ImgURL  string `json:"img_url"`
	LinkURL string `json:"link_url"`
	MpURL   string `json:"mp_url"`
}
type Share struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Poster  string `json:"poster"`
	Cover   string `json:"cover"`
}
type ArticleReq struct {
	ID               string `json:"id"`
	IncludeNeighbors bool   `json:"include_neighbors"`
	IsFreelyread     bool   `json:"is_freelyread"`
}

type ArticleResp struct {
	Error interface{} `json:"error"`
	Extra interface{} `json:"extra"`
	Data  ArticleData `json:"data"`
	Code  int         `json:"code"`
}

type Article struct {
	Id    int    `json:"id"`
	Title string `json:"title"`

	Content string `json:"content"`
}

var articleUrl = "https://time.geekbang.org/serv/v1/article"
var commentUrl = "https://time.geekbang.org/serv/v1/comments"
var columnBaseUrl = "https://time.geekbang.org/column/article/"
var columnUrl = "https://time.geekbang.org/serv/v1/column/articles"

func GetArticle(articleId string) (a Article) {
	jsonStr, _ := json.Marshal(ArticleReq{
		ID:               articleId,
		IncludeNeighbors: true,
		IsFreelyread:     true,
	})
	var resp ArticleResp

	Post(articleUrl, jsonStr, &resp, articleId)
	fmt.Println("content", resp.Data.ArticleContent)
	biz := resp.Data

	a = Article{
		biz.ID, biz.ArticleTitle, biz.ArticleContent,
	}
	return

}
