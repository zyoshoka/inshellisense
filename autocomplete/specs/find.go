// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["find"] = model.Subcommand{
		Name:        []string{"find"},
		Description: `Walk a file hierarchy`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFolders},
			Name:       "path",
			IsOptional: true,
			IsVariadic: true,
		}, {
			Name:        "expression",
			Description: `Composition of primaries and operands`,
			IsOptional:  true,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-E"},
			Description: `Interpret regular expressions followed by -regex and -iregex primaries as extended`,
		}, {
			Name:        []string{"-H"},
			Description: `Cause the file information and file type returned for each symbolic link specified to be those referenced by the link`,
			ExclusiveOn: []string{"-L", "-P"},
		}, {
			Name:        []string{"-L"},
			Description: `Cause the file information and file type returned for each symbolic link to be those of the file referenced by the link`,
			ExclusiveOn: []string{"-H", "-P"},
		}, {
			Name:        []string{"-P"},
			Description: `Cause the file information and file type returned for each symbolic link to be those for the link itself`,
			ExclusiveOn: []string{"-H", "-L"},
		}, {
			Name:        []string{"-X"},
			Description: `Permit find to be safely used in conjunction with xargs`,
		}, {
			Name:        []string{"-d"},
			Description: `Cause find to perform a depth-first traversal`,
		}, {
			Name:        []string{"-f"},
			Description: `Specify a file hierarch for find to traverse`,
			Args: []model.Arg{{
				Name: "path",
			}},
		}, {
			Name:        []string{"-s"},
			Description: `Cause find to traverse the file hierarchies in lexicographical order`,
		}, {
			Name:        []string{"-x"},
			Description: `Prevent find from descending into directories that have a device number different than that of the file from which the descent began`,
		}},
	}
}
