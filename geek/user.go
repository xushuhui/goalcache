package geek

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Resp struct {
	Error interface{} `json:"error"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
}
type LoginReq struct {
	Country   string `json:"country"`
	Cellphone string `json:"cellphone"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	Remember  int    `json:"remember"`
	Platform  int    `json:"platform"`
	Appid     int    `json:"appid"`
}
type LoginData struct {
	UID               int         `json:"uid"`
	Ucode             string      `json:"ucode"`
	UIDStr            string      `json:"uid_str"`
	Type              int         `json:"type"`
	Cellphone         string      `json:"cellphone"`
	Country           string      `json:"country"`
	Nickname          string      `json:"nickname"`
	Avatar            string      `json:"avatar"`
	Gender            string      `json:"gender"`
	Birthday          string      `json:"birthday"`
	Graduation        string      `json:"graduation"`
	Profession        string      `json:"profession"`
	Industry          string      `json:"industry"`
	Description       string      `json:"description"`
	Overdue           int         `json:"overdue"`
	Province          string      `json:"province"`
	City              string      `json:"city"`
	Mail              string      `json:"mail"`
	Wechat            string      `json:"wechat"`
	GithubName        string      `json:"github_name"`
	GithubEmail       string      `json:"github_email"`
	Company           string      `json:"company"`
	Post              string      `json:"post"`
	ExpirenceYears    string      `json:"expirence_years"`
	School            string      `json:"school"`
	RealName          string      `json:"real_name"`
	Openid            string      `json:"openid"`
	Euid              string      `json:"euid"`
	Subtype           int         `json:"subtype"`
	Role              int         `json:"role"`
	Name              string      `json:"name"`
	Address           string      `json:"address"`
	Mobile            string      `json:"mobile"`
	Contact           string      `json:"contact"`
	Position          string      `json:"position"`
	Passworded        bool        `json:"passworded"`
	CReateTime        int         `json:"create_time"`
	JoinInfoq         string      `json:"join_infoq"`
	Actives           interface{} `json:"actives"`
	IsStudent         int         `json:"is_student"`
	StudentExpireTime int         `json:"student_expire_time"`
	OssToken          string      `json:"oss_token"`
}

func Login(phone, password string) (err error) {

	jsonStr, _ := json.Marshal(LoginReq{
		Country: "86", Cellphone: phone, Password: password,
		Remember: 1, Platform: 3, Appid: 1,
	})
	var resp Resp
	cookies, err := Send(loginUrl, jsonStr, &resp)
	if err != nil {
		return
	}
	if resp.Code != 0 {
		e := resp.Error.(map[string]interface{})
		err = errors.New(e["msg"].(string))
		return
	}

	for _, v := range cookies {
		c := v.Name + "=" + v.Value + ";"
		cookie += c
	}

	err = ioutil.WriteFile("./cookie.txt", []byte(cookie), 777)
	return
}

type IsLoginResp struct {
	Error interface{} `json:"error"`
	Code  int         `json:"code"`
	Data  struct {
		Result bool `json:"result"`
	} `json:"data"`
}

//Referer https://time.geekbang.org/dashboard/course
func IsLogin() bool {
	var resp IsLoginResp
	Get(checkUserUrl, &resp)
	return resp.Data.Result

}
