// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["clilol"] = model.Subcommand{
		Name:        []string{"clilol"},
		Description: `A cli for omg.lol`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"create"},
			Description: `Create things`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"dns"},
				Description: `Create a DNS record`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the DNS record`,
				}, {
					Name:        "type",
					Description: `Type of the DNS record`,
				}, {
					Name:        "data",
					Description: `Data of the DNS record`,
				}},
				Options: []model.Option{{
					Name:        []string{"--priority", "-p"},
					Description: `Priority of the DNS record`,
					Args: []model.Arg{{
						Name: "priority",
					}},
				}, {
					Name:        []string{"--ttl", "-T"},
					Description: `Time to live of the DNS record`,
					Args: []model.Arg{{
						Name: "ttl",
					}},
				}},
			}, {
				Name:        []string{"paste"},
				Description: `Create or update a paste`,
				Args: []model.Arg{{
					Name:        "title",
					Description: `Title of the paste`,
				}},
				Options: []model.Option{{
					Name:        []string{"--filename", "-f"},
					Description: `File to read paste from (default stdin)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}, {
					Name:        []string{"--listed", "-l"},
					Description: `Create paste as listed (default false)`,
				}},
			}, {
				Name:        []string{"picture"},
				Description: `Upload a picture to some.pics`,
				Args: []model.Arg{{
					Name:        "filename",
					Description: `Filename of the image file`,
				}},
				Options: []model.Option{{
					Name:        []string{"--description", "-d"},
					Description: `Description of the picture (default empty/unlisted)`,
				}},
			}, {
				Name:        []string{"purl"},
				Description: `Create a PURL`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the PURL`,
				}, {
					Name:        "url",
					Description: `URL that the PURL redirects to`,
				}},
				Options: []model.Option{{
					Name:        []string{"--listed", "-l"},
					Description: `Create as listed (default false)`,
				}},
			}, {
				Name:        []string{"status"},
				Description: `Create a status`,
				Args: []model.Arg{{
					Name:        "text",
					Description: `Text of the status`,
				}},
				Options: []model.Option{{
					Name:        []string{"--emoji", "-e"},
					Description: `Emoji to add to status (default sparkles)`,
					Args: []model.Arg{{
						Name: "emoji",
					}},
				}, {
					Name:        []string{"--skip-mastodon-post"},
					Description: `Do not cross-post to Mastodon`,
				}},
			}, {
				Name:        []string{"weblog"},
				Description: `Create a weblog entry`,
				Options: []model.Option{{
					Name:        []string{"--filename", "-f"},
					Description: `File to read entry from (default stdin)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}},
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Delete things`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"account"},
				Description: `Delete information about your account`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"session"},
					Description: `Delete a session`,
					Args: []model.Arg{{
						Name:        "id",
						Description: `ID of the session to delete`,
					}},
				}},
			}, {
				Name:        []string{"dns"},
				Description: `Delete a DNS record`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the record to delete`,
				}},
			}, {
				Name:        []string{"paste"},
				Description: `Delete a paste`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the paste to delete`,
				}},
			}, {
				Name:        []string{"picture"},
				Description: `Delete a picture from some.pics`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the picture to delete`,
				}},
			}, {
				Name:        []string{"purl"},
				Description: `Delete a PURL`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the PURL to delete`,
				}},
			}, {
				Name:        []string{"status"},
				Description: `Delete a status`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the status to delete`,
				}},
			}, {
				Name:        []string{"weblog"},
				Description: `Delete a weblog entry`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the weblog entry to delete`,
				}},
			}},
		}, {
			Name:        []string{"get"},
			Description: `Get things`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"account"},
				Description: `Get information about your account`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"info"},
					Description: `Get info about your account`,
				}, {
					Name:        []string{"name"},
					Description: `Get your account name`,
				}, {
					Name:        []string{"settings"},
					Description: `Get your account settings`,
				}},
			}, {
				Name:        []string{"address"},
				Description: `Get information about an address`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"availability"},
					Description: `Get address availability`,
					Args: []model.Arg{{
						Name:        "address",
						Description: `Address to get`,
					}},
				}, {
					Name:        []string{"expiration"},
					Description: `Get address expiration`,
					Args: []model.Arg{{
						Name:        "address",
						Description: `Address to get`,
					}},
				}, {
					Name:        []string{"info"},
					Description: `Get information about an address`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"private"},
						Description: `Get private information about an address`,
						Args: []model.Arg{{
							Name:        "address",
							Description: `Address to get`,
						}},
					}, {
						Name:        []string{"public"},
						Description: `Get public information about an address`,
						Args: []model.Arg{{
							Name:        "address",
							Description: `Address to get`,
						}},
					}},
				}},
			}, {
				Name:        []string{"email"},
				Description: `Get email forwarding address(es)`,
			}, {
				Name:        []string{"now"},
				Description: `Get a Now page`,
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose Now page to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}, {
					Name:        []string{"--filename", "-f"},
					Description: `File to write Now page to (default stdout)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}},
			}, {
				Name:        []string{"paste"},
				Description: `Get a paste`,
				Args: []model.Arg{{
					Name:        "title",
					Description: `Title of the paste`,
				}},
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose paste to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}, {
					Name:        []string{"--filename", "-f"},
					Description: `File to write paste to (default stdout)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}},
			}, {
				Name:        []string{"purl"},
				Description: `Get a PURL`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the PURL`,
				}},
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose PURL to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}},
			}, {
				Name:        []string{"service"},
				Description: `Get service stats`,
			}, {
				Name:        []string{"status"},
				Description: `Get status`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the status`,
				}},
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose status to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}},
			}, {
				Name:        []string{"status-bio"},
				Description: `Get status bio`,
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose status bio to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}},
			}, {
				Name:        []string{"theme"},
				Description: `Get theme information`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the theme`,
				}},
				Subcommands: []model.Subcommand{{
					Name:        []string{"preview"},
					Description: `Get theme preview`,
					Args: []model.Arg{{
						Name:        "name",
						Description: `Name of the theme`,
					}},
					Options: []model.Option{{
						Name:        []string{"--filename", "-f"},
						Description: `File to write preview to (default stdout)`,
						Args: []model.Arg{{
							Templates: []model.Template{model.TemplateFilepaths},
							Name:      "filename",
						}},
					}},
				}},
			}, {
				Name:        []string{"web"},
				Description: `Get your webpage content`,
				Options: []model.Option{{
					Name:        []string{"--filename", "-f"},
					Description: `File to write webpage to (default stdout)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}},
			}, {
				Name:        []string{"weblog"},
				Description: `Get a weblog entry`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `ID of the weblog entry`,
				}},
				Subcommands: []model.Subcommand{{
					Name:        []string{"config"},
					Description: `Get your weblog config`,
					Options: []model.Option{{
						Name:        []string{"--filename", "-f"},
						Description: `File to write configuration to (default stdout)`,
						Args: []model.Arg{{
							Templates: []model.Template{model.TemplateFilepaths},
							Name:      "filename",
						}},
					}},
				}, {
					Name:        []string{"latest"},
					Description: `Get the latest weblog entry`,
				}, {
					Name:        []string{"template"},
					Description: `Get your weblog template`,
					Options: []model.Option{{
						Name:        []string{"--filename", "-f"},
						Description: `File to write template to (default stdout)`,
						Args: []model.Arg{{
							Templates: []model.Template{model.TemplateFilepaths},
							Name:      "filename",
						}},
					}},
				}},
			}},
		}, {
			Name:        []string{"list"},
			Description: `List things`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"account"},
				Description: `List information about your account`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"addresses"},
					Description: `List your addresses`,
				}, {
					Name:        []string{"sessions"},
					Description: `List your sessions`,
				}},
			}, {
				Name:        []string{"directory"},
				Description: `List the address directory`,
			}, {
				Name:        []string{"dns"},
				Description: `List your dns records`,
			}, {
				Name:        []string{"now"},
				Description: `List Now pages`,
			}, {
				Name:        []string{"paste"},
				Description: `List pastes`,
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose pastes to list`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}},
			}, {
				Name:        []string{"picture"},
				Description: `List pictures`,
			}, {
				Name:        []string{"purl"},
				Description: `List all PURLs`,
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose PURLs to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}},
			}, {
				Name:        []string{"status"},
				Description: `List statuses`,
				Options: []model.Option{{
					Name:        []string{"--address", "-a"},
					Description: `Address whose status(es) to get`,
					Args: []model.Arg{{
						Name: "address",
					}},
				}, {
					Name:        []string{"--limit", "-l"},
					Description: `How many status(es) to get (default all)`,
					Args: []model.Arg{{
						Name: "limit",
					}},
				}},
			}, {
				Name:        []string{"statuslog"},
				Description: `List the statuslog`,
				Options: []model.Option{{
					Name:        []string{"--all", "-A"},
					Description: `Get the entire statuslog (default is latest statuses only)`,
				}},
			}, {
				Name:        []string{"theme"},
				Description: `List profile themes`,
			}, {
				Name:        []string{"weblog"},
				Description: `List all weblog entries`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Update things`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"account"},
				Description: `Update information about your account`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"name"},
					Description: `Set the name on your account`,
				}, {
					Name:        []string{"settings"},
					Description: `Set the settings on your account`,
					Options: []model.Option{{
						Name:        []string{"--communication", "-c"},
						Description: `Communication preference`,
						Args: []model.Arg{{
							Name: "communication",
						}},
					}, {
						Name:        []string{"--date-format", "-d"},
						Description: `Date format preference`,
						Args: []model.Arg{{
							Name: "date-format",
						}},
					}, {
						Name:        []string{"--web-editor", "-w"},
						Description: `Web editor preference`,
						Args: []model.Arg{{
							Name: "web-editor",
						}},
					}},
				}},
			}, {
				Name:        []string{"dns"},
				Description: `Update a DNS record`,
				Options: []model.Option{{
					Name:        []string{"--priority", "-p"},
					Description: `Updated priority`,
					Args: []model.Arg{{
						Name: "priority",
					}},
				}, {
					Name:        []string{"--ttl", "-T"},
					Description: `Updated TTL`,
					Args: []model.Arg{{
						Name: "ttl",
					}},
				}},
			}, {
				Name:        []string{"email"},
				Description: `Set email forwarding address(es)`,
				Options: []model.Option{{
					Name:        []string{"--destination", "-d"},
					Description: `Address(es) to forward to`,
					Args: []model.Arg{{
						Name: "destination",
					}},
				}},
			}, {
				Name:        []string{"preference"},
				Description: `Set a preference`,
			}, {
				Name:        []string{"set"},
				Description: `Set Now page content`,
				Options: []model.Option{{
					Name:        []string{"--filename", "-f"},
					Description: `File to read Now page from (default stdin)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}, {
					Name:        []string{"--listed", "-l"},
					Description: `Create Now page as listed (default false)`,
				}},
			}, {
				Name:        []string{"status"},
				Description: `Update a status`,
				Options: []model.Option{{
					Name:        []string{"--emoji", "-e"},
					Description: `Emoji to add to status (default sparkles)`,
					Args: []model.Arg{{
						Name: "emoji",
					}},
				}},
			}, {
				Name:        []string{"status-bio"},
				Description: `Update your status bio`,
			}, {
				Name:        []string{"web"},
				Description: `Set webpage content`,
				Options: []model.Option{{
					Name:        []string{"--filename", "-f"},
					Description: `File to read webpage from (default stdin)`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "filename",
					}},
				}, {
					Name:        []string{"--publish", "-p"},
					Description: `Publish the updated page (default false)`,
				}},
				Subcommands: []model.Subcommand{{
					Name:        []string{"pfp"},
					Description: `Set your profile picture`,
				}},
			}, {
				Name:        []string{"weblog"},
				Description: `Set your weblog config`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"config"},
					Description: `Set your weblog config`,
					Options: []model.Option{{
						Name:        []string{"--filename", "-f"},
						Description: `File to read config from (default stdin)`,
						Args: []model.Arg{{
							Templates: []model.Template{model.TemplateFilepaths},
							Name:      "filename",
						}},
					}},
				}, {
					Name:        []string{"template"},
					Description: `Set your weblog template`,
					Options: []model.Option{{
						Name:        []string{"--filename", "-f"},
						Description: `File to read template from (default stdin)`,
						Args: []model.Arg{{
							Templates: []model.Template{model.TemplateFilepaths},
							Name:      "filename",
						}},
					}},
				}},
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create things`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"dns"},
					Description: `Create a DNS record`,
				}, {
					Name:        []string{"paste"},
					Description: `Create or update a paste`,
				}, {
					Name:        []string{"purl"},
					Description: `Create a PURL`,
				}, {
					Name:        []string{"status"},
					Description: `Create a status`,
				}, {
					Name:        []string{"weblog"},
					Description: `Create a weblog entry`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete things`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"account"},
					Description: `Delete information about your account`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"session"},
						Description: `Delete a session`,
					}},
				}, {
					Name:        []string{"dns"},
					Description: `Delete a DNS record`,
				}, {
					Name:        []string{"paste"},
					Description: `Delete a paste`,
				}, {
					Name:        []string{"purl"},
					Description: `Delete a PURL`,
				}, {
					Name:        []string{"weblog"},
					Description: `Delete a weblog entry`,
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get things`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"account"},
					Description: `Get information about your account`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"info"},
						Description: `Get info about your account`,
					}, {
						Name:        []string{"name"},
						Description: `Get your account name`,
					}, {
						Name:        []string{"settings"},
						Description: `Get your account settings`,
					}},
				}, {
					Name:        []string{"address"},
					Description: `Get information about an address`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"availability"},
						Description: `Get address availability`,
					}, {
						Name:        []string{"expiration"},
						Description: `Get address expiration`,
					}, {
						Name:        []string{"info"},
						Description: `Get information about an address`,
						Subcommands: []model.Subcommand{{
							Name:        []string{"private"},
							Description: `Get private information about an address`,
						}, {
							Name:        []string{"public"},
							Description: `Get public information about an address`,
						}},
					}},
				}, {
					Name:        []string{"email"},
					Description: `Get email forwarding address(es)`,
				}, {
					Name:        []string{"now"},
					Description: `Get a Now page`,
				}, {
					Name:        []string{"paste"},
					Description: `Get a paste`,
				}, {
					Name:        []string{"purl"},
					Description: `Get a PURL`,
				}, {
					Name:        []string{"service"},
					Description: `Get service stats`,
				}, {
					Name:        []string{"status"},
					Description: `Get status`,
				}, {
					Name:        []string{"status-bio"},
					Description: `Get status bio`,
				}, {
					Name:        []string{"theme"},
					Description: `Get theme information`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"preview"},
						Description: `Get theme preview`,
					}},
				}, {
					Name:        []string{"web"},
					Description: `Get your webpage content`,
				}, {
					Name:        []string{"weblog"},
					Description: `Get a weblog entry`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"config"},
						Description: `Get your weblog config`,
					}, {
						Name:        []string{"latest"},
						Description: `Get the latest weblog entry`,
					}, {
						Name:        []string{"template"},
						Description: `Get your weblog template`,
					}},
				}},
			}, {
				Name:        []string{"list"},
				Description: `List things`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"account"},
					Description: `List information about your account`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"addresses"},
						Description: `List your addresses`,
					}, {
						Name:        []string{"sessions"},
						Description: `List your sessions`,
					}},
				}, {
					Name:        []string{"directory"},
					Description: `List the address directory`,
				}, {
					Name:        []string{"dns"},
					Description: `List your dns records`,
				}, {
					Name:        []string{"now"},
					Description: `List Now pages`,
				}, {
					Name:        []string{"paste"},
					Description: `List pastes`,
				}, {
					Name:        []string{"purl"},
					Description: `List all PURLs`,
				}, {
					Name:        []string{"status"},
					Description: `List statuses`,
				}, {
					Name:        []string{"statuslog"},
					Description: `List the statuslog`,
				}, {
					Name:        []string{"theme"},
					Description: `List profile themes`,
				}, {
					Name:        []string{"weblog"},
					Description: `List all weblog entries`,
				}},
			}, {
				Name:        []string{"update"},
				Description: `Update things`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"account"},
					Description: `Update information about your account`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"name"},
						Description: `Set the name on your account`,
						Args: []model.Arg{{
							Name:        "name",
							Description: `The name to set`,
						}},
					}, {
						Name:        []string{"settings"},
						Description: `Set the settings on your account`,
					}},
				}, {
					Name:        []string{"dns"},
					Description: `Update a DNS record`,
					Args: []model.Arg{{
						Name:        "id",
						Description: `The ID of the record to update`,
					}, {
						Name:        "name",
						Description: `Updated name`,
					}, {
						Name:        "type",
						Description: `Updated type`,
					}, {
						Name:        "data",
						Description: `Updated data`,
					}},
				}, {
					Name:        []string{"email"},
					Description: `Set email forwarding address(es)`,
					Args: []model.Arg{{
						Name:        "address",
						Description: `Address(es) to forward to`,
					}},
				}, {
					Name:        []string{"preference"},
					Description: `Set a preference`,
					Args: []model.Arg{{
						Name:        "item",
						Description: `Preferences item to set`,
					}, {
						Name:        "value",
						Description: `Value to set`,
					}},
				}, {
					Name:        []string{"now"},
					Description: `Update Now page content`,
				}, {
					Name:        []string{"status"},
					Description: `Update a status`,
					Args: []model.Arg{{
						Name:        "id",
						Description: `The ID of the status to update`,
					}, {
						Name:        "text",
						Description: `New text for the status`,
					}},
				}, {
					Name:        []string{"status-bio"},
					Description: `Update your status bio`,
					Args: []model.Arg{{
						Name:        "text",
						Description: `New text for the status bio`,
					}},
				}, {
					Name:        []string{"web"},
					Description: `Set webpage content`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"pfp"},
						Description: `Set your profile picture`,
						Args: []model.Arg{{
							Templates:   []model.Template{model.TemplateFilepaths},
							Name:        "filename",
							Description: `The filename of the image to set`,
						}},
					}},
				}, {
					Name:        []string{"weblog"},
					Description: `Set your weblog config`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"config"},
						Description: `Set your weblog config`,
					}, {
						Name:        []string{"template"},
						Description: `Set your weblog template`,
					}},
				}},
			}},
		}},
	}
}
