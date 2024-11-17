package dtos

type Tenders struct {
	Limit    int  `form:"limit"`
	Offset   int  `form:"offset"`
	ClientID uint `form:"-"`
}

type Bids struct {
	Limit        int  `form:"limit"`
	Offset       int  `form:"offset"`
	ContractorID uint `form:"-"`
	TenderID     uint `form:"tender_id"`
}
