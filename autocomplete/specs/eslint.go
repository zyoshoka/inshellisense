// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["eslint"] = model.Subcommand{
		Name:        []string{"eslint"},
		Description: `Pluggable JavaScript linter`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:        "file|dir|glob",
			Description: `File(s) to lint`,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--no-eslintrc"},
			Description: `Disable use of configuration from .eslintrc.*`,
		}, {
			Name:        []string{"-c", "--config"},
			Description: `Use this configuration, overriding .eslintrc.* config options if present`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"--env"},
			Description: `Specify environments`,
			Args: []model.Arg{{
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--ext"},
			Description: `Specify JavaScript file extensions`,
			Args: []model.Arg{{
				Name: "Extension",
			}},
		}, {
			Name:        []string{"--global"},
			Description: `Define global variables`,
			Args: []model.Arg{{
				Name: "Variables",
			}},
		}, {
			Name:        []string{"--parser"},
			Description: `Specify the parser to be used`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--parser-options"},
			Description: `Specify parser options`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--resolve-plugins-relative-to"},
			Description: `A folder where plugins should be resolved from`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
			}},
		}, {
			Name:        []string{"--rulesdir"},
			Description: `Use additional rules from this directory`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
			}},
		}, {
			Name:        []string{"--plugin"},
			Description: `Specify plugins`,
			Args: []model.Arg{{
				Name:      "Plugin",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--rule"},
			Description: `Specify rules`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--fix"},
			Description: `Automatically fix problems`,
		}, {
			Name:        []string{"--fix-dry-run"},
			Description: `Automatically fix problems without saving the changes to the file system`,
		}, {
			Name:        []string{"--fix-type"},
			Description: `Specify the types of fixes to apply`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`problem`}}, {Name: []string{`suggestion`}}, {Name: []string{`layout`}}},
			}},
		}, {
			Name:        []string{"--ignore-path"},
			Description: `Specify path of ignore file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"--no-ignore"},
			Description: `Disable use of ignore files and patterns`,
		}, {
			Name:        []string{"--ignore-pattern"},
			Description: `Pattern of files to ignore (in addition to those in .eslintignore)`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--stdin"},
			Description: `Lint code provided on <STDIN>`,
		}, {
			Name:        []string{"--stdin-filename"},
			Description: `Specify filename to process STDIN as`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"--quiet"},
			Description: `Report errors only`,
		}, {
			Name:        []string{"--max-warnings"},
			Description: `Number of warnings to trigger nonzero exit code`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"-o", "--output-file"},
			Description: `Specify file to write report to`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"-f", "--format"},
			Description: `Use a specific output format`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`checkstyle`}}, {Name: []string{`codeframe`}}, {Name: []string{`compact`}}, {Name: []string{`html`}}, {Name: []string{`jslint-xml`}}, {Name: []string{`json`}}, {Name: []string{`junit`}}, {Name: []string{`stylish`}}, {Name: []string{`table`}}, {Name: []string{`tap`}}, {Name: []string{`unix`}}, {Name: []string{`visualstudio`}}},
			}},
		}, {
			Name:        []string{"--color"},
			Description: `Force enabling of color`,
			ExclusiveOn: []string{"--no-color"},
		}, {
			Name:        []string{"--no-color"},
			Description: `Force disabling of color`,
			ExclusiveOn: []string{"--color"},
		}, {
			Name:        []string{"--no-inline-config"},
			Description: `Prevent comments from changing config or rules`,
		}, {
			Name:        []string{"--report-unused-disable-directives"},
			Description: `Adds reported errors for unused eslint-disable-directives`,
		}, {
			Name:        []string{"--cache"},
			Description: `Only check changed files`,
		}, {
			Name:        []string{"--cache-location"},
			Description: `Path to the cache file or directory`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			}},
		}, {
			Name:        []string{"--cache-strategy"},
			Description: `Strategy to use for detecting changed files`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`metadata`}}, {Name: []string{`content`}}},
			}},
		}, {
			Name:        []string{"--init"},
			Description: `Run config initialization wizard`,
		}, {
			Name:        []string{"--env-info"},
			Description: `Output execution environment information`,
		}, {
			Name:        []string{"--no-error-on-unmatched-pattern"},
			Description: `Prevent errors when pattern is unmatched`,
		}, {
			Name:        []string{"--debug"},
			Description: `Output debugging information`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Show help`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Output the version number`,
		}, {
			Name:        []string{"--print-config"},
			Description: `Print the configuration for the give file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}},
	}
}
