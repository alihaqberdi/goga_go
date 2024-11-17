package dtos

type Tenders struct {
	Limit    int  `form:"limit"`
	Offset   int  `form:"offset"`
	ClientID uint `form:"-"`
}
