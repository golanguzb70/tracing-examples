package models

type ProductCreateReq struct {
	ProductName string `json:"product_name"`
}

type ProductUpdateReq struct {
	Id           int    `json:"id"`
	ProductName string `json:"product_name"`
}

type ProductGetReq struct {
	Id int `json:"id"`
}

type ProductFindReq struct {
	Page             int    `json:"page"`
	Limit            int    `json:"limit"`
	OrderByCreatedAt uint64 `json:"order_by_created_at"`
	Search           string `json:"search"`
}

type ProductDeleteReq struct {
	Id int `json:"id"`
}

type ProductFindResponse struct {
	Products []*ProductResponse `json:"products"`
	Count     int                 `json:"count"`
}

type ProductResponse struct {
	Id           int    `json:"id"`
	ProductName string `json:"product_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ProductApiResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *ProductResponse
}

type ProductApiFindResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *ProductFindResponse
}
