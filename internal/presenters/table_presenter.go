package presenters

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

	table.Append([]string{
		p.ServiceReport.Organization.Name,
		p.ServiceReport.Space.Name,
		p.ServiceReport.Service.Name,
	})

	table.Render()

}
