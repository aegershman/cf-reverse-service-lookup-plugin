package presenters

import "github.com/aegershman/cf-service-reverse-lookup-plugin/models"

// Presenter -
type Presenter struct {
	ServiceReport models.ServiceReport `json:"service_report"`
	Format        string               `json:"format"`
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "json":
		p.asJSON()
	default:
		p.asJSON()
	}
}
