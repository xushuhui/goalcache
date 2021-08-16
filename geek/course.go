package geek

import "encoding/json"

type CourseResp struct {
	Data CourseData `json:"data"`
}
type CourseReq struct {
	//{"cid":139,"with_groupbuy":false}
	WithGroupbuy bool `json:"with_groupbuy"`
	Cid          int  `json:"cid"`
}
type Nps struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}
type FirstPromo struct {
	Price     int  `json:"price"`
	CouldJoin bool `json:"could_join"`
}
type GroupbuyForGift struct {
	SuccessUcount int  `json:"success_ucount"`
	CouldGroupbuy bool `json:"could_groupbuy"`
	Price         int  `json:"price"`
}
type Groupbuy struct {
	SuccessUcount int    `json:"success_ucount"`
	JoinCode      string `json:"join_code"`
	EndTime       int    `json:"end_time"`
	CouldGroupbuy bool   `json:"could_groupbuy"`
	HadJoin       bool   `json:"had_join"`
	Price         int    `json:"price"`
}
type CourseData struct {
	IsOnborad            bool            `json:"is_onborad"`
	ColumnUtime          int             `json:"column_utime"`
	ColumnPoster         string          `json:"column_poster"`
	ProductType          string          `json:"product_type"`
	ColumnVideoCover     string          `json:"column_video_cover"`
	IsSharesale          bool            `json:"is_sharesale"`
	Share                Share           `json:"share"`
	IsOnboard            bool            `json:"is_onboard"`
	SaleType             int             `json:"sale_type"`
	ColumnVideoMedia     string          `json:"column_video_media"`
	ColumnPrice          int             `json:"column_price"`
	ColumnCoverInner     string          `json:"column_cover_inner"`
	ColumnCoverWxlite    string          `json:"column_cover_wxlite"`
	ColumnPriceMarket    int             `json:"column_price_market"`
	AuthorInfo           string          `json:"author_info"`
	ArticleTotalCount    int             `json:"article_total_count"`
	IsShareget           bool            `json:"is_shareget"`
	ColumnEndTime        int             `json:"column_end_time"`
	ColumnTitle          string          `json:"column_title"`
	LastAid              int             `json:"last_aid"`
	IsIncludeAudio       bool            `json:"is_include_audio"`
	IsIncludePreview     bool            `json:"is_include_preview"`
	ColumnCtime          int             `json:"column_ctime"`
	ColumnBeginTime      int             `json:"column_begin_time"`
	ArticleCount         int             `json:"article_count"`
	IsChannel            bool            `json:"is_channel"`
	ColumnIntro          string          `json:"column_intro"`
	IsFinish             bool            `json:"is_finish"`
	AuthorName           string          `json:"author_name"`
	GotoCs               int             `json:"goto_cs"`
	ColumnName           string          `json:"column_name"`
	ColumnCover          string          `json:"column_cover"`
	ColumnCoverExplore   string          `json:"column_cover_explore"`
	ColumnShareTitle     string          `json:"column_share_title"`
	ColumnUnit           string          `json:"column_unit"`
	ColumnSubtitle       string          `json:"column_subtitle"`
	FooterCoverData      string          `json:"footer_cover_data"`
	ColumnSku            int             `json:"column_sku"`
	Nps                  Nps             `json:"nps"`
	FreelyreadTotal      int             `json:"freelyread_total"`
	IsPreorder           bool            `json:"is_preorder"`
	SubCount             int             `json:"sub_count"`
	FirstPromo           FirstPromo      `json:"first_promo"`
	ColumnType           int             `json:"column_type"`
	ID                   int             `json:"id"`
	ShowChapter          bool            `json:"show_chapter"`
	GroupbuyForGift      GroupbuyForGift `json:"groupbuy_for_gift"`
	Groupbuy             Groupbuy        `json:"groupbuy"`
	IsSaleProduct        bool            `json:"is_sale_product"`
	FreelyreadCount      int             `json:"freelyread_count"`
	HadFaved             bool            `json:"had_faved"`
	IsMemberSub          int             `json:"is_member_sub"`
	IsExperience         bool            `json:"is_experience"`
	HadSub               bool            `json:"had_sub"`
	ChannelBackAmount    int             `json:"channel_back_amount"`
	ColumnPosterWxlite   string          `json:"column_poster_wxlite"`
	NavID                int             `json:"nav_id"`
	LectureURL           string          `json:"lecture_url"`
	ColumnBgcolor        string          `json:"column_bgcolor"`
	LastChapterID        int             `json:"last_chapter_id"`
	AuthorIntro          string          `json:"author_intro"`
	UpdateFrequency      string          `json:"update_frequency"`
	ColumnSharesale      int             `json:"column_sharesale"`
	ColumnWxliteCode     string          `json:"column_wxlite_code"`
	ColumnSharesaleData  string          `json:"column_sharesale_data"`
	ArticleReqTotalCount int             `json:"article_req_total_count"`
	FlashPromoEtime      int             `json:"flash_promo_etime"`
	AuthorHeader         string          `json:"author_header"`
}
type Course struct {
	Id    int
	Title string
	Sku   int
}

func GetCourse(id int) Course {
	jsonStr, _ := json.Marshal(CourseReq{
		Cid: id,
	})
	var resp CourseResp
	Post(courseUrl, jsonStr, &resp)
	return Course{Sku: resp.Data.ColumnSku, Title: resp.Data.ColumnTitle}
}
