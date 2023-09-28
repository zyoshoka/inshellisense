// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["palera1n"] = model.Subcommand{
		Name:        []string{"palera1n"},
		Description: `Jailbreaking tool for iOS/iPadOS 15.x-16.x arm64`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `Print version`,
		}, {
			Name:        []string{"--force-revert"},
			Description: `Remove jailbreak`,
		}, {
			Name:        []string{"--setup-partial-fakefs", "-B"},
			Description: `Setup partial fakefs`,
		}, {
			Name:        []string{"--setup-fakefs", "-c"},
			Description: `Setup fakefs`,
		}, {
			Name:        []string{"--demote", "-d"},
			Description: `Demote`,
		}, {
			Name:        []string{"--dfuhelper", "-D"},
			Description: `Exit after entering DFU`,
		}, {
			Name:        []string{"--boot-args", "-e"},
			Description: `XNU boot arguments`,
			Args: []model.Arg{{
				Name: "BOOT_ARGUMENTS",
			}},
		}, {
			Name:        []string{"--enter-recovery", "-E"},
			Description: `Enter recovery mode`,
		}, {
			Name:        []string{"--fakefs", "-f"},
			Description: `Boot fakefs`,
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Show help for palera1n`,
		}, {
			Name:        []string{"--override-checkra1n", "-i"},
			Description: `Override checkra1n`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"--override-pongo", "-k"},
			Description: `Override Pongo image`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"--override-kpf", "-K"},
			Description: `Override kernel patchfinder`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"--rootless", "-l"},
			Description: `Boot rootless. This is the default`,
		}, {
			Name:        []string{"--jbinit-log-to-file", "-L"},
			Description: `Make jbinit log to /cores/jbinit.log (can be read from sandbox while jailbroken)`,
		}, {
			Name:        []string{"--exit-recovery", "-n"},
			Description: `Exit recovery mode`,
		}, {
			Name:        []string{"--device-info", "-I"},
			Description: `Print info about the connected device`,
		}, {
			Name:        []string{"--override-overlay", "-o"},
			Description: `Override overlay`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"--pongo-shell", "-p"},
			Description: `Boots to PongoOS shell`,
		}, {
			Name:        []string{"--pongo-full", "-P"},
			Description: `Boots to a PongoOS shell with default images already uploaded`,
		}, {
			Name:        []string{"--override-ramdisk", "-r"},
			Description: `Override ramdisk`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"--reboot-device", "-R"},
			Description: `Reboot connected device in normal mode`,
		}, {
			Name:        []string{"--safe-mode", "-s"},
			Description: `Enter safe mode`,
		}, {
			Name:        []string{"--no-colors", "-S"},
			Description: `Disable colors on the command line`,
		}, {
			Name:        []string{"--debug-logging", "-v"},
			Description: `Enable debug logging. This option can be repeated for extra verbosity`,
		}, {
			Name:        []string{"--verbose-boot", "-V"},
			Description: `Verbose boot`,
		}},
	}
}
