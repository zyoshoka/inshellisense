// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["vercel"] = model.Subcommand{
		Name:        []string{"vercel"},
		Description: `CLI Interface for Vercel.com`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFolders},
			Name:      "path to project",
		}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Output usage information`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Output the version number`,
		}, {
			Name:        []string{"-V", "--platform-version"},
			Description: `Set the platform version to deploy to`,
		}, {
			Name:        []string{"-A", "--local-config"},
			Description: `Path to the local 'vercel.json' file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"-Q", "--global-config"},
			Description: `Path to the global '.vercel' directory`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
			}},
		}, {
			Name:        []string{"-d", "--debug"},
			Description: `Provides more verbose output`,
		}, {
			Name:        []string{"-f", "--force"},
			Description: `Force a new deployment even if nothing has changed`,
		}, {
			Name:        []string{"-with-cache"},
			Description: `Retain build cache when using --force`,
		}, {
			Name:        []string{"-t", "--token"},
			Description: `Execute command with an auth token`,
			Args: []model.Arg{{
				Name:        "auth token",
				Description: `Auth token to add to your command`,
			}},
		}, {
			Name:        []string{"-p", "--public"},
			Description: `Deployment is public ('/_src' is exposed)`,
		}, {
			Name:        []string{"-e", "--env"},
			Description: `Include an env var during run time (e.g.: '-e KEY=value'). Can appear many times`,
		}, {
			Name:        []string{"-b", "--build-env"},
			Description: `Similar to "--env" but for build time only`,
		}, {
			Name:        []string{"-m", "--meta"},
			Description: `Add metadata for the deployment (e.g.: "-m KEY=value"). Can appear many times`,
		}, {
			Name:        []string{"-C", "--no-clipboard"},
			Description: `Do not attempt to copy URL to clipboard`,
		}, {
			Name:        []string{"-S", "--scope"},
			Description: `Set a custom scope`,
			Args: []model.Arg{{
				Name:        "team name",
				Description: `Team to execute commands from`,
			}},
		}, {
			Name:        []string{"--regions"},
			Description: `Set default regions to enable the deployment on`,
		}, {
			Name:        []string{"--prod"},
			Description: `Create a production deployment`,
		}, {
			Name:        []string{"-c", "--confirm"},
			Description: `Confirm default options and skip questions`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"deploy"},
			Description: `Performs a deployment (default)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
			}},
		}, {
			Name:        []string{"dev"},
			Description: `Start a local development server`,
			Options: []model.Option{{
				Name:        []string{"--listen"},
				Description: `Specifies which port to run on`,
			}},
		}, {
			Name:        []string{"env"},
			Description: `Manages the Environment Variables for your current Project`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add"},
				Description: `Add an environment variable`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the env variable to add`,
				}, {
					Name:        "environment",
					Description: `Environment to add the variable to`,
					Suggestions: []model.Suggestion{{
						Name: []string{`production`},
					}, {
						Name: []string{`preview`},
					}, {
						Name: []string{`development`},
					}},
				}},
			}, {
				Name:        []string{"rm"},
				Description: `Remove an environment variable`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the variable to remove`,
				}, {
					Name:        "environment",
					Description: `Environment to remove from`,
					Suggestions: []model.Suggestion{{
						Name: []string{`production`},
					}, {
						Name: []string{`preview`},
					}, {
						Name: []string{`development`},
					}},
				}},
			}, {
				Name:        []string{"pull"},
				Description: `Download dev env variables from cloud and write to .env`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "file",
					Description: `The file to write downloaded variables to`,
				}},
			}, {
				Name:        []string{"ls"},
				Description: `List environment variables for a specific environment`,
				Args: []model.Arg{{
					Name:        "environment",
					Description: `Environment to list variables for`,
				}},
			}},
		}, {
			Name:        []string{"init"},
			Description: `Initialize an example project`,
			Args: []model.Arg{{
				Name:        "project name",
				Description: `Project name to initialize locally`,
			}, {
				Name:        "new project name",
				Description: `Initialize specific project locally and rename directory`,
			}},
		}, {
			Name:        []string{"list"},
			Description: `Lists deployments`,
			Args: []model.Arg{{
				Name:        "project name",
				Description: `View deployments for specified project`,
			}},
			Options: []model.Option{{
				Name:        []string{"-m", "--meta"},
				Description: `Filters results based on project metadata`,
				Args: []model.Arg{{
					IsVariadic: true,
				}},
			}},
		}, {
			Name:        []string{"ls"},
			Description: `Lists deployments`,
			Args: []model.Arg{{
				Name:        "project name",
				Description: `View deployments for specified project`,
			}},
		}, {
			Name:        []string{"inspect"},
			Description: `Displays information related to a deployment`,
			Args: []model.Arg{{
				Name:        "url",
				Description: `The URL of the deployment to inspect`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `Logs into your account or creates a new one`,
		}, {
			Name:        []string{"logout"},
			Description: `Logs out of your account`,
		}, {
			Name:        []string{"switch"},
			Description: `Switches between teams and your personal account`,
			Args: []model.Arg{{
				Name:        "team name",
				Description: `Team to switch to`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Displays complete help for [cmd]`,
			Args: []model.Arg{{
				Name:        "command",
				Description: `Command to detailed information about`,
			}},
		}, {
			Name:        []string{"rm"},
			Description: `Removes a deployment`,
			Args: []model.Arg{{
				Name:        "deployment url",
				Description: `URL of the deployment to remove`,
			}},
			Options: []model.Option{{
				Name:        []string{"-s", "--safe"},
				Description: `Skip removal of deployments with an active preview URL or production domain`,
			}, {
				Name:        []string{"-y", "--yes"},
				Description: `Skip confirmation step for a deployment or project removal`,
			}},
		}, {
			Name:        []string{"remove"},
			Description: `Removes a deployment`,
			Args: []model.Arg{{
				Name:        "deployment url",
				Description: `URL of the deployment to remove`,
			}},
		}, {
			Name:        []string{"domains"},
			Description: `Manages your domain names`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"ls"},
				Description: `List all domains under an account`,
			}, {
				Name:        []string{"inspect"},
				Description: `Retrieves information about a domain`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to inspect`,
				}},
			}, {
				Name:        []string{"add"},
				Description: `Add a domain to an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to add`,
				}},
			}, {
				Name:        []string{"rm"},
				Description: `Removes a domain from an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to remove`,
				}},
			}, {
				Name:        []string{"buy"},
				Description: `Buy a domain for an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to buy`,
				}},
			}, {
				Name:        []string{"move"},
				Description: `Removes a domain from an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to move`,
				}, {
					Name:        "account name",
					Description: `Account to move the domain to`,
				}},
			}, {
				Name:        []string{"transfer-in"},
				Description: `Transfers in a domain to an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to transfer in`,
				}},
			}, {
				Name:        []string{"verify"},
				Description: `Verifies a domain for an account`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to verify`,
				}},
			}},
		}, {
			Name:        []string{"dns"},
			Description: `Manages your DNS records`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add"},
				Description: `Add DNS record for a domain`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `Domain to add record to`,
				}, {
					Name:        "subdomain",
					Description: `Subdomain to add record to`,
				}, {
					Name:        "record type",
					Description: `Type of record to add`,
					Suggestions: []model.Suggestion{{
						Name:        []string{`A`},
						Description: `A record`,
					}, {
						Name:        []string{`AAAA`},
						Description: `AAAA record`,
					}, {
						Name:        []string{`ALIAS`},
						Description: `ALIAS record`,
					}, {
						Name:        []string{`CNAME`},
						Description: `CName record`,
					}, {
						Name:        []string{`TXT`},
						Description: `TXT record`,
					}},
				}, {
					Name:        "value",
					Description: `Record value`,
				}},
			}},
		}, {
			Name:        []string{"certs"},
			Description: `Manages your SSL certificates`,
			Options: []model.Option{{
				Name:        []string{"--challenge-only"},
				Description: `Show challenges needed to issue a certificate`,
			}, {
				Name:        []string{"--crt"},
				Description: `Include path to .crt`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
				}},
			}, {
				Name:        []string{"--key"},
				Description: `Include path to .key`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
				}},
			}, {
				Name:        []string{"--ca"},
				Description: `Include path to .ca`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"ls"},
				Description: `List all certificates under an account`,
			}, {
				Name:        []string{"issue"},
				Description: `Issue certificates for multiple domains`,
				Args: []model.Arg{{
					Name:        "Domains",
					Description: `List of domains separated by commas to issue certificates for`,
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"rm"},
				Description: `Remove a certificate by id`,
				Args: []model.Arg{{
					Name:        "certificate id",
					Description: `Id of the certificate to remove`,
				}},
			}},
		}, {
			Name:        []string{"secrets"},
			Description: `Manages your global Secrets, for use in Environment Variables`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List all secrets`,
			}, {
				Name:        []string{"add"},
				Description: `Add a new secret`,
				Args: []model.Arg{{
					Name:        "secret name",
					Description: `Name of the secret to add`,
				}, {
					Name:        "secret value",
					Description: `Value of the secret to add`,
				}},
			}, {
				Name:        []string{"rename"},
				Description: `Rename a secret`,
				Args: []model.Arg{{
					Name:        "old name",
					Description: `Old name of the secret to rename`,
				}, {
					Name:        "new name",
					Description: `New name of the secret`,
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove a secret`,
				Args: []model.Arg{{
					Name:        "secret name",
					Description: `Name of the seret to remove`,
				}},
			}},
		}, {
			Name:        []string{"logs"},
			Description: `Displays the logs for a deployment`,
			Args: []model.Arg{{
				Name:        "deployment url",
				Description: `Get logs for specified deployment`,
			}},
			Options: []model.Option{{
				Name:        []string{"-a", "--all"},
				Description: `Receive access logs in addition to regular logs`,
			}, {
				Name:        []string{"-f", "--follow"},
				Description: `Watch for additional logs output`,
			}, {
				Name:        []string{"-n", "--number"},
				Description: `Specify number of log lines to output`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"-o", "--output"},
				Description: `Specifies format of logs output as 'short' or 'raw'`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--since"},
				Description: `Return logs after a specific ISO 8601 date`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"-q", "--query"},
				Description: `Return logs against a search query`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--until"},
				Description: `Return logs up to a specific ISO 8601 date`,
				Args:        []model.Arg{{}},
			}},
		}, {
			Name:        []string{"teams"},
			Description: `Manages your teams`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List all teams under an account`,
			}, {
				Name:        []string{"add"},
				Description: `Create a new team`,
			}, {
				Name:        []string{"invite"},
				Description: `Invite a new member to the team`,
				Args: []model.Arg{{
					Name:        "email",
					Description: `Email of member to invite to team`,
				}},
			}},
		}, {
			Name:        []string{"whoami"},
			Description: `Shows the username of the currently logged in user`,
		}, {
			Name:        []string{"alias"},
			Description: `Apply custom domains based on git branches, or other heuristics`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"set"},
				Description: `Assign a custom domain to a deployment`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `The URL of the deployment to assign to a domain`,
				}, {
					Name:        "domain",
					Description: `The domain you want to assign to`,
				}},
			}, {
				Name:        []string{"rm"},
				Description: `Remove a custom domain from a deployment`,
				Args: []model.Arg{{
					Name:        "domain",
					Description: `The domain to remove`,
				}},
			}, {
				Name:        []string{"ls"},
				Description: `List custom domains that were assigned to deployments`,
			}},
		}, {
			Name:        []string{"link"},
			Description: `Links your local directory to a Project`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFolders},
				Name:       "directory",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"billing"},
			Description: `Manage payment methods`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"ls"},
				Description: `List all the payment methods of the active account`,
			}, {
				Name:        []string{"add"},
				Description: `Interactively add a new credit card`,
			}, {
				Name:        []string{"rm"},
				Description: `Remove a credit card by ID`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `The id of the card to remove`,
				}},
			}, {
				Name:        []string{"set-default"},
				Description: `Select which credit card should be default`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `The id of the card to set default`,
				}},
			}},
		}},
	}
}
