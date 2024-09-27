package parser

type ProductInfo struct {
	Success    bool   `json:"success"`
	Errors     []any  `json:"errors"`
	Goods      Goods  `json:"goods"`
	GoodsIDAlt string `json:"goodsIdAlt"`
}

type Goods struct {
	GoodsID     string `json:"goodsId"`
	Title       string `json:"title"`
	TitleImage  string `json:"titleImage"`
	WebURL      string `json:"webUrl"`
	Slug        string `json:"slug"`
	CategoryID  string `json:"categoryId"`
	Brand       string `json:"brand"`
	Stocks      int    `json:"stocks"`
	PhotosCount int    `json:"photosCount"`
	OffersCount int    `json:"offersCount"`
}
