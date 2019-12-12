package v2client

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func (p *Presenter) asTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"org",
		"space",
		"service",
	})

	for _, report := range p.ServiceReport {
		table.Append([]string{
			report.Organization.Name,
			report.Space.Name,
			report.Service.Name,
		})
	}

	table.Render()

}
