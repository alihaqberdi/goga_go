package app_errors

var (
	TenderInvalidInput  = newErr("Invalid input", 400)
	TenderInvalidData   = newErr("Invalid tender data", 400)
	TenderInvalidStatus = newErr("Invalid tender status", 400)
)
