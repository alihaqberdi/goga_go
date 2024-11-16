package types

type TenderStatus string

const (
	TenderStatusOpen    TenderStatus = "open"
	TenderStatusClosed  TenderStatus = "closed"
	TenderStatusAwarded TenderStatus = "awarded"
)
