package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/v2client"
)

type reverseServiceLookupCmd struct{}

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

// reverseServiceLookupCommand is the "real" main entrypoint into program execution
func (cmd *reverseServiceLookupCmd) reverseServiceLookupCommand(cli plugin.CliConnection, args []string) {
	var (
		formatFlag      formatFlag
		serviceGUIDFlag serviceGUIDFlag
		trimPrefixFlag  string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&formatFlag, "format", "")
	flagss.Var(&serviceGUIDFlag, "s", "")
	flagss.StringVar(&trimPrefixFlag, "trim-prefix", "service-instance_", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	if len(serviceGUIDFlag.guids) == 0 {
		log.Fatalln("please provide at least one -s service-instance_GUID")
	}

	cf, err := v2client.NewClient(cli)
	if err != nil {
		log.Fatalln(err)
	}

	var trimmedServiceGUIDs []string
	for _, service := range serviceGUIDFlag.guids {
		trimmedServiceGUID := strings.TrimPrefix(service, trimPrefixFlag)
		trimmedServiceGUIDs = append(trimmedServiceGUIDs, trimmedServiceGUID)
	}

	serviceReports, err := cf.ServiceReportService.GetServiceReportsFromServiceGUIDs(trimmedServiceGUIDs...)
	if err != nil {
		log.Fatalln(err)
	}

	presenter := v2client.NewPresenter(formatFlag.formats, serviceReports, os.Stdout)
	presenter.Render()
}

// GetMetadata -
func (cmd *reverseServiceLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-reverse-service-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "rsl",
				HelpText: "reverse-service-lookup (rsl) against service_instance GUIDs",
				UsageDetails: plugin.Usage{
					Usage: "cf rsl [-s service_instance-xyzabc...]",
					Options: map[string]string{
						"format":      "format to present (options: table,json,plain-text) (default: plain-text)",
						"s":           "service_instance-GUID to look up. Can be of form 'service_instance-xyzguid123' or just 'xyzguid123'",
						"trim-prefix": "if your services are prefixed with something besides BOSH defaults, change this to be the string prefix before the service GUID... also, if you have that use-case, definitely let me know, I'm intrigued. (default: service_instance-)",
					},
				},
			},
		},
	}
}

// Run -
func (cmd *reverseServiceLookupCmd) Run(cli plugin.CliConnection, args []string) {
	switch args[0] {
	case "rsl":
		cmd.reverseServiceLookupCommand(cli, args)
	default:
		return
	}
}

func main() {
	plugin.Start(new(reverseServiceLookupCmd))
}
