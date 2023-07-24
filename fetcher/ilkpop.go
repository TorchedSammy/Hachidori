package fetcher

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"github.com/michiwend/gomusicbrainz"
)

// Ilkpop is a fetcher 
type Ilkpop struct{}

func (Ilkpop) Fetch(music *gomusicbrainz.Release) []Result {
	resp, err := http.Get(fmt.Sprintf("https://ilkpop.com/site_59.xhtml?get-q=%s&get-type=kpop", music.Title))
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom != atom.A {
			return false
		}

		if n.Parent == nil {
			return false
		}

		albumText := n.Parent.NextSibling
		if albumText == nil || albumText.DataAtom != atom.P {
			return false
		}
		
		albm := strings.Split(scrape.Text(albumText), " • ")[0]
		fmt.Println(scrape.Text(albumText))
		fmt.Println(music.Title)
		fmt.Println(albm)
		fmt.Println(albm == music.Title)
		return albm == music.Title
	}

	songs := scrape.FindAll(root, matcher)
	artist := artistString(music.ReleaseGroup.ArtistCredit)

	// TODO: properly handle multi cd here
	fmt.Println(music)
	fmt.Println(music.Mediums[0])
	downloads := make([]Download, len(music.Mediums[0].Tracks))

	var _ sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(songs))
	for _, t := range music.Mediums[0].Tracks {
		fmt.Println(t.Recording.Title)
	}
	for i, song := range songs {
		//fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
		var track *gomusicbrainz.Track
		for _, t := range music.Mediums[0].Tracks {
			fmt.Println(artist)
			fmt.Println(scrape.Text(song))
			songName := strings.TrimPrefix(scrape.Text(song), artist + " - ")
			fmt.Println(t.Recording.Title, "|", songName)
			if strings.ToLower(t.Recording.Title) == strings.ToLower(songName) {
				track = t
				break
			}
		}
		fmt.Printf("%+q\n", track)
		if track == nil {
			wg.Done()
			continue
		}
	//	go func() {
			resp, err := http.Get(fmt.Sprintf("https://ilkpop.com/%s", scrape.Attr(song, "href")))
			if err != nil {
				panic(err)
			}

			root, err := html.Parse(resp.Body)
			if err != nil {
				panic(err)
			}

			matcher := func(n *html.Node) bool {
				// must check for nil values
				return n.DataAtom == atom.Source
			}

			// grab all articles and print them
			audio, _ := scrape.Find(root, matcher)
			fmt.Println(track.Recording.Title)
			downloads[i] = Download{
				Name: track.Recording.Title,
				Artist: artist,
				Album: music.Title,
				URL: scrape.Attr(audio, "src"),
			}
			wg.Done()
	//	}()
	}

	wg.Wait()
	return []Result{
		{
			Name: music.Title,
			Artist: artist,
			Album: music.Title,
			Downloads: downloads,
		},
	}
}
