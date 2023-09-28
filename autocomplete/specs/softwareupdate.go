// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["softwareupdate"] = model.Subcommand{
		Name:        []string{"softwareupdate"},
		Description: `Software Update checks for new and updated versions of your software based on information about your computer and current software`,
		Options: []model.Option{{
			Name:         []string{"--no-scan"},
			Description:  `Do not scan when listing or installing updates (use available updates previously scanned)`,
			IsPersistent: true,
		}, {
			Name:        []string{"--product-types"},
			Description: `Limit a scan to a particular product type only - ignoring all others`,
			Args: []model.Arg{{
				Name: "type",
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--products"},
			Description:  `A comma-separated (no spaces) list of product keys to operate on`,
			IsPersistent: true,
		}, {
			Name:         []string{"--force"},
			Description:  `Force an operation to complete.  Use with --background to trigger a background scan regardless of "Automatically check" pref`,
			IsPersistent: true,
		}, {
			Name:         []string{"--agree-to-license"},
			Description:  `Agree to the software license agreement without user interaction`,
			IsPersistent: true,
		}, {
			Name:         []string{"--verbose"},
			Description:  `Enable verbose output`,
			IsPersistent: true,
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Show help for softwareupdate`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"-l", "--list"},
			Description: `List all appropriate update labels`,
		}, {
			Name:        []string{"-d", "--download"},
			Description: `Download Only`,
		}, {
			Name:        []string{"-i", "--install"},
			Description: `Install`,
			Args: []model.Arg{{
				Name:       "label",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-a", "--all"},
				Description: `All appropriate updates`,
				ExclusiveOn: []string{"-r", "--recommended"},
			}, {
				Name:        []string{"-R", "--restart"},
				Description: `Automatically restart (or shut down) if required to complete installation`,
			}, {
				Name:        []string{"-r", "--recommended"},
				Description: `Only recommended updates`,
				ExclusiveOn: []string{"-a", "--all"},
			}},
		}, {
			Name:        []string{"--list-full-installers"},
			Description: `List the available macOS Installers`,
		}, {
			Name:        []string{"--fetch-full-installer"},
			Description: `Install the latest recommended macOS Installer`,
		}, {
			Name:        []string{"--full-installer-version"},
			Description: `The version of macOS to install`,
			Args: []model.Arg{{
				Name: "version",
			}},
		}, {
			Name:        []string{"--install-rosetta"},
			Description: `Install Rosetta 2`,
		}, {
			Name:        []string{"--background"},
			Description: `Trigger a background scan and update operation`,
		}, {
			Name:        []string{"--dump-state"},
			Description: `Log the internal state of the SU daemon to /var/log/install.log`,
		}, {
			Name:        []string{"--evaluate-products"},
			Description: `Evaluate a list of product keys specified by the --products option`,
		}, {
			Name:        []string{"--history"},
			Description: `Show the install history.  By default, only displays updates installed by softwareupdate`,
			Options: []model.Option{{
				Name:        []string{"--all"},
				Description: `Include all processes in history (including App installs)`,
			}},
		}},
	}
}
