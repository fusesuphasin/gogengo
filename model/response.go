package model

type ShotResponse struct {
	Status        string          `json:"status"`
	ProcessStatus string          `json:"processStatus,omitempty"`
	AppTokenId    string          `json:"appTokenId,omitempty"`
	Timestamp     string          `json:"timestamp,omitempty"`
	Error         []ErrorResponse `json:"errorMessage,omitempty"`
}

type ErrorResponse struct {
	FailedField  string `json:"failedField,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Value        string `json:"value,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
