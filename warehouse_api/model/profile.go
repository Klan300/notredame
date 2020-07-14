package model

type CompanyProfile struct {
	// Address of company's headquarter.
	Address string `json:"address,omitempty"`
	// City of company's headquarter.
	City string `json:"city,omitempty"`
	// Country of company's headquarter.
	Country string `json:"country,omitempty"`
	// Currency used in company filings.
	Currency string `json:"currency,omitempty"`
	// CUSIP number.
	Cusip string `json:"cusip,omitempty"`
	// Sedol number.
	Sedol int64 `json:"sedol,omitempty"`
	// Company business summary.
	Description string `json:"description,omitempty"`
	// Listed exchange.
	Exchange string `json:"exchange,omitempty"`
	// GICS industry group.
	Ggroup string `json:"ggroup,omitempty"`
	// GICS industry.
	Gind string `json:"gind,omitempty"`
	// GICS sector.
	Gsector string `json:"gsector,omitempty"`
	// GICS sub-industry.
	Gsubind string `json:"gsubind,omitempty"`
	// ISIN number.
	Isin string `json:"isin,omitempty"`
	// NAICS national industry.
	NaicsNationalIndustry string `json:"naicsNationalIndustry,omitempty"`
	// NAICS industry.
	Naics string `json:"naics,omitempty"`
	// NAICS sector.
	NaicsSector string `json:"naicsSector,omitempty"`
	// NAICS subsector.
	NaicsSubsector string `json:"naicsSubsector,omitempty"`
	// Company name.
	Name string `json:"name,omitempty"`
	// Company phone number.
	Phone string `json:"phone,omitempty"`
	// State of company's headquarter.
	State string `json:"state,omitempty"`
	// Company symbol/ticker as used on the listed exchange.
	Ticker string `json:"ticker,omitempty"`
	// Company website.
	Weburl string `json:"weburl,omitempty"`
	// IPO date.
	Ipo string `json:"ipo,omitempty"`
	// Market Capitalization.
	MarketCapitalization int64 `json:"marketCapitalization,omitempty"`
	// Number of oustanding shares.
	ShareOutstanding float32 `json:"shareOutstanding,omitempty"`
	// Number of employee.
	EmployeeTotal int64 `json:"employeeTotal,omitempty"`
	// Logo image.
	Logo string `json:"logo,omitempty"`
	// Finnhub industry classification.
	FinnhubIndustry string `json:"finnhubIndustry,omitempty"`
}

type Data struct {
	Symbol string
	Data map[string]string 
}