package scrape

import (
	"context"
	"findX/graph"
	"findX/model"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// Scraper holds chromedp context and optional auth token.
type Scraper struct {
	Ctx       context.Context
	AuthToken string
}

// NewScraper initializes a new Scraper.
func NewScraper(ctx context.Context, authToken string) *Scraper {
	return &Scraper{
		Ctx:       ctx,
		AuthToken: authToken,
	}
}

// GetUserData scrapes a users's info, followers, and following list.
func (s *Scraper) GetUserData(username string, g *graph.SocialGraph) error {
	profileURL := fmt.Sprintf("https://x.com/%s", username)

	var name, bio, location string
	log.Printf("visiting profile: %s", profileURL)
	tasks := chromedp.Tasks{
		chromedp.Navigate(profileURL),
		chromedp.Sleep(4 * time.Second),
		chromedp.Text(`header h1 span`, &name, chromedp.ByQuery, chromedp.NodeVisible),
		chromedp.Text(`div[data-testid="UserDescription"]`, &bio, chromedp.ByQuery),
		chromedp.Text(`div[data-testid="UserProfileHeader_Items"]`, &location, chromedp.ByQuery),
	}

	if err := chromedp.Run(s.Ctx, tasks); err != nil {
		return fmt.Errorf("failed to scrape profile: %w", err)
	}

	mainUser := &model.User{
		ID:        model.NodeID(username),
		Name:      strings.TrimSpace(name),
		Platform:  "X",
		Handle:    username,
		Bio:       strings.TrimSpace(bio),
		Location:  strings.TrimSpace(location),
		CreatedAt: time.Now(),
	}
	g.AddUser(mainUser)
	log.Printf("✅ Scraped user: %s (%s)", mainUser.Name, username)

	// --- Followers ---
	followersURL := fmt.Sprintf("https://x.com/%s/followers", username)
	var followers []string
	if err := s.extractList(followersURL, &followers); err == nil {
		for _, f := range followers {
			fUser := &model.User{
				ID:       model.NodeID(f),
				Handle:   f,
				Platform: "X",
			}
			g.AddUser(fUser)
			g.AddRelationship(fUser.ID, mainUser.ID, graph.Followers, 1)
		}
		log.Printf("👥 Followers scraped: %d", len(followers))
	}

	// --- Following ---
	followingURL := fmt.Sprintf("https://x.com/%s/following", username)
	var following []string
	if err := s.extractList(followingURL, &following); err == nil {
		for _, f := range following {
			fUser := &model.User{
				ID:       model.NodeID(f),
				Handle:   f,
				Platform: "X",
			}
			g.AddUser(fUser)
			g.AddRelationship(mainUser.ID, fUser.ID, graph.Following, 1)
		}
		log.Printf("➡️ Following scraped: %d", len(following))
	}
	return nil
}

// extractList scrapes a few usernames from /followers or /following pages.
func (s *Scraper) extractList(url string, out *[]string) error {
	log.Printf("🔗 Loading list: %s", url)
	if err := chromedp.Run(s.Ctx, chromedp.Navigate(url)); err != nil {
		return err
	}
	time.Sleep(4 * time.Second) // wait for DOM to render
	var handles []string
	err := chromedp.Run(s.Ctx, chromedp.Evaluate(`
		Array.from(document.querySelectorAll('a[href^="/"]'))
			.map(a => a.getAttribute("href"))
			.filter(v => v && !v.includes("/status/") && !v.includes("/photo/"))
	`, &handles))
	if err != nil {
		return fmt.Errorf("extract failed: %w", err)
	}
	seen := map[string]bool{}
	for _, h := range handles {
		h = strings.TrimPrefix(h, "/")
		if h != "" && !strings.Contains(h, "/") && !seen[h] {
			seen[h] = true
			*out = append(*out, h)
		}
	}
	log.Printf("🧾 Extracted %d accounts from %s", len(*out), url)
	return nil
}
