package utils

import (
	"encoding/json"
	"fmt"
	"github.com/biu7/biu.to/conf"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func UrlCheck(unCheckedUrl string) (bool, string) {
	googleCheckResult := googleUrlCheck(unCheckedUrl)
	if !googleCheckResult {
		return false, "Google 安全网址检测未通过"
	}
	return true, ""
}

func googleUrlCheck(unCheckedUrl string) bool {
	if !conf.UseGoogleUrlCheck {
		return true
	}

	googleApiUrl := fmt.Sprintf("https://safebrowsing.googleapis.com/v4/threatMatches:find?key=%s", conf.GoogleApiKey)
	postBody := map[string]interface{}{
		"client": map[string]string{
			"clientId":      "biu.to",
			"clientVersion": "1.0",
		},
		"threatInfo": map[string][]interface{}{
			"threatTypes":      {"MALWARE", "SOCIAL_ENGINEERING"},
			"platformTypes":    {"WINDOWS"},
			"threatEntryTypes": {"URL"},
			"threatEntries": {map[string]string{
				"url": unCheckedUrl,
			}},
		},
	}
	b, _ := json.Marshal(postBody)
	fmt.Println(string(b))
	resp, _ := http.Post(googleApiUrl, "application/json", strings.NewReader(string(b)))

	switch resp.StatusCode {
	case 200:
	case 400:
		log.Error("Google 网站检测错误：无效的参数！")
		return false
	case 403:
		log.Error("Google 网站检测错误：API 密钥权限被拒绝！")
		return false
	case 429:
		log.Error("Google 网站检测错误：超出资源配额或达到速率限制！")
		return false
	case 500:
		log.Error("Google 网站检测错误：内部服务器错误！")
		return false
	case 503:
		log.Error("Google 网站检测错误：服务不可用！")
		return false
	case 504:
		log.Error("Google 网站检测错误：超时！")
		return false
	default:
		return false
	}

	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	if strings.TrimSpace(string(result)) == "{}" {
		return true
	}
	return false
}
