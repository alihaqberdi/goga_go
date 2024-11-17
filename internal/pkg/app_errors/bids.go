package app_errors

var (
	TenderNotFound  = newErr("Tender not found", 404)
	TenderNotClosed = newErr("Tender is not closed", 400)
	BidNotFound     = newErr("Bid not found", 404)
	BidNotPending   = newErr("Bid is not pending", 400)
)
