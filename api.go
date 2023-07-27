package main

import (
	"encoding/json"
	"fmt"

	"github.com/TorchedSammy/Hachidori/fetcher"
	"github.com/gofiber/fiber/v2"
	"github.com/michiwend/gomusicbrainz"
)

var mb, _ = gomusicbrainz.NewWS2Client(
    "https://musicbrainz.org/ws/2",
    "Hachidori",
    "0.0.1-beta",
    "http://github.com/TorchedSammy/Hachidori")

type Release struct{
	mbr *gomusicbrainz.Release
}

func (r *Release) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		ID gomusicbrainz.MBID `json:"mbid"`
		Title string `json:"title"`
	}{
		ID: r.mbr.ID,
		Title: r.mbr.Title,
	})
}

func initAPI(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/search", func(c *fiber.Ctx) error {
		resp, _ := mb.SearchRelease(c.Query("str"), -1, -1)
		fmt.Println(resp)

		releases := []Release{}
		for _, rel := range resp.Releases {
			releases = append(releases, Release{rel})
		}

		return c.JSON(releases)
	})

	api.Get("/fetchRelease", func(c *fiber.Ctx) error {
		results := []fetcher.Result{}
		release, _ := mb.LookupRelease(gomusicbrainz.MBID(c.Query("id")), "media", "recordings", "artist-credits", "release-groups")

		for _, fetcher := range fetcher.All {
			result := fetcher.Fetch(release)
			results = append(results, result...)
		}

		return c.JSON(results)
	})
}
