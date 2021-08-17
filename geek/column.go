package geek

import (
	"encoding/json"
)

type ColumnResp struct {
	Error []ClientError `json:"error"`
	Extra []Extra       `json:"extra"`
	Data  ColumnData    `json:"data"` //columndata
	Code  int           `json:"code"`
}

type ClientError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
type Extra struct {
	Internal []interface{} `json:"internal"`
}
type ColumnReq struct {
	Cid        int      `json:"cid"`
	Size       int      `json:"size"`
	Prev       int      `json:"prev"`
	Order      string   `json:"order"`
	Sample     bool     `json:"sample"`
	ChapterIds []string `json:"chapter_ids"`
}

type ColumnData struct {
	List []List `json:"list"`
	Page Page   `json:"page"`
}
type List struct {
	ArticleSubtitle     string        `json:"article_subtitle"`
	AudioDownloadURL    string        `json:"audio_download_url,omitempty"`
	ID                  int           `json:"id"`
	ArticleSharetitle   string        `json:"article_sharetitle,omitempty"`
	ArticleTitle        string        `json:"article_title"`
	ArticleCover        string        `json:"article_cover"`
	Subtitles           []interface{} `json:"subtitles"`
	HadFreelyread       bool          `json:"had_freelyread"`
	IsVideoPreview      bool          `json:"is_video_preview"`
	ArticleSummary      string        `json:"article_summary"`
	ColumnHadSub        bool          `json:"column_had_sub"`
	AudioDubber         string        `json:"audio_dubber,omitempty"`
	ColumnCover         string        `json:"column_cover,omitempty"`
	AuthorName          string        `json:"author_name,omitempty"`
	ColumnID            int           `json:"column_id,omitempty"`
	ChapterID           string        `json:"chapter_id"`
	AudioTime           string        `json:"audio_time,omitempty"`
	AudioTitle          string        `json:"audio_title,omitempty"`
	AudioMd5            string        `json:"audio_md5,omitempty"`
	AudioTimeArr        AudioTimeArr  `json:"audio_time_arr,omitempty"`
	ArticleCouldPreview bool          `json:"article_could_preview"`
	AudioSize           int           `json:"audio_size,omitempty"`
	ColumnSku           int           `json:"column_sku"`
	Offline             Offline       `json:"offline"`
	AudioURL            string        `json:"audio_url,omitempty"`
	IsRequired          bool          `json:"is_required"`
	ColumnBgcolor       string        `json:"column_bgcolor,omitempty"`
	Score               int64         `json:"score"`
	ArticleCtime        int           `json:"article_ctime"`
	IncludeAudio        bool          `json:"include_audio"`
}
type Page struct {
	Count int  `json:"count"`
	More  bool `json:"more"`
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

func ListColumn(courseId int) (list []Article) {
	//{"cid":166,"size":100,"prev":0,"order":"earliest","sample":false,
	//"chapter_ids":["916","917","918","919","950","1056","1181","1240","1238"]}
	jsonStr, _ := json.Marshal(ColumnReq{
		Cid:   courseId,
		Size:  100,
		Order: "earliest",
	})

	var resp ColumnResp
	Send(columnUrl, jsonStr, &resp)

	for _, v := range resp.Data.List {
		a := Article{
			v.ID, v.ArticleTitle, "",
		}
		list = append(list, a)
	}
	return
}
