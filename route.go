package phony


type Route struct {
	Path string `json:"path"`
	Method string `json:"method"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func (r *Route) GetStatus() int {
	status := r.Status
	if status == 0 {
		status = 200
	}
	return status
}
