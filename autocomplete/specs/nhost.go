// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["nhost"] = model.Subcommand{
		Name:        []string{"nhost"},
		Description: `Nhost's command-line`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"deploy"},
			Description: `Deploy local migrations and metadata changes to Nhost production`,
		}, {
			Name:        []string{"dev"},
			Description: `Start Nhost project for local development`,
		}, {
			Name:        []string{"down"},
			Description: `Stop and remove local Nhost backend started by "nhost dev"`,
		}, {
			Name:        []string{"env", "env:ls"},
			Description: `List environment variables`,
		}, {
			Name:        []string{"env:pull"},
			Description: `Sync remote environment variables to your local environment`,
		}, {
			Name:        []string{"help"},
			Description: `Display help for nhost`,
		}, {
			Name:        []string{"init"},
			Description: `Initialize current working directory as a Nhost project`,
		}, {
			Name:        []string{"link"},
			Description: `Link Nhost Project`,
		}, {
			Name:        []string{"login"},
			Description: `Login to your Nhost account`,
			Options: []model.Option{{
				Name:        []string{"--email", "-e"},
				Description: `Email address`,
				Args: []model.Arg{{
					Name: "email",
				}},
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Logout from your Nhost account`,
		}},
	}
}
