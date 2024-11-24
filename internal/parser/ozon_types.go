package parser

type OzonProductInfo struct {
	Context         string `json:"@context"`
	Type            string `json:"@type"`
	AggregateRating struct {
		Type        string `json:"@type"`
		RatingValue string `json:"ratingValue"`
		ReviewCount string `json:"reviewCount"`
	} `json:"aggregateRating"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Offers      struct {
		Type          string `json:"@type"`
		Availability  string `json:"availability"`
		Price         string `json:"price"`
		PriceCurrency string `json:"priceCurrency"`
		URL           string `json:"url"`
	} `json:"offers"`
	Sku string `json:"sku"`
}

type OzonInfo struct {
	Seo struct {
		Script []struct {
			InnerHTML string `json:"innerHTML"`
			Type      string `json:"type"`
		} `json:"script"`
	} `json:"seo"`
}
