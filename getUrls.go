package main
import (
  "os"
  "fmt"
	"regexp"
  "encoding/json"
	"github.com/gocolly/colly"
  "net/http"
  "bytes"
  "io/ioutil"
  "time"
  "log"
  "strings"
)

func main() {
  defer func() {
    if err := recover(); err != nil {
      log.Println("panic occurred:", err)
      e := err.(error)
    }
  }()
  start := time.Now()
  pageUrl := os.Args[1]

  urls := getUrlsFromPage(pageUrl)
  log.Printf("found %s urls on '%s'" , fmt.Sprint(len(urls)), pageUrl)

  urlsJson, _ := json.Marshal(urls)
  urlsObj :=  "{\"urls\":"+string(urlsJson)+"}"
  fmt.Printf(urlsObj)
  log.Printf("took %s", time.Since(start))
}

func getUrlsFromPage (pageUrl string) []string {
	c := colly.NewCollector()
  var urls []string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    url := e.Attr("href")
    r, _ := regexp.Compile(".*.html$")
    r2, _ := regexp.Compile("(/news/|/test/)")
    if r.MatchString(url) && r2.MatchString(url) {
      urls = append(urls, url)
    }
	})
  c.Visit(pageUrl)
  return urls
}
