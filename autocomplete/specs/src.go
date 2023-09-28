// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["src"] = model.Subcommand{
		Name:        []string{"src"},
		Description: `Interact with Sourcegraph from the command line`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for src`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"search"},
			Description: `Run a code search`,
			Options: []model.Option{{
				Name:        []string{"-display"},
				Description: `Limit the number of results that are displayed. Only supported together with stream flag. Statistics continue to report all results`,
			}, {
				Name:        []string{"-dump-requests"},
				Description: `Log GraphQL requests and responses to stdout`,
			}, {
				Name:        []string{"-explain-json"},
				Description: `Explain the JSON output schema and exit`,
			}, {
				Name:        []string{"-get-curl"},
				Description: `Print the curl command for executing this query and exit. (WARNING: Includes printing your access token!)`,
			}, {
				Name:        []string{"-insecure-skip-verify"},
				Description: `Skip validation of TLD certificates against trusted chains`,
			}, {
				Name:        []string{"-json"},
				Description: `Whether or not to output results as JSON`,
			}, {
				Name:        []string{"-less"},
				Description: `Pipe output to "less -R" (only if stdout is terminal, and not json flag)`,
			}, {
				Name:        []string{"-stream"},
				Description: `Consume results as stream. Streaming search only supports a subset of flags and parameters: trace, insecure-skip-verify, display, json`,
			}, {
				Name:        []string{"-trace"},
				Description: `Log the trace ID for requests`,
			}, {
				Name:        []string{"-user-agent-telemetry"},
				Description: `Include the operating system and architecture in the User-Agent sent with requests to Sourcegraph`,
			}},
		}, {
			Name:        []string{"api"},
			Description: `Sourcegraph API Access`,
			Options: []model.Option{{
				Name:        []string{"-get-curl"},
				Description: `Print the curl command for executing this query and exit (WARNING: includes printing your access token!)`,
			}, {
				Name:        []string{"-dump-requests"},
				Description: `Log GraphQL requests and responses to stdout`,
			}, {
				Name:        []string{"-insecure-skip-verify"},
				Description: `Skip validation of TLS certificates against trusted chains`,
			}, {
				Name:        []string{"-query"},
				Description: `GraphQL query to execute, e.g. ‘query { currentUser { username } }’ (stdin otherwise)`,
			}, {
				Name:        []string{"-trace"},
				Description: `Log the trace ID for requests`,
			}, {
				Name:        []string{"-user-agent-telemetry"},
				Description: `Include the operating system and architecture in the User-Agent sent with requests to Sourcegraph`,
			}, {
				Name:        []string{"-vars"},
				Description: `GraphQL query variables to include as JSON string, e.g. ‘{“var”: “val”, “var2”: “val2”}’`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `'src login' helps you authenticate 'src' to access a Sourcegraph instance with your user credentials`,
			Options: []model.Option{{
				Name:        []string{"-dump-requests"},
				Description: `Log GraphQL requests and responses to stdout`,
			}, {
				Name:        []string{"-get-curl"},
				Description: `Print the curl command for executing this query and exit (WARNING: includes printing your access token!)`,
			}, {
				Name:        []string{"-insecure-skip-verify"},
				Description: `Skip validation of TLS certificates against trusted chains`,
			}, {
				Name:        []string{"-trace"},
				Description: `Log the trace ID for requests`,
			}, {
				Name:        []string{"-user-agent-telemetry"},
				Description: `Include the operating system and architecture in the User-Agent sent with requests to Sourcegraph`,
			}},
		}, {
			Name:        []string{"batch"},
			Description: `'Batch gives you a declarative structure for finding and modifying code across all of your repositories`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"apply"},
				Description: `Apply batch`,
			}, {
				Name:        []string{"exec"},
				Description: `Execute batch`,
			}, {
				Name:        []string{"new"},
				Description: `New batch`,
			}, {
				Name:        []string{"preview"},
				Description: `Preview batch`,
			}, {
				Name:        []string{"repositories"},
				Description: `Repositories to batch`,
			}, {
				Name:        []string{"validate"},
				Description: `Validate batch`,
			}},
		}, {
			Name:        []string{"config"},
			Description: `'src config' helps you configure 'src'`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"edit"},
				Description: `Edit config`,
			}, {
				Name:        []string{"get"},
				Description: `Get configs`,
			}, {
				Name:        []string{"list"},
				Description: `List configs`,
			}},
		}, {
			Name:        []string{"extsvc"},
			Description: `Edit or view external service configuration on the Sourcegraph instance`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"edit"},
				Description: `Edit external service configuration`,
			}, {
				Name:        []string{"list"},
				Description: `List external service configurations`,
			}},
		}, {
			Name:        []string{"lsif"},
			Description: `Upload an LSIF dumps`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"upload"},
				Description: `Upload LSIF dump`,
			}},
		}, {
			Name:        []string{"orgs"},
			Description: `Create, edit, view, or delete organizations and members`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create an organization`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete an organization`,
			}, {
				Name:        []string{"get"},
				Description: `Get an organization`,
			}, {
				Name:        []string{"list"},
				Description: `List organizations`,
			}, {
				Name:        []string{"members"},
				Description: `List organization members`,
			}},
		}, {
			Name:        []string{"repos"},
			Description: `View or delete repositories`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"delete"},
				Description: `Delete repository`,
			}, {
				Name:        []string{"get"},
				Description: `Get repository`,
			}, {
				Name:        []string{"list"},
				Description: `List repositories`,
			}},
		}, {
			Name:        []string{"serve-git"},
			Description: `By default 'src serve-git' will recursively serve your current directory on the address ':3434'`,
			Options: []model.Option{{
				Name:        []string{"-addr"},
				Description: `Address on which to server (end with : for unused port)`,
			}, {
				Name:        []string{"-list"},
				Description: `List found repository names`,
			}},
		}, {
			Name:        []string{"users"},
			Description: `Create, edit, view, tag, or delete users`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create user`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete user`,
			}, {
				Name:        []string{"get"},
				Description: `Get user`,
			}, {
				Name:        []string{"list"},
				Description: `List users`,
			}, {
				Name:        []string{"tag"},
				Description: `Tag user`,
			}},
		}, {
			Name:        []string{"validate"},
			Description: `EXPERIMENTAL: Instance validation provides a quick way to check that a Sourcegraph instance functions properly after a fresh install or an update`,
			Options: []model.Option{{
				Name:        []string{"-context"},
				Description: `Comma-separated list of key=value pairs to add to the script execution context`,
			}, {
				Name:        []string{"-dump-requests"},
				Description: `Log GraphQL requests and responses to stdout`,
			}, {
				Name:        []string{"-get-curl"},
				Description: `Print the curl command for executing this query and exit (WARNING: includes printing your access token!)`,
			}, {
				Name:        []string{"-insecure-skip-verify"},
				Description: `Skip validation of TLS certificates against trusted chains`,
			}, {
				Name:        []string{"-secrets"},
				Description: `Path to a file containing key=value lines. The key value pairs will be added to the script context`,
			}, {
				Name:        []string{"-trace"},
				Description: `Log the trace ID for requests`,
			}, {
				Name:        []string{"-user-agent-telemetry"},
				Description: `Include the operating system and architecture in the User-Agent sent with requests to Sourcegraph`,
			}},
		}},
	}
}
