// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["jmeter"] = model.Subcommand{
		Name:        []string{"jmeter"},
		Description: `Apache JMeter - 100% Java Load Testing Tool`,
		Options: []model.Option{{
			Name:        []string{"-v", "--version"},
			Description: `Print the JMeter version information and exit`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print usage information and exit`,
		}, {
			Name:        []string{"-p", "--propfile"},
			Description: `The jmeter property file to use`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "property",
				Description: `Your jmeter property file`,
			}},
		}, {
			Name:        []string{"-q", "--addprop"},
			Description: `Additional JMeter property file(s)`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "files...",
				Description: `Additional JMeter property file(s)`,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"-t", "--testfile"},
			Description: `The JMeter test(.jmx) file to run. "-t LAST" will load last used file`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "testfile",
				Description: `The JMeter test(.jmx) file to run. "-t LAST" will load last used file`,
			}},
		}, {
			Name:        []string{"-l", "--logfile"},
			Description: `The file to log samples to`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "logfile",
				Description: `The file to log samples to`,
			}},
		}, {
			Name:        []string{"-i", "--jmeterlogconf"},
			Description: `JMeter logging configuration file`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "jmeterlogconf",
				Description: `Jmeter logging configuration file`,
			}},
		}, {
			Name:        []string{"-j", "--jmeterlogfile"},
			Description: `JMeter run log file`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "jmeterlogfile",
				Description: `JMeter run log file`,
			}},
		}, {
			Name:        []string{"-n", "--nongui"},
			Description: `Run JMeter in nongui mode`,
		}, {
			Name:        []string{"-s", "--server"},
			Description: `Run the JMeter server`,
		}, {
			Name:        []string{"-E", "--proxyScheme"},
			Description: `Set a proxy scheme to use for the proxy server`,
			Args: []model.Arg{{
				Name:        "proxyScheme",
				Description: `Set a proxy scheme to use for the proxy server`,
			}},
		}, {
			Name:        []string{"-H", "--proxyHost"},
			Description: `Set a proxy server for JMeter to use`,
			Args: []model.Arg{{
				Name: "server",
			}},
		}, {
			Name:        []string{"-P", "--proxyPort"},
			Description: `Set proxy server port for JMeter to use`,
			Args: []model.Arg{{
				Name: "port",
			}},
		}, {
			Name:        []string{"-N", "--nonProxyHosts"},
			Description: `Set nonproxy host list (e.g. *.apache.org|localhost)`,
			Args: []model.Arg{{
				Name:        "nonProxyHosts",
				Description: `Set nonproxy host list (e.g. *.apache.org|localhost)`,
			}},
		}, {
			Name:        []string{"-u", "--username"},
			Description: `Set username for proxy server that JMeter is to use`,
			Args: []model.Arg{{
				Name:        "username",
				Description: `Set username for proxy server that JMeter is to use`,
			}},
		}, {
			Name:        []string{"-a", "--password"},
			Description: `Set password for proxy server that JMeter is to use`,
			Args: []model.Arg{{
				Name:        "password",
				Description: `Set password for proxy server that JMeter is to use`,
			}},
		}, {
			Name:        []string{"-J", "--jmeterproperty"},
			Description: `Define additional JMeter properties <argument>=<value>`,
			Args: []model.Arg{{
				Name:        "jmeterproperty",
				Description: `Define additional JMeter properties <argument>=<value>`,
			}},
		}, {
			Name:        []string{"-G", "--globalproperty"},
			Description: `Define Global properties (sent to servers) e.g. -Gport=123 or -Gglobal.properties`,
			Args: []model.Arg{{
				Name:        "globalproperty",
				Description: `Define Global properties (sent to servers) e.g. -Gport=123 or -Gglobal.properties`,
			}},
		}, {
			Name:        []string{"-D", "--systemproperty"},
			Description: `Define additional system properties <argument>=<value>`,
			Args: []model.Arg{{
				Name:        "systemproperty",
				Description: `Define additional system properties <argument>=<value>`,
			}},
		}, {
			Name:        []string{"-S", "--systemPropertyFile"},
			Description: `Additional system property file(s)`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "systemPropertyFile",
				Description: `Additional system property file(s)`,
			}},
		}, {
			Name:        []string{"-f", "--forceDeleteResultFile"},
			Description: `Force delete existing results files and web report folder`,
		}, {
			Name:        []string{"-L", "--loglevel"},
			Description: `[category=]level e.g. jorphan=INFO, jmeter.util=DEBUG or com.example.foo=WARN`,
			Args: []model.Arg{{
				Name:        "loglevel",
				Suggestions: []model.Suggestion{{Name: []string{`OFF`}}, {Name: []string{`FATAL`}}, {Name: []string{`WARN`}}, {Name: []string{`INFO`}}, {Name: []string{`DEBUG`}}, {Name: []string{`TRACE`}}, {Name: []string{`ALL`}}},
			}},
		}, {
			Name:        []string{"-r", "--runremote"},
			Description: `Start remote servers (as defined in remote_hosts)`,
		}, {
			Name:        []string{"-R", "--remotestart"},
			Description: `Start these remote servers (overrides remote_hosts)`,
			Args: []model.Arg{{
				Name:       "servers",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"-d", "--homedir"},
			Description: `The JMeter home directory to use`,
			Args: []model.Arg{{
				Name: "homedir",
			}},
		}, {
			Name:        []string{"-X", "--remoteexit"},
			Description: `Exit the remote servers at end  of test (non-GUI)`,
		}, {
			Name:        []string{"-g", "--reportonly"},
			Description: `Generate report dashboard only, from a test results file`,
		}, {
			Name:        []string{"-e", "--reportatendofloadtests"},
			Description: `Generate report dashboard after load test`,
		}, {
			Name:        []string{"-o", "--reportoutputfolder"},
			Description: `Output folder for report dashboard`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "reportoutputfolder",
				Description: `Output folder for report dashboard`,
			}},
		}},
	}
}
