package parser

type OfferYM struct {
	Location             string   `json:"location"`
	SkuID                string   `json:"skuId"`
	MarketSkuCreator     string   `json:"marketSkuCreator"`
	Price                int      `json:"price"`
	OldPrice             int      `json:"oldPrice"`
	VendorID             int      `json:"vendorId"`
	Hid                  int      `json:"hid"`
	Nid                  int      `json:"nid"`
	IsDigital            bool     `json:"isDigital"`
	DeliveryPartnerTypes []string `json:"deliveryPartnerTypes"`
	OfferColor           string   `json:"offerColor"`
	ProductID            int      `json:"productId"`
	WareID               string   `json:"wareId"`
	FeedID               int      `json:"feedId"`
	AvailableCount       int      `json:"availableCount"`
	ShopID               int      `json:"shopId"`
	SupplierID           int      `json:"supplierId"`
	ShopSku              string   `json:"shop_sku"`
	IsEda                bool     `json:"isEda"`
	IsExpress            bool     `json:"isExpress"`
	DeliveryOptions      []struct {
		PartnerType string `json:"partnerType"`
		DayFrom     int    `json:"dayFrom"`
		DayTo       int    `json:"dayTo"`
		Price       int    `json:"price"`
	} `json:"deliveryOptions"`
	WarehouseID        int    `json:"warehouseId"`
	IsAnyExpress       bool   `json:"isAnyExpress"`
	IsBnpl             bool   `json:"isBnpl"`
	IsInstallments     bool   `json:"isInstallments"`
	BusinessID         string `json:"businessId"`
	IsFoodtech         bool   `json:"isFoodtech"`
	AllDeliveryOptions []struct {
		Price       int    `json:"price"`
		PartnerType string `json:"partnerType"`
		DayFrom     int    `json:"dayFrom"`
		DayTo       int    `json:"dayTo"`
		Type        string `json:"type"`
	} `json:"allDeliveryOptions"`
	PaymentTypes []string `json:"paymentTypes"`
	IsOnDemand   bool     `json:"isOnDemand"`
	Benefit      struct {
		Type        string `json:"type"`
		IsPrimary   bool   `json:"isPrimary"`
		Description string `json:"description"`
	} `json:"benefit"`
	YandexBnplInfo struct {
		Enabled bool `json:"enabled"`
	} `json:"yandexBnplInfo"`
	YaBankPrice struct {
		Price struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"price"`
		Type string `json:"type"`
	} `json:"yaBankPrice"`
	IsDSBS bool `json:"isDSBS"`
	Promos []struct {
		Key         string `json:"key"`
		Type        string `json:"type"`
		ShopPromoID string `json:"shopPromoId"`
		LandingURL  string `json:"landingUrl"`
		IsPersonal  bool   `json:"isPersonal"`
	} `json:"promos"`
	PriceDetails struct {
		Price struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"price"`
		PriceWithoutVat struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"priceWithoutVat"`
		GreenPrice struct {
			Price struct {
				Value    int    `json:"value"`
				Currency string `json:"currency"`
			} `json:"price"`
			Type string `json:"type"`
		} `json:"greenPrice"`
		DiscountedPrice struct {
			Price struct {
				Value    int    `json:"value"`
				Currency string `json:"currency"`
			} `json:"price"`
			Discount struct {
				Value    int    `json:"value"`
				Currency string `json:"currency"`
			} `json:"discount"`
			Percent int `json:"percent"`
		} `json:"discountedPrice"`
	} `json:"priceDetails"`
	Type            string `json:"type"`
	SnippetType     string `json:"snippet_type"`
	ShowUID         string `json:"showUid"`
	PromoAttributes []struct {
		ShopPromoID   string `json:"shopPromoId"`
		PromoType     string `json:"promoType"`
		PromoKey      string `json:"promoKey"`
		ParentPromoID string `json:"parentPromoId,omitempty"`
	} `json:"promo_attributes"`
	ModelID          int    `json:"modelId"`
	MarketSku        string `json:"marketSku"`
	AdditionalPrices []struct {
		PriceType  string `json:"priceType"`
		PriceValue int    `json:"priceValue"`
	} `json:"additionalPrices"`
	PaymentMethodTypes []string `json:"paymentMethodTypes"`
	IsCrossBorder      string   `json:"isCrossBorder"`
	Pp                 int      `json:"pp"`
	Gci                string   `json:"gci"`
	ArticleNumber      int      `json:"articleNumber"`
	ShopSku0           string   `json:"shopSku"`
	SupplierType       int      `json:"supplierType"`
	StockItemCount     int      `json:"stockItemCount"`
}
