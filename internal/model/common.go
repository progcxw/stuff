package model

type ListParam struct {
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type ListResp struct {
	Lists interface{} `json:"lists"`
	Total int64       `json:"total"`
}
