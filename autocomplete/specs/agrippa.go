// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["agrippa"] = model.Subcommand{
		Name:        []string{"agrippa"},
		Description: ``,
		Args: []model.Arg{{
			Name: "component's name",
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Generates a basic .agripparc.json file, with some default values for options that agrippa accepts`,
		}, {
			Name:        []string{"generate", "gen"},
			Description: `This is the core of the CLI - this command generates a new React component, based on the name and options passed to it`,
			Args: []model.Arg{{
				Name: "component",
			}},
			Options: []model.Option{{
				Name:        []string{"--props"},
				Description: `Determines which prop declaration/validation solution to generate. Defaults to "ts" if "typescript" flag is "true", "none" otherwise`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{
						Name:        []string{`ts`},
						Description: `Generate a TypeScript interface for props (this requires the typescript option to be true)`,
					}, {
						Name:        []string{`jsdoc`},
						Description: `Generates JSDoc type hints for props (a @typedef and a @type)`,
					}, {
						Name:        []string{`prop-types`},
						Description: `Generates prop-types for props`,
					}, {
						Name:        []string{`none`},
						Description: `Generates no props`,
					}},
				}},
			}, {
				Name:        []string{"--children"},
				Description: `Whether the component is meant to have children or not. Defaults to "false"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--typescript", "--ts"},
				Description: `Option to generate Typescript components, defaults to "true" if the CLI manages to find a "tsconfig.json" file, false otherwise`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--import-react"},
				Description: `Option to import React at the top of the component. Defaults to "true" if the CLI manages to find a "tsconfig.json" file and it has a "jsx" field under "compilerOptions" with "react-jsx" or "react-jsxdev" as the value, "false" otherwise`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--export-type"},
				Description: `Whether to export the component as a named export or a default export. Defaults to "named"`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{
						Name: []string{`named`},
					}, {
						Name: []string{`default`},
					}},
				}},
			}, {
				Name:        []string{"--declaration"},
				Description: `Whether to declare the component as a "const" with an arrow function, or as a "function" declaration. Defaults to "const"`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{
						Name: []string{`const`},
					}, {
						Name: []string{`function`},
					}},
				}},
			}, {
				Name:        []string{"--memo"},
				Description: `Generates a component as a memo component. Defaults to "false"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--react-native"},
				Description: `Generates a React Native component. Defaults to "true" if "react-native" is a dependency in the project's "package.json", "false" otherwise`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--styling"},
				Description: `Which styling solution to generate. Defaults to "none"`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{
						Name:        []string{`none`},
						Description: `Generates no style`,
					}, {
						Name:        []string{`css`},
						Description: `Generates a "css" file`,
					}, {
						Name:        []string{`scss`},
						Description: `Generates a "scss" file`,
					}, {
						Name:        []string{`jss`},
						Description: `Generates a "useStyles" hook above the component with "react-jss"'s "createUseStyles"`,
					}, {
						Name:        []string{`mui`},
						Description: `Generates a "useStyles" hook above the component with "material-ui"'s "makeStyles". Note that this generates styles for "material-ui" v4`,
					}, {
						Name:        []string{`react-native`},
						Description: `Generates a React Native "StyleSheet"`,
					}},
				}},
			}, {
				Name:        []string{"--styling-module"},
				Description: `Whether to generate a scoped "module" stylesheet. This option is only relevant for "css" or "scss" styling, and will be ignored for other values. Defaults to "true"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--base-dir"},
				Description: `Path to a base directory which components should be generated inside of (directly or nested). Defaults to the current working directory`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "path",
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--dest", "--destination"},
				Description: `Allows the user to generate the component at a path relative to "baseDir". It's mostly used to augment or modify the generation path in the command-line when creating the component itself. "dest" is resolved relative to "baseDir"`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "path",
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--allow-outside-base"},
				Description: `Overrides default Agrippa behaviour to allow components to be specified outside "baseDir". Defaults to current working directory`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "path",
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--flat"},
				Description: `By default, Agrippa generates a dedicated directory for the generated component; this directory has the same name as the component, and will contain the component file (as "index.tsx" or "index.jsx"), and the style file (if one is generated). Defaults to "false"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--separate-index"},
				Description: `In the "separateIndex" generation scheme, Agrippa creates one file for the component's code, and one "index" file, whose purpose is to allow importing exports from the component's directory elegantly (the import path can then be "<path>/<to>/<directory>/Component" instead of "<path>/<to>/<directory>/Component/Component"). Defaults to "true"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--post-command", "--postCommand"},
				Description: `Allows a custom command to be run after having generated a component. This allows for some cool functionality, such as opening an IDE at the newly generated files automatically or linting them. Defaults to "none"`,
				Args: []model.Arg{{
					IsCommand: true,
				}},
			}, {
				Name:        []string{"--overwrite"},
				Description: `By default, Agrippa prevents the generation of a component over an existing components, to prevent a loss of code. "--overwrite" tells Agrippa to replace existing files, if any are found. Defaults to "false"`,
				Args:        []model.Arg{{}},
			}, {
				Name:        []string{"--debug"},
				Description: `The "debug" flag makes Agrippa print out additional debug information. It's quite useful when debugging. Defaults to "false"`,
				Args:        []model.Arg{{}},
			}},
		}},
	}
}
