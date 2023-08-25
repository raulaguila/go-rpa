package rpa

import (
	"time"

	"github.com/go-rod/rod"

	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type RPA struct {
	Page     *rod.Page
	Launcher *launcher.Launcher
	Browser  *rod.Browser
}

func (s *RPA) Init(visible bool) {
	s.Launcher = launcher.New().
		Headless(!visible).
		Devtools(true)

	s.Browser = rod.New().
		ControlURL(s.Launcher.MustLaunch()).
		// Trace(true).
		SlowMotion(1 * time.Second).
		MustConnect()
}

func (s *RPA) CloseAll() {
	// s.Launcher.Cleanup()
	s.Browser.Close()
}

func (s *RPA) ClosePage() {
	s.Page.Close()
	s.Page = s.Browser.MustPages()[0]
}

func (s *RPA) HasByXPath(xPath string) bool {
	return s.Page.MustHasX(xPath)
}

func (s *RPA) FindElementByXPath(xPath string) (*rod.Element, error) {
	return s.Page.ElementX(xPath)
}

func (s *RPA) ClickElementByXPath(xPath string) error {
	el, err := s.FindElementByXPath(xPath)
	if err != nil {
		return err
	}

	return el.Click(proto.InputMouseButtonLeft, 1)
}

func (s *RPA) InputElementByXPath(xPath string, text string) error {
	el, err := s.FindElementByXPath(xPath)
	if err != nil {
		return err
	}

	return el.Input(text)
}

func (s *RPA) NewPage(url string, closeActual bool) {
	if s.Page != nil && closeActual {
		s.Page.Close()
	}

	s.Page = s.Browser.MustPage(url)
	s.Page.MustWaitNavigation()()
}

func (s *RPA) Screenshot() {
	s.Page.MustScreenshot("")
}

func (s *RPA) OpenInNewPageElement(xPath string) error {
	el, err := s.FindElementByXPath(xPath)
	if err != nil {
		return err
	}

	s.Page.Keyboard.Press(input.ControlLeft)
	return el.Click(proto.InputMouseButtonLeft, 1)
}

func (s *RPA) URL() string {
	return s.Page.MustInfo().URL
}

func (s *RPA) Title() string {
	return s.Page.MustInfo().Title
}
