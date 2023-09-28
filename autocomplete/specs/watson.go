// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["watson"] = model.Subcommand{
		Name:        []string{"watson"},
		Description: `A wonderful CLI to track your time`,
		Options: []model.Option{{
			Name:         []string{"--help"},
			Description:  `Show help for watson`,
			IsPersistent: true,
		}, {
			Name:        []string{"--version"},
			Description: `Show the version for watson`,
		}, {
			Name:        []string{"--color"},
			Description: `Color output`,
			ExclusiveOn: []string{"--no-color"},
		}, {
			Name:        []string{"--no-color"},
			Description: `No color output`,
			ExclusiveOn: []string{"--color"},
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"add"},
			Description: `Add time to a project with tag(s) that was not tracked live`,
			Args: []model.Arg{{
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--from", "-f"},
				Description: `Date and time of start of tracked activity`,
				Args: []model.Arg{{
					Name:        "DATETIME",
					Description: `The start time. Ex: 2022-03-04 12:00:00`,
				}},
			}, {
				Name:        []string{"--to", "-t"},
				Description: `Date and time of end of tracked activity`,
				Args: []model.Arg{{
					Name:        "DATETIME",
					Description: `The start time. Ex: 2022-03-04 12:00:00`,
				}},
			}, {
				Name:        []string{"--confirm-new-project", "-c"},
				Description: `Confirm creation of new tag`,
			}},
		}, {
			Name:        []string{"aggregate"},
			Description: `Display a report of the time spent on each project aggregated by day`,
			Options: []model.Option{{
				Name:        []string{"--current", "-c"},
				Description: `Include currently running frame in report`,
				ExclusiveOn: []string{"-C"},
			}, {
				Name:        []string{"--no-current", "-C"},
				Description: `Don't include currently running frame in report`,
				ExclusiveOn: []string{"--current"},
			}, {
				Name:        []string{"--from", "-f"},
				Description: `The date from when the reports should start. Defaults to seven days ago`,
				Args: []model.Arg{{
					Name:        "DATETIME",
					Description: `The start time. Ex: 2022-03-04 12:00:00`,
				}},
			}, {
				Name:        []string{"--to", "-t"},
				Description: `The date at which the report should stop (inclusive). Defaults to tomorrow`,
				Args: []model.Arg{{
					Name:        "DATETIME",
					Description: `The start time. Ex: 2022-03-04 12:00:00`,
				}},
			}, {
				Name:        []string{"--project", "-p"},
				Description: `Reports activity only for the given project. You can add other projects by using this option several times`,
				Args: []model.Arg{{
					Name:      "PROJECT",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--tag", "-T"},
				Description: `Reports activity only for frames containing the given tag. You can add several tags by using this option multiple times`,
				Args: []model.Arg{{
					Name:      "TAG",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--json", "-j"},
				Description: `Format output in JSON instead of plain text`,
				ExclusiveOn: []string{"--csv", "--pager", "--no-pager"},
			}, {
				Name:        []string{"--csv", "-s"},
				Description: `Format output in CSV instead of plain text`,
				ExclusiveOn: []string{"--json", "--pager", "--no-pager"},
			}, {
				Name:        []string{"--pager", "-g"},
				Description: `View output through a pager`,
				ExclusiveOn: []string{"--csv", "--json", "--no-pager"},
			}, {
				Name:        []string{"--no-pager", "-G"},
				Description: `Don't view output through a pager`,
				ExclusiveOn: []string{"--csv", "--pager", "--json"},
			}},
		}, {
			Name:        []string{"cancel"},
			Description: `Cancel the last call to the start command. The time will not be recorded`,
		}, {
			Name:        []string{"config"},
			Description: `Get and set configuration options`,
			Args: []model.Arg{{
				Name: "SECTION.OPTION",
				Suggestions: []model.Suggestion{{
					Name:        []string{`backend.url`},
					Description: `This is the API root url of your repository, e.g. https://my.server.com/api/`,
				}, {
					Name:        []string{`backend.token`},
					Description: `To authenticate watson as an API client, once generated, you will need to set up your API token in your configuration, e.g. 7e329263e329`,
				}, {
					Name:        []string{`options.confirm_new_project`},
					Description: `If true, the user will be asked for confirmation before creating a new project. The option can be overriden in the command line with --confirm-new-project flag. Default: false`,
				}, {
					Name:        []string{`options.confirm_new_tag`},
					Description: `If true, the user will be asked for confirmation before creating a new tag. The option can be overriden in the command line with --confirm-new-tag flag. Default: false`,
				}, {
					Name:        []string{`options.date_format`},
					Description: `Globally configure how dates should be formatted. All python’s strftime directives are supported. Default: %Y.%m.%d`,
				}, {
					Name:        []string{`options.log_current`},
					Description: `If true, the output of the log command will include the currently running frame (if any) by default. The option can be overridden on the command line with the -c/-C resp. --current/--no-current flags. Default: false`,
				}, {
					Name:        []string{`options.pager`},
					Description: `If true, the output of the log and report command will be run through a pager by default. The option can be overridden on the command line with the -g/-G or --pager/--no-pager flags. If other commands output in colour, but log or report do not, try disabling the pager. Default: true`,
				}, {
					Name:        []string{`options.report_current`},
					Description: `If true, the output of the report command will include the currently running frame (if any) by default. The option can be overridden on the command line with the -c/-C resp. --current/--no-current flags. Default: false`,
				}, {
					Name:        []string{`options.reverse_log`},
					Description: `If true, the output of the log command will reverse the order of the days to display the latest day’s entries on top and the oldest day’s entries at the bottom. The option can be overridden on the command line with the -r/-R resp. --reverse/--no-reverse flags. Default: true`,
				}, {
					Name:        []string{`options.stop_on_start`},
					Description: `If true, starting a new project will stop running projects. Default false`,
				}, {
					Name:        []string{`options.stop_on_restart`},
					Description: `Similar to the options.stop_on_start option, but for the restart command. Default: false`,
				}, {
					Name:        []string{`options.time_format`},
					Description: `Globally configure how time should be formatted. All python’s strftime directives are supported. Default: %H.%M:%S%z`,
				}, {
					Name:        []string{`options.week_start`},
					Description: `Globally configure which day corresponds to the start of a week. Allowable values are monday, tuesday, wednesday, thursday, friday, saturday, and sunday`,
				}},
			}, {
				Name:       "VALUE",
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--edit", "-e"},
				Description: `Edit the configuration file with an editor`,
			}},
		}, {
			Name:        []string{"edit"},
			Description: `Edit a frame`,
			Args: []model.Arg{{
				Name:      "FRAME ID",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--confirm-new-project", "-c"},
				Description: `Confirm addition of new project`,
			}, {
				Name:        []string{"--confirm-new-tag", "-b"},
				Description: `Confirm creation of new tag`,
			}},
		}, {
			Name:        []string{"frames"},
			Description: `Display the list of all frame IDs`,
		}, {
			Name:        []string{"help"},
			Description: `Display help information`,
			Args: []model.Arg{{
				Name: "COMMAND",
			}},
		}, {
			Name:        []string{"log"},
			Description: `Display each recorded session during the given timespan`,
			Options: []model.Option{{
				Name:        []string{"--current", "-c"},
				Description: `Include currently running frame in report`,
				ExclusiveOn: []string{"--no-current"},
			}, {
				Name:        []string{"--no-current", "-C"},
				Description: `Don't include currently running frame in report`,
				ExclusiveOn: []string{"--current"},
			}, {
				Name:        []string{"--reverse", "-r"},
				Description: `Reverse the order of the days in output`,
				ExclusiveOn: []string{"--no-reverse"},
			}, {
				Name:        []string{"--no-reverse", "-R"},
				Description: `Don't Reverse the order of the days in output`,
				ExclusiveOn: []string{"--reverse"},
			}, {
				Name:        []string{"--from", "-f"},
				Description: `The date from when the log should start. Defaults to seven days ago`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--to", "-t"},
				Description: `The date at which the log should stop (inclusive). Defaults to tomorrow`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--year", "-y"},
				Description: `Reports activity for the current year`,
				ExclusiveOn: []string{"--month", "--luna", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--month", "-m"},
				Description: `Reports activity for the current month`,
				ExclusiveOn: []string{"--year", "--luna", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--luna", "-l"},
				Description: `Reports activity for the current moon cycle`,
				ExclusiveOn: []string{"--year", "--month", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--week", "-w"},
				Description: `Reports activity for the current week`,
				ExclusiveOn: []string{"--year", "--month", "--luna", "--day", "--all"},
			}, {
				Name:        []string{"--day", "-d"},
				Description: `Reports activity for the current day`,
				ExclusiveOn: []string{"--year", "--month", "--luna", "--week", "--all"},
			}, {
				Name:        []string{"--all", "-a"},
				Description: `Reports all activities`,
				ExclusiveOn: []string{"--year", "--month", "--luna", "--week", "--day"},
			}, {
				Name:        []string{"--project", "-p"},
				Description: `Logs activity only for the given project. You can add other projects by using this option several times`,
				Args: []model.Arg{{
					Name:       "PROJECT",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--tag", "-T"},
				Description: `Logs activity only for frames containing the given tag. You can add several tags by using this option multiple times`,
				Args: []model.Arg{{
					Name:       "TAG",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--ignore-project"},
				Description: `Logs activity for all projects but the given ones. You can ignore several projects by using the option multiple times. Any given project will be ignored`,
				Args: []model.Arg{{
					Name:       "PROJECT",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--ignore-tag"},
				Description: `Logs activity for all tags but the given ones. You can ignore several tags by using the option multiple times. Any given tag will be ignored`,
				Args: []model.Arg{{
					Name:       "TAG",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--json", "-j"},
				Description: `Format output in JSON instead of plain text`,
				ExclusiveOn: []string{"--csv"},
			}, {
				Name:        []string{"--csv", "-s"},
				Description: `Format output in CSV instead of plain text`,
				ExclusiveOn: []string{"--json"},
			}, {
				Name:        []string{"--pager", "-g"},
				Description: `View output through a pager`,
				ExclusiveOn: []string{"--no-pager"},
			}, {
				Name:        []string{"--no-pager", "-G"},
				Description: `Don't view output through a pager`,
				ExclusiveOn: []string{"--pager"},
			}},
		}, {
			Name:        []string{"merge"},
			Description: `Perform a merge of the existing frames with a conflicting frames file.  When storing the frames on a file hosting service, there is the possibility that the frame file goes out-of-sync due to one or more of the connected clients going offline. This can cause the frames to diverge`,
			Args: []model.Arg{{
				Name: "FRAMES_WITH_CONFLICT",
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `If specified, then the merge will automatically be performed`,
			}},
		}, {
			Name:        []string{"projects"},
			Description: `Display the list of all the existing projects`,
		}, {
			Name:        []string{"remove"},
			Description: `Remove a frame. You can specify the frame either by id or by position (ex: -1 for the last frame)`,
			Args: []model.Arg{{
				Name:      "ID",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `Don’t ask for confirmation`,
			}},
		}, {
			Name:        []string{"rename"},
			Description: `Rename a project or tag`,
			Args: []model.Arg{{
				Name:        "type",
				Description: `Project or tag`,
				Suggestions: []model.Suggestion{{
					Name: []string{`project`},
				}, {
					Name: []string{`tag`},
				}},
			}, {
				Name:      "OLD_NAME",
				Generator: nil, // TODO: port over generator
			}, {
				Name: "NEW_NAME",
			}},
		}, {
			Name:        []string{"report"},
			Description: `Display a report of the time spent on each project`,
			Options: []model.Option{{
				Name:        []string{"--current", "-c"},
				Description: `Include currently running frame in report`,
				ExclusiveOn: []string{"--no-current"},
			}, {
				Name:        []string{"--no-current", "-C"},
				Description: `Don't include currently running frame in report`,
				ExclusiveOn: []string{"--current"},
			}, {
				Name:        []string{"--from", "-f"},
				Description: `The date from when the report should start. Defaults to seven days ago`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--to", "-t"},
				Description: `The date at which the report should stop (inclusive). Defaults to tomorrow`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--year", "-y"},
				Description: `Reports activity for the current year`,
				ExclusiveOn: []string{"--month", "--luna", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--month", "-m"},
				Description: `Reports activity for the current month`,
				ExclusiveOn: []string{"--year", "--luna", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--luna", "-l"},
				Description: `Reports activity for the current moon cycle`,
				ExclusiveOn: []string{"--year", "--month", "--week", "--day", "--all"},
			}, {
				Name:        []string{"--week", "-w"},
				Description: `Reports activity for the current week`,
				ExclusiveOn: []string{"--year", "--month", "--luna", "--day", "--all"},
			}, {
				Name:        []string{"--day", "-d"},
				Description: `Reports activity for the current day`,
				ExclusiveOn: []string{"--year", "--month", "--luna", "--week", "--all"},
			}, {
				Name:        []string{"--project", "-p"},
				Description: `Reports activity only for the given project. You can add other projects by using this option several times`,
				Args: []model.Arg{{
					Name:      "PROJECT",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--tag", "-T"},
				Description: `Reports activity only for frames containing the given tag. You can add several tags by using this option multiple times`,
				Args: []model.Arg{{
					Name:      "TAG",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--ignore-project"},
				Description: `Reports activity for all projects but the given ones. You can ignore several projects by using the option multiple times. Any given project will be ignored`,
				Args: []model.Arg{{
					Name:      "PROJECT",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--ignore-tag"},
				Description: `Reports activity for all tags but the given ones. You can ignore several tags by using the option multiple times. Any given tag will be ignored`,
				Args: []model.Arg{{
					Name:      "TAG",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--json", "-j"},
				Description: `Format output in JSON instead of plain text`,
				ExclusiveOn: []string{"--csv"},
			}, {
				Name:        []string{"--csv", "-s"},
				Description: `Format output in CSV instead of plain text`,
				ExclusiveOn: []string{"--json"},
			}, {
				Name:        []string{"--pager", "-g"},
				Description: `View output through a pager`,
				ExclusiveOn: []string{"--no-pager"},
			}, {
				Name:        []string{"--no-pager", "-G"},
				Description: `Don't view output through a pager`,
				ExclusiveOn: []string{"--pager"},
			}},
		}, {
			Name:        []string{"restart"},
			Description: `Restart monitoring time for a previously stopped project`,
			Args: []model.Arg{{
				Name:       "FRAME",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--at"},
				Description: `Start frame at this time. Must be in (YYYY-MM-DDT)?HH:MM(:SS)? format`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--stop", "-s"},
				Description: `Stop an already running project`,
				ExclusiveOn: []string{"--no-stop"},
			}, {
				Name:        []string{"--no-stop", "-S"},
				Description: `Don't stop an already running project`,
				ExclusiveOn: []string{"--stop"},
			}},
		}, {
			Name:        []string{"start"},
			Description: `Start monitoring time for the given project. You can add tags indicating more specifically what you are working on with +tag`,
			Args: []model.Arg{{
				Name:       "ARGS",
				Generator:  nil, // TODO: port over generator
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--at"},
				Description: `Start frame at this time. Must be in (YYYY-MM-DDT)?HH:MM(:SS)? format`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}, {
				Name:        []string{"--gap", "-g"},
				Description: `Leave gap between end time of previous project and start time of the current`,
				ExclusiveOn: []string{"--no-gap"},
			}, {
				Name:        []string{"--no-gap", "-G"},
				Description: `Don't leave gap between end time of previous project and start time of the current`,
				ExclusiveOn: []string{"--gap"},
			}, {
				Name:        []string{"--confirm-new-project", "-c"},
				Description: `Confirm addition of new project`,
			}, {
				Name:        []string{"--confirm-new-tag", "-b"},
				Description: `Confirm creation of new tag`,
			}},
		}, {
			Name:        []string{"status"},
			Description: `Display when the current project was started and the time spent since`,
			Options: []model.Option{{
				Name:        []string{"--project", "-p"},
				Description: `Only output project`,
				ExclusiveOn: []string{"--tags", "--elapsed"},
			}, {
				Name:        []string{"--tags", "-t"},
				Description: `Only show tags`,
				ExclusiveOn: []string{"--project", "--elapsed"},
			}, {
				Name:        []string{"--elapsed", "-e"},
				Description: `Only show time elapsed`,
				ExclusiveOn: []string{"--project", "--tags"},
			}},
		}, {
			Name:        []string{"stop"},
			Description: `Stop monitoring time for the current project`,
			Options: []model.Option{{
				Name:        []string{"--at"},
				Description: `Stop frame at this time. Must be in (YYYY-MM-DDT)?HH:MM(:SS)? format`,
				Args: []model.Arg{{
					Name: "DATETIME",
				}},
			}},
		}, {
			Name:        []string{"sync"},
			Description: `Get the frames from the server and push the new ones`,
		}, {
			Name:        []string{"tags"},
			Description: `Display the list of all the tags`,
		}},
	}
}
