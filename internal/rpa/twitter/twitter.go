package twitter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/raulaguila/go-rpa/internal/rpa"
	"golang.org/x/exp/slices"
)

type RPATwitter struct {
	rpa.RPA
}

func (s *RPATwitter) Login(user string, pass string) error {
	s.NewPage("https://twitter.com/login", true)

	if err := s.InputElementByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/label/div/div[2]/div/input", user); err != nil {
		return err
	}

	if err := s.ClickElementByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[6]"); err != nil {
		return err
	}

	if err := s.InputElementByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div/div[3]/div/label/div/div[2]/div[1]/input", pass); err != nil {
		return err
	}

	if err := s.ClickElementByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div"); err != nil {
		return err
	}

	return nil
}

func (s *RPATwitter) ClickElement(el *rod.Element) error {
	s.Page.Keyboard.Press(input.ControlLeft)
	return el.Click(proto.InputMouseButtonLeft, 1)
}

func (s *RPATwitter) ListTweets() {
	fmt.Println("Listing tweets")
	count := 0

	for {
		time.Sleep(10 * time.Second)
		if err := s.Page.Reload(); err != nil {
			log.Printf("Error to reload: %v\n", err.Error())
			continue
		}
		s.Page.MustWaitNavigation()()
		tweets := s.Page.MustElementX("/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/div[5]/div/section/div/div/div/div")

		for index, tweetText := range tweets.MustElements("div[data-testid=tweetText]") {
			images := []string{}
			links := []string{}

			lnks, err := tweetText.Elements("a")
			if err != nil {
				log.Printf("Erro links: %v\n", err.Error())
			} else {
				for _, lnk := range lnks {
					link := *lnk.MustAttribute("href")
					if !slices.Contains(links, link) && (strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")) {
						links = append(links, link)
					}
				}
			}

			fmt.Printf("Index: %v\n", index)
			if err := s.ClickElement(tweetText); err != nil {
				log.Printf("Erro ClickElement: %v\n", err.Error())
				break
			}

			println("Iniciando!")
			time.Sleep(2 * time.Second)
			s.Page = s.Browser.MustPages().First()

			user, err := s.FindElementByXPath("/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div/div/div[1]/div/div/article/div/div/div[2]/div[2]/div/div/div[1]/div/div/div[1]/div/a/div/div[1]/span/span")
			if err != nil {
				log.Printf("Erro user: %v\n", err.Error())
			}

			date, err := s.FindElementByXPath("/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div/div/div[1]/div/div/article/div/div/div[3]/div[4]/div/div[1]/div/div[1]/a/time")
			if err != nil {
				log.Printf("Erro date: %v\n", err.Error())
			}

			contentXPath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div/div/div[1]/div/div/article/div/div/div[3]/div[2]/div"
			if s.HasByXPath(contentXPath) {
				content, err := s.FindElementByXPath(contentXPath)
				if err != nil {
					log.Printf("Erro content: %v\n", err.Error())
				}

				imgs, err := content.Elements("img")
				if err != nil {
					log.Printf("Erro image: %v\n", err.Error())
				} else {
					for _, img := range imgs {
						if !slices.Contains(images, *img.MustAttribute("src")) {
							images = append(images, *img.MustAttribute("src"))
						}
					}
				}

				lnks, err := content.Elements("a")
				if err != nil {
					log.Printf("Erro links: %v\n", err.Error())
				} else {
					for _, lnk := range lnks {
						link := *lnk.MustAttribute("href")
						if !slices.Contains(links, link) && (strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")) {
							links = append(links, link)
						}
					}
				}
			}

			fmt.Printf("Tweet url: %v\n", s.URL())
			fmt.Printf("Tweet txt: %v\n", strings.ReplaceAll(strings.TrimSpace(tweetText.MustText()), "\n", ""))
			fmt.Printf("Tweet usr: %v\n", user.MustText())
			fmt.Printf("Tweet images: %v\n", images)
			fmt.Printf("Tweet links: %v\n", links)
			fmt.Printf("Tweet date: %v\n\n", *date.MustAttribute("datetime"))

			s.Page.Close()
			s.Page = s.Browser.MustPages().First()
			count += 1
		}

		fmt.Printf("Total atual: %v\n", count)
	}
}
