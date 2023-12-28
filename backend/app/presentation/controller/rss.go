package controller

import (
	"context"
	"log"

	"github.com/minomusigk/tech-blog-summary/backend/oapi"
	"github.com/mmcdole/gofeed"
)

type RssController struct{}

func NewRssController() *RssController {
	return &RssController{}
}

type HatenaItem struct {
	Title string
	Link  string
}

func (*RssController) GetRss(c context.Context, _ oapi.GetRssRequestObject) (oapi.GetRssResponseObject, error) {
	url := "http://b.hatena.ne.jp/hotentry/it.rss"

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	items := feed.Items

	itemCount := len(items)

	log.Println("Number of items:", itemCount)

	return oapi.GetRss200JSONResponse{Title: items[0].Title, Link: items[0].Link}, nil
}
