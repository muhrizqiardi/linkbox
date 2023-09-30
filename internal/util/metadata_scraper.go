package util

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/muhrizqiardi/linkbox/internal/entities"
)

func MetadataScraper(res *http.Response) (entities.LinkMetadata, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return entities.LinkMetadata{}, err
	}

	result := entities.LinkMetadata{
		OG: entities.OpenGraph{},
	}

	metaTitle := doc.Find(`meta[property="og:title"]`)
	if metaTitle.Length() >= 1 {
		for _, e := range metaTitle.Get(0).Attr {
			if e.Key == "content" {
				result.OG.Title = e.Val
			}
		}
	}
	metaType := doc.Find(`meta[property="og:type"]`)
	if metaType.Length() >= 1 {
		for _, e := range metaType.Get(0).Attr {
			if e.Key == "content" {
				result.OG.Type = e.Val
			}

		}
	}
	metaURL := doc.Find(`meta[property="og:url"]`)
	if metaURL.Length() >= 1 {
		for _, e := range metaURL.Get(0).Attr {
			if e.Key == "content" {
				result.OG.URL = e.Val
			}
		}
	}
	metaDescription := doc.Find(`meta[property="og:description"]`)
	if metaDescription.Length() >= 1 {
		for _, e := range metaDescription.Get(0).Attr {
			if e.Key == "content" {
				result.OG.Description = e.Val
			}
		}
	}
	metaImageURL := doc.Find(`meta[property="og:image"]`)
	if metaImageURL.Length() >= 1 {
		for _, e := range metaImageURL.Get(0).Attr {
			if e.Key == "content" {
				result.OG.OGImage = make([]entities.OGImage, 1)
				result.OG.OGImage[0] = entities.OGImage{
					URL: e.Val,
				}
			}
		}
	}

	return result, nil
}
