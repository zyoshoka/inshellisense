// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["touch"] = model.Subcommand{
		Name:        []string{"touch"},
		Description: `Change file access and modification times`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFolders},
			Name:       "file",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-A"},
			Description: `Adjust the access and modification time stamps for the file by the specified value`,
			Args: []model.Arg{{
				Name:        "time",
				Description: `[-][[hh]mm]SS`,
			}},
		}, {
			Name:        []string{"-a"},
			Description: `Change the access time of the file`,
		}, {
			Name:        []string{"-c"},
			Description: `Do not create the file if it does not exist`,
		}, {
			Name:        []string{"-f"},
			Description: `Attempt to force the update, even if the file permissions do not currently permit it`,
		}, {
			Name:        []string{"-h"},
			Description: `If the file is a symbolic link, change the times of the link itself rather than the file that the link points to`,
		}, {
			Name:        []string{"-m"},
			Description: `Change the modification time of the file`,
		}, {
			Name:        []string{"-r"},
			Description: `Use the access and modifications times from the specified file instead of the current time of day`,
			Args: []model.Arg{{
				Name: "file",
			}},
		}, {
			Name:        []string{"-t"},
			Description: `Change the access and modification times to the specified time instead of the current time of day`,
			Args: []model.Arg{{
				Name:        "timestamp",
				Description: `[[CC]YY]MMDDhhmm[.SS]`,
			}},
		}},
	}
}
