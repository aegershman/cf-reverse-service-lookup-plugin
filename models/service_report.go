package models

// ServiceReport -
type ServiceReport struct {
	Organization `json:"organization"`
	Service      `json:"service"`
	Space        `json:"space"`
}
