// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["steadybit"] = model.Subcommand{
		Name:        []string{"steadybit"},
		Description: `Command-line interface to interact with the Steadybit API`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Display usage information`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"config"},
			Description: `Show/modify the CLI configuration and authentication profiles`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"show"},
				Description: `Show the active CLI configuration. Warning: Prints secrets!`,
			}, {
				Name:        []string{"profile"},
				Description: `Configure authentication profiles`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"add"},
					Description: `Interactively configure a new profile`,
				}, {
					Name:        []string{"list", "ls"},
					Description: `List all configured profiles`,
				}, {
					Name:        []string{"remove"},
					Description: `Interactively remove an existing profile`,
				}, {
					Name:        []string{"select"},
					Description: `Interactively change the currently active profile`,
				}},
			}},
		}, {
			Name:        []string{"def-repo"},
			Description: `Change versions and verify a task/policy definition repository state`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"set-version"},
				Description: `Set the versions in policies and task references`,
				Options: []model.Option{{
					Name:        []string{"-v", "--version"},
					Description: `Version to set`,
					Args: []model.Arg{{
						Name: "version",
					}},
				}, {
					Name:        []string{"-d", "--directory"},
					Description: `The directory to search for task and policy files`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFolders},
						Name:      "directory",
					}},
				}},
			}, {
				Name:        []string{"check"},
				Description: `Checks that the tasks and policies are valid`,
				Options: []model.Option{{
					Name:        []string{"-v", "--version"},
					Description: `The version to check`,
					Args: []model.Arg{{
						Name: "version",
					}},
				}, {
					Name:        []string{"-d", "--directory"},
					Description: `The directory to search for task and policy files`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFolders},
						Name:      "directory",
					}},
				}},
			}},
		}, {
			Name:        []string{"service", "service-definition"},
			Description: `Configure or verify service definitions`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"apply"},
				Description: `Upload a service definition`,
				Options: []model.Option{{
					Name:        []string{"-f", "--file"},
					Description: `Path to the service definition file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "file",
					}},
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete a service definition`,
				Options: []model.Option{{
					Name:        []string{"-f", "--file"},
					Description: `Path to the service definition file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "file",
					}},
				}, {
					Name:        []string{"-i", "--id"},
					Description: `ID of the service definition to delete`,
					Args: []model.Arg{{
						Name: "id",
					}},
				}},
			}, {
				Name:        []string{"init"},
				Description: `Initialize a service definition file`,
			}, {
				Name:        []string{"open"},
				Description: `Open the service in the Steadybit UI`,
				Options: []model.Option{{
					Name:        []string{"-f", "--file"},
					Description: `Path to the service definition file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "file",
					}},
				}},
			}, {
				Name:        []string{"verify"},
				Description: `Read the current service definition and state`,
				Options: []model.Option{{
					Name:        []string{"-f", "--file"},
					Description: `Path to the service definition file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "file",
					}},
				}, {
					Name:        []string{"-pp", "--print-parameters"},
					Description: `Print task parameters when listing tasks`,
				}, {
					Name:        []string{"-pm", "--print-matrix-context"},
					Description: `Print the matrix execution context information when listing tasks`,
				}},
			}, {
				Name:        []string{"show"},
				Description: `Show a list of tasks and policies referenced by this service`,
				Options: []model.Option{{
					Name:        []string{"-f", "--file"},
					Description: `Path to the service definition file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "file",
					}},
				}, {
					Name:        []string{"-n", "--name"},
					Description: `Optional task name to filter the result list`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"-v", "--version"},
					Description: `Optional task version to filter the result list`,
					Args: []model.Arg{{
						Name: "version",
					}},
				}},
			}},
		}},
	}
}
