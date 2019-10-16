package utils

import (
	"fmt"
	"github.com/biu7/biu.to/conf"
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
	fmt.Println(googleApiUrl)
	return true

}
