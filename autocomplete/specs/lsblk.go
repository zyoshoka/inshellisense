// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["lsblk"] = model.Subcommand{
		Name:        []string{"lsblk"},
		Description: `List block devices`,
		Args: []model.Arg{{
			Name:        "device",
			Description: `Device to list`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for lsblk`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Show version for lsblk`,
		}, {
			Name:        []string{"--all", "-a"},
			Description: `Also list empty devices and RAM disk devices`,
		}, {
			Name:        []string{"--bytes", "-b"},
			Description: `Print the SIZE column in bytes`,
		}, {
			Name:        []string{"--discard", "-D"},
			Description: `Print information about the discarding capabilities (TRIM, UNMAP) for each device`,
		}, {
			Name:        []string{"--nodeps", "-d"},
			Description: `Do not print holder devices or slaves`,
		}, {
			Name:        []string{"--dedup", "-E"},
			Description: `Use column as a de-duplication key to de-duplicate output tree`,
			Args: []model.Arg{{
				Name: "column",
			}},
		}, {
			Name:        []string{"--exclude", "-e"},
			Description: `Exclude the devices specified by the comma-separated list of major device numbers`,
			Args: []model.Arg{{
				Name: "list",
			}},
		}, {
			Name:        []string{"--fs", "-f"},
			Description: `Output info about filesystems`,
		}, {
			Name:        []string{"--include", "-I"},
			Description: `Include devices specified by the comma-separated list of major device numbers`,
			Args: []model.Arg{{
				Name: "list",
			}},
		}, {
			Name:        []string{"--ascii", "-i"},
			Description: `Use ASCII characters for tree formatting`,
		}, {
			Name:        []string{"--json", "-J"},
			Description: `Use JSON output format`,
		}, {
			Name:        []string{"--list", "-l"},
			Description: `Produce output in the form of a list`,
		}, {
			Name:        []string{"--merge", "-M"},
			Description: `Group parents of sub-trees to provide more readable output for RAIDs and Multi-path devices`,
		}, {
			Name:        []string{"--perms", "-m"},
			Description: `Output info about device owner, group and mode`,
		}, {
			Name:        []string{"--noheadings", "-n"},
			Description: `Do not print a header line`,
		}, {
			Name:        []string{"--output", "-o"},
			Description: `Specify which output columns to print`,
			Args: []model.Arg{{
				Name:       "list",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--output-all", "-O"},
			Description: `Output all available columns`,
		}, {
			Name:        []string{"--pairs", "-P"},
			Description: `Produce output in the form of key-value pairs`,
		}, {
			Name:        []string{"--raw", "-r"},
			Description: `Produce output in raw format`,
		}, {
			Name:        []string{"--scsi", "-S"},
			Description: `Output info about SCSI devices only`,
		}, {
			Name:        []string{"--inverse", "-s"},
			Description: `Print dependencies in inverse order`,
		}, {
			Name:        []string{"--tree", "-T"},
			Description: `Force tree-like output format`,
			Args: []model.Arg{{
				Name: "column",
			}},
		}, {
			Name:        []string{"--topology", "-t"},
			Description: `Output info about block-device topology`,
		}, {
			Name:        []string{"--width", "-w"},
			Description: `Specifies output width as a number of characters`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"--sort", "-x"},
			Description: `Sort output lines by column`,
			Args: []model.Arg{{
				Name: "column",
			}},
		}, {
			Name:        []string{"--zoned", "-z"},
			Description: `Print the zone model for each device`,
		}, {
			Name:        []string{"--sysroot"},
			Description: `Gather data for a Linux instance other than the instance from which the lsblk command is issued`,
			Args: []model.Arg{{
				Name: "directory",
			}},
		}},
	}
}
