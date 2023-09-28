// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["awsume"] = model.Subcommand{
		Name:        []string{"awsume"},
		Description: `Awsume`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateHistory},
			Name:      "profile",
		}},
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for awsume`,
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Display the current version of awsume`,
		}, {
			Name:        []string{"--output-profile", "-o"},
			Description: `A profile to output credentials to`,
			Args: []model.Arg{{
				Name: "output_profile",
			}},
		}, {
			Name:        []string{"--clean"},
			Description: `Clean expired output profiles`,
		}, {
			Name:        []string{"--refresh", "-r"},
			Description: `Force refresh credentials`,
		}, {
			Name:        []string{"--show-commands", "-s"},
			Description: `Show the commands to set the credentials`,
		}, {
			Name:        []string{"--unset", "-u"},
			Description: `Unset your aws environment variables`,
		}, {
			Name:        []string{"--auto-refresh", "-a"},
			Description: `Auto refresh credentials`,
		}, {
			Name:        []string{"--kill-refresher", "-k"},
			Description: `Kill autoawsume`,
		}, {
			Name:        []string{"--list-profiles", "-l"},
			Description: `List profiles, "more" for detail (slow)`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--refresh-autocomplete"},
			Description: `Refresh all plugin autocomplete profiles`,
		}, {
			Name:        []string{"--role-arn"},
			Description: `Role ARN or <partition>:<account_id>:<role_name>`,
			Args: []model.Arg{{
				Name: "role_arn",
			}},
		}, {
			Name:        []string{"--principal-arn"},
			Description: `Principal ARN or <partition>:<account_id>:<provider_name>`,
			Args: []model.Arg{{
				Name: "principal_arn",
			}},
		}, {
			Name:        []string{"--source-profile"},
			Description: `Source_profile to use(role-arn only)`,
			Args: []model.Arg{{
				Name: "source_profile",
			}},
		}, {
			Name:        []string{"--external-id"},
			Description: `External ID to pass to the assume_role`,
			Args: []model.Arg{{
				Name: "external_id",
			}},
		}, {
			Name:        []string{"--mfa-token"},
			Description: `Your mfa token`,
			Args: []model.Arg{{
				Name: "mfa-token",
			}},
		}, {
			Name:        []string{"--region"},
			Description: `The region you want to awsume into`,
			Args: []model.Arg{{
				Name: "region",
			}},
		}, {
			Name:        []string{"--session-name"},
			Description: `Set a custom role session name`,
			Args: []model.Arg{{
				Name: "session_name",
			}},
		}, {
			Name:        []string{"--role-duration"},
			Description: `Seconds to get role creds for`,
			Args: []model.Arg{{
				Name: "role_duration",
			}},
		}, {
			Name:        []string{"--with-saml"},
			Description: `Use saml (requires plugin)`,
		}, {
			Name:        []string{"--with-web-identity"},
			Description: `Use web identity (requires plugin)`,
		}, {
			Name:        []string{"--json"},
			Description: `Use json credentials`,
			Args: []model.Arg{{
				Name: "json",
			}},
		}, {
			Name:        []string{"--credentials-file"},
			Description: `Target a shared credentials file`,
			Args: []model.Arg{{
				Name: "credentials_file",
			}},
		}, {
			Name:        []string{"--config-file"},
			Description: `Target a config file`,
			Args: []model.Arg{{
				Name: "config_file",
			}},
		}, {
			Name:        []string{"--config"},
			Description: `Configure awsume`,
			Args: []model.Arg{{
				Name:       "option",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--list-plugins"},
			Description: `List installed plugins`,
		}, {
			Name:        []string{"--info"},
			Description: `Print any info logs to stderr`,
		}, {
			Name:        []string{"--debug"},
			Description: `Print any debug logs to stderr`,
		}, {
			Name:        []string{"--console", "-c"},
			Description: `Open AWS console`,
		}, {
			Name:        []string{"--console-link", "-cl"},
			Description: `Get a sign-on url`,
		}, {
			Name:        []string{"--console-service", "-cs"},
			Description: `Open the console to a specific service`,
			Args: []model.Arg{{
				Name: "service",
			}},
		}, {
			Name:        []string{"-cls", "-csl"},
			Description: `Get a sign-on url to a specific service`,
			Args: []model.Arg{{
				Name: "service",
			}},
		}},
	}
}
