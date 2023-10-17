package job

import (
    "net/http"
	"strings"
	"strconv"
    "github.com/PuerkitoBio/goquery"
	"NewProjectSearchApp/constants"
)

type JobInfo struct {
    Name, URL string
}

// フリーランススタート
func GetFreelanceStartDetails() ([]JobInfo, error) {
    var jobInfoSlice []JobInfo

    // サイト情報取得
    resp, err := http.Get("https://fa-works.com/offerlist/?skills=74")
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    doc.Find("div.border-b-gray-300 div.bg-primary.bg-opacity-5.border-indigo-100.mb-8.rounded").Each(func(i int, s *goquery.Selection) {

        pElements := s.Find("p.line-clamp-1")
		priceText := pElements.Eq(0).Text()
		priceText = strings.Replace(priceText, ",", "", -1)
        priceText = strings.Replace(priceText, " 〜", "", -1)
        priceText = strings.Replace(priceText, "円/月", "", -1)
		price, _ := strconv.Atoi(priceText)

        secondPText := pElements.Eq(1).Text()

		// 新規案件 && リモート案件 && 指定単価以上
        if
            secondPText == "フルリモート" &&
			price >= constants.UnitPrice {

            url, _ := s.Find("h1 a").Attr("href")
            name := s.Find("h1 a").Text()
            jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://fa-works.com"+url})
        }
    })

    return jobInfoSlice, nil
}

// レバテック
func GetLevtechDetails() ([]JobInfo, error) {
    var jobInfoSlice []JobInfo

    // サイト情報取得
    resp, err := http.Get("https://freelance.levtech.jp/project/skill-10/")
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    doc.Find("ul.prjList li").Each(func(i int, s *goquery.Selection) {

		priceText := s.Find("li.prjContent__summary__price span").Text()
		priceText = strings.Replace(priceText, "円", "", -1)
		priceText = strings.Replace(priceText, ",", "", -1)
		price, _ := strconv.Atoi(priceText)

		// 新規案件 && リモート案件 && 指定単価以上
        if
			s.Find("p.prjLabel__txt:contains('New')").Length() > 0 &&
			s.Find("li.prjContent__feature__item:contains('リモートOK')").Length() > 0 &&
			price >= constants.UnitPrice {
				
            url, _ := s.Find("a.js-link_rel").Attr("href")
            name := s.Find("a.js-link_rel").Text()
            jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://freelance.levtech.jp/"+url})
        }
    })

    return jobInfoSlice, nil
}
