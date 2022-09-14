package helpers

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseJSON(message string, code int, status string, data interface{}) Response {

	meta := Meta{Message: message, Code: code, Status: status}

	res := Response{Meta: meta, Data: data}

	return res
}
