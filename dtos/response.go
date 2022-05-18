package dtos

import "net/http"

const (
	ContextResponse = "x-response"
)

type Meta struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ExtraData interface{} `json:"extraData,omitempty"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data" swaggertype:"object"`
}

type Error struct {
	Meta Meta `json:"meta"`
}

// New returns new Meta.
func NewResponse(code int, extra interface{}, msg ...string) Meta {
	cd := 200
	if code > 0 {
		cd = code
	}

	m := http.StatusText(cd)
	if len(msg) > 0 {
		m = msg[0]
	}

	return Meta{
		Code:      cd,
		Message:   m,
		ExtraData: extra,
	}
}

type HealthCheckResponse struct {
	Meta Meta `json:"meta"`
}

type OSInfoResponse struct {
	Meta Meta       `json:"meta"`
	Data OSInfoData `json:"data"`
}
type OSInfoData struct {
	OSVersion string `json:"OSVersion"`
}

type UserInfoResponse struct {
	Meta Meta         `json:"meta"`
	Data UserInfoData `json:"data"`
}

type UserInfoData struct {
	Info string `json:"info"`
}

type MetricInfoResponse struct {
	Meta Meta           `json:"meta"`
	Data MetricInfoData `json:"data"`
}

type MetricInfoData struct {
	CPUUsed    string `json:"cpu_used"`
	MemoryUsed string `json:"memory_used"`
}
