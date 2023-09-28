// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["typeorm"] = model.Subcommand{
		Name:        []string{"typeorm"},
		Description: `TypeORM CLI`,
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Show help for command`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Show the version`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"schema:sync"},
			Description: `Synchronizes your entities with database schema`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which schema synchronization needs to to run`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"schema:log"},
			Description: `Shows sql to be executed by schema:sync command`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection of which schema sync log should be shown`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"schema:drop"},
			Description: `Drops all tables in the database on your default connection`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to drop all tables`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"query"},
			Description: `Executes given SQL query on a default connection`,
			Args: []model.Arg{{
				Name:        "query",
				Description: `The SQL Query to run`,
			}},
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"entity:create"},
			Description: `Generates a new entity`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name of the entity class`,
				Args: []model.Arg{{
					Name: "entity",
				}},
			}, {
				Name:        []string{"-d", "--dir"},
				Description: `Directory where entity should be created`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "directory",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"subscriber:create"},
			Description: `Generates a new subscriber`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name of the subscriber class`,
				Args: []model.Arg{{
					Name: "subscriber",
				}},
			}, {
				Name:        []string{"-d", "--dir"},
				Description: `Directory where subscriber should be created`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "directory",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"migration:create"},
			Description: `Creates a new migration file`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name of the migration class`,
				Args: []model.Arg{{
					Name: "migration",
				}},
			}, {
				Name:        []string{"-d", "--dir"},
				Description: `Directory where migration should be created`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "directory",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-o", "--outputJs"},
				Description: `Generate a migration file on Javascript instead of Typescript`,
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"migration:generate"},
			Description: `Generates a new migration file with sql needs to be executed to update schema`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name of the migration class`,
				Args: []model.Arg{{
					Name: "migration",
				}},
			}, {
				Name:        []string{"-d", "--dir"},
				Description: `Directory where migration should be created`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "directory",
				}},
			}, {
				Name:        []string{"-p", "--pretty"},
				Description: `Pretty-print generated SQL`,
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-o", "--outputJs"},
				Description: `Generate a migration file on Javascript instead of Typescript`,
			}, {
				Name:        []string{"--dr", "--dryrun"},
				Description: `Prints out the contents of the migration instead of writing it to a file`,
			}, {
				Name:        []string{"--ch", "--check"},
				Description: `Verifies that the current database is up to date and that no migrations are needed`,
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"migration:run"},
			Description: `Runs all pending migrations`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-t", "--transaction"},
				Description: `Indicates if transaction should be used or not for migration run`,
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"migration:show"},
			Description: `Show all migrations and whether they have been run or not`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"migration:revert"},
			Description: `Reverts last executed migration`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-t", "--transaction"},
				Description: `Indicates if transaction should be used or not for migration revert`,
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Prints TypeORM version this project uses`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"cache:clear"},
			Description: `Clears all data stored in query runner cache`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-f", "--config"},
				Description: `Name of the file with connection configuration`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}, {
			Name:        []string{"init"},
			Description: `Generates initial TypeORM project structure`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help for command`,
			}, {
				Name:        []string{"-c", "--connection"},
				Description: `Name of the connection on which to run a query`,
				Args: []model.Arg{{
					Name: "connection",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name of the project directory`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"--db", "--database"},
				Description: `Database type you'll use in your project`,
				Args: []model.Arg{{
					Name:        "database",
					Suggestions: []model.Suggestion{{Name: []string{`mysql`}}, {Name: []string{`mariadb`}}, {Name: []string{`postgres`}}, {Name: []string{`cockroachdb`}}, {Name: []string{`sqlite`}}, {Name: []string{`mssql`}}, {Name: []string{`oracle`}}, {Name: []string{`cordova`}}, {Name: []string{`nativescript`}}, {Name: []string{`react-native`}}, {Name: []string{`expo`}}, {Name: []string{`mongodb`}}},
				}},
			}, {
				Name:        []string{"--express"},
				Description: `Indicates if express should be included in the project`,
			}, {
				Name:        []string{"--docker"},
				Description: `Set to true if docker-compose must be generated as well`,
			}, {
				Name:        []string{"--pm", "--manager"},
				Description: `Install packages`,
				Args: []model.Arg{{
					Name:        "manager",
					Suggestions: []model.Suggestion{{Name: []string{`npm`}}, {Name: []string{`yarn`}}},
				}},
			}, {
				Name:        []string{"-v", "--version"},
				Description: `Show the version`,
			}},
		}},
	}
}
