package common

type WebBaseResponse struct {
	Msg  string `json:"message"`
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type WebResponseCreate struct {
	Status *WebBaseResponse `json:"status"`
	Data   string           `json:"data"`
}

type WebResponseRead struct {
	Status *WebBaseResponse `json:"status"`
	Data   []Book           `json:"data"`
}

type WebResponseUpdate struct {
	Status *WebBaseResponse  `json:"status"`
	Data   *WebRequestUpdate `json:"data"`
}

type WebResponseDelete struct {
	Status *WebBaseResponse `json:"status"`
	Data   string           `json:"data"`
}

type WebResponseError struct {
	Status *WebBaseResponse `json:"status"`
	Data   interface{}      `json:"data"`
}
