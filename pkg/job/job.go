package job

import (
    "net/http"
	"strings"
	"strconv"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/PuerkitoBio/goquery"
	"NewProjectSearchApp/constants"
    "NewProjectSearchApp/database"
)

type JobInfo struct {
    Name, URL string
}

// FAworks
func GetFaworksDetails(db *sql.DB) ([]JobInfo, error) {
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
            if database.IsNewJobUnique(db, name) {
                jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://fa-works.com/"+url})
                database.InsertJob(db, name, url)
            }
        }
    })

    return jobInfoSlice, nil
}

// レバテック
func GetLevtechDetails(db *sql.DB) ([]JobInfo, error) {
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
            if database.IsNewJobUnique(db, name) {
                jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://freelance.levtech.jp/"+url})
                database.InsertJob(db, name, url)
            }
        }
    })

    return jobInfoSlice, nil
}

// AKKODIS
func GetAkkodisDetails(db *sql.DB) ([]JobInfo, error) {
    var jobInfoSlice []JobInfo

    // サイト情報取得
    resp, err := http.Get("https://freelance.akkodis.co.jp/projects/?skill=446&sort=created&direction=desc")
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    doc.Find("ul.main-offer-list li.main-offer-list-item.new_project").Each(func(i int, s *goquery.Selection) {

		priceText := s.Find("p.main-offer-list-item-overview-reward em").First().Text()
		priceText = strings.Replace(priceText, ",", "", -1)
		price, _ := strconv.Atoi(priceText)

		// 新規案件 && リモート案件 && 指定単価以上
        if
            (s.Find("div.main-offer-list-item-overview-table-inner.lh-2:contains('リモート')").Length() > 0 ||
             s.Find("div.upper a h2:contains('リモート')").Length() > 0 ) &&
			price >= constants.UnitPrice {
				
            url, _ := s.Find("div.upper a").Attr("href")
            name := s.Find("div.upper a h2").Text()
            if database.IsNewJobUnique(db, name) {
                jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://freelance.akkodis.co.jp/"+url})
                database.InsertJob(db, name, url)
            }
        }
    })

    return jobInfoSlice, nil
}

// geechs
func GetGeechsDetails(db *sql.DB) ([]JobInfo, error) {
    var jobInfoSlice []JobInfo

    // サイト情報取得
    resp, err := http.Get("https://geechs-job.com/project/go")
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    doc.Find("ul.p-article-project li.c-card.p-card-project").Each(func(i int, s *goquery.Selection) {

		priceText := s.Find("span.c-text_price").First().Text()
        priceText += "0000"
		price, _ := strconv.Atoi(priceText)

		// 新規案件 && リモート案件 && 指定単価以上
        if
             s.Find("span.c-newLabel:contains('New')").Length() > 0 &&
			price >= constants.UnitPrice {
				
            url, _ := s.Find("a.c-card_title_link").Attr("href")
            name := s.Find("a.c-card_title_link").Text()
            if database.IsNewJobUnique(db, name) {
                jobInfoSlice = append(jobInfoSlice, JobInfo{Name: name, URL: "https://geechs-job.com/"+url})
                database.InsertJob(db, name, url)
            }
        }
    })

    return jobInfoSlice, nil
}
