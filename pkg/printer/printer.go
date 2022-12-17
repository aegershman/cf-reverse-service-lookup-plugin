package printer

import (
	"encoding/json"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/details"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/v2client"
	"io"
	"log"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Printer struct {
	client                 *v2client.Client
	writer                 io.Writer
	serviceInstanceDetails []details.ServiceInstanceDetails
}

func NewPrinter(client *v2client.Client, writer io.Writer) *Printer {
	return &Printer{
		client:                 client,
		writer:                 writer,
		serviceInstanceDetails: make([]details.ServiceInstanceDetails, 0),
	}
}

func (p *Printer) Print(formats []string, serviceInstanceDetails []details.ServiceInstanceDetails) {
	p.serviceInstanceDetails = serviceInstanceDetails
	if len(formats) == 0 {
		p.asPlainText()
	}

	for _, f := range formats {
		switch f {
		case "table":
			p.asTable()
		case "json":
			p.asJSON()
		case "plain-text":
			p.asPlainText()
		}
	}
}

func (p *Printer) asPlainText() {
	for _, report := range p.serviceInstanceDetails {
		fieldsJoined := strings.Join([]string{
			report.Service.GUID,
			report.Service.Name,
			report.Organization.Name,
			report.Space.Name,
		}, "\n")
		p.writer.Write([]byte(fieldsJoined))
	}
}

func (p *Printer) asTable() {
	table := tablewriter.NewWriter(p.writer)
	table.SetHeader([]string{
		"service_guid",
		"service",
		"org",
		"space",
	})

	for _, report := range p.serviceInstanceDetails {
		table.Append([]string{
			report.Service.GUID,
			report.Service.Name,
			report.Organization.Name,
			report.Space.Name,
		})
	}

	table.Render()
}

func (p *Printer) asJSON() {
	j, err := json.Marshal(p.serviceInstanceDetails)
	if err != nil {
		log.Fatalln(err)
	}
	p.writer.Write(j)
}
