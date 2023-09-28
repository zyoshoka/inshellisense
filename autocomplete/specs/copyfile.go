// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["copyfile"] = model.Subcommand{
		Name:        []string{"copyfile"},
		Description: `Oh-My-Zsh plugin that copies the contents of a file to the clipboard`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
			Name:      "file",
		}},
	}
}
