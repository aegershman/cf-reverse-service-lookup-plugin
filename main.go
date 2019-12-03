package main

import (
	"flag"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/apihelper"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/models"
	"github.com/aegershman/cf-service-reverse-lookup-plugin/presenters"
	log "github.com/sirupsen/logrus"
)

// ServiceReverseLookupCmd -
type ServiceReverseLookupCmd struct {
	apiHelper apihelper.CFAPIHelper
}

// GetMetadata -
func (cmd *ServiceReverseLookupCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-service-reverse-lookup-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "service-reverse-lookup",
				HelpText: "perform reverse lookups against service instance GUIDs",
				UsageDetails: plugin.Usage{
					Usage: "cf service-reverse-lookup [-s serviceGuid]",
					Options: map[string]string{
						"s":         "serviceGuid to reverse-lookup. Can be of form 'service_instance-xyzguid123' or just 'xyzguid123'",
						"log-level": "(options: info,debug,trace) (default: info)",
					},
				},
			},
		},
	}
}

// ServiceReverseLookupCommand -
func (cmd *ServiceReverseLookupCmd) ServiceReverseLookupCommand(args []string) {
	var (
		serviceGUIDFlag string
		logLevelFlag    string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.StringVar(&serviceGUIDFlag, "s", "", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	serviceInstance, err := cmd.apiHelper.GetServiceInstance(serviceGUIDFlag)
	if err != nil {
		log.Fatalln(err)
	}

	serviceSpace, err := cmd.apiHelper.GetSpace(serviceInstance.SpaceURL)
	if err != nil {
		log.Fatalln(err)
	}

	serviceOrganization, err := cmd.apiHelper.GetOrganization(serviceSpace.OrganizationURL)
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
		Format:        "",
	}

	presenter.Render()

}

// Run -
func (cmd *ServiceReverseLookupCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "service-reverse-lookup" {
		cmd.apiHelper = apihelper.New(cli)
		cmd.ServiceReverseLookupCommand(args)
	}
}

func main() {
	plugin.Start(new(ServiceReverseLookupCmd))
}
