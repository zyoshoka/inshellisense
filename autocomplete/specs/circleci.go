// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["circleci"] = model.Subcommand{
		Name:        []string{"circleci"},
		Description: `CircleCI CLI`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Show help for CircleCI`,
			IsPersistent: true,
		}, {
			Name:         []string{"--skip-update-check"},
			Description:  `Skip update check before every command`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"completion"},
			Description: `Generate shell completion scripts`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate bash completion scripts`,
			}, {
				Name:        []string{"zsh"},
				Description: `Generate zsh completion scripts`,
			}},
		}, {
			Name:        []string{"config"},
			Description: `Operate on build config files`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"pack"},
				Description: `Pack CircleCI config files into a single file`,
			}, {
				Name:        []string{"process"},
				Description: `Validate and display extended config`,
			}, {
				Name:        []string{"validate"},
				Description: `Checks that config is valid`,
			}},
		}, {
			Name:        []string{"context"},
			Description: `Secure and share environment variables across projects`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new context`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
					IsOptional:  true,
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
					IsOptional:  true,
				}, {
					Name:        "context-name",
					Description: `The name for your context`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete the named context`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
				}, {
					Name:        "context-name",
					Description: `The name for your context`,
				}},
			}, {
				Name:        []string{"list"},
				Description: `List all contexts`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
				}},
			}, {
				Name:        []string{"remove-secret"},
				Description: `Remove environment variable from a context`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
				}, {
					Name:        "context-name",
					Description: `The name for your context`,
				}, {
					Name:        "secret name",
					Description: `The name of the env variable to remove`,
				}},
			}, {
				Name:        []string{"show"},
				Description: `Show a context`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
				}, {
					Name:        "context-name",
					Description: `The name for your context`,
				}},
			}, {
				Name:        []string{"store-secret"},
				Description: `Store environment variables`,
				Args: []model.Arg{{
					Name:        "vcs-type",
					Description: `Your VCS provider, can be either 'github' or 'bitbucket'`,
					Suggestions: []model.Suggestion{{Name: []string{`github`}}, {Name: []string{`bitbucket`}}},
				}, {
					Name:        "org-name",
					Description: `The name of your organization`,
				}, {
					Name:        "context-name",
					Description: `The name for your context`,
				}, {
					Name:        "secret name",
					Description: `The name of the env variable to store`,
				}},
			}},
		}, {
			Name:        []string{"diagnostic"},
			Description: `Check the status of your CircleCI CLI`,
		}, {
			Name:        []string{"follow"},
			Description: `Attempt to follow the project for the current git repo`,
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
		}, {
			Name:        []string{"local"},
			Description: `Debug jobs on the local machine`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"execute"},
				Description: `Run a job in a container on the local machine`,
			}},
		}, {
			Name:        []string{"namespace"},
			Description: `Operate on namespaces`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a namespace`,
			}},
		}, {
			Name:        []string{"open"},
			Description: `Open the current project in the browser`,
		}, {
			Name:        []string{"orb"},
			Description: `Operate on orbs`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add-to-category"},
				Description: `Add an orb to a category`,
				Args: []model.Arg{{
					Name:        "namespace/orb",
					Description: `The namespace and orb to add to a category`,
				}, {
					Name:        "category name",
					Description: `The name of the category to add the orb to, in quotes`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create an orb in a namespace`,
				Args: []model.Arg{{
					Name:        "namespace/orb",
					Description: `Create an orb in the specified namespace`,
				}},
				Options: []model.Option{{
					Name:        []string{"--private"},
					Description: `Specify that this orb is for private use within your org, unlisted from the public registry`,
				}},
			}, {
				Name:        []string{"info"},
				Description: `Show metadata of an orb`,
				Args: []model.Arg{{
					Name:        "orb",
					Description: `The namespace and orb to show metadata for`,
				}},
			}, {
				Name:        []string{"init"},
				Description: `Initialize a new orb`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `The /path/to/myProject-orb`,
				}},
				Options: []model.Option{{
					Name:        []string{"--private"},
					Description: `Specify that this orb is for private use within your org, unlisted from the public registry`,
				}},
			}, {
				Name:        []string{"list"},
				Description: `List orbs`,
				Args: []model.Arg{{
					Name:        "namespace",
					Description: `The namespace used for the orb (i.e. circleci)`,
				}},
				Options: []model.Option{{
					Name:        []string{"--private"},
					Description: `Specify that this orb is for private use within your org, unlisted from the public registry`,
				}, {
					Name:        []string{"--sort"},
					Description: `Specify the sorting`,
					Args: []model.Arg{{
						Suggestions: []model.Suggestion{{Name: []string{`builds`}}, {Name: []string{`projects`}}, {Name: []string{`orgs`}}},
					}},
				}, {
					Name:        []string{"-u", "--uncertified"},
					Description: `Include uncertified orbs`,
				}},
			}, {
				Name:        []string{"list-categories"},
				Description: `List orb categories`,
			}, {
				Name:        []string{"pack"},
				Description: `Pack an orb with local scripts`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `The /path/to/myProject-orb`,
				}},
			}, {
				Name:        []string{"process"},
				Description: `Validate an orb and print its form after all pre-registration processing is complete`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `The path to your orb (use '-' for STDIN)`,
				}},
			}, {
				Name:        []string{"publish"},
				Description: `Publish an orb to the registry`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `The /path/to/myProject-orb`,
				}, {
					Name:        "orb",
					Description: `A fully-qualified reference to an orb, i.e. namespace/orb@version`,
				}},
			}, {
				Name:        []string{"remove-from-category"},
				Description: `Remove an orb from a category`,
				Args: []model.Arg{{
					Name:        "namespace/orb",
					Description: `The namespace and orb to add to a category`,
				}, {
					Name:        "category name",
					Description: `The name of the category to add the orb to, in quotes`,
				}},
			}, {
				Name:        []string{"source"},
				Description: `Show source code of an orb`,
				Args: []model.Arg{{
					Name:        "orb",
					Description: `A fully-qualified reference to an orb, i.e. namespace/orb@version`,
				}},
			}, {
				Name:        []string{"unlist"},
				Description: `Disable/enable an orb's listing in the registry`,
				Args: []model.Arg{{
					Name:        "namespace/orb",
					Description: `The namespace and orb to unlist/list from the registry`,
				}, {
					Name:        "condition",
					Description: `Use either true|false`,
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"validate"},
				Description: `Validate an orb.yml`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `The /path/to/myProject-orb`,
				}},
			}},
		}, {
			Name:        []string{"policy"},
			Description: `Manage security policies`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"decide"},
				Description: `Make a decision`,
				Args: []model.Arg{{
					Name:        "path",
					Description: `Policy file or directory path`,
				}},
				Options: []model.Option{{
					Name:        []string{"--input"},
					Description: `Path to input file, i.e. ./.circleci/config.yml`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--metafile"},
					Description: `Path to decision metadata file`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--strict"},
					Description: `Return non-zero status code for decision resulting in HARD_FAIL`,
				}, {
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"diff"},
				Description: `Get diff between local and remote policy bundles`,
				Args: []model.Arg{{
					Name:        "policy_dir_path",
					Description: `Policy file or directory path`,
				}},
				Options: []model.Option{{
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"eval"},
				Description: `Perform raw opa evaluation locally`,
				Args: []model.Arg{{
					Name:        "policy_dir_path",
					Description: `Policy file or directory path`,
				}},
				Options: []model.Option{{
					Name:        []string{"--input"},
					Description: `Path to input file, i.e. ./.circleci/config.yml`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--metafile"},
					Description: `Path to decision metadata file`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--query"},
					Description: `Policy decision query, default is 'data'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"fetch"},
				Description: `Fetch policy bundle (or a single policy)`,
				Args: []model.Arg{{
					Name:        "policy name",
					Description: `Name of policy to fetch`,
				}},
				Options: []model.Option{{
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
				}},
			}, {
				Name:        []string{"logs"},
				Description: `Get policy decision logs / decision log (or policy bundle) by decision ID`,
				Args: []model.Arg{{
					Name:        "decision ID",
					Description: `Decision ID to get logs for`,
				}},
				Options: []model.Option{{
					Name:        []string{"--after"},
					Description: `Filter decision logs triggered AFTER this datetime`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--before"},
					Description: `Filter decision logs triggered BEFORE this datetime`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--branch"},
					Description: `Filter decision logs based on branch name`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--out"},
					Description: `Specify output file name`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--policy-bundle"},
					Description: `Get only the policy bundle for given decisionID`,
				}, {
					Name:        []string{"--project-id"},
					Description: `Filter decision logs based on project-id`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--status"},
					Description: `Filter decision logs based on their status`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"push"},
				Description: `Push policy bundle (or a single policy)`,
				Args: []model.Arg{{
					Name:        "policy_dir_path",
					Description: `Policy file or directory path`,
				}},
				Options: []model.Option{{
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--no-prompt"},
					Description: `Removes the prompt`,
				}},
			}, {
				Name:        []string{"settings"},
				Description: `Get/set policy decision settings (To read settings: run command without any settings flags)`,
				Options: []model.Option{{
					Name:        []string{"--context"},
					Description: `Policy context for decision, default is 'config'`,
				}, {
					Name:        []string{"--enabled"},
					Description: `Enable/disable policy decision evaluation in build pipeline`,
					Args: []model.Arg{{
						Name:        "boolean",
						Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
						IsOptional:  true,
					}},
				}, {
					Name:        []string{"--owner-id"},
					Description: `The id of the policy's owner`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}},
		}, {
			Name:        []string{"runner"},
			Description: `Operate on runners`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"instance"},
				Description: `Operate on runner instances`,
			}, {
				Name:        []string{"resource-class"},
				Description: `Operate on runner resource-classes`,
			}, {
				Name:        []string{"token"},
				Description: `Operate on runner tokens`,
			}},
		}, {
			Name:        []string{"setup"},
			Description: `Setup CLI with your credentials`,
		}, {
			Name:        []string{"update"},
			Description: `Update and switch to new CLI version`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"check"},
				Description: `Check for new CLI version`,
			}, {
				Name:        []string{"install"},
				Description: `Install new CLI version`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Display CircleCI CLI version`,
		}},
	}
}
