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

type ProductOffers struct {
	Success     bool        `json:"success"`
	Errors      []any       `json:"errors"`
	Offers      []Offers    `json:"offers"`
	Merchants   []Merchants `json:"merchants"`
	IsAvailable bool        `json:"isAvailable"`
}

type Offers struct {
	ID                       string  `json:"id"`
	Price                    int     `json:"price"`
	Score                    int     `json:"score"`
	IsFavorite               bool    `json:"isFavorite"`
	MerchantID               string  `json:"merchantId"`
	FinalPrice               int     `json:"finalPrice"`
	BonusPercent             int     `json:"bonusPercent"`
	BonusAmount              int     `json:"bonusAmount"`
	AvailableQuantity        int     `json:"availableQuantity"`
	BonusAmountFinalPrice    int     `json:"bonusAmountFinalPrice"`
	DeliveryDate             string  `json:"deliveryDate"`
	PickupDate               string  `json:"pickupDate"`
	MerchantOfferID          string  `json:"merchantOfferId"`
	MerchantName             string  `json:"merchantName"`
	MerchantLogoURL          string  `json:"merchantLogoUrl"`
	MerchantURL              string  `json:"merchantUrl"`
	MerchantSummaryRating    float64 `json:"merchantSummaryRating"`
	IsBpgByMerchant          bool    `json:"isBpgByMerchant"`
	Nds                      int     `json:"nds"`
	OldPrice                 int     `json:"oldPrice"`
	OldPriceChangePercentage int     `json:"oldPriceChangePercentage"`
	MaxDeliveryDays          any     `json:"maxDeliveryDays"`
	BpgType                  string  `json:"bpgType"`
	CreditPaymentAmount      int     `json:"creditPaymentAmount"`
	InstallmentPaymentAmount int     `json:"installmentPaymentAmount"`
	ShowMerchant             any     `json:"showMerchant"`
	DueDate                  string  `json:"dueDate"`
	DueDateText              string  `json:"dueDateText"`
	LocationID               string  `json:"locationId"`
	SpasiboIsAvailable       bool    `json:"spasiboIsAvailable"`
	IsShowcase               bool    `json:"isShowcase"`
	SuperPrice               int     `json:"superPrice"`
	WarehouseID              string  `json:"warehouseId"`
	BnplPaymentParams        any     `json:"bnplPaymentParams"`
	InstallmentPaymentParams any     `json:"installmentPaymentParams"`
	CalculatedDeliveryDate   string  `json:"calculatedDeliveryDate"`
	GoodsID                  string  `json:"goodsId"`
}

type Merchants struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	MinDeliveryDate       string `json:"minDeliveryDate"`
	PartnerID             string `json:"partnerId"`
	Slug                  string `json:"slug"`
	URL                   string `json:"url"`
	SiteURL               string `json:"siteUrl"`
	ConfirmationTimeLimit string `json:"confirmationTimeLimit"`
	FullName              string `json:"fullName"`
}
