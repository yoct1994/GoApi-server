package response

type ErrorRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
