package v2client

import (
	"encoding/json"
	"io"
	"log"

	"github.com/olekukonko/tablewriter"
)

// Presenter -
type Presenter struct {
	formats        []string
	serviceReports []ServiceReport
	writer         io.Writer
}

// NewPresenter -
func NewPresenter(formats []string, serviceReports []ServiceReport, writer io.Writer) *Presenter {
	return &Presenter{
		formats:        formats,
		serviceReports: serviceReports,
		writer:         writer,
	}
}

// Render -
func (p *Presenter) Render() {
	if len(p.formats) == 0 {
		p.asTable()
	}

	for _, f := range p.formats {
		switch f {
		case "table":
			p.asTable()
		case "json":
			p.asJSON()
		}
	}
}

func (p *Presenter) asTable() {
	table := tablewriter.NewWriter(p.writer)
	table.SetHeader([]string{
		"service_guid",
		"service",
		"org",
		"space",
	})

	for _, report := range p.serviceReports {
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
	j, err := json.Marshal(p.serviceReports)
	if err != nil {
		log.Fatalln(err)
	}
	p.writer.Write(j)
}
