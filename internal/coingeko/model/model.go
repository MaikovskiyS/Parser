package model

type Result struct {
	Coins []Coin
}
type Coin struct {
	Id                               string
	Symbol                           string
	Name                             string
	Image                            string
	Current_price                    float64
	Market_cap                       float64
	Market_cap_rank                  int
	Fully_diluted_valuation          float64
	Total_volume                     float64
	High_24h                         float64
	Low_24h                          float64
	Price_change_24h                 float64
	Price_change_percentage_24h      float64
	Market_cap_change_24h            float64
	Market_cap_change_percentage_24h float64
	Circulating_supply               float64
	Total_supply                     float64
	Max_supply                       float64
	Ath                              float64
	Ath_change_percentage            float64
	Ath_date                         string
	Atl                              float64
	Atl_change_percentage            float64
	Atl_date                         string
	Last_updated                     string
}
