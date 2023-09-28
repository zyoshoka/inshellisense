// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["esbuild"] = model.Subcommand{
		Name:        []string{"esbuild"},
		Description: `An extremely fast JavaScript bundler`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "entry points",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--bundle"},
			Description: `Bundle all dependencies into the output files`,
		}, {
			Name:        []string{"--define"},
			Description: `Replace variable names with a literal value, eg. --define:DEBUG=true`,
			Args: []model.Arg{{
				Name: "name=value",
			}},
		}, {
			Name:        []string{"--external"},
			Description: `Exclude modules from the build`,
			Args: []model.Arg{{
				Name: "module specifier",
			}},
		}, {
			Name:        []string{"--format"},
			Description: `The output format`,
			Args: []model.Arg{{
				Name:        "format",
				Suggestions: []model.Suggestion{{Name: []string{`iife`}}, {Name: []string{`cjs`}}, {Name: []string{`esm`}}},
			}},
		}, {
			Name:        []string{"--loader"},
			Description: `For a given file extension, specify a loader`,
			Args: []model.Arg{{
				Name:      "loaders",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--minify"},
			Description: `Minify the output (sets all the --minify-* options)`,
		}, {
			Name:        []string{"--outdir"},
			Description: `The output directory for multiple entrypoints`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--outfile"},
			Description: `The output file for one entrypoint`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--platform"},
			Description: `The platform target`,
			Args: []model.Arg{{
				Name:        "name",
				Suggestions: []model.Suggestion{{Name: []string{`browser`}}, {Name: []string{`node`}}, {Name: []string{`neutral`}}},
			}},
		}, {
			Name:        []string{"--serve"},
			Description: `Start a local HTTP server on this host:port`,
			Args: []model.Arg{{
				Name:       "[address:]port",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--splitting"},
			Description: `Enable code splitting`,
		}, {
			Name:        []string{"--target"},
			Description: `Set the environment target. Can be a particular ES version or browser version, eg. chrome101`,
			Args: []model.Arg{{
				Name:      "target",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--watch"},
			Description: `Rebuild on file system changes`,
			Args: []model.Arg{{
				Name:        "forever",
				Suggestions: []model.Suggestion{{Name: []string{`forever`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--allow-overwrite"},
			Description: `Allow output files to overwrite input files`,
		}, {
			Name:        []string{"--analyze"},
			Description: `Print a report about the contents of the bundle`,
			Args: []model.Arg{{
				Name:        "verbose",
				Suggestions: []model.Suggestion{{Name: []string{`verbose`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--asset-names"},
			Description: `Path template for 'file' loader files`,
			Args: []model.Arg{{
				Name: "template",
			}},
		}, {
			Name:        []string{"--banner"},
			Description: `Text to be prepended to each output file type`,
			Args: []model.Arg{{
				Name:      "ext=text[,ext=text...]",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--charset"},
			Description: `Use UTF-8 instead of escaped codepoints in ASCII`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"--chunk-names"},
			Description: `Path template to use for code splitting chunks`,
			Args: []model.Arg{{
				Name: "template",
			}},
		}, {
			Name:        []string{"--color"},
			Description: `Force use of terminal colors`,
			Args: []model.Arg{{
				Name:        "enabled",
				Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
			}},
		}, {
			Name:        []string{"--drop"},
			Description: `Remove certain constructs`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`console`}}, {Name: []string{`debugger`}}},
			}},
		}, {
			Name:        []string{"--entry-names"},
			Description: `Path template to use for entry point output paths`,
			Args: []model.Arg{{
				Name: "template",
			}},
		}, {
			Name:        []string{"--footer"},
			Description: `Text to be appended to each file type`,
			Args: []model.Arg{{
				Name:      "ext=text",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--global-name"},
			Description: `The name of the global if using --format=iife`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--ignore-annotations"},
			Description: `Enable this to work with packages that have incorrect tree-shaking annotations`,
		}, {
			Name:        []string{"--inject"},
			Description: `Import the file into all input files, automatically replace matching globals`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "import",
			}},
		}, {
			Name:        []string{"--jsx-factory"},
			Description: `What to use for the JSX factory`,
			Args: []model.Arg{{
				Name:        "factory",
				Suggestions: []model.Suggestion{{Name: []string{`React.createElement`}}, {Name: []string{`h`}}, {Name: []string{`preact.h`}}},
			}},
		}, {
			Name:        []string{"--jsx-fragment"},
			Description: `What to use for the JS Fragment factory`,
			Args: []model.Arg{{
				Name:        "fragment",
				Suggestions: []model.Suggestion{{Name: []string{`React.Fragment`}}, {Name: []string{`Fragment`}}},
			}},
		}, {
			Name:        []string{"--jsx"},
			Description: `Preserve JSX instead of transforming`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name:        []string{`preserve`},
					Description: `Preserve JSX instead of transforming`,
				}, {
					Name:        []string{`automatic`},
					Description: `Use React's new automatic JSX transform`,
				}},
			}},
		}, {
			Name:        []string{"--jsx-dev"},
			Description: `Toggles development mode for the automatic runtime`,
		}, {
			Name:        []string{"--jsx-import-source"},
			Description: `Overrides the root import for runtime functions (default: react)`,
			Args: []model.Arg{{
				Name: "source",
			}},
		}, {
			Name:        []string{"--keep-names"},
			Description: `Preserve 'name' on functions and classes`,
		}, {
			Name:        []string{"--legal-comments"},
			Description: `Where to place legal comments`,
			Args: []model.Arg{{
				Name:        "location",
				Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`inline`}}, {Name: []string{`eof`}}, {Name: []string{`linked`}}, {Name: []string{`external`}}},
			}},
		}, {
			Name:        []string{"--log-level"},
			Description: `Set the log level`,
			Args: []model.Arg{{
				Name:        "level",
				Suggestions: []model.Suggestion{{Name: []string{`verbose`}}, {Name: []string{`debug`}}, {Name: []string{`info`}}, {Name: []string{`warning`}}, {Name: []string{`error`}}, {Name: []string{`silent`}}},
			}},
		}, {
			Name:        []string{"--log-limit"},
			Description: `Maximum message count, 0 to disable`,
			Args: []model.Arg{{
				Name: "count",
			}},
		}, {
			Name:        []string{"--log-override"},
			Description: `For a particular identifier, set the log level`,
			Args: []model.Arg{{
				Name:      "identifier:level",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--main-fields"},
			Description: `Override the main file order in package.json`,
			Args: []model.Arg{{
				Name: "field order",
			}},
		}, {
			Name:        []string{"--mangle-cache"},
			Description: `Save 'mangle props' decisions to a JSON file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--mangle-props"},
			Description: `Rename all properties matching a regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"--mangle-quoted"},
			Description: `Enable mangling (renaming) quoted properties`,
			Args: []model.Arg{{
				Name:        "status",
				Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
			}},
		}, {
			Name:        []string{"--metafile"},
			Description: `Write metadata about the build to a JSON file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--minify-whitespace"},
			Description: `Remove unnecessary whitespace in output files`,
		}, {
			Name:        []string{"--minify-identifiers"},
			Description: `Shorten identifiers in output files`,
		}, {
			Name:        []string{"--minify-syntax"},
			Description: `Use equivalent but shorter syntax in output files`,
		}, {
			Name:        []string{"--out-extension"},
			Description: `Use a custom output extension for each extension`,
			Args: []model.Arg{{
				Name:      "ext=new",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--outbase"},
			Description: `Base path used to determine entrypoint output paths, for multiple entrypoints`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--preserve-symlinks"},
			Description: `Disable symlink resolution`,
		}, {
			Name:        []string{"--public-path"},
			Description: `Set the base URL for the 'file' loader`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--pure"},
			Description: `Mark the name as a pure function for tree shaking`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--reserve-props"},
			Description: `Do not mangle these properties`,
			Args: []model.Arg{{
				Name: "properties",
			}},
		}, {
			Name:        []string{"--resolve-extensions"},
			Description: `Comma-separated list of implicit extensions`,
			Args: []model.Arg{{
				Name:      "extensions",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--servedir"},
			Description: `What to serve in addition to the generated output files`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--source-root"},
			Description: `Set the sourceRoot field in generated source maps`,
			Args: []model.Arg{{
				Name: "URL",
			}},
		}, {
			Name:        []string{"--sourcefile"},
			Description: `Set the source file for the source map if there's no file name to use`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--sourcemap"},
			Description: `Generate source maps?`,
			Args: []model.Arg{{
				Name: "options",
				Suggestions: []model.Suggestion{{
					Name:        []string{`external`},
					Description: `Generate a separate .map.js file with no comment`,
				}, {
					Name:        []string{`inline`},
					Description: `Append source maps to JS files`,
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--sources-content"},
			Description: `Omit the sourcesContent field in generated source maps`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`false`}}},
			}},
		}, {
			Name:        []string{"--supported"},
			Description: `Consider a given syntax to be supported`,
			Args: []model.Arg{{
				Name:      "syntax=status",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--tree-shaking"},
			Description: `Force tree shaking on or off`,
			Args: []model.Arg{{
				Name:        "status",
				Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
			}},
		}, {
			Name:        []string{"--tsconfig"},
			Description: `Use this TypeScript config instead of the default`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--version"},
			Description: `Print the current version and exit`,
		}},
	}
}
