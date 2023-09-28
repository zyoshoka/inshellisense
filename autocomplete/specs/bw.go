// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["bw"] = model.Subcommand{
		Name: []string{"bw"},
		Options: []model.Option{{
			Name:         []string{"--pretty"},
			Description:  `Format output. JSON is tabbed with two spaces`,
			IsPersistent: true,
		}, {
			Name:         []string{"--raw"},
			Description:  `Return raw output instead of a descriptive message`,
			IsPersistent: true,
		}, {
			Name:         []string{"--response"},
			Description:  `Return a JSON formatted version of response output`,
			IsPersistent: true,
		}, {
			Name:         []string{"--cleanexit"},
			Description:  `Exit with a success exit code (0) unless an error is thrown`,
			IsPersistent: true,
		}, {
			Name:         []string{"--quiet"},
			Description:  `Don't return anything to stdout`,
			IsPersistent: true,
		}, {
			Name:         []string{"--nointeraction"},
			Description:  `Do not prompt for interactive user input`,
			IsPersistent: true,
		}, {
			Name:        []string{"--session"},
			Description: `Pass session key instead of reading from env`,
			Args: []model.Arg{{
				Name:        "session",
				Description: `Session Key`,
			}},
			IsPersistent: true,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Output BW CLI tool version number`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Display help for command`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"login"},
			Description: `Log into a user account`,
			Args: []model.Arg{{
				Name:        "email",
				Description: `Email Address of Bitwarden Account`,
				IsOptional:  true,
			}, {
				Name:        "password",
				Description: `Master Password of Bitwarden Vault`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--method"},
				Description: `Two-step login method`,
				Args: []model.Arg{{
					Name:        "method",
					Description: `Authenticator = 0, Email = 1, YubiKey = 3`,
					Suggestions: []model.Suggestion{{
						Name:        []string{`0`},
						Description: `Authenticator`,
					}, {
						Name:        []string{`1`},
						Description: `Email`,
					}, {
						Name:        []string{`3`},
						Description: `YubiKey`,
					}},
				}},
				ExclusiveOn: []string{"--sso", "--apikey", "--check"},
			}, {
				Name:        []string{"--code"},
				Description: `Two-step login code`,
				Args: []model.Arg{{
					Name:        "code",
					Description: `Valid TOTP Code`,
				}},
				ExclusiveOn: []string{"--method", "--apikey", "--check"},
			}, {
				Name:        []string{"--sso"},
				Description: `Log in with Single-Sign On`,
				ExclusiveOn: []string{"--method", "--code", "--apikey", "--passwordenv", "--passwordfile", "--check"},
			}, {
				Name:        []string{"--apikey"},
				Description: `Log in with an Api Key`,
				ExclusiveOn: []string{"--method", "--code", "--sso", "--passwordenv", "--passwordfile", "--check"},
			}, {
				Name:        []string{"--passwordenv"},
				Description: `Environment variable storing your password`,
				Args: []model.Arg{{
					Name:        "passwordenv",
					Description: `Environmental Variable Name`,
				}},
				ExclusiveOn: []string{"--method", "--code", "--sso", "--apikey", "--passwordfile", "--check"},
			}, {
				Name:        []string{"--passwordfile"},
				Description: `Path to a file containing your password as its first line`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "passwordfile",
					Description: `Path to a file containing your password`,
				}},
				ExclusiveOn: []string{"--passwordenv", "--sso", "--apikey", "--check"},
			}, {
				Name:        []string{"--check"},
				Description: `Check login status`,
				ExclusiveOn: []string{"--method", "--code", "--sso", "--apikey", "--passwordenv", "--passwordfile"},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for login command`,
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Log out of the current user account`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for logout command`,
			}},
		}, {
			Name:        []string{"lock"},
			Description: `Lock the vault and destroy active session keys`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for lock command`,
			}},
		}, {
			Name:        []string{"unlock"},
			Description: `Unlock the vault and return a new session key`,
			Args: []model.Arg{{
				Name:        "password",
				Description: `Master Password for the Vault to be unlocked`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--check"},
				Description: `Check lock status`,
				ExclusiveOn: []string{"--passwordenv", "--passwordfile"},
			}, {
				Name:        []string{"--passwordenv"},
				Description: `Supply an environment variable storing your password`,
				Args: []model.Arg{{
					Name:        "passwordenv",
					Description: `Environmental variable containing your password`,
				}},
				ExclusiveOn: []string{"--passwordfile"},
			}, {
				Name:        []string{"--passwordfile"},
				Description: `Path to a file containing your password as its first line`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "passwordfile",
					Description: `File containing your password`,
				}},
				ExclusiveOn: []string{"--passwordenv"},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for unlock command`,
				ExclusiveOn: []string{"--check", "--passwordenv", "--passwordfile"},
			}},
		}, {
			Name:        []string{"sync"},
			Description: `Pull the latest vault data from server`,
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `Force a full sync`,
				ExclusiveOn: []string{"--last"},
			}, {
				Name:        []string{"--last"},
				Description: `Show details of the last sync`,
				ExclusiveOn: []string{"--force", "-f"},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for sync command`,
			}},
		}, {
			Name:        []string{"generate"},
			Description: `Generate a password/passphrase`,
			Options: []model.Option{{
				Name:        []string{"-u", "--uppercase"},
				Description: `Include uppercase characters`,
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-h", "--help"},
			}, {
				Name:        []string{"-l", "--lowercase"},
				Description: `Include lowercase characters`,
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-h", "--help"},
			}, {
				Name:        []string{"-n", "--number"},
				Description: `Include numeric characters`,
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-h", "--help"},
			}, {
				Name:        []string{"-s", "--special"},
				Description: `Include special characters`,
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-h", "--help"},
			}, {
				Name:        []string{"--length"},
				Description: `Define password length`,
				Args: []model.Arg{{
					Name:        "length",
					Description: `Output length`,
				}},
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-h", "--help"},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for generate command`,
				ExclusiveOn: []string{"-p", "--passphrase", "--words", "-c", "--capitalize", "--separator", "--includeNumber", "-u", "--uppercase", "-l", "--lowercase", "-n", "--number", "-s", "--special", "--length"},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"-p", "--passphrase"},
				Description: `Generate a passphrase`,
				Options: []model.Option{{
					Name:        []string{"--words"},
					Description: `Number of words to include in passphrase`,
					Args: []model.Arg{{
						Name:        "Number of words",
						Description: `Number of words to include in passphrase`,
					}},
					ExclusiveOn: []string{"-u", "--uppercase", "-l", "--lowercase", "-n", "--number", "-s", "--special", "--length", "-h", "--help"},
				}, {
					Name:        []string{"-c", "--capitalize"},
					Description: `Title case passphrase`,
					ExclusiveOn: []string{"-u", "--uppercase", "-l", "--lowercase", "-n", "--number", "-s", "--special", "--length", "-h", "--help"},
				}, {
					Name:        []string{"--separator"},
					Description: `Define word separator`,
					Args: []model.Arg{{
						Name:        "separator",
						Description: `Character to delineate words in generated passphrase`,
					}},
					ExclusiveOn: []string{"-u", "--uppercase", "-l", "--lowercase", "-n", "--number", "-s", "--special", "--length", "-h", "--help"},
				}, {
					Name:        []string{"--includeNumber"},
					Description: `Include a number in output`,
					ExclusiveOn: []string{"-u", "--uppercase", "-l", "--lowercase", "-n", "--number", "-s", "--special", "--length", "-h", "--help"},
				}},
			}},
		}, {
			Name:        []string{"encode"},
			Description: `Base 64 encode stdin`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for encode command`,
			}},
		}, {
			Name:        []string{"config"},
			Description: `Configure CLI settings`,
			Args: []model.Arg{{
				Name:        "setting",
				Description: `Setting to configure`,
				Suggestions: []model.Suggestion{{Name: []string{`server`}}},
			}, {
				Name:        "value",
				Description: `Value to set`,
			}},
			Options: []model.Option{{
				Name:        []string{"--web-vault"},
				Description: `Provides a custom web vault URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `URL for a self-hosted Bitwarden Server`,
				}},
			}, {
				Name:        []string{"--api"},
				Description: `Provides a custom API URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `API URL for a self-hosted Bitwarden Server`,
				}},
			}, {
				Name:        []string{"--identity"},
				Description: `Provides a custom identity URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `Custom Identity URL that differs from the base URL`,
				}},
			}, {
				Name:        []string{"--icons"},
				Description: `Provides a custom icon service URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `Custom icon service URL that differs from the base URL`,
				}},
			}, {
				Name:        []string{"--notifications"},
				Description: `Provides a custom notifications URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `Custom notifications URL that differs from the base URL`,
				}},
			}, {
				Name:        []string{"--events"},
				Description: `Provides a custom events URL that differs from the base URL`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `Custom events URL that differs from the base URL`,
				}},
			}, {
				Name:        []string{"--key-connector"},
				Description: `Provides the URL for your Key Connector server`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `Custom Key Connector URL that differs from the base URL`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the config command`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Check for updates`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for update command`,
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate shell completions`,
			Options: []model.Option{{
				Name:        []string{"--shell"},
				Description: `Shell to generate completions for`,
				Args: []model.Arg{{
					Name:        "shell",
					Description: `Shell to generate completions for`,
					Suggestions: []model.Suggestion{{Name: []string{`zsh`}}},
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for completion command`,
			}},
		}, {
			Name:        []string{"status"},
			Description: `Show server details, last sync time & date, user information, and vault status`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for status command`,
			}},
		}, {
			Name:        []string{"serve"},
			Description: `Start a RESTful API webserver`,
			Options: []model.Option{{
				Name:        []string{"--hostname"},
				Description: `The hostname to bind your API webserver to`,
				Args: []model.Arg{{
					Name: "hostname",
				}},
			}, {
				Name:        []string{"--port"},
				Description: `The port to run your API webserver on`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for serve command`,
			}},
		}, {
			Name:        []string{"list"},
			Description: `Retrieves an array of objects (items, folders, collections, org-collections, org-members, organizations) from your Vault`,
			Args: []model.Arg{{
				Name:        "object",
				Suggestions: []model.Suggestion{{Name: []string{`items`}}, {Name: []string{`folders`}}, {Name: []string{`collections`}}, {Name: []string{`org-collections`}}, {Name: []string{`org-members`}}, {Name: []string{`organizations`}}},
			}},
			Options: []model.Option{{
				Name:        []string{"--search"},
				Description: `Perform a search on the listed objects`,
				Args: []model.Arg{{
					Name: "search",
				}},
			}, {
				Name:        []string{"--url"},
				Description: `Filter items of type login with a url-match search`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--folderid"},
				Description: `Filter items by folder id`,
				Args: []model.Arg{{
					Name: "folderid",
				}},
			}, {
				Name:        []string{"--collectionid"},
				Description: `Filter items by collection id`,
				Args: []model.Arg{{
					Name: "collectionid",
				}},
			}, {
				Name:        []string{"--organizationid"},
				Description: `Filter items or collections by organization id`,
				Args: []model.Arg{{
					Name: "organizationid",
				}},
			}, {
				Name:        []string{"--trash"},
				Description: `Filter items that are in the trash`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the list command`,
			}},
		}, {
			Name:        []string{"get"},
			Description: `Get an object from the vault`,
			Args: []model.Arg{{
				Name:        "object",
				Description: `Type of object to get from the Vault`,
				Suggestions: []model.Suggestion{{Name: []string{`item`}}, {Name: []string{`username`}}, {Name: []string{`password`}}, {Name: []string{`uri`}}, {Name: []string{`totp`}}, {Name: []string{`exposed`}}, {Name: []string{`folder`}}, {Name: []string{`collection`}}},
			}, {
				Name:        "id",
				Description: `Search term or object's globally unique id`,
			}},
			Options: []model.Option{{
				Name:        []string{"--id"},
				Description: `Exact id of the object`,
				Args: []model.Arg{{
					Name:        "id",
					Description: `Object id`,
				}},
			}, {
				Name:        []string{"--organizationid"},
				Description: `Organization id for an organization object`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization id`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the get command`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"template"},
				Description: `Type of template to get`,
				Args: []model.Arg{{
					Name:        "object",
					Suggestions: []model.Suggestion{{Name: []string{`item`}}, {Name: []string{`item.field`}}, {Name: []string{`item.login`}}, {Name: []string{`item.login.uri`}}, {Name: []string{`item.card`}}, {Name: []string{`item.identity`}}, {Name: []string{`item.securenote`}}, {Name: []string{`folder`}}, {Name: []string{`collection`}}, {Name: []string{`item-collections`}}, {Name: []string{`org-collection`}}, {Name: []string{`send.text`}}, {Name: []string{`send.file`}}},
				}},
			}, {
				Name: []string{"attachment"},
				Options: []model.Option{{
					Name:        []string{"--itemid"},
					Description: `Vault Item ID containing attachment`,
					Args: []model.Arg{{
						Name:        "itemid",
						Description: `ID of Vault Item containing attachment`,
					}},
				}, {
					Name:        []string{"--output"},
					Description: `Output directory and/or filename`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFolders},
						Name:        "output",
						Description: `Output directory and/or filename`,
					}},
				}},
			}},
		}, {
			Name:        []string{"create"},
			Description: `Create an object in the vault`,
			Args: []model.Arg{{
				Name:        "object",
				Description: `Object to be created`,
				Suggestions: []model.Suggestion{{Name: []string{`item`}}, {Name: []string{`attachment`}}, {Name: []string{`folder`}}, {Name: []string{`org-collection`}}},
			}, {
				Name:        "encodedJson",
				Description: `Encoded Json instead of uploading to Vault`,
			}},
			Options: []model.Option{{
				Name:        []string{"--file"},
				Description: `Path to file for attachment`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "file",
					Description: `File to attach`,
				}},
			}, {
				Name:        []string{"--itemid"},
				Description: `Attachment's item id`,
				Args: []model.Arg{{
					Name:        "itemid",
					Description: `Item ID`,
				}},
			}, {
				Name:        []string{"--organizationid"},
				Description: `Organization id for an organization object`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization ID`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the create command`,
			}},
		}, {
			Name:        []string{"edit"},
			Description: `Edit an object from the vault`,
			Args: []model.Arg{{
				Name:        "object",
				Description: `Type of Vault Object to edit`,
				Suggestions: []model.Suggestion{{Name: []string{`item`}}, {Name: []string{`item-collections`}}, {Name: []string{`folder`}}, {Name: []string{`org-collection`}}},
			}, {
				Name:        "id",
				Description: `Object's globally unique id`,
			}, {
				Name:        "encodedJson",
				Description: `Encoded json of the object to create`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--organizationid"},
				Description: `Organization id for an organization object`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization ID`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the edit command`,
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Delete an object from the vault`,
			Args: []model.Arg{{
				Name:        "object",
				Description: `Object to delete`,
				Suggestions: []model.Suggestion{{Name: []string{`item`}}, {Name: []string{`attachment`}}, {Name: []string{`folder`}}, {Name: []string{`org-collection`}}},
			}, {
				Name:        "id",
				Description: `Unique Vault ID`,
			}},
			Options: []model.Option{{
				Name:        []string{"--itemid"},
				Description: `Attachment's item id`,
				Args: []model.Arg{{
					Name:        "itemid",
					Description: `Item Id`,
				}},
			}, {
				Name:        []string{"--organizationid"},
				Description: `Organization id for an organization object`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization Id`,
				}},
			}, {
				Name:        []string{"-p", "--permanent"},
				Description: `Permanently deletes the item instead of soft-deleting it (item only)`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the delete command`,
				ExclusiveOn: []string{"--itemid", "--organizationid", "-p", "--permanent"},
			}},
		}, {
			Name:        []string{"restore"},
			Description: `Restores an object from the trash`,
			Args: []model.Arg{{
				Name: "object",
			}, {
				Name: "id",
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for the restore command`,
			}},
		}, {
			Name:        []string{"move"},
			Description: `Transfers a Vault item to an Organization`,
			Args: []model.Arg{{
				Name: "id",
			}, {
				Name: "organizationId",
			}, {
				Name:       "encodedJson",
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for the move command`,
			}},
		}, {
			Name:        []string{"confirm"},
			Description: `Confirm an object to the organization`,
			Args: []model.Arg{{
				Name:        "object",
				Description: `Valid objects are: org-member`,
			}, {
				Name:        "id",
				Description: `Organization id for an organization object`,
			}},
			Options: []model.Option{{
				Name:        []string{"--organizationid"},
				Description: `Organization id for an organization object`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization ID`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the confirm command`,
			}},
		}, {
			Name:        []string{"import"},
			Description: `Import vault data from a file`,
			Args: []model.Arg{{
				Name:        "format",
				Description: `Format of import file`,
				Suggestions: []model.Suggestion{{Name: []string{`1password1pif`}}, {Name: []string{`1password1pux`}}, {Name: []string{`1passwordmaccsv`}}, {Name: []string{`1passwordwincsv`}}, {Name: []string{`ascendocsv`}}, {Name: []string{`avastcsv`}}, {Name: []string{`avastjson`}}, {Name: []string{`aviracsv`}}, {Name: []string{`bitwardencsv`}}, {Name: []string{`bitwardenjson`}}, {Name: []string{`blackberrycsv`}}, {Name: []string{`blurcsv`}}, {Name: []string{`buttercupcsv`}}, {Name: []string{`chromecsv`}}, {Name: []string{`clipperzhtml`}}, {Name: []string{`codebookcsv`}}, {Name: []string{`dashlanecsv`}}, {Name: []string{`dashlanejson`}}, {Name: []string{`encryptrcsv`}}, {Name: []string{`enpasscsv`}}, {Name: []string{`enpassjson`}}, {Name: []string{`firefoxcsv`}}, {Name: []string{`fsecurefsk`}}, {Name: []string{`gnomejson`}}, {Name: []string{`kasperskytxt`}}, {Name: []string{`keepass2xml`}}, {Name: []string{`keepassxcsv`}}, {Name: []string{`keepercsv`}}, {Name: []string{`lastpasscsv`}}, {Name: []string{`logmeoncecsv`}}, {Name: []string{`meldiumcsv`}}, {Name: []string{`msecurecsv`}}, {Name: []string{`mykicsv`}}, {Name: []string{`nordpasscsv`}}, {Name: []string{`operacsv`}}, {Name: []string{`padlockcsv`}}, {Name: []string{`passboltcsv`}}, {Name: []string{`passkeepcsv`}}, {Name: []string{`passmanjson`}}, {Name: []string{`passpackcsv`}}, {Name: []string{`passwordagentcsv`}}, {Name: []string{`passwordbossjson`}}, {Name: []string{`passworddragonxml`}}, {Name: []string{`passwordwallettxt`}}, {Name: []string{`pwsafexml`}}, {Name: []string{`remembearcsv`}}, {Name: []string{`roboformcsv`}}, {Name: []string{`safaricsv`}}, {Name: []string{`safeincloudxml`}}, {Name: []string{`saferpasscsv`}}, {Name: []string{`securesafecsv`}}, {Name: []string{`splashidcsv`}}, {Name: []string{`stickypasswordxml`}}, {Name: []string{`truekeycsv`}}, {Name: []string{`upmcsv`}}, {Name: []string{`vivaldicsv`}}, {Name: []string{`yoticsv`}}, {Name: []string{`zohovaultcsv`}}},
			}, {
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "input",
				Description: `File to be imported`,
			}},
			Options: []model.Option{{
				Name:        []string{"--formats"},
				Description: `List supported file formats`,
			}, {
				Name:        []string{"--organizationid"},
				Description: `ID of the organization to import to`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization ID to be imported to`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the import command`,
			}},
		}, {
			Name:        []string{"export"},
			Description: `Export vault data to a CSV or JSON file`,
			Options: []model.Option{{
				Name:        []string{"--output"},
				Description: `Set output directory or filename`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFolders},
					Name:        "output",
					Description: `Path to export to`,
				}},
			}, {
				Name:        []string{"--format"},
				Description: `Export file format`,
				Args: []model.Arg{{
					Name:        "format",
					Description: `Export format`,
					Suggestions: []model.Suggestion{{Name: []string{`csv`}}, {Name: []string{`json`}}, {Name: []string{`encrypted_json`}}},
				}},
			}, {
				Name:        []string{"--password"},
				Description: `Use a password to encrypt instead of your Bitwarden account encryption key. Only applies to the encrypted_json format`,
				Args: []model.Arg{{
					Name:        "password",
					Description: `Password that will encrypt this export`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"--organizationid"},
				Description: `Enter organization id for an organization`,
				Args: []model.Arg{{
					Name:        "organizationid",
					Description: `Organization id`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the export command`,
			}},
		}, {
			Name:        []string{"share"},
			Description: `--DEPRECATED-- Move an item to an organization`,
			Args: []model.Arg{{
				Name:        "id",
				Description: `Object's globally unique ID`,
			}, {
				Name:        "organizationId",
				Description: `Organization's globally unique ID`,
			}, {
				Name:        "encodedJson",
				Description: `Encoded json of an array of collection ids. Can also be piped into stdin`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Display help for the share command`,
			}},
		}, {
			Name:        []string{"send"},
			Description: `Work with Bitwarden Sends. A Send can be quickly created using this command or subcommands can be used to fine-tune the Send`,
			Args: []model.Arg{{
				Name:        "data",
				Description: `Encoded Json data`,
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--file"},
				Description: `Specifies that <data> is a filepath`,
			}, {
				Name:        []string{"-d", "--deleteInDays"},
				Description: `The number of days in the future to set deletion date (defaults to 7 if not specified)`,
				Args: []model.Arg{{
					Name:       "days",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"-a", "--maxAccessCount"},
				Description: `The max number of accesses allowed for this Send`,
				Args: []model.Arg{{
					Name:        "amount",
					Description: `Max number of accesses allowed`,
				}},
			}, {
				Name:        []string{"--hidden"},
				Description: `Hide <data> in web by default. Valid only if --file is not set`,
			}, {
				Name:        []string{"-n", "--name"},
				Description: `The name of the Send. Defaults to a guid for text Sends and the filename for files`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Send name`,
				}},
			}, {
				Name:        []string{"--notes"},
				Description: `Notes to add to the Send`,
				Args: []model.Arg{{
					Name:        "notes",
					Description: `Notes to add to the Send`,
				}},
			}, {
				Name:        []string{"--fullObject"},
				Description: `Specifies that the full Send object should be returned rather than just the access url`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the Send command`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List all the Sends owned by you`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Display help for the Send command`,
				}},
			}, {
				Name:        []string{"template"},
				Description: `Get json templates for Send objects`,
				Args: []model.Arg{{
					Name:        "object",
					Description: `Json template to fetch`,
					Suggestions: []model.Suggestion{{Name: []string{`send.text`}}, {Name: []string{`send.file`}}},
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Display help for the template sub-command`,
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get Sends owned by you`,
				Args: []model.Arg{{
					Name:        "Send ID",
					Description: `Send ID`,
				}},
				Options: []model.Option{{
					Name:        []string{"--output"},
					Description: `Output directory or filename for attachment`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFolders},
						Name:        "output",
						Description: ``,
					}},
				}, {
					Name:        []string{"--text"},
					Description: `Specifies to return the text content of a Send`,
				}, {
					Name:        []string{"-h", "--help"},
					Description: `Display help for the Send get sub-command`,
				}},
			}, {
				Name:        []string{"receive"},
				Description: `Access a Bitwarden Send from a url`,
				Args: []model.Arg{{
					Name:        "Send URL",
					Description: `URL of Send to download`,
				}},
				Options: []model.Option{{
					Name:        []string{"--password"},
					Description: `Password needed to access the Send`,
					Args: []model.Arg{{
						Name:        "password",
						Description: `Password for Send`,
					}},
					ExclusiveOn: []string{"--passwordenv", "--passwordfile"},
				}, {
					Name:        []string{"--passwordenv"},
					Description: `Environment variable storing the Send's password`,
					Args: []model.Arg{{
						Name:        "passwordenv",
						Description: `Environment variable containing Send's password`,
					}},
					ExclusiveOn: []string{"--password", "--passwordfile"},
				}, {
					Name:        []string{"--passwordfile"},
					Description: `Path to a file containing the Send's password as its first line`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFilepaths},
						Name:        "passwordfile",
						Description: `File containing Send's password`,
					}},
					ExclusiveOn: []string{"--passwordenv", "--password"},
				}, {
					Name:        []string{"--obj"},
					Description: `Return the Send's json object rather than the Send's content`,
				}, {
					Name:        []string{"--output"},
					Description: `Specify a file path to save a File-type Send to`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFolders},
						Name:        "location",
						Description: `Path to save File-type Send`,
					}},
				}, {
					Name:        []string{"-h", "--help"},
					Description: `Display help for the receive sub-command`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a Send`,
				Args: []model.Arg{{
					Name:        "encodedJson",
					Description: `Encoded Json to be uploaded`,
				}},
				Options: []model.Option{{
					Name:        []string{"--file"},
					Description: `File to Send. Can also be specified in parent's JSON`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFilepaths},
						Name:        "path",
						Description: `Path to file to attach as content`,
					}},
					ExclusiveOn: []string{"--text", "--hidden"},
				}, {
					Name:        []string{"--text"},
					Description: `Text to Send. Can also be specified in parent's JSON`,
					Args: []model.Arg{{
						Name:        "text",
						Description: `Text to include as content`,
					}},
					ExclusiveOn: []string{"--file"},
				}, {
					Name:        []string{"--hidden"},
					Description: `Text hidden flag. Valid only with the --text option`,
				}, {
					Name:        []string{"--password"},
					Description: `Optional password to access this Send. Can also be specified in JSON`,
					Args: []model.Arg{{
						Name:        "password",
						Description: `Password to secure the Send`,
					}},
				}, {
					Name:        []string{"-h", "--help"},
					Description: `Display help for the create sub-command`,
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit a Send`,
				Args: []model.Arg{{
					Name:        "encodedJson",
					Description: `Encoded Json`,
				}},
				Options: []model.Option{{
					Name:        []string{"--itemid"},
					Description: `Overrides the itemId provided in [encodedJson]`,
					Args: []model.Arg{{
						Name:        "itemid",
						Description: `ItemID`,
					}},
				}, {
					Name:        []string{"-h", "--help"},
					Description: `Display help for the edit sub-command`,
				}},
			}, {
				Name:        []string{"remove-password"},
				Description: `Removes the saved password from a Send`,
				Args: []model.Arg{{
					Name:        "Send ID",
					Description: `ID of the Send to remove password from`,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Display help for the remove-password sub-command`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete a Send`,
				Args: []model.Arg{{
					Name:        "Send ID",
					Description: `ID of the Send to delete`,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Display help for the delete sub-command`,
				}},
			}, {
				Name:        []string{"help"},
				Description: `Display help for the Send command`,
				Args: []model.Arg{{
					Name:       "command",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"receive"},
			Description: `Access a Bitwarden Send from a url`,
			Args: []model.Arg{{
				Name:        "Send URL",
				Description: `URL of Send to download`,
			}},
			Options: []model.Option{{
				Name:        []string{"--password"},
				Description: `Password needed to access the Send`,
				Args: []model.Arg{{
					Name:        "password",
					Description: `Password for Send`,
				}},
				ExclusiveOn: []string{"--passwordenv", "--passwordfile"},
			}, {
				Name:        []string{"--passwordenv"},
				Description: `Environment variable storing the Send's password`,
				Args: []model.Arg{{
					Name:        "passwordenv",
					Description: `Environment variable containing Send's password`,
				}},
				ExclusiveOn: []string{"--password", "--passwordfile"},
			}, {
				Name:        []string{"--passwordfile"},
				Description: `Path to a file containing the Send's password as its first line`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "passwordfile",
					Description: `File containing Send's password`,
				}},
				ExclusiveOn: []string{"--passwordenv", "--password"},
			}, {
				Name:        []string{"--obj"},
				Description: `Return the Send's json object rather than the Send's content`,
			}, {
				Name:        []string{"--output"},
				Description: `Specify a file path to save a File-type Send to`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFolders},
					Name:        "location",
					Description: `Path to save File-type Send`,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Display help for the receive sub-command`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Display help for the Bitwarden CLI tool`,
			Args: []model.Arg{{
				Name:        "command",
				Description: `Subcommand to display help for`,
				Suggestions: []model.Suggestion{{Name: []string{`login`}}, {Name: []string{`logout`}}, {Name: []string{`lock`}}, {Name: []string{`unlock`}}, {Name: []string{`sync`}}, {Name: []string{`generate`}}, {Name: []string{`encode`}}, {Name: []string{`config`}}, {Name: []string{`update`}}, {Name: []string{`completion`}}, {Name: []string{`status`}}, {Name: []string{`serve`}}, {Name: []string{`list`}}, {Name: []string{`get`}}, {Name: []string{`create`}}, {Name: []string{`edit`}}, {Name: []string{`delete`}}, {Name: []string{`restore`}}, {Name: []string{`move`}}, {Name: []string{`confirm`}}, {Name: []string{`import`}}, {Name: []string{`export`}}, {Name: []string{`share`}}, {Name: []string{`send`}}, {Name: []string{`receive`}}},
				IsOptional:  true,
			}},
		}},
	}
}
