// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["for"] = model.Subcommand{
		Name:        []string{"for"},
		Description: `Perform a set of commands multiple times`,
		Args: []model.Arg{{
			Name: "var",
		}, {
			Suggestions: []model.Suggestion{{Name: []string{`in`}}},
		}, {
			Name:       "values",
			IsVariadic: true,
		}},
	}
}
