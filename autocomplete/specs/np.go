// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["np"] = model.Subcommand{
		Name:        []string{"np"},
		Description: `A better npm publish`,
		Args: []model.Arg{{
			Name:        "version",
			Description: `Version to publish`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--any-branch"},
			Description: `Allow publishing from any branch`,
		}, {
			Name:        []string{"--branch"},
			Description: `Name of the release branch (default: main | master)`,
			Args: []model.Arg{{
				Name: "branch",
			}},
		}, {
			Name:        []string{"--no-cleanup"},
			Description: `Skips cleanup of node_modules`,
		}, {
			Name:        []string{"--no-tests"},
			Description: `Skips tests`,
		}, {
			Name:        []string{"--yolo"},
			Description: `Skips cleanup and testing`,
		}, {
			Name:        []string{"--no-publish"},
			Description: `Skips publishing`,
		}, {
			Name:        []string{"--preview"},
			Description: `Show tasks without actually executing them`,
		}, {
			Name:        []string{"--tag"},
			Description: `Publish under a given dist-tag`,
			Args: []model.Arg{{
				Name: "tag",
			}},
		}, {
			Name:        []string{"--no-yarn"},
			Description: `Don't use Yarn`,
		}, {
			Name:        []string{"--contents"},
			Description: `Subdirectory to publish`,
			Args: []model.Arg{{
				Name: "directory",
			}},
		}, {
			Name:        []string{"--no-release-draft"},
			Description: `Skips opening a GitHub release draft`,
		}, {
			Name:        []string{"--release-draft-only"},
			Description: `Only opens a GitHub release draft`,
		}, {
			Name:        []string{"--test-script"},
			Description: `Name of npm run script to run tests before publishing (default: test)`,
			Args: []model.Arg{{
				Name: "script",
			}},
		}, {
			Name:        []string{"--no-2fa"},
			Description: `Don't enable 2FA on new packages (not recommended)`,
		}, {
			Name:        []string{"--message"},
			Description: `Version bump commit message. "%s" will be replaced with version. (default: '%s' with npm and 'v%s' with yarn)`,
			Args: []model.Arg{{
				Name: "message",
			}},
		}},
	}
}
