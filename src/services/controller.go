package services

type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type ResponseAll struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type ResponseDetail struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
