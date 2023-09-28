// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["bat"] = model.Subcommand{
		Name:        []string{"bat"},
		Description: `A cat(1) clone with syntax highlighting and Git integration`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths},
			Description: `File(s)`,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-A", "--show-all"},
			Description: `Show non-printable characters`,
		}, {
			Name:        []string{"-p", "--plain"},
			Description: `Show plain style, no decorations`,
		}, {
			Name:        []string{"-l", "--language"},
			Description: `Explicitly set the language for syntax highlighting`,
			Args: []model.Arg{{
				Name:      "<language>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-H", "--highlight-line"},
			Description: `Highlight the specified line ranges`,
			Args: []model.Arg{{
				Name:        "<N:M> ",
				Description: `Range of line`,
			}},
		}, {
			Name:        []string{"--file-name"},
			Description: `Specify the name to display for a file. Useful when piping data to bat from STDIN when bat does not otherwise know the filename`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Description: `File(s)`,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"-d", "--diff"},
			Description: `Show lines that have been added/removed/modified with respect to the Git index`,
		}, {
			Name:        []string{"--diff-context"},
			Description: `Include N lines of context around added/removed/modified lines when using '--diff'`,
			Args: []model.Arg{{
				Name:        "<N> ",
				Description: `Lines of context`,
			}},
		}, {
			Name:        []string{"--tabs"},
			Description: `Set the tab width to T spaces. Use a width of 0 to pass tabs through directly`,
			Args: []model.Arg{{
				Name:        "<T> ",
				Description: `Spaces of tab width`,
			}},
		}, {
			Name:        []string{"--wrap"},
			Description: `Specify when to use colored output`,
			Args: []model.Arg{{
				Name:      "<mode>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--terminal-width"},
			Description: `Explicitly set the width of the terminal instead of determining it automatically`,
			Args: []model.Arg{{
				Name:        "<width> ",
				Description: `Width to wrap`,
			}},
		}, {
			Name:        []string{"-n", "--number"},
			Description: `Show line numbers, no other decorations`,
		}, {
			Name:        []string{"--color"},
			Description: `Specify when to use colored output`,
			Args: []model.Arg{{
				Name:      "<when>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--italic-text"},
			Description: `Specify when to use ANSI sequences for italic text in the output`,
			Args: []model.Arg{{
				Name:      "<when>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--decorations"},
			Description: `Specify when to use the decorations that have been specified via '--style'`,
			Args: []model.Arg{{
				Name:      "<when>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-f", "--force-colorization"},
			Description: `Alias for '--decorations=always --color=always'`,
		}, {
			Name:        []string{"--paging"},
			Description: `Specify when to use the pager`,
			Args: []model.Arg{{
				Name:      "<when>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--pager"},
			Description: `Determine which pager is used`,
			Args: []model.Arg{{
				Name:        "<command>",
				Description: `Command to pager`,
			}},
		}, {
			Name:        []string{"-m", "--map-syntax"},
			Description: `Map a glob pattern to an existing syntax name`,
			Args: []model.Arg{{
				Name:        "<glob:syntax>",
				Description: `Full path or the filename:language that highlighted`,
			}},
		}, {
			Name:        []string{"--ignored-suffix"},
			Description: `Ignore extension`,
			Args: []model.Arg{{
				Name:        "<ignored-suffix>",
				Description: `Extension`,
			}},
		}, {
			Name:        []string{"--theme"},
			Description: `Set the theme for syntax highlighting`,
			Args: []model.Arg{{
				Name:      "<theme>",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--list-themes"},
			Description: `Display a list of supported themes for syntax highlighting`,
		}, {
			Name:        []string{"--style"},
			Description: `Display a list of supported themes for syntax highlighting`,
		}, {
			Name:        []string{"-r", "--line-range"},
			Description: `Only print the specified range of lines for each file`,
			Args: []model.Arg{{
				Name:        "<N:M> ",
				Description: `Range of line`,
			}},
		}, {
			Name:        []string{"-L", "--list-languages"},
			Description: `Display a list of supported languages for syntax highlighting`,
		}, {
			Name:        []string{"-u", "--unbuffered"},
			Description: `Make output unbuffered (exists for POSIX-compliance reasons and is simply ignored)`,
		}, {
			Name:        []string{"--diagnostic"},
			Description: `Show diagnostic information for bug reports`,
		}, {
			Name:        []string{"--acknowledgements"},
			Description: `Show acknowledgements`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print help message`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Show version information`,
		}},
	}
}
