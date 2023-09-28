// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["fin"] = model.Subcommand{
		Name:        []string{"fin"},
		Description: `Docksal command line utility`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Shows help`,
		}, {
			Name:        []string{"addon"},
			Description: `Addons management commands: install, remove (fin help addon)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"install"},
				Description: `Install addon`,
				Args: []model.Arg{{
					Name:        "Name",
					Description: `See available addons in the Addons Repository https://github.com/docksal/addons`,
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove addon`,
				Args: []model.Arg{{
					Name: "Name",
				}},
			}},
		}, {
			Name:        []string{"alias"},
			Description: `Manage aliases that allow fin @alias execution (fin help alias). Create/update alias with <alias_name> that links to <path>`,
			Args: []model.Arg{{
				Name: "path",
			}, {
				Name: "alias_name",
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `Show aliases list`,
			}, {
				Name:        []string{"remove"},
				Description: `Remove alias`,
				Args: []model.Arg{{
					Name:      "alias_name",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"db"},
			Description: `Manage databases (fin help db)`,
			Options: []model.Option{{
				Name:        []string{"--db"},
				Description: `Use another database (default is the one set with 'MYSQL_DATABASE')`,
				Args: []model.Arg{{
					Name: "database",
				}},
			}, {
				Name:        []string{"--db-user"},
				Description: `Use another mysql username (default is 'root')`,
				Args: []model.Arg{{
					Name: "user",
				}},
			}, {
				Name:        []string{"--db-password"},
				Description: `Use another database password (default is the one set with 'MYSQL_ROOT_PASSWORD', see fin config)`,
				Args: []model.Arg{{
					Name: "password",
				}},
			}, {
				Name:        []string{"--db-charset"},
				Description: `Override charset when creating a database (default is utf8)`,
				Args: []model.Arg{{
					Name: "charset",
				}},
			}, {
				Name:        []string{"--db-collation"},
				Description: `Override collation when creating a database (default is utf8_general_ci)`,
				Args: []model.Arg{{
					Name: "collation",
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"import"},
				Description: `Truncate the database and import from SQL dump file or stdin`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
				Options: []model.Option{{
					Name:        []string{"--progress"},
					Description: `Show import progess (requires pv)`,
				}, {
					Name:        []string{"--no-truncate"},
					Description: `Do no truncate database before import`,
				}},
			}, {
				Name:        []string{"dump"},
				Description: `Dump a database into an SQL dump file or stdout`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `Show list of existing databases`,
			}, {
				Name:        []string{"cli"},
				Description: `Open command line interface to the DB server (and execute query if provided)`,
				Args: []model.Arg{{
					Name:       "query",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a database`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"drop"},
				Description: `Delete a database`,
				Args: []model.Arg{{
					Name:      "name",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"truncate"},
				Description: `Truncate a database (defaults to the "default")`,
				Args: []model.Arg{{
					Name:       "name",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"hosts"},
			Description: `Hosts file commands: add, remove, list (fin help hosts)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add"},
				Description: `Add hostname to hosts file. If none provided uses VIRTUAL_HOST`,
				Args: []model.Arg{{
					Name: "hostname",
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove lines containing hostname from hosts file. If none provided uses VIRTUAL_HOST`,
				Args: []model.Arg{{
					Name:      "hostname",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"list"},
				Description: `Output hosts file`,
			}},
		}, {
			Name:        []string{"project"},
			Description: `Manage project(s) (fin help project)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Start project services (alias: fin start)`,
			}, {
				Name:        []string{"up"},
				Description: `Configuration re-read and start project services (alias: fin up)`,
			}, {
				Name:        []string{"stop"},
				Description: `Stop all or specified project services (alias: fin stop)`,
				Args: []model.Arg{{
					Name:       "service",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--all", "-a"},
					Description: `Stop all services on all Docksal projects`,
				}},
			}, {
				Name:        []string{"status"},
				Description: `List project services (alias: fin ps)`,
			}, {
				Name:        []string{"restart"},
				Description: `Restart project services (alias: fin restart)`,
			}, {
				Name:        []string{"reset"},
				Description: `Recreate all or specified project services, their containers and volumes`,
				Args: []model.Arg{{
					Name:      "service",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"remove", "rm"},
				Description: `Remove all project services, networks and all their volumes, or specified services only`,
				Args: []model.Arg{{
					Name:      "service",
					Generator: nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"--force", "-f"},
					Description: `Do not ask for confirmation when deleting all project services`,
				}},
			}, {
				Name:        []string{"list"},
				Description: `List running Docksal projects (alias: fin pl)`,
				Options: []model.Option{{
					Name:        []string{"--all", "-a"},
					Description: `List all Docksal projects (stopped as well)`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a new project with a pre-configured boilerplate: Drupal, Wordpress, Magento, Laravel, Backdrop, Hugo, Gatsby, and others`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Provide project name upfront`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--choice"},
					Description: `Provide software choice number upfront`,
					Args: []model.Arg{{
						Name: "#",
					}},
				}, {
					Name:        []string{"--repo"},
					Description: `Clone from a custom repo: name (--choice is set to '0' automatically)`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--branch"},
					Description: `Clone from a custom repo: branch name (optional)`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--yes", "-y"},
					Description: `Avoid confirmation`,
				}},
			}, {
				Name:        []string{"config"},
				Description: `Show project configuration`,
			}, {
				Name:        []string{"build"},
				Description: `Build or rebuild services (alias for 'docker-compose build')`,
			}},
		}, {
			Name:        []string{"ssh-key"},
			Description: `Manage SSH keys (fin help ssh-key)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add"},
				Description: `Add a private SSH key from $HOME/.ssh by file name. Adds all default keys (id_rsa/id_dsa/id_ecdsa/id_ed25519) if no file name is given`,
				Args: []model.Arg{{
					Name:       "key-name",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--quiet"},
					Description: `Suppress key already loaded notifications`,
				}},
			}, {
				Name:        []string{"ls"},
				Description: `List SSH keys loaded in the docksal-ssh-agent`,
			}, {
				Name:        []string{"rm"},
				Description: `Remove all keys from the docksal-ssh-agent`,
			}, {
				Name:        []string{"new"},
				Description: `Generate a new SSH key pair`,
				Args: []model.Arg{{
					Name: "key-name",
				}},
			}},
		}, {
			Name:        []string{"system"},
			Description: `Manage Docksal state (fin help system)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"reset"},
				Description: `Reset Docksal`,
			}, {
				Name:        []string{"start"},
				Description: `Start Docksal`,
			}, {
				Name:        []string{"stop"},
				Description: `Stop Docksal`,
			}, {
				Name:        []string{"status"},
				Description: `Check Docksal status`,
			}},
		}, {
			Name:        []string{"vm"},
			Description: `Manage Docksal VM (fin help vm)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Start the machine (create if needed)`,
			}, {
				Name:        []string{"stop"},
				Description: `Stop the machine`,
			}, {
				Name:        []string{"kill"},
				Description: `Forcibly stop the machine`,
			}, {
				Name:        []string{"restart"},
				Description: `Restart the machine`,
			}, {
				Name:        []string{"status"},
				Description: `Get the status`,
			}, {
				Name:        []string{"ssh"},
				Description: `Log into ssh or run a command via ssh`,
				Args: []model.Arg{{
					Name: "command",
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove Docksal machine and cleanup after it`,
			}, {
				Name:        []string{"ip"},
				Description: `Show the machine IP address`,
			}, {
				Name:        []string{"ls"},
				Description: `List all docker machines`,
			}, {
				Name:        []string{"env"},
				Description: `Display the commands to set up the shell for direct use of Docker client`,
			}, {
				Name:        []string{"mount"},
				Description: `Try remounting host filesystem (NFS on macOS, SMB on Windows)`,
			}, {
				Name:        []string{"ram"},
				Description: `Show memory size`,
				Args: []model.Arg{{
					Name:        "megabyte",
					Description: `Set memory size. Default is 1024 (requires vm restart)`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"hdd"},
				Description: `Show disk size and usage`,
			}, {
				Name:        []string{"stats"},
				Description: `Show CPU and network usage`,
			}, {
				Name:        []string{"regenerate-certs"},
				Description: `Regenerate TLS certificates and restart the VM`,
			}},
		}, {
			Name:        []string{"bash"},
			Description: `Open shell into service's container. Defaults to cli`,
			Args: []model.Arg{{
				Name:      "service",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"config"},
			Description: `Show or change configuration (fin help config)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"show"},
				Description: `Display configuration for the current project`,
				Options: []model.Option{{
					Name:        []string{"--show-secrets"},
					Description: `Do not truncate value of SECRET_* environment vars`,
				}},
			}, {
				Name:        []string{"env"},
				Description: `Display only environment variables section`,
			}, {
				Name:        []string{"yml"},
				Description: `Display static YML project config suitable for export (NOTE: SECRET_* values will not be hidden)`,
			}, {
				Name:        []string{"generate"},
				Description: `Generate empty Docksal configuration for the project`,
				Options: []model.Option{{
					Name:        []string{"--stack"},
					Description: `Set non-default DOCKSAL_STACK during config generate`,
					Args: []model.Arg{{
						Name: "stack",
					}},
				}, {
					Name:        []string{"--docroot"},
					Description: `Set non-default DOCROOT during config generate`,
					Args: []model.Arg{{
						Name: "directory",
					}},
				}},
			}, {
				Name:        []string{"set"},
				Description: `Set value(s) for the variable(s) in project ENV file`,
				Args: []model.Arg{{
					Name:       "VAR=VAL",
					IsVariadic: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--global"},
					Description: `Set for global ENV file`,
				}, {
					Name:        []string{"--env"},
					Description: `Set in environment specific project ENV file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "name",
					}},
				}},
			}, {
				Name:        []string{"remove", "rm"},
				Description: `Remove variable(s) from project ENV file`,
				Args: []model.Arg{{
					Name:       "VAR",
					IsVariadic: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--global"},
					Description: `Remove from global ENV file`,
				}, {
					Name:        []string{"--env"},
					Description: `Remove from environment specific project ENV file`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get the value of the single variable from project ENV file`,
				Args: []model.Arg{{
					Name: "VAR",
				}},
				Options: []model.Option{{
					Name:        []string{"--global"},
					Description: `Get value from global ENV file`,
				}, {
					Name:        []string{"--env"},
					Description: `Get value from environment specific project ENV file`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "name",
					}},
				}},
			}},
		}, {
			Name:        []string{"exec"},
			Description: `Execute a command or a script in cli`,
			Args: []model.Arg{{
				Name:      "command|file",
				IsCommand: true,
			}},
		}, {
			Name:        []string{"exec-url"},
			Description: `Download script from URL and run it on host (URL should be public)`,
			Args: []model.Arg{{
				Name: "url",
			}},
		}, {
			Name:        []string{"init"},
			Description: `Initialize a project (override it with your own automation, see fin help init)`,
		}, {
			Name:        []string{"image"},
			Description: `Image management commands: registry, save, load (fin help image)`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"registry"},
				Description: `Show all Docksal images on Docker Hub`,
				Args: []model.Arg{{
					Name:        "image name",
					Description: `Show all tags for a certain image`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"save"},
				Description: `Save docker images into a tar archive`,
				Options: []model.Option{{
					Name:        []string{"--system"},
					Description: `Save Docksal system images`,
					ExclusiveOn: []string{"--project", "--all"},
				}, {
					Name:        []string{"--project"},
					Description: `Save current project's images`,
					ExclusiveOn: []string{"--system", "--all"},
				}, {
					Name:        []string{"--all"},
					Description: `Save all images available on the host`,
					ExclusiveOn: []string{"--system", "--project"},
				}},
			}, {
				Name:        []string{"load"},
				Description: `Load docker images from a tar archive`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}},
		}, {
			Name:        []string{"logs"},
			Description: `Show service logs (e.g., Apache logs, MySQL logs) and Unison logs (fin help logs)`,
			Args: []model.Arg{{
				Name:      "service",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--follow"},
				Description: `Follow log output`,
			}, {
				Name:        []string{"--no-color"},
				Description: `Produce monochrome output`,
			}, {
				Name:        []string{"--no-log-prefix"},
				Description: `Don't print prefix in logs`,
			}, {
				Name:        []string{"--since"},
				Description: `Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--tail"},
				Description: `Number of lines to show from the end of the logs for each container. (default "all")`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-t", "--timestamps"},
				Description: `Show timestamps`,
			}, {
				Name:        []string{"--until"},
				Description: `Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}},
		}, {
			Name:        []string{"run-cli", "rc"},
			Description: `Run a command in a standalone cli container in the current directory (fin help run-cli)`,
			Args: []model.Arg{{
				Name: "command",
			}},
			Options: []model.Option{{
				Name:        []string{"--clean"},
				Description: `Run command with a non-persistent $HOME directory`,
			}, {
				Name:        []string{"--clenup"},
				Description: `Clean the persistent $HOME directory and run command`,
			}, {
				Name:        []string{"--debug"},
				Description: `Print container debug output`,
			}, {
				Name:        []string{"--image"},
				Description: `Override default container image`,
				Args: []model.Arg{{
					Name: "IMAGE",
				}},
			}, {
				Name:        []string{"-e"},
				Description: `Pass environment variable(s) to the container`,
				Args: []model.Arg{{
					Name: "VAR=VALUE",
				}},
			}, {
				Name:        []string{"-T"},
				Description: `Disable pseudo-tty allocation (useful to get clean stdout)`,
			}},
		}, {
			Name:        []string{"share"},
			Description: `Create temporary public url for current project using ngrok`,
			Options: []model.Option{{
				Name:        []string{"--host"},
				Description: `Override a hostname for ngrok to route to (default is $VIRTUAL_HOST)`,
				Args: []model.Arg{{
					Name: "host",
				}},
			}},
		}, {
			Name:        []string{"share-v2"},
			Description: `Create a temporary public URL for the project using Cloudflare Tunnel`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Start tunnel and print public URL`,
			}, {
				Name:        []string{"stop"},
				Description: `Stop tunnel`,
			}, {
				Name:        []string{"status"},
				Description: `Prints tunnel status and public URL (if Active)`,
			}, {
				Name:        []string{"url"},
				Description: `Prints tunnel public URL`,
			}, {
				Name:        []string{"logs"},
				Description: `Prints tunnel container logs`,
			}},
		}, {
			Name:        []string{"vhosts"},
			Description: `List all virtual *.docksal hosts registered in Docksal proxy`,
		}, {
			Name:        []string{"docker", "d"},
			Description: `Run Docker commands directly`,
		}, {
			Name:        []string{"docker-compose", "dc"},
			Description: `Run Docker Compose commands directly`,
		}, {
			Name:        []string{"docker-machine", "dm"},
			Description: `Run Docker Machine commands directly`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"composer"},
			Description: `Run Composer commands`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"drush"},
			Description: `Drush command (requires Drupal)`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"drupal"},
			Description: `Drupal Console command (requires Drupal 8)`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"platform"},
			Description: `Platform.sh's CLI (requires docksal/cli 2.3+)`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"terminus"},
			Description: `Pantheon's Terminus (requires docksal/cli 2.1+)`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"wp"},
			Description: `WordPress CLI command (requires WordPress)`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"cleanup"},
			Description: `Remove all unused Docker images, unused Docksal volumes and containers`,
			Options: []model.Option{{
				Name:        []string{"--images"},
				Description: `Docker images cleanup Wizard`,
			}, {
				Name:        []string{"--hard"},
				Description: `Remove ALL stopped containers even unrelated to Docksal (potentially destructive operation)`,
			}},
		}, {
			Name:        []string{"diagnose"},
			Description: `Show diagnostic information for troubleshooting and bug reporting`,
		}, {
			Name:        []string{"sysinfo"},
			Description: `Show system information`,
		}, {
			Name:        []string{"update"},
			Description: `Update Docksal`,
			Options: []model.Option{{
				Name:        []string{"DOCKSAL_VERSION"},
				Description: `Update Docksal to the latest development version`,
				Args: []model.Arg{{
					Name: "develop",
				}},
			}},
		}, {
			Name:        []string{"version", "--version", "v", "-v"},
			Description: `Print fin version. [v, -v] prints short version`,
		}},
	}
}
