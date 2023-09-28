// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["snaplet"] = model.Subcommand{
		Name:        []string{"snaplet"},
		Description: `Create and share PostgreSQL snapshots with schema, data transformation, and preview databases for collaborative development (see: https://docs.snaplet.dev)`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Show help for snaplet`,
			IsPersistent: true,
		}, {
			Name:         []string{"--version", "-v"},
			Description:  `Show version number`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"auth"},
			Description: `Manage auth state`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"login", "setup"},
				Description: `Login with an access token`,
				Args: []model.Arg{{
					Name:       "access-token",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"config"},
			Description: `Manage configuration`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"generate"},
				Description: `Generate transform files`,
				Options: []model.Option{{
					Name: []string{"--dry-run"},
				}, {
					Name: []string{"--type", "-t"},
					Args: []model.Arg{{
						Suggestions: []model.Suggestion{{Name: []string{`typedefs`}}, {Name: []string{`transform`}}, {Name: []string{`keys`}}, {Name: []string{`seed`}}},
					}},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List config variables`,
			}, {
				Name:        []string{"pull"},
				Description: `Pull cloud project config to local`,
				Options: []model.Option{{
					Name: []string{"--type", "-t"},
					Args: []model.Arg{{
						Suggestions: []model.Suggestion{{Name: []string{`typedefs`}}, {Name: []string{`transform`}}, {Name: []string{`keys`}}, {Name: []string{`seed`}}},
					}},
				}},
			}, {
				Name:        []string{"push"},
				Description: `Push local project config to cloud`,
				Options: []model.Option{{
					Name: []string{"--type", "-t"},
					Args: []model.Arg{{
						Suggestions: []model.Suggestion{{Name: []string{`typedefs`}}, {Name: []string{`transform`}}, {Name: []string{`keys`}}, {Name: []string{`seed`}}},
					}},
				}},
			}, {
				Name:        []string{"setup"},
				Description: `Setup local project configuration`,
				Args: []model.Arg{{
					Name:       "targetDatabaseUrl",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name: []string{"--generate"},
				}, {
					Name: []string{"--no-generate"},
				}, {
					Name:        []string{"--yes"},
					Description: `Automatically accept any prompt, useful for scripts`,
				}},
			}},
		}, {
			Name:        []string{"database", "db"},
			Description: `Manage preview database`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"cache", "ca"},
				Description: `Cache a snapshot into the preview database server`,
				Args: []model.Arg{{
					Name:       "snapshot",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--tags"},
					Description: `Filter snapshots by tags`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}, {
					Name:        []string{"--latest"},
					Description: `Select the latest snapshot captured to cache in the preview database server`,
				}, {
					Name:        []string{"--clear"},
					Description: `Remove the snapshot from the preview database server cache`,
				}},
			}, {
				Name:        []string{"create", "c"},
				Description: `Create a preview database from a snapshot`,
				Args: []model.Arg{{
					Name: "database-name",
				}, {
					Name:       "snapshot",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--tags"},
					Description: `Filter snapshots by tags`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}, {
					Name:        []string{"--latest"},
					Description: `Restore the latest snapshot`,
				}, {
					Name:        []string{"--git"},
					Description: `Infer the database name from the current git branch`,
				}, {
					Name:        []string{"--no-snapshot"},
					Description: `Create an empty database not based on a snapshot`,
				}, {
					Name:        []string{"--force"},
					Description: `Force the database creation, it will drop and recreate the existing database if present`,
				}},
			}, {
				Name:        []string{"drop", "delete", "d"},
				Description: `Drop a preview database`,
				Args: []model.Arg{{
					Name:       "database-name",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--git"},
					Description: `Infer the database name from the current git branch`,
				}},
			}, {
				Name:        []string{"url", "u"},
				Description: `Show a preview database url`,
				Args: []model.Arg{{
					Name:       "database-name",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--git"},
					Description: `Infer the database name from the current git branch`,
				}},
			}, {
				Name:        []string{"destroy", "ds"},
				Description: `Destroy the database server`,
				Options: []model.Option{{
					Name:        []string{"--yes"},
					Description: `Automatically accept any prompt, useful for scripts`,
				}},
			}, {
				Name:        []string{"setup", "s"},
				Description: `Create a preview database server`,
				Options: []model.Option{{
					Name:        []string{"--region"},
					Description: `Fly.io region where the preview database server should be created`,
					Args: []model.Arg{{
						Suggestions: []model.Suggestion{{
							Name:        []string{`ams`},
							Description: `Amsterdam, Netherlands`,
						}, {
							Name:        []string{`cdg`},
							Description: `Paris, France`,
						}, {
							Name:        []string{`den`},
							Description: `Denver, Colorado (US)`,
						}, {
							Name:        []string{`dfw`},
							Description: `Dallas, Texas (US)`,
						}, {
							Name:        []string{`ewr`},
							Description: `Secaucus, NJ (US)`,
						}, {
							Name:        []string{`fra`},
							Description: `Frankfurt, Germany`,
						}, {
							Name:        []string{`gru`},
							Description: `São Paulo`,
						}, {
							Name:        []string{`hkg`},
							Description: `Hong Kong, Hong Kong`,
						}, {
							Name:        []string{`iad`},
							Description: `Ashburn, Virginia (US)`,
						}, {
							Name:        []string{`jnb`},
							Description: `Johannesburg, South Africa`,
						}, {
							Name:        []string{`lax`},
							Description: `Los Angeles, California (US)`,
						}, {
							Name:        []string{`lhr`},
							Description: `London, United Kingdom`,
						}, {
							Name:        []string{`maa`},
							Description: `Chennai (Madras), India`,
						}, {
							Name:        []string{`mad`},
							Description: `Madrid, Spain`,
						}, {
							Name:        []string{`mia`},
							Description: `Miami, Florida (US)`,
						}, {
							Name:        []string{`nrt`},
							Description: `Tokyo, Japan`,
						}, {
							Name:        []string{`ord`},
							Description: `Chicago, Illinois (US)`,
						}, {
							Name:        []string{`otp`},
							Description: `Bucharest, Romania`,
						}, {
							Name:        []string{`scl`},
							Description: `Santiago, Chile`,
						}, {
							Name:        []string{`sea`},
							Description: `Seattle, Washington (US)`,
						}, {
							Name:        []string{`sin`},
							Description: `Singapore`,
						}, {
							Name:        []string{`sjc`},
							Description: `Sunnyvale, California (US)`,
						}, {
							Name:        []string{`syd`},
							Description: `Sydney, Australia`,
						}, {
							Name:        []string{`waw`},
							Description: `Warsaw, Poland`,
						}, {
							Name:        []string{`yul`},
							Description: `Montreal, Canada`,
						}, {
							Name:        []string{`yyz`},
							Description: `Toronto, Canada`,
						}},
					}},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List preview databases`,
			}},
		}, {
			Name:        []string{"discord", "chat"},
			Description: `Open the Snaplet Discord chat window in your browser`,
		}, {
			Name:        []string{"documentation", "docs"},
			Description: `Opens the Snaplet Documentation in your browser`,
		}, {
			Name:        []string{"project"},
			Description: `Manage project configuration`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new project`,
				Args: []model.Arg{{
					Name: "name",
				}},
				Options: []model.Option{{
					Name: []string{"--team", "-t"},
					Args: []model.Arg{{
						Name: "team-name",
					}},
				}},
			}, {
				Name:        []string{"invite"},
				Description: `Create an invite URL for this project`,
			}, {
				Name:        []string{"setup"},
				Description: `Set up a project`,
				Args: []model.Arg{{
					Name:       "project-id",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"proxy", "dev"},
			Description: `Start a database proxy`,
		}, {
			Name:        []string{"seed", "gen", "generate"},
			Description: `Populates an empty database with generated data`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Show interactive prompts to customise seeding`,
			}, {
				Name: []string{"-n", "--entriesPerTable"},
				Args: []model.Arg{{
					Name: "number",
				}},
			}},
		}, {
			Name:        []string{"snapshot", "ss"},
			Description: `Manage snapshots`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"capture", "c"},
				Description: `Capture a new snapshot`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "destination-path",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--env", "--environment"},
					Description: `Environment to use when capturing the snapshot`,
					Args: []model.Arg{{
						Name:        "environment-name",
						Suggestions: []model.Suggestion{{Name: []string{`cloud`}}, {Name: []string{`local`}}},
					}},
				}, {
					Name:        []string{"-m", "--message"},
					Description: `Attach a message to the snapshot`,
					Args: []model.Arg{{
						Name: "message",
					}},
				}, {
					Name:        []string{"--subset-path", "--subset"},
					Description: `Path to a subset config file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "path",
					}},
				}, {
					Name:        []string{"--tags"},
					Description: `Attach tags to the snapshot`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}, {
					Name:        []string{"-t", "--transform-mode"},
					Description: `Transformation mode to apply to the snapshot`,
					Args: []model.Arg{{
						Name:        "transformation-mode",
						Suggestions: []model.Suggestion{{Name: []string{`strict`}}, {Name: []string{`unsafe`}}, {Name: []string{`auto`}}},
					}},
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a snapshot in cloud`,
				Options: []model.Option{{
					Name: []string{"--json"},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List all snapshots`,
				Options: []model.Option{{
					Name:        []string{"--tags"},
					Description: `Filter snapshots by tags`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}},
			}, {
				Name:        []string{"restore", "r"},
				Description: `Restore a snapshot`,
				Args: []model.Arg{{
					Name:       "snapshot-name",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"data", "--no-data"},
					Description: `Restore data on the database (skip with --no-data)`,
				}, {
					Name:        []string{"schema", "--no-schema"},
					Description: `Restore schema on the database (skip with --no-schema)`,
				}, {
					Name:        []string{"reset", "--no-reset"},
					Description: `Drop destination database before restoring schemas (skip with --no-reset)`,
				}, {
					Name:        []string{"--tags"},
					Description: `Filter snapshots by tags`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}, {
					Name:        []string{"--latest"},
					Description: `Restore the latest snapshot`,
				}, {
					Name:        []string{"-y", "--yes"},
					Description: `Performs a restore without a confirmation message`,
				}, {
					Name:        []string{"--schema-only"},
					Description: `Skip data, only restore schema`,
				}, {
					Name:        []string{"--data-only"},
					Description: `Restore data only (keep the current schema and indexes)`,
				}},
			}, {
				Name:        []string{"share", "upload"},
				Description: `Share a snapshot`,
				Args: []model.Arg{{
					Name:       "snapshot-name",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--no-encrypt"},
					Description: `Disable encryption`,
				}, {
					Name:        []string{"--tags"},
					Description: `Filter snapshots by tags`,
					Args: []model.Arg{{
						Name:        "tags-list",
						Description: `Tag1,tag2,tag3`,
					}},
				}, {
					Name:        []string{"--latest"},
					Description: `Share the latest snapshot`,
				}},
			}, {
				Name:        []string{"sample", "s"},
				Description: `Create a sample of a database`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "destination-path",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--env", "--environment"},
					Description: `Environment to use when slicing the snapshot`,
					Args: []model.Arg{{
						Name:        "environment-name",
						Suggestions: []model.Suggestion{{Name: []string{`cloud`}}, {Name: []string{`local`}}},
					}},
				}, {
					Name:        []string{"--sample-config-path", "--subset-path"},
					Description: `Path to a sample config file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
					}},
				}, {
					Name:        []string{"-f", "--force"},
					Description: `Force overwrite existing sample file`,
				}},
			}},
		}, {
			Name:        []string{"subset", "subsetting"},
			Description: `Manage subsetting`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"setup", "configure", "config"},
				Description: `Configure subsetting`,
			}},
		}, {
			Name:        []string{"team"},
			Description: `Manage team configuration`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new team`,
				Args: []model.Arg{{
					Name: "team-name",
				}},
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Upgrade this binary`,
		}, {
			Name:        []string{"completion"},
			Description: `Generate completion script`,
		}},
	}
}
