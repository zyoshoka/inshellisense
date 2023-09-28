// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["envchain"] = model.Subcommand{
		Name:        []string{"envchain"},
		Description: `Set environment variables with macOS keychain or D-Bus secret service`,
		Args: []model.Arg{{
			Name:      "NAMESPACE",
			Generator: nil, // TODO: port over generator
		}, {
			Name:      "CMD",
			IsCommand: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"-s", "--set"},
			Description: `Add keychain item of environment variable +ENV+ for namespace +NAMESPACE+`,
			Args: []model.Arg{{
				Name:      "NAMESPACE",
				Generator: nil, // TODO: port over generator
			}, {
				Name:       "ENV",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-n", "--noecho"},
				Description: `Do not echo user input`,
			}, {
				Name:        []string{"-p", "--require-passphrase"},
				Description: `Always ask for keychain passphrase`,
				ExclusiveOn: []string{"--no-require-passphrase"},
			}, {
				Name:        []string{"-P", "--no-require-passphrase"},
				Description: `Do not ask for keychain passphrase`,
				ExclusiveOn: []string{"--require-passphrase"},
			}},
		}, {
			Name:        []string{"-l", "--list"},
			Description: `List namespaces that have been created`,
		}},
	}
}
