// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["stepzen"] = model.Subcommand{
		Name:        []string{"StepZen"},
		Description: `The StepZen CLI is the primary way to build, deploy and test your schemas on StepZen`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Display help for StepZen`,
			Args: []model.Arg{{
				Name:        "command",
				Description: `Command name for which to display help`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `Login to StepZen`,
		}, {
			Name:        []string{"logout"},
			Description: `Logout of StepZen`,
		}, {
			Name:        []string{"start"},
			Description: `Deploy, watch and develop your API`,
			Options: []model.Option{{
				Name:        []string{"--dir"},
				Description: `The working directory for StepZen assets`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFolders},
					Name:        "path",
					Description: `Path to StepZen directory`,
				}},
			}, {
				Name:        []string{"--endpoint"},
				Description: `The folder/endpoint to deploy to`,
				Args: []model.Arg{{
					Name:        "endpoint",
					Description: `The StepZen endpoint`,
					Generator:   nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--port"},
				Description: `The port number to use for the GraphiQL explorer`,
				Args: []model.Arg{{
					Name:        "port",
					Description: `A port to run on`,
				}},
			}},
		}, {
			Name:        []string{"import"},
			Description: `Import pre-configured schemas to Your API`,
			Args: []model.Arg{{
				Name:        "name",
				Description: `The name of the generator to import`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--dir"},
				Description: `The directory to which the schema will be imported`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFolders},
					Name:        "path",
					Description: `Path to directory`,
				}},
			}},
		}, {
			Name:        []string{"list"},
			Description: `List the assets of a specified type that are linked to the StepZen account`,
			Args: []model.Arg{{
				Name:        "type",
				Description: `The type of asset to list (schemas or configurationsets)`,
				Suggestions: []model.Suggestion{{Name: []string{`schemas`}}, {Name: []string{`configurationsets`}}},
			}},
		}},
	}
}
