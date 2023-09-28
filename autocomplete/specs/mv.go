// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["mv"] = model.Subcommand{
		Name:        []string{"mv"},
		Description: `Move & rename files and folders`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "source",
			IsVariadic: true,
		}, {
			Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:      "target",
		}},
		Options: []model.Option{{
			Name:        []string{"-f"},
			Description: `Do not prompt for confirmation before overwriting the destination path`,
			ExclusiveOn: []string{"-i", "-n"},
		}, {
			Name:        []string{"-i"},
			Description: `Cause mv to write a prompt to standard error before moving a file that would overwrite an existing file`,
			ExclusiveOn: []string{"-f", "-n"},
		}, {
			Name:        []string{"-n"},
			Description: `Do not overwrite existing file`,
			ExclusiveOn: []string{"-f", "-i"},
		}, {
			Name:        []string{"-v"},
			Description: `Cause mv to be verbose, showing files after they are moved`,
		}},
	}
}
