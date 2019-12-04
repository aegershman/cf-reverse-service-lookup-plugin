package main

import (
	"flag"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/apihelper"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/models"
	"github.com/aegershman/cf-reverse-service-lookup-plugin/presenters"
	log "github.com/sirupsen/logrus"
)

// ReverseServiceLookupCmd -
type ReverseServiceLookupCmd struct {
	apiHelper apihelper.CFAPIHelper
}

// ReverseServiceLookupCommand -
func (cmd *ReverseServiceLookupCmd) ReverseServiceLookupCommand(args []string) {
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

	serviceInstance, err := cmd.apiHelper.GetServiceInstanceByGUID(trimmedServiceGUID)
	if err != nil {
		log.Fatalln(err)
	}

	serviceSpace, err := cmd.apiHelper.GetSpaceByGUID(serviceInstance.SpaceGUID)
	if err != nil {
		log.Fatalln(err)
	}

	serviceOrganization, err := cmd.apiHelper.GetOrganizationByGUID(serviceSpace.OrganizationGUID)
	if err != nil {
		log.Fatalln(err)
	}

	serviceReport := models.ServiceReport{
		Service:      serviceInstance,
		Space:        serviceSpace,
		Organization: serviceOrganization,
	}

	presenter := presenters.Presenter{
		ServiceReport: serviceReport,
		Format:        formatFlag,
	}

	presenter.Render()

}

// GetMetadata -
func (cmd *ReverseServiceLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-reverse-service-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 2,
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
func (cmd *ReverseServiceLookupCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "reverse-service-lookup" {
		cmd.apiHelper = apihelper.New(cli)
		cmd.ReverseServiceLookupCommand(args)
	}
}

func main() {
	plugin.Start(new(ReverseServiceLookupCmd))
}
