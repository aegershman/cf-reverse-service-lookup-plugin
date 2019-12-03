package presenters

import (
	"github.com/aegershman/cf-service-reverse-lookup-plugin/models"
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
	case "json":
		p.asJSON()
	default:
		log.Debugf("unknown format [%s], using default\n", p.Format)
		p.asJSON()
	}
}
