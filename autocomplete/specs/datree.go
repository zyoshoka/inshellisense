// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["datree"] = model.Subcommand{
		Name:        []string{"datree"},
		Description: `Datree can be used on the command line to run policies against Kubernetes manifests YAML files or Helm charts`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for datree`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"completion"},
			Description: `Generate completion script for bash,zsh,fish,powershell`,
		}, {
			Name:        []string{"config"},
			Description: `Internal configuration management for datree config file`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Help for config`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"get"},
				Description: `Get configuration value`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Get value for specific key from datree config.yaml file. Defaults to $HOME/.datree/config.yaml`,
				}},
			}, {
				Name:        []string{"set"},
				Description: `Set configuration value`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Apply value for specific key in datree config.yaml file. Defaults to $HOME/.datree/config.yaml`,
				}},
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateHelp},
			}},
		}, {
			Name:        []string{"kustomize"},
			Description: `Generate kustomization files from manifests`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Help for kustomize`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"test"},
				Description: `Test kustomization files`,
				Options: []model.Option{{
					Name:        []string{"--help", "-h"},
					Description: `Help for 'test'`,
				}, {
					Name:        []string{"--ignore-missing-schemas"},
					Description: `Ignore missing schemas when executing schema validation step`,
				}, {
					Name:        []string{"--no-record"},
					Description: `Do not send policy checks metadata to the backend`,
				}, {
					Name:        []string{"--only-k8s-files"},
					Description: `Evaluate only valid yaml files with the properties 'apiVersion' and 'kind'. Ignore everything else`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Define output format (simple, yaml, json, xml, JUnit)`,
					Args: []model.Arg{{
						Name:        "output",
						Description: `Output format`,
						Suggestions: []model.Suggestion{{
							Name:        []string{`simple`},
							Description: `Simple output without colors`,
						}, {
							Name:        []string{`yaml`},
							Description: `YAML output`,
						}, {
							Name:        []string{`json`},
							Description: `JSON output`,
						}, {
							Name:        []string{`xml`},
							Description: `XML output`,
						}, {
							Name:        []string{`junit`},
							Description: `JUnit output`,
						}},
					}},
				}, {
					Name:        []string{"-p", "--policy"},
					Description: `Policy name to run against`,
					Args: []model.Arg{{
						Name:        "policy",
						Description: `String`,
					}},
				}, {
					Name:        []string{"--policy-config"},
					Description: `Path for local policies configuration file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
					}},
				}, {
					Name:        []string{"--schema-location"},
					Description: `Override schemas location search path (can be specified multiple times)`,
					Args: []model.Arg{{
						Name:        "stringArray",
						Description: `String array`,
					}},
				}, {
					Name:        []string{"-s", "--schema-version"},
					Description: `Set kubernetes version to validate against. Defaults to 1.19.0`,
				}, {
					Name:        []string{"--verbose"},
					Description: `Display 'How to Fix' link`,
				}},
			}},
		}, {
			Name:        []string{"publish"},
			Description: `Publish policies configuration for given <fileName>. Input should be the path to the Policy-as-Code yaml configuration file. ## Note You need to first enable Policy-as-Code (PaC) on the settings page in the dashboard`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file path",
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Help for publish`,
			}},
		}, {
			Name:        []string{"test"},
			Description: `Trigger a policy check, provide a Kubernetes configuration file path or a glob pattern`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file path",
			}},
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Help for 'test'`,
			}, {
				Name:        []string{"--ignore-missing-schemas"},
				Description: `Ignore missing schemas when executing schema validation step`,
			}, {
				Name:        []string{"--no-record"},
				Description: `Do not send policy checks metadata to the backend`,
			}, {
				Name:        []string{"--only-k8s-files"},
				Description: `Evaluate only valid yaml files with the properties 'apiVersion' and 'kind'. Ignore everything else`,
			}, {
				Name:        []string{"-o", "--output"},
				Description: `Define output format (simple, yaml, json, xml, JUnit)`,
				Args: []model.Arg{{
					Name:        "output",
					Description: `Output format`,
					Suggestions: []model.Suggestion{{
						Name:        []string{`simple`},
						Description: `Simple output without colors`,
					}, {
						Name:        []string{`yaml`},
						Description: `YAML output`,
					}, {
						Name:        []string{`json`},
						Description: `JSON output`,
					}, {
						Name:        []string{`xml`},
						Description: `XML output`,
					}, {
						Name:        []string{`junit`},
						Description: `JUnit output`,
					}},
				}},
			}, {
				Name:        []string{"-p", "--policy"},
				Description: `Policy name to run against`,
				Args: []model.Arg{{
					Name:        "policy",
					Description: `String`,
				}},
			}, {
				Name:        []string{"--policy-config"},
				Description: `Path for local policies configuration file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
				}},
			}, {
				Name:        []string{"--schema-location"},
				Description: `Override schemas location search path (can be specified multiple times)`,
				Args: []model.Arg{{
					Name:        "stringArray",
					Description: `String array`,
				}},
			}, {
				Name:        []string{"-s", "--schema-version"},
				Description: `Set kubernetes version to validate against. Defaults to 1.19.0`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Display 'How to Fix' link`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print the version number`,
		}},
	}
}
