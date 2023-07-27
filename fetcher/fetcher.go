package fetcher

import (
	"strings"

	"github.com/michiwend/gomusicbrainz"
)

var All = []Fetcher{Ilkpop{}}

// Fetcher defines a standard interface for all music fetchers.
type Fetcher interface{
	Fetch(music *gomusicbrainz.Release) []Result
}

type MusicInfo struct{
	Name string
	Artist string
	Album string
}

type Download struct{
	Name string `json:"name"`
	Artist string
	Album string
	URL string `json:"url"`
}

type Result struct{
	Name string
	Artist string
	Album string
	Downloads []Download `json:"downloads"`
}

func artistString(ac gomusicbrainz.ArtistCredit) string {
	sb := strings.Builder{}

	for _, credit := range ac.NameCredits {
		// TODO: get joinphrase (requires fork)
		sb.WriteString(credit.Artist.Name)
	}

	return sb.String()
}
