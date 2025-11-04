package scrape

import (
	"context"
	"findX/graph"
	"log"

	"github.com/chromedp/chromedp"
)

func x_scrape() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	g := graph.NewSocialGraph()
	scraper := NewScraper(ctx, "")

	if err := scraper.GetUserData("nasa", g); err != nil {
		log.Fatal(err)
	}

}
