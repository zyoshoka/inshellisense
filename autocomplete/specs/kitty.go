// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["kitty"] = model.Subcommand{
		Name: []string{"kitty"},
		Args: []model.Arg{{
			IsCommand: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-T", "--title"},
			Description: `Set the OS window title`,
			Args: []model.Arg{{
				Name: "TITLE",
			}},
		}, {
			Name:        []string{"-C", "--config"},
			Description: `Specify a path to the configuration file(s) to use`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "CONFIG",
			}},
		}, {
			Name:        []string{"-o", "--override"},
			Description: `Override individual configuration options`,
		}, {
			Name:        []string{"-d", "--directory", "--working-directory"},
			Description: `Change to the specified directory when launching`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "DIRECTORY",
			}},
		}, {
			Name:        []string{"--session"},
			Description: `Path to a file containing the startup session`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "SESSION",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-`},
					Description: `Read from stdin`,
				}},
			}},
		}, {
			Name:        []string{"--hold"},
			Description: `Remain open after child process exits`,
		}, {
			Name:        []string{"-1", "--single-instance"},
			Description: `Only a single instance of kitty will run`,
		}, {
			Name:        []string{"--instance-group"},
			Description: `Kitty will open a new window in an existing instance and quit immediately`,
			Args: []model.Arg{{
				Name: "INSTANCE_GROUP",
			}},
		}, {
			Name:        []string{"--wait-for-single-instance-window-close"},
			Description: `The new window will not quit till the newly opened window is closed`,
		}, {
			Name:        []string{"--listen-on"},
			Description: `Tell kitty to listen on the specified address for control messages`,
			Args: []model.Arg{{
				Name: "LISTEN_ON",
			}},
		}, {
			Name:        []string{"--start-as"},
			Description: `Control how the initial kitty window is created`,
			Args: []model.Arg{{
				Name:        "START_AS",
				Suggestions: []model.Suggestion{{Name: []string{`normal`}}, {Name: []string{`fullscreen`}}, {Name: []string{`maximized`}}, {Name: []string{`minimized`}}},
			}},
		}, {
			Name:        []string{"-v", "--version"},
			Description: `The current kitty version`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Display this help message`,
		}},
		Subcommands: []model.Subcommand{{
			Name: []string{"@"},
			Options: []model.Option{{
				Name:        []string{"--to"},
				Description: `An address for the kitty instance to control`,
				Args: []model.Arg{{
					Name: "TO",
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"close-tab"},
				Description: `Close the specified tab(s)`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Close the tab of the window this command is run in, rather than the active tab`,
				}, {
					Name:        []string{"--target-group"},
					Description: `Close the specified group of tabs`,
					Args: []model.Arg{{
						Name: "TARGET_GROUP",
					}},
				}},
			}, {
				Name:        []string{"close-window"},
				Description: `Close the specified window(s)`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Close the window this command is run in, rather than the active window`,
				}},
			}, {
				Name: []string{"create-marker"},
				Args: []model.Arg{{
					Name: "MARKER",
				}, {
					Name: "SPECIFICATION",
				}},
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Close the window this command is run in, rather than the active window`,
				}},
			}, {
				Name:        []string{"detach-tab"},
				Description: `Detach the specified tab`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"-t", "--target-tab"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "TARGET_TAB",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Detach the tab this command is run in, rather than the active tab`,
				}},
			}, {
				Name:        []string{"detach-window"},
				Description: `Detach the specified window`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"-t", "--target-tab"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "TARGET_TAB",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Detach the window this command is run in, rather than the active window`,
				}},
			}, {
				Name:        []string{"disable-ligatures"},
				Description: `Control ligature rendering for the specified windows/tabs`,
				Args: []model.Arg{{
					Name:        "STRATEGY",
					Suggestions: []model.Suggestion{{Name: []string{`never`}}, {Name: []string{`always`}}, {Name: []string{`cursor`}}},
				}},
				Options: []model.Option{{
					Name:        []string{"-a", "--all"},
					Description: `Disable in all windows`,
				}, {
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"-t", "--match-tab"},
					Description: `The tab to match`,
				}},
			}, {
				Name:        []string{"env"},
				Description: `Change the environment variables seen by processing in newly launched windows`,
				Args: []model.Arg{{
					Name:       "ENV",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"focus-tab"},
				Description: `The active window in the specified tab will be focused`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--no-response"},
					Description: `Don't wait for a response indicating the success of the action`,
				}},
			}, {
				Name:        []string{"focus-window"},
				Description: `Focus the specified window, if no window is specified, focus the window this command is run inside`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--no-response"},
					Description: `Don't wait for a response indicating the success of the action`,
				}},
			}, {
				Name:        []string{"get-colors"},
				Description: `Get the terminal colors for the specified window`,
				Options: []model.Option{{
					Name:        []string{"-c", "--configured"},
					Description: `Instead of outputting the colors for the specified window, output the currently configured colors`,
				}, {
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
				}},
			}, {
				Name: []string{"get-text"},
			}, {
				Name: []string{"goto-layout"},
			}, {
				Name: []string{"kitten"},
			}, {
				Name: []string{"last-used-layout"},
			}, {
				Name: []string{"launch"},
			}, {
				Name: []string{"ls"},
			}, {
				Name: []string{"new-window"},
			}, {
				Name: []string{"remove-marker"},
			}, {
				Name: []string{"resize-os-window"},
			}, {
				Name: []string{"resize-window"},
			}, {
				Name: []string{"scroll-window"},
			}, {
				Name: []string{"select-window"},
			}, {
				Name: []string{"send-text"},
			}, {
				Name: []string{"set-background-image"},
			}, {
				Name: []string{"set-background-opacity"},
			}, {
				Name: []string{"set-colors"},
			}, {
				Name: []string{"set-enabled-layouts"},
			}, {
				Name: []string{"set-font-size"},
			}, {
				Name: []string{"set-spacing"},
			}, {
				Name:        []string{"set-tab-color"},
				Description: `The foreground and background colors when active and inactive can be overridden using this command`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The tab to match`,
					Args: []model.Arg{{
						Name: "MATCH",
					}},
				}, {
					Name:        []string{"--self"},
					Description: `Close the window this command is run in, rather than the active window`,
				}},
			}, {
				Name:        []string{"set-tab-title"},
				Description: `Set the title for the specified tab(s)`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The tab to match`,
				}},
			}, {
				Name: []string{"set-window-logo"},
			}, {
				Name:        []string{"set-window-title"},
				Description: `Set the title of the specified window(s)`,
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
				}, {
					Name:        []string{"--temporary"},
					Description: `The title can be overwritten by escape sequences`,
				}},
			}, {
				Name:        []string{"signal-child"},
				Description: `Send one or more signals to the foreground process in the specified window(s)`,
				Args: []model.Arg{{
					Name:       "SIGNAL",
					IsVariadic: true,
				}},
				Options: []model.Option{{
					Name:        []string{"-m", "--match"},
					Description: `The window to match`,
				}},
			}},
		}, {
			Name:        []string{"icat"},
			Description: `A cat like utility to display images in the terminal`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "image-file-or-url-or-directory",
			}},
			Options: []model.Option{{
				Name:        []string{"--align"},
				Description: `Horizontal alignment for the displayed image`,
				Args: []model.Arg{{
					Name:        "ALIGN",
					Suggestions: []model.Suggestion{{Name: []string{`center`}}, {Name: []string{`left`}}, {Name: []string{`right`}}},
				}},
			}, {
				Name:        []string{"--place"},
				Description: `Choose where on the screen to display the image`,
				Args: []model.Arg{{
					Name:        "PLACE",
					Description: `<width>x<height>@<left>x<top>`,
				}},
			}, {
				Name:        []string{"--scale-up"},
				Description: `Images that are smaller than the specified area to be scaled up to use as much of the specified area as possible`,
			}, {
				Name:        []string{"--background"},
				Description: `Specify a background color, this will cause transparent images to be composited on top of the specified color`,
				Args: []model.Arg{{
					Name: "COLOR",
					Suggestions: []model.Suggestion{{
						Name: []string{`none`},
					}, {
						Name: []string{`black`},
					}, {
						Name: []string{`white`},
					}, {
						Name: []string{`gray`},
					}, {
						Name: []string{`red`},
					}, {
						Name: []string{`green`},
					}, {
						Name: []string{`blue`},
					}, {
						Name: []string{`yellow`},
					}, {
						Name: []string{`magenta`},
					}, {
						Name: []string{`cyan`},
					}},
				}},
			}, {
				Name:        []string{"--mirror"},
				Description: `Mirror the image about a horizontal or vertical axis or both`,
				Args: []model.Arg{{
					Name:        "AXIS",
					Suggestions: []model.Suggestion{{Name: []string{`horizontal`}}, {Name: []string{`both`}}, {Name: []string{`none`}}, {Name: []string{`vertical`}}},
				}},
			}, {
				Name:        []string{"--clear"},
				Description: `Remove all images currently displayed on the screen`,
			}, {
				Name:        []string{"--transfer-mode"},
				Description: `Which mechanism to use to transfer images to the terminal`,
				Args: []model.Arg{{
					Name:        "TRANSFER_MODE",
					Suggestions: []model.Suggestion{{Name: []string{`file`}}, {Name: []string{`detect`}}, {Name: []string{`stream`}}},
				}},
			}, {
				Name:        []string{"--detect-support"},
				Description: `Detect support for image display in the terminal`,
			}, {
				Name:        []string{"--detection-timeout"},
				Description: `How long to wait for detection to complete before aborting`,
				Args: []model.Arg{{
					Name: "TIMEOUT",
				}},
			}, {
				Name:        []string{"--print-window-size"},
				Description: `Print the current terminal window size in pixels`,
			}, {
				Name:        []string{"--stdin"},
				Description: `Read an image from stdin`,
			}},
		}, {
			Name:        []string{"+icat"},
			Description: `A cat like utility to display images in the terminal`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "image-file-or-url-or-directory",
			}},
			Options: []model.Option{{
				Name:        []string{"--align"},
				Description: `Horizontal alignment for the displayed image`,
				Args: []model.Arg{{
					Name:        "ALIGN",
					Suggestions: []model.Suggestion{{Name: []string{`center`}}, {Name: []string{`left`}}, {Name: []string{`right`}}},
				}},
			}, {
				Name:        []string{"--place"},
				Description: `Choose where on the screen to display the image`,
				Args: []model.Arg{{
					Name:        "PLACE",
					Description: `<width>x<height>@<left>x<top>`,
				}},
			}, {
				Name:        []string{"--scale-up"},
				Description: `Images that are smaller than the specified area to be scaled up to use as much of the specified area as possible`,
			}, {
				Name:        []string{"--background"},
				Description: `Specify a background color, this will cause transparent images to be composited on top of the specified color`,
				Args: []model.Arg{{
					Name: "COLOR",
					Suggestions: []model.Suggestion{{
						Name: []string{`none`},
					}, {
						Name: []string{`black`},
					}, {
						Name: []string{`white`},
					}, {
						Name: []string{`gray`},
					}, {
						Name: []string{`red`},
					}, {
						Name: []string{`green`},
					}, {
						Name: []string{`blue`},
					}, {
						Name: []string{`yellow`},
					}, {
						Name: []string{`magenta`},
					}, {
						Name: []string{`cyan`},
					}},
				}},
			}, {
				Name:        []string{"--mirror"},
				Description: `Mirror the image about a horizontal or vertical axis or both`,
				Args: []model.Arg{{
					Name:        "AXIS",
					Suggestions: []model.Suggestion{{Name: []string{`horizontal`}}, {Name: []string{`both`}}, {Name: []string{`none`}}, {Name: []string{`vertical`}}},
				}},
			}, {
				Name:        []string{"--clear"},
				Description: `Remove all images currently displayed on the screen`,
			}, {
				Name:        []string{"--transfer-mode"},
				Description: `Which mechanism to use to transfer images to the terminal`,
				Args: []model.Arg{{
					Name:        "TRANSFER_MODE",
					Suggestions: []model.Suggestion{{Name: []string{`file`}}, {Name: []string{`detect`}}, {Name: []string{`stream`}}},
				}},
			}, {
				Name:        []string{"--detect-support"},
				Description: `Detect support for image display in the terminal`,
			}, {
				Name:        []string{"--detection-timeout"},
				Description: `How long to wait for detection to complete before aborting`,
				Args: []model.Arg{{
					Name: "TIMEOUT",
				}},
			}, {
				Name:        []string{"--print-window-size"},
				Description: `Print the current terminal window size in pixels`,
			}, {
				Name:        []string{"--stdin"},
				Description: `Read an image from stdin`,
			}},
		}, {
			Name: []string{"+list-fonts"},
		}, {
			Name: []string{"+hold"},
		}, {
			Name: []string{"+complete"},
		}, {
			Name: []string{"+runpy"},
		}, {
			Name: []string{"+launch"},
		}, {
			Name: []string{"+open"},
		}, {
			Name: []string{"+kitten"},
			Subcommands: []model.Subcommand{{
				Name:        []string{"icat"},
				Description: `A cat like utility to display images in the terminal`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "image-file-or-url-or-directory",
				}},
				Options: []model.Option{{
					Name:        []string{"--align"},
					Description: `Horizontal alignment for the displayed image`,
					Args: []model.Arg{{
						Name:        "ALIGN",
						Suggestions: []model.Suggestion{{Name: []string{`center`}}, {Name: []string{`left`}}, {Name: []string{`right`}}},
					}},
				}, {
					Name:        []string{"--place"},
					Description: `Choose where on the screen to display the image`,
					Args: []model.Arg{{
						Name:        "PLACE",
						Description: `<width>x<height>@<left>x<top>`,
					}},
				}, {
					Name:        []string{"--scale-up"},
					Description: `Images that are smaller than the specified area to be scaled up to use as much of the specified area as possible`,
				}, {
					Name:        []string{"--background"},
					Description: `Specify a background color, this will cause transparent images to be composited on top of the specified color`,
					Args: []model.Arg{{
						Name: "COLOR",
						Suggestions: []model.Suggestion{{
							Name: []string{`none`},
						}, {
							Name: []string{`black`},
						}, {
							Name: []string{`white`},
						}, {
							Name: []string{`gray`},
						}, {
							Name: []string{`red`},
						}, {
							Name: []string{`green`},
						}, {
							Name: []string{`blue`},
						}, {
							Name: []string{`yellow`},
						}, {
							Name: []string{`magenta`},
						}, {
							Name: []string{`cyan`},
						}},
					}},
				}, {
					Name:        []string{"--mirror"},
					Description: `Mirror the image about a horizontal or vertical axis or both`,
					Args: []model.Arg{{
						Name:        "AXIS",
						Suggestions: []model.Suggestion{{Name: []string{`horizontal`}}, {Name: []string{`both`}}, {Name: []string{`none`}}, {Name: []string{`vertical`}}},
					}},
				}, {
					Name:        []string{"--clear"},
					Description: `Remove all images currently displayed on the screen`,
				}, {
					Name:        []string{"--transfer-mode"},
					Description: `Which mechanism to use to transfer images to the terminal`,
					Args: []model.Arg{{
						Name:        "TRANSFER_MODE",
						Suggestions: []model.Suggestion{{Name: []string{`file`}}, {Name: []string{`detect`}}, {Name: []string{`stream`}}},
					}},
				}, {
					Name:        []string{"--detect-support"},
					Description: `Detect support for image display in the terminal`,
				}, {
					Name:        []string{"--detection-timeout"},
					Description: `How long to wait for detection to complete before aborting`,
					Args: []model.Arg{{
						Name: "TIMEOUT",
					}},
				}, {
					Name:        []string{"--print-window-size"},
					Description: `Print the current terminal window size in pixels`,
				}, {
					Name:        []string{"--stdin"},
					Description: `Read an image from stdin`,
				}},
			}, {
				Name: []string{"diff"},
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file1",
				}, {
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file2",
				}},
			}, {
				Name: []string{"show_key"},
			}, {
				Name: []string{"clipboard"},
			}, {
				Name: []string{"unicode_input"},
			}, {
				Name: []string{"panel"},
			}, {
				Name: []string{"transfer"},
			}, {
				Name: []string{"query_terminal"},
			}, {
				Name: []string{"broadcast"},
			}, {
				Name: []string{"hyperlinked_grep"},
			}, {
				Name: []string{"ssh"},
			}, {
				Name: []string{"choose"},
			}, {
				Name: []string{"ask"},
			}, {
				Name:        []string{"themes"},
				Description: `Change the kitty theme`,
			}, {
				Name: []string{"hints"},
			}, {
				Name: []string{"remote_file"},
			}, {
				Name: []string{"show_error"},
			}, {
				Name: []string{"resize_window"},
			}, {
				Name: []string{"mouse_demo"},
			}},
		}, {
			Name: []string{"+edit-config"},
		}, {
			Name: []string{"+shebang"},
		}},
	}
}
