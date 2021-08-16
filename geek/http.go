package geek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var header = map[string]string{
	"Content-Type": "application/json",
	"Cookie":       "gksskpitn=aa908a7b-5c78-430e-b2c9-7a160c8bfd05; _ga=GA1.2.825834839.1606976856; LF_ID=1606976856481-902469-9666286; Hm_lvt_f83d162eeb1abea371d41d4b60d345db=1614754394; Hm_lpvt_f83d162eeb1abea371d41d4b60d345db=1614754394; MEIQIA_TRACK_ID=1n2hHVZFwnhuy09uFZfQgJeZhL0; GCID=d425958-34c12e6-ca8861f-db1d427; GRID=d425958-34c12e6-ca8861f-db1d427; _gid=GA1.2.267155807.1629075893; _gat=1; GCESS=BggBAwIE.DwaYQsCBgABCG3ODwAAAAAABgQj1SMOBQQAAAAADAEBDQEBAwT4PBphCgQAAAAABAQALw0ABwQCpFA1CQEB; Hm_lvt_59c4ff31a9ee6263811b23eb921a5083=1629075893,1629109498; Hm_lpvt_59c4ff31a9ee6263811b23eb921a5083=1629109498; Hm_lvt_022f847c4e3acd44d4a2481d9187f1e6=1629075893,1629109498; Hm_lpvt_022f847c4e3acd44d4a2481d9187f1e6=1629109498; gk_process_ev={%22referrer%22:%22https://time.geekbang.org/column/article/68319%22%2C%22utime%22:1629109496401%2C%22count%22:29%2C%22target%22:%22%22}; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221035885%22%2C%22first_id%22%3A%2217b4e5bf7d344a-05f0a34fb2345-c791e37-1327104-17b4e5bf7d5694%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E5%BC%95%E8%8D%90%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Faccount.infoq.cn%2F%22%2C%22%24latest_landing_page%22%3A%22https%3A%2F%2Ftime.geekbang.org%2Fcolumn%2Farticle%2F68319%22%2C%22%24latest_utm_source%22%3A%22wenzhangditu%22%2C%22%24latest_utm_medium%22%3A%22geektime%22%2C%22%24latest_utm_campaign%22%3A%22100073201%22%2C%22%24latest_utm_content%22%3A%22houduanhejiagou%22%2C%22%24latest_utm_term%22%3A%22zeus8VLFU%22%7D%2C%22%24device_id%22%3A%221762748e146419-0a647f7f5485e7-c791e37-2073600-1762748e1479f9%22%7D; SERVERID=3431a294a18c59fc8f5805662e2bd51e|1629109499|1629099837",
	"Referer":      columnBaseUrl,
	"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
}

func Post(url string, jsonStr []byte, resp interface{}, id string) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, _ := client.Do(req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("err ", err, url)

	}

	//fmt.Println(string(body))

	return
}
