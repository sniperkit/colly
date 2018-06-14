// +build advanced

package colly

import (
	"strings"

	// internal - core
	glod "github.com/sniperkit/glod/pkg"

	// internal - plugins
	"github.com/sniperkit/glod/plugin/provider/chiasenhac"  // TODO: `code` chiesenhac
	"github.com/sniperkit/glod/plugin/provider/dailymotion" // TODO: `code` dailymotion
	"github.com/sniperkit/glod/plugin/provider/facebook"    // TODO: `code` facebook
	"github.com/sniperkit/glod/plugin/provider/instagram"   // TODO: `code` instagram
	"github.com/sniperkit/glod/plugin/provider/nhaccuatui"  // TODO: `code` nhaccuatui
	"github.com/sniperkit/glod/plugin/provider/soundcloud"  // TODO: `code` soundcloud
	"github.com/sniperkit/glod/plugin/provider/twitter"     // TODO: `code` twitter
	"github.com/sniperkit/glod/plugin/provider/vimeo"       // TODO: `code` vimeo
	"github.com/sniperkit/glod/plugin/provider/youtube"     // TODO: `test` dailymotion
	"github.com/sniperkit/glod/plugin/provider/zing"        // TODO: `test` zing
)

const (
	initFacebook    string = "facebook"
	initInstagram   string = "instagram"
	initTwitter     string = "twitter"
	initYoutube     string = "youtube"
	initDailymotion string = "dailymotion"
	initVimeo       string = "vimeo"
	initSoundCloud  string = "soundcloud"
	initChiaSeNhac  string = "chiasenhac"
	initNhacCuaTui  string = "nhaccuatui"
	initZingMp3     string = "zing"
)

type ContentTypeMediaProvider struct {
	resp *Response
	err  error
}

func (c *ContentTypeMediaProvider) Check(resp *Response) error { return ErrNotImplementedYet }

func getGlod(link string) glod.Source {
	switch {

	case strings.Contains(link, initFacebook):
		// Facebook
		return &facebook.Facebook{}

	case strings.Contains(link, initInstagram):
		// Instagram
		return &instagram.Instagram{}

	case strings.Contains(link, initTwitter):
		// Twitter
		return &twitter.Twitter{}

	case strings.Contains(link, initYoutube):
		// Youtube
		return &youtube.Youtube{}

	case strings.Contains(link, initDailymotion):
		// Dailymotion
		return &dailymotion.Dailymotion{}

	case strings.Contains(link, initVimeo):
		// Vimeo
		return &vimeo.Vimeo{}

	case strings.Contains(link, initSoundCloud):
		// SoundCloud
		return &soundcloud.SoundCloud{}

	case strings.Contains(link, initNhacCuaTui):
		// nhaccuatui
		return &nhaccuatui.NhacCuaTui{}

	case strings.Contains(link, initChiaSeNhac):
		// ChiaSeNhac
		return &chiasenhac.ChiaSeNhac{}

	case strings.Contains(link, initZingMp3):
		// Zing
		return &zing.Zing{}
	}

	return nil
}
