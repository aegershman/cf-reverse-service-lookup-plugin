package presenters

import (
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/models"
	log "github.com/sirupsen/logrus"
)

// Presenter -
type Presenter struct {
	ServiceReport models.ServiceReport
	Format        string
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "table":
		p.asTable()
	case "json":
		p.asJSON()
	default:
		log.Debugf("unknown format [%s], using default\n", p.Format)
		p.asJSON()
	}
}
