package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChapterResp struct {
	Error []interface{} `json:"error"`
	Extra []interface{} `json:"extra"`
	Data  []ChapterData `json:"data"`
	Code  int           `json:"code"`
}
type ChapterData struct {
	SourceID     int    `json:"source_id"`
	Title        string `json:"title"`
	ArticleCount int    `json:"article_count"`
	ID           string `json:"id"`
}
type ArticleResp struct {
	Error []interface{} `json:"error"`
	Extra []interface{} `json:"extra"`
	Data  ArticleData   `json:"data"`
	Code  int           `json:"code"`
}

type ArticleData struct {
	VideoID             string          `json:"video_id"`
	Sku                 string          `json:"sku"`
	VideoCover          string          `json:"video_cover"`
	AuthorName          string          `json:"author_name"`
	TextReadVersion     int             `json:"text_read_version"`
	AudioSize           int             `json:"audio_size"`
	ArticleCover        string          `json:"article_cover"`
	Subtitles           []interface{}   `json:"subtitles"`
	HadFreelyread       bool            `json:"had_freelyread"`
	AudioURL            string          `json:"audio_url"`
	ChapterID           string          `json:"chapter_id"`
	ColumnHadSub        bool            `json:"column_had_sub"`
	AudioDubber         string          `json:"audio_dubber"`
	ColumnCover         string          `json:"column_cover"`
	Like                Like            `json:"like"`
	Neighbors           Neighbors       `json:"neighbors"`
	AudioTime           string          `json:"audio_time"`
	VideoHeight         int             `json:"video_height"`
	FreeGet             bool            `json:"free_get"`
	ArticleContent      string          `json:"article_content"`
	FooterCoverData     FooterCoverData `json:"footer_cover_data"`
	FreelyreadCount     int             `json:"freelyread_count"`
	HlsVideos           []interface{}   `json:"hls_videos"`
	ArticleCoverHidden  bool            `json:"article_cover_hidden"`
	ColumnIsExperience  bool            `json:"column_is_experience"`
	IsRequired          bool            `json:"is_required"`
	Share               Share           `json:"share"`
	ProductID           int             `json:"product_id"`
	FreelyreadTotal     int             `json:"freelyread_total"`
	LikeCount           int             `json:"like_count"`
	HadLiked            bool            `json:"had_liked"`
	ArticleSubtitle     string          `json:"article_subtitle"`
	AudioDownloadURL    string          `json:"audio_download_url"`
	ID                  int             `json:"id"`
	AudioTimeArr        AudioTimeArr    `json:"audio_time_arr"`
	ArticleTitle        string          `json:"article_title"`
	ColumnBgcolor       string          `json:"column_bgcolor"`
	ArticlePosterWxlite string          `json:"article_poster_wxlite"`
	ArticleFeatures     int             `json:"article_features"`
	IsVideoPreview      bool            `json:"is_video_preview"`
	ArticleSummary      string          `json:"article_summary"`
	Score               string          `json:"score"`
	ColumnSaleType      int             `json:"column_sale_type"`
	OfflinePackage      string          `json:"offline_package"`
	ArticleCouldPreview bool            `json:"article_could_preview"`
	ColumnID            int             `json:"column_id"`
	VideoTime           string          `json:"video_time"`
	ArticleSharetitle   string          `json:"article_sharetitle"`
	AudioTitle          string          `json:"audio_title"`
	AudioMd5            string          `json:"audio_md5"`
	VideoSize           int             `json:"video_size"`
	TextReadPercent     int             `json:"text_read_percent"`
	CommentCount        int             `json:"comment_count"`
	Cid                 int             `json:"cid"`
	Offline             Offline         `json:"offline"`
	ColumnIsOnboard     bool            `json:"column_is_onboard"`
	ArticleCshort       string          `json:"article_cshort"`
	VideoWidth          int             `json:"video_width"`
	ColumnCouldSub      bool            `json:"column_could_sub"`
	ArticleCtime        int             `json:"article_ctime"`
	ProductType         string          `json:"product_type"`
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
type AudioTimeArr struct {
	M string `json:"m"`
	S string `json:"s"`
	H string `json:"h"`
}
type Offline struct {
	Size        int    `json:"size"`
	FileName    string `json:"file_name"`
	DownloadURL string `json:"download_url"`
}

var header = map[string]string{
	"Content-Type": "application/json",
	"Cookie":       `gksskpitn=de393850-d059-4951-b42f-7dec52e181b5; SERVERID=1fa1f330efedec1559b3abbcb6e30f50|1629041696|1629041677; LF_ID=1629041679367-896198-3518063; gk_process_ev={%22count%22:1%2C%22target%22:%22%22}; Hm_lvt_59c4ff31a9ee6263811b23eb921a5083=1629041679; Hm_lpvt_59c4ff31a9ee6263811b23eb921a5083=1629041679; Hm_lvt_022f847c4e3acd44d4a2481d9187f1e6=1629041679; Hm_lpvt_022f847c4e3acd44d4a2481d9187f1e6=1629041679; GRID=ab42c0d-2f9da3f-9301a29-c5ce5ac; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2217b4a735e622cf-0a708c8fa7890a-4c3e257a-1065024-17b4a735e6340f%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_landing_page%22%3A%22https%3A%2F%2Ftime.geekbang.org%2Fcolumn%2Farticle%2F67888%22%7D%2C%22%24device_id%22%3A%2217b4a735e622cf-0a708c8fa7890a-4c3e257a-1065024-17b4a735e6340f%22%7D; sajssdk_2015_cross_new_user=1; _ga=GA1.2.2116361856.1629041681; _gid=GA1.2.837279957.1629041681; _gat=1`,
	"Referer":      columnBaseUrl,
	"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
}

func Post(url string, jsonStr []byte, resp interface{}) {
	client := &http.Client{}
	fmt.Print(string(jsonStr))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &resp)
	fmt.Print(resp)
	return
}
