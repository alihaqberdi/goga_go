package types

type BidStatus string

const (
	BidStatusPending BidStatus = "pending"
	BidStatusAwarded BidStatus = "awarded"
)
