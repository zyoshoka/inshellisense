// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["volta"] = model.Subcommand{
		Name:        []string{"volta"},
		Description: `The JavaScript Launcher`,
		Options: []model.Option{{
			Name:        []string{"--verbose"},
			Description: `Enables verbose diagnostics`,
		}, {
			Name:        []string{"--quiet"},
			Description: `Prevents unnecessary output`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Prints help information`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Prints the current version of Volta`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"fetch"},
			Description: `Fetches a tool to the local machine`,
			Args: []model.Arg{{
				Name:       "tool | tool@version",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"install"},
			Description: `Installs a tool in your toolchain`,
			Args: []model.Arg{{
				Name:       "tool@version",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"uninstall"},
			Description: `Uninstalls a tool from your toolchain`,
			Args: []model.Arg{{
				Name: "tool",
			}},
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"pin"},
			Description: `Pins your project's runtime or package manager`,
			Args: []model.Arg{{
				Name:       "tool@version",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"list"},
			Description: `Displays the current toolchain`,
			Args: []model.Arg{{
				Name: "tool",
			}},
			Options: []model.Option{{
				Name:        []string{"-c", "--current"},
				Description: `Show the currently-active tool(s)`,
			}, {
				Name:        []string{"-d", "--default"},
				Description: `Show your default tool(s)`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--format"},
				Description: `Specify output format`,
				Args: []model.Arg{{
					Name:        "output format",
					Suggestions: []model.Suggestion{{Name: []string{`human`}}, {Name: []string{`plain`}}},
				}},
			}},
		}, {
			Name:        []string{"completions"},
			Description: `Generates Volta completions`,
			Args: []model.Arg{{
				Name:        "shell",
				Description: `Shell to generate completions for`,
				Suggestions: []model.Suggestion{{Name: []string{`zsh`}}, {Name: []string{`bash`}}, {Name: []string{`fish`}}, {Name: []string{`powershell`}}, {Name: []string{`elivsh`}}},
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `Write over an existing file, if any`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"-o", "--output"},
				Description: `File to write generated completions to`,
				Args: []model.Arg{{
					Name: "file",
				}},
			}},
		}, {
			Name:        []string{"which"},
			Description: `Locates the actual binary that will be called by Volta`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "binary",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"setup"},
			Description: `Enables Volta for the current user`,
			Options: []model.Option{{
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run a command with custom Node, npm, and/or Yarn versions`,
			Args: []model.Arg{{
				Name: "command",
			}},
			Options: []model.Option{{
				Name:        []string{"--bundle"},
				Description: `Forces npm to be the version bundled with Node`,
			}, {
				Name:        []string{"--no-yarn"},
				Description: `Disables Yarn`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Enables verbose diagnostics`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Prevents unnecessary output`,
			}, {
				Name:        []string{"--node"},
				Description: `Set the custom Node version`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"--npm"},
				Description: `Set the custom npm version`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"--yarn"},
				Description: `Set the custom Yarn version`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"--env"},
				Description: `Set an environment variable (can be used multiple times)`,
				Args: []model.Arg{{
					Name: "NAME=value",
				}},
			}},
		}, {
			Name:        []string{"help"},
			Description: `Prints this message or the help of the given subcommand(s)`,
			Args: []model.Arg{{
				Name: "subcommand",
			}},
		}},
	}
}
