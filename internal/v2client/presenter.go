package v2client

import (
	log "github.com/sirupsen/logrus"
)

// Presenter -
type Presenter struct {
	ServiceReport ServiceReport
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
