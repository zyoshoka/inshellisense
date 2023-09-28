// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["fisher"] = model.Subcommand{
		Name:        []string{"fisher"},
		Description: `A plugin manager for Fish`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for fisher`,
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Show fisher version`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"install"},
			Description: `Install plugin`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"IlanCosman/tide@v5"},
				Description: `[Prompt] - 🌊 The ultimate Fish prompt`,
			}, {
				Name:        []string{"pure-fish/pure"},
				Description: `[Prompt] - Pretty, minimal, and fast prompt for Fish from Zsh`,
			}, {
				Name:        []string{"jorgebucaran/hydro"},
				Description: `[Prompt] - Minimal, lag-free prompt with async Git status`,
			}, {
				Name:        []string{"jethrokuan/z"},
				Description: `[Plugin] - Pure-fish z directory jumping`,
			}, {
				Name:        []string{"PatrickF1/fzf.fish"},
				Description: `[Plugin] - Augment your fish command line with fzf key bindings`,
			}, {
				Name:        []string{"jorgebucaran/nvm.fish"},
				Description: `[Plugin] -  Node.js version manager lovingly made for Fish`,
			}, {
				Name:        []string{"franciscolourenco/done"},
				Description: `[Plugin] - A fish-shell package to automatically receive notifications when long processes finish`,
			}, {
				Name:        []string{"jorgebucaran/replay.fish"},
				Description: `[Plugin] - Run Bash commands replaying changes in Fish. 🍤`,
			}, {
				Name:        []string{"jorgebucaran/spark.fish"},
				Description: `[Plugin] - Visualize a range of numbers right in your terminal`,
			}, {
				Name:        []string{"joseluisq/gitnow@2.10.0"},
				Description: `[Plugin] - Speed up your Git workflow. 🐠`,
			}},
		}, {
			Name:        []string{"remove"},
			Description: `Remove plugins`,
			Args: []model.Arg{{
				Name:        "installed plugins",
				Description: `The plugin you want to remove`,
				Generator:   nil, // TODO: port over generator
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Update plugins`,
			Args: []model.Arg{{
				Name:        "installed plugins",
				Description: `The plugin you want to update`,
				Generator:   nil, // TODO: port over generator
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"list"},
			Description: `List plugins`,
			Args: []model.Arg{{
				Name:        "RegEx",
				Description: `Search in list with regular expression`,
				IsOptional:  true,
			}},
		}},
	}
}
