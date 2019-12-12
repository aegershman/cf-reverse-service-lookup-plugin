package main

import (
	"flag"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/internal/v2client"
	log "github.com/sirupsen/logrus"
)

type reverseServiceLookupCmd struct{}

// ReverseServiceLookupCommand -
func (cmd *reverseServiceLookupCmd) ReverseServiceLookupCommand(cli plugin.CliConnection, args []string) {
	var (
		formatFlag      string
		logLevelFlag    string
		serviceGUIDFlag string
		trimPrefixFlag  string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.StringVar(&formatFlag, "format", "json", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")
	flagss.StringVar(&serviceGUIDFlag, "service-guid", "", "")
	flagss.StringVar(&trimPrefixFlag, "trim-prefix", "service-instance_", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	trimmedServiceGUID := strings.TrimPrefix(serviceGUIDFlag, trimPrefixFlag)

	cf := v2client.NewClient(cli)

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

	presenter := v2client.Presenter{
		ServiceReport: serviceReport,
		Format:        formatFlag,
	}

	presenter.Render()

}

// GetMetadata -
func (cmd *reverseServiceLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-reverse-service-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 4,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "reverse-service-lookup",
				HelpText: "perform reverse lookups against service instance GUIDs",
				UsageDetails: plugin.Usage{
					Usage: "cf reverse-service-lookup --service-guid service_instance-xyzabc]",
					Options: map[string]string{
						"format":       "format to present (options: table,json) (default: json)",
						"log-level":    "(options: info,debug,trace) (default: info)",
						"service-guid": "GUID of service instance to reverse-lookup. Can be of form 'service_instance-xyzguid123' or just 'xyzguid123'",
						"trim-prefix":  "if your services are prefixed with something besides BOSH defaults, change this to be the string prefix before the service GUID... also, if you have that use-case, definitely let me know, I'm intrigued. (default: service_instance-)",
					},
				},
			},
		},
	}
}

// Run -
func (cmd *reverseServiceLookupCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "reverse-service-lookup" {
		cmd.ReverseServiceLookupCommand(cli, args)
	}
}

func main() {
	plugin.Start(new(reverseServiceLookupCmd))
}
