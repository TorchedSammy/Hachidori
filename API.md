# API Docs

The Hachidori API is very simple. At the moment there are only 2 endpoint.
Accessing the API is done via the `/api/:route` endpoint.
For example, search would be `/api/search`.

## /search
Searches via MusicBrainz for a release that matches the search query
- `?str=` provides the query to search for.

This endpoint returns an array/list of releases.

#### Fields
- mbid (string) - The MusicBrainz ID of the release. This can be used
in the fetchRelease endpoint below
- title (string) - Title of the release (for example, the name of an album)
- artist (string) - Name of the artist
- type (string) - The type of release (album, single)
- disambiguation (string) - What specifically is different about the release.
Most releases *won't* have this set, even if there are multiple releases that
seem the same. What differs releases most of the time is the region they're from.

#### Example
`/api/search?str=misamo%20do%20not%20touch` results in:  

```json
[
	{
		"mbid": "472c5a0d-12cc-4cf3-872b-59e458c8f5d6",
		"title": "Do not touch",
		"artist": "MISAMO",
		"type": "Single",
		"disambiguation": ""
	},
	{
		"mbid": "119d6651-41c0-4392-90d7-e220f109e989",
		"title": "Do not touch",
		"artist": "MISAMO",
		"type": "Single",
		"disambiguation": "Dolby Atmos mix"
	},
	{
		"mbid": "bac0e655-f245-430b-86fa-b834a4e44b1a",
		"title": "Doctor's Orders: Do Not Touch!",
		"artist": "Stranguliatorius",
		"type": "Album",
		"disambiguation": ""
	},
	{
		"mbid": "2f238f91-5007-4c6d-93f8-72f12ceb3e2c",
		"title": "Do Not",
		"artist": "Kleistwahr",
		"type": "Album",
		"disambiguation": ""
	},
	{
		"mbid": "52e8c07b-f692-4311-b315-8228da47bc41",
		"title": "DO NOT",
		"artist": "藤井フミヤ",
		"type": "Single",
		"disambiguation": ""
	}
]
```

## /fetchRelease
Retrieves the download links for the release from the
builtin fetchers.

The current available ones are:
- Ilkpop

This returns an array/list of downloads from each available fetcher.
#### Fields
- name (string) - Name of the release
- artist (string) - Name of the artist
- downloads (array) - List of downloads for this release. In an album, there
will be multiple entries here for each album track.

#### Fields (downloads)
- name (string) - Name of the *track*. API consumers should use this field to name
the music file, for example.
- artist (string) - Name of the artist. This is still useful, as it includes
who is featured in the specific track.
- album (string) - Name of the album.
- url (string) - A URL to a direct download of the music file.

#### Example
`/api/fetchRelease?id=472c5a0d-12cc-4cf3-872b-59e458c8f5d6` returns:  

```
[
	{
		"Name": "Do not touch",
		"Artist": "MISAMO",
		"Album": "Do not touch",
		"downloads": [
			{
				"name": "Do not touch",
				"Artist": "MISAMO",
				"Album": "Do not touch",
				"url": "https://pub-ae08218a46e24102994285e8d1eb6a3c.r2.dev/files/NTkyMDc2MzY.mp3"
			}
		]
	}
]
```
