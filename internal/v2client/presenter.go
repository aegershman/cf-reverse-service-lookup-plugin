package v2client

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/olekukonko/tablewriter"
)

// Presenter -
type Presenter struct {
	ServiceReport []ServiceReport
	Format        []string
}

// Render -
func (p *Presenter) Render() {
	if len(p.Format) == 0 {
		p.asTable()
	}

	for _, f := range p.Format {
		switch f {
		case "table":
			p.asTable()
		case "json":
			p.asJSON()
		default:
			log.Debugf("unknown format [%s], using default\n", p.Format)
		}
	}
}

func (p *Presenter) asTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"service_guid",
		"service",
		"org",
		"space",
	})

	for _, report := range p.ServiceReport {
		table.Append([]string{
			report.Service.GUID,
			report.Service.Name,
			report.Organization.Name,
			report.Space.Name,
		})
	}

	table.Render()
}

func (p *Presenter) asJSON() {
	j, err := json.Marshal(p.ServiceReport)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(j))
}
