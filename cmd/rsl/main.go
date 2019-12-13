package main

import (
	"flag"
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/v2client"
	log "github.com/sirupsen/logrus"
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
		logLevelFlag    string
		serviceGUIDFlag serviceGUIDFlag
		trimPrefixFlag  string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&formatFlag, "format", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")
	flagss.Var(&serviceGUIDFlag, "s", "")
	flagss.StringVar(&trimPrefixFlag, "trim-prefix", "service-instance_", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	if len(serviceGUIDFlag.guids) == 0 {
		log.Fatalln("please provide at least one -s service-instance_GUID")
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	cf, err := v2client.NewClient(cli)
	if err != nil {
		log.Fatalln(err)
	}

	var serviceReports []v2client.ServiceReport
	for _, service := range serviceGUIDFlag.guids {
		trimmedServiceGUID := strings.TrimPrefix(service, trimPrefixFlag)

		serviceInstance, err := cf.Services.GetServiceInstanceByGUID(trimmedServiceGUID)
		if err != nil {
			log.Fatalln(err)
		}

		serviceSpace, err := cf.Spaces.GetSpaceByGUID(serviceInstance.SpaceGUID)
		if err != nil {
			log.Fatalln(err)
		}

		serviceOrganization, err := cf.Orgs.GetOrganizationByGUID(serviceSpace.OrganizationGUID)
		if err != nil {
			log.Fatalln(err)
		}

		serviceReport := v2client.ServiceReport{
			Service:      serviceInstance,
			Space:        serviceSpace,
			Organization: serviceOrganization,
		}

		serviceReports = append(serviceReports, serviceReport)
	}

	presenter := v2client.Presenter{
		ServiceReport: serviceReports,
		Format:        formatFlag.formats,
	}

	presenter.Render()
}

// GetMetadata -
func (cmd *reverseServiceLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-reverse-service-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 5,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "rsl",
				HelpText: "reverse-service-lookup (rsl) against service_instance GUIDs",
				UsageDetails: plugin.Usage{
					Usage: "cf rsl [-s service_instance-xyzabc...]",
					Options: map[string]string{
						"format":      "format to present (options: table,json) (default: table)",
						"log-level":   "(options: info,debug,trace) (default: info)",
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
		log.Debugln("did you know plugin commands can still get ran when uninstalling a plugin? interesting, right?")
		return
	}
}

func main() {
	plugin.Start(new(reverseServiceLookupCmd))
}
