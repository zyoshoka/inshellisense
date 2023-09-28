// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["elm-format"] = model.Subcommand{
		Name:        []string{"elm-format"},
		Description: `Format your code in the Elm idiomatic way`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
			Name:      "INPUT",
		}},
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for elm-format`,
		}, {
			Name:        []string{"--output"},
			Description: `Write output to FILE instead of overwriting the given source file`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "FILE",
				Description: `Name and location of output file`,
			}},
		}, {
			Name:        []string{"--yes"},
			Description: `Reply 'yes' to all automated prompts`,
		}, {
			Name:        []string{"--validate"},
			Description: `Check if files are formatted without changing them`,
		}, {
			Name:        []string{"--stdin"},
			Description: `Read from stdin, output to stdout`,
		}, {
			Name:        []string{"--elm-version"},
			Description: `The Elm version of the source files being formatted`,
			Args: []model.Arg{{
				Name:        "VERSION",
				Description: `Valid values: 0.18, 0.19. Default: auto`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
		}},
	}
}
