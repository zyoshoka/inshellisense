// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["else"] = model.Subcommand{
		Name:        []string{"else"},
		Description: `Execute this command if the test returned 1`,
		Args: []model.Arg{{
			IsCommand: true,
		}},
	}
}
