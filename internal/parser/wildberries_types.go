package parser

type WBResponse struct {
	State          int `json:"state"`
	PayloadVersion int `json:"payloadVersion"`
	Data           struct {
		Products []struct {
			ID          int    `json:"id"`
			Root        int    `json:"root"`
			KindID      int    `json:"kindId"`
			Brand       string `json:"brand"`
			BrandID     int    `json:"brandId"`
			SiteBrandID int    `json:"siteBrandId"`
			Colors      []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"colors"`
			SubjectID       int     `json:"subjectId"`
			SubjectParentID int     `json:"subjectParentId"`
			Name            string  `json:"name"`
			Entity          string  `json:"entity"`
			Supplier        string  `json:"supplier"`
			SupplierID      int     `json:"supplierId"`
			SupplierRating  float64 `json:"supplierRating"`
			SupplierFlags   int     `json:"supplierFlags"`
			Pics            int     `json:"pics"`
			Rating          int     `json:"rating"`
			ReviewRating    float64 `json:"reviewRating"`
			NmReviewRating  float64 `json:"nmReviewRating"`
			Feedbacks       int     `json:"feedbacks"`
			NmFeedbacks     int     `json:"nmFeedbacks"`
			PanelPromoID    int     `json:"panelPromoId"`
			PromoTextCard   string  `json:"promoTextCard"`
			PromoTextCat    string  `json:"promoTextCat"`
			Volume          int     `json:"volume"`
			ViewFlags       int     `json:"viewFlags"`
			Promotions      []int   `json:"promotions"`
			Sizes           []struct {
				Name     string `json:"name"`
				OrigName string `json:"origName"`
				Rank     int    `json:"rank"`
				OptionID int    `json:"optionId"`
				Stocks   []struct {
					Wh       int `json:"wh"`
					Dtype    int `json:"dtype"`
					Qty      int `json:"qty"`
					Priority int `json:"priority"`
					Time1    int `json:"time1"`
					Time2    int `json:"time2"`
				} `json:"stocks"`
				Time1 int `json:"time1"`
				Time2 int `json:"time2"`
				Wh    int `json:"wh"`
				Dtype int `json:"dtype"`
				Price struct {
					Basic     int `json:"basic"`
					Product   int `json:"product"`
					Total     int `json:"total"`
					Logistics int `json:"logistics"`
					Return    int `json:"return"`
				} `json:"price"`
				SaleConditions int    `json:"saleConditions"`
				Payload        string `json:"payload"`
			} `json:"sizes"`
			TotalQuantity int `json:"totalQuantity"`
			Time1         int `json:"time1"`
			Time2         int `json:"time2"`
			Wh            int `json:"wh"`
			Dtype         int `json:"dtype"`
		} `json:"products"`
	} `json:"data"`
}
