package main

import (
	"flag"

	"code.cloudfoundry.org/cli/plugin"
	log "github.com/sirupsen/logrus"
)

// ServiceReverseLookupCmd -
type ServiceReverseLookupCmd struct{}

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
	flagss.StringVar(&serviceGUIDFlag, "serviceGuid", "", "")
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

}

// Run -
func (cmd *ServiceReverseLookupCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "service-reverse-lookup" {
		cmd.ServiceReverseLookupCommand(args)
	}
}

func main() {
	plugin.Start(new(ServiceReverseLookupCmd))
}
