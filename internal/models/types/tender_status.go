package types

type TenderStatus string

func (r TenderStatus) Valid() bool {
	return validTenderStatus[r]
}

const (
	TenderStatusOpen    TenderStatus = "open"
	TenderStatusClosed  TenderStatus = "closed"
	TenderStatusAwarded TenderStatus = "awarded"
)

var validTenderStatus = map[TenderStatus]bool{
	TenderStatusOpen:    true,
	TenderStatusClosed:  true,
	TenderStatusAwarded: true,
}
