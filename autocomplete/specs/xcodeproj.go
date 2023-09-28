// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["xcodeproj"] = model.Subcommand{
		Name:        []string{"xcodeproj"},
		Description: `Xcodeproj lets you create and modify Xcode projects`,
		Options: []model.Option{{
			Name:         []string{"--verbose"},
			Description:  `Show more debugging information`,
			IsPersistent: true,
		}, {
			Name:        []string{"--version"},
			Description: `Show the version of the tool`,
		}, {
			Name:         []string{"--no-ansi"},
			Description:  `Show output without ANSI codes`,
			IsPersistent: true,
		}, {
			Name:         []string{"--help"},
			Description:  `Show help banner of specified command`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"config-dump"},
			Description: `Dumps the build settings of all project targets for all configurations in directories named by the target in the given output directory`,
			Args: []model.Arg{{
				Name:       "PROJECT",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}, {
				Templates:  []model.Template{model.TemplateFolders},
				Name:       "OUTPUT",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"project-diff"},
			Description: `Shows the difference between two projects`,
			Args: []model.Arg{{
				Name:      "PROJECT1",
				Generator: nil, // TODO: port over generator
			}, {
				Name:      "PROJECT2",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--ignore"},
				Description: `A key to ignore in the comparison. Can be specified multiple times`,
				Args: []model.Arg{{
					Name: "KEY",
				}},
			}},
		}, {
			Name:        []string{"show"},
			Description: `Shows an overview of a project in a YAML representation`,
			Args: []model.Arg{{
				Name:       "PROJECT",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--format"},
				Description: `YAML output format`,
				Args: []model.Arg{{
					Name:        "FORMAT",
					Suggestions: []model.Suggestion{{Name: []string{`hash`}}, {Name: []string{`tree_hash`}}, {Name: []string{`raw`}}},
				}},
			}},
		}, {
			Name:        []string{"sort"},
			Description: `Sorts the given project`,
			Args: []model.Arg{{
				Name:       "PROJECT",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--group-option"},
				Description: `The position of the groups when sorting. If no option is specified sorting will interleave groups and files`,
				Args: []model.Arg{{
					Name:        "POSITION",
					Suggestions: []model.Suggestion{{Name: []string{`above`}}, {Name: []string{`below`}}},
				}},
			}},
		}, {
			Name:        []string{"target-diff"},
			Description: `Shows the difference between two targets`,
			Args: []model.Arg{{
				Name: "TARGET1",
			}, {
				Name: "TARGET2",
			}},
			Options: []model.Option{{
				Name:        []string{"--project"},
				Description: `The Xcode project document to use`,
				Args: []model.Arg{{
					Name:      "PATH",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}},
	}
}
