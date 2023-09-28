// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["xed"] = model.Subcommand{
		Name:        []string{"xed"},
		Description: `Xcode text editor invocation tool`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths},
			Name:        "file",
			Description: `A list of file paths. If no files are passed, then standard input will       be read and piped into a new untitled document`,
			IsOptional:  true,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--launch", "-x"},
			Description: `Launches Xcode, opening a new empty unsaved file`,
		}, {
			Name:        []string{"--create", "-c"},
			Description: `Creates any non-existent files in the file list. If used without --launch, standard input will be read and piped to the last file created`,
		}, {
			Name:        []string{"--wait", "-w"},
			Description: `Wait for the files to be closed before exiting. xed will idle and will only terminate when all files are closed`,
		}, {
			Name:        []string{"--line", "-l"},
			Description: `Selects the given line in the last file opened`,
			Args: []model.Arg{{
				Name:        "number",
				Description: `The line number to select`,
			}},
		}, {
			Name:        []string{"--background", "-b"},
			Description: `Opens Xcode without activating it; the process that invoked xed remains in front`,
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Show help for xed`,
			ExclusiveOn: []string{"-x", "-c", "-w", "-l", "-b", "-v"},
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Prints the version number of xed`,
		}},
	}
}
