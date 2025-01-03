package utils

import "github.com/realtemirov/go-sqlc-grpc-http/generated/general"

func ConvertPageSizeToLimitOffset(req *general.GetAllRequest) {
	req.Limit = req.GetPageSize()
	req.Offset = req.GetPageSize() * (req.GetPage() - 1)
}
