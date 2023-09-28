// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["tldr"] = model.Subcommand{
		Name:        []string{"tldr"},
		Description: `A simpler man page than the existing man page`,
		Args: []model.Arg{{
			Generator: nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Display help for command`,
		}, {
			Name:        []string{"-s", "--search"},
			Description: `Search all pages for the query`,
			Args: []model.Arg{{
				Name: "query",
			}},
		}, {
			Name:        []string{"--linux"},
			Description: `Show command page for Linux`,
			Args: []model.Arg{{
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--osx"},
			Description: `Show command page for OSX`,
			Args: []model.Arg{{
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--sunos"},
			Description: `Show command page for SunOS`,
			Args: []model.Arg{{
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-l", "--list"},
			Description: `Show all pages for current platform`,
		}, {
			Name:        []string{"-u", "--update"},
			Description: `Download the latest pages and generate search index`,
		}, {
			Name:        []string{"-c", "--clear-cache"},
			Description: `Delete the entire local cache`,
		}, {
			Name:        []string{"--platform", "-p"},
			Description: `Select platform`,
			Args: []model.Arg{{
				Name:        "platform",
				Suggestions: []model.Suggestion{{Name: []string{`linux`}}, {Name: []string{`osx`}}, {Name: []string{`sunos`}}, {Name: []string{`windows`}}, {Name: []string{`common`}}},
			}},
		}},
	}
}
