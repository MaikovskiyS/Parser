package parser

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type parser struct {
	timeout time.Duration
	url     string
}

func New() *parser {

	return &parser{
		url:     "https://hypeauditor.com/top-instagram-all-russia/",
		timeout: 120 * time.Second,
	}
}
func (h *parser) GetData() ([]string, error) {
	cl := http.DefaultClient
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, "GET", h.url, nil)
	if err != nil {
		return nil, err
	}
	res, err := cl.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("data.csv", 1, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Find the review items
	// doc.Find("div.table").Each(func(i int, s *goquery.Selection) {
	// 	title := s.Find(".contributor__title").Text()
	// 	content := s.Find(".tag__content ellipsis").Text()
	// 	fmt.Printf("Review %d: %s %s\n", i, title, content)
	// })

	// s := doc.Find("div.table").Find("div.row").Map(func(i int, s *goquery.Selection) string {
	// 	return s.Find(".contributor__title").Text()
	// })
	arr := []string{}
	doc.Find("div.table").Find("div.row").Each(func(i int, s *goquery.Selection) {
		//sC := s.Children()
		field1 := s.Find(".row-cell rank").AddBack().Text()
		//field := s.Find(".contributor__title").Text()
		//fmt.Println(field1)
		arr = append(arr, field1)
		//fmt.Println(s.Text())
	})

	return arr, nil
}
