package main

import (
	"flag"
	"fmt"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/printer"
	"log"
	"os"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/pkg/v2client"
)

type serviceGUIDFlag struct {
	guids []string
}

func (s *serviceGUIDFlag) String() string {
	return fmt.Sprint(s.guids)
}

func (s *serviceGUIDFlag) Set(value string) error {
	s.guids = append(s.guids, value)
	return nil
}

type formatFlag struct {
	formats []string
}

func (f *formatFlag) String() string {
	return fmt.Sprint(f.formats)
}

func (f *formatFlag) Set(value string) error {
	f.formats = append(f.formats, value)
	return nil
}

type reverseServiceLookupCmd struct{}

func (cmd *reverseServiceLookupCmd) Run(cli plugin.CliConnection, args []string) {
	var (
		format      formatFlag
		serviceGuid serviceGUIDFlag
	)

	flagz := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagz.Var(&format, "format", "")
	flagz.Var(&serviceGuid, "s", "")

	err := flagz.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	if len(serviceGuid.guids) == 0 {
		log.Fatalln("please provide at least one -s service-instance_GUID")
	}

	apiEndpoint, err := cli.ApiEndpoint()
	if err != nil {
		log.Fatalln(err)
	}
	accessToken, err := cli.AccessToken()
	if err != nil {
		log.Fatalln(err)
	}

	accessToken = strings.TrimPrefix(accessToken, "bearer ")

	client, err := v2client.NewClient(apiEndpoint, accessToken)
	if err != nil {
		log.Fatalln(err)
	}

	var guids []string
	for _, service := range serviceGuid.guids {
		guid := strings.TrimPrefix(service, "service-instance_")
		guids = append(guids, guid)
	}

	serviceInstancesDetails, err := client.GetServiceInstancesDetails(guids...)
	if err != nil {
		log.Fatalln(err)
	}

	p := printer.NewPrinter(client, os.Stdout)
	p.Print(format.formats, serviceInstancesDetails)
}

func (cmd *reverseServiceLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-reverse-service-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 8,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "rsl",
				HelpText: "reverse-service-lookup (rsl) against service_instance GUIDs",
				UsageDetails: plugin.Usage{
					Usage: "cf rsl [-s service_instance-xyzabc...]",
					Options: map[string]string{
						"format": "format to present (options: table,json,plain-text) (default: plain-text)",
						"s":      "service_instance-GUID to look up. Can be of form 'service_instance-xyzguid123' or just 'xyzguid123'",
					},
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(reverseServiceLookupCmd))
}
