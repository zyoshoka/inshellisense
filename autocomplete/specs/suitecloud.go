// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["suitecloud"] = model.Subcommand{
		Name:        []string{"suitecloud"},
		Description: `SuiteCloud CLI`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `Outputs the version number`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Displays help for the command`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"account:manageauth"},
			Description: `Manages authentication IDs (authID) for all your projects. An authentication ID is a custom alias you gave to a specific account-role combination`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--info"},
				Description: `Prints the following information for the specified authentication ID (authID): account ID, role ID, and url. Usage: account:manageauth --info {authID}`,
				Args: []model.Arg{{
					Name: "authID",
				}},
			}, {
				Name:        []string{"--list"},
				Description: `Prints a list of all the configured authentication IDs (authID). Usage: account:manageauth --list`,
			}, {
				Name:        []string{"--remove"},
				Description: `Removes an authentication ID (authID). Usage: account:manageauth --remove {authID}`,
				Args: []model.Arg{{
					Name: "authID",
				}},
			}, {
				Name:        []string{"--rename"},
				Description: `Renames an authentication ID (authID). You must specify it together with the renameto option. Usage: account:manageauth --rename {authID} --renameto {newauthID}`,
				Args: []model.Arg{{
					Name: "authID",
				}},
			}, {
				Name:        []string{"--renameto"},
				Description: `Specifies the new ID for an authentication ID (authID). You must specify it together with the rename option. Usage: account:manageauth --rename {authID} --renameto {newauthID}`,
				Args: []model.Arg{{
					Name: "authID",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"account:savetoken"},
			Description: `Saves a TBA token that you issued previously in NetSuite`,
			Options: []model.Option{{
				Name:        []string{"--account"},
				Description: `Specifies the ID of the account to log in to`,
				Args: []model.Arg{{
					Name: "id",
				}},
			}, {
				Name:        []string{"--authid"},
				Description: `References the custom name you gave to a specific account-role combination`,
				Args: []model.Arg{{
					Name: "auth id",
				}},
			}, {
				Name:        []string{"--tokenid"},
				Description: `Specifies the token ID of a TBA token you issued previously in NetSuite`,
				Args: []model.Arg{{
					Name: "token id",
				}},
			}, {
				Name:        []string{"--tokensecret"},
				Description: `Specifies the token secret of a TBA token you issued previously in NetSuite`,
				Args: []model.Arg{{
					Name: "secret",
				}},
			}, {
				Name:        []string{"--url"},
				Description: `Specifies the NetSuite domain of the account you want to use. You only need to specify the NetSuite domain if you are not using a production account`,
				Args: []model.Arg{{
					Name: "domain",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"account:setup"},
			Description: `Sets up an account to use with the SuiteCloud CLI for Node.js`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"-d", "--dev"},
				Description: `Specifies the NetSuite domain of the account you want to use. You only need to specify the NetSuite domain if you are not using a production account`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"file:create"},
			Description: `Creates a SuiteScript file`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--module"},
				Description: `Specifies the SuiteScript modules you want to add to the SuiteScript file. For example, "N/record"`,
				Args: []model.Arg{{
					Name: "module",
				}},
			}, {
				Name:        []string{"--path"},
				Description: `Specifies the File Cabinet path of the SuiteScript file to create. For example, "/SuiteScripts/ClientScipt.js"`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}, {
				Name:        []string{"--type"},
				Description: `Specifies the type of the SuiteScript file that you want to create. For example, "ClientScript"`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"file:import"},
			Description: `Imports files from an account to your account customization project. You cannot import files from a SuiteApp`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--path"},
				Description: `Specifies the File Cabinet paths of the files to import. For example, "/SuiteScripts/file.js"`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}, {
				Name:        []string{"--calledfromcomparefiles"},
				Description: `Message displayed should be different if called from Compare Files`,
			}, {
				Name:        []string{"--excludeproperties"},
				Description: `Excludes all file properties within the .attributes folder`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"file:list"},
			Description: `Lists the files in the File Cabinet of your account`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--folder"},
				Description: `Specifies the File Cabinet path, for example, "/SuiteScripts". All files within subfolders are included`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "cabinet",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"file:upload"},
			Description: `Uploads files from your project to an account`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--paths"},
				Description: `Specifies the file cabinet paths of the files to upload. To specify multiple paths, enter a space between paths and enclose the entire argument in double quotes. For example, "/SuiteScripts/file.js" for account customization projects, and "/SuiteApps/com.project.example/script.js" for SuiteApp projects`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFilepaths},
					Name:       "files",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"object:import"},
			Description: `Imports custom objects from your NetSuite account to the SuiteCloud project. In account customization projects (ACP), if SuiteScript files are referenced in the custom objects you import, these files get imported by default`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--appid"},
				Description: `Specifies your application ID. If specified, only custom objects with that application ID are imported. Otherwise, only custom objects with no application ID are imported`,
				Args: []model.Arg{{
					Name: "app id",
				}},
			}, {
				Name:        []string{"--destinationfolder"},
				Description: `Specifies the project folder where objects will be stored. It must be within the Objects folder of your project. For example, /Objects/MyObjects`,
			}, {
				Name:        []string{"--excludefiles"},
				Description: `Indicates that the SuiteScript files referenced in custom objects are not imported. It can only be used in account customization projects (ACP)`,
			}, {
				Name:        []string{"--scriptid"},
				Description: `Specifies the script ID. To specify multiple IDs, enter the IDs separated by spaces. Enter "ALL" to import all custom objects of the specified type`,
			}, {
				Name:        []string{"--type"},
				Description: `Specifies the type of custom objects to import. Enter "ALL" to import all custom objects. To see what custom objects are supported by SDF, see https://system.netsuite.com/app/help/helpcenter.nl?fid=sdfxml.html`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"object:list"},
			Description: `Lists the custom objects deployed in an account`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--appid"},
				Description: `Specifies your application ID. If specified, only custom objects with that application ID are listed. Otherwise, only custom objects with no application ID are listed`,
				Args: []model.Arg{{
					Name: "app id",
				}},
			}, {
				Name:        []string{"--scriptid"},
				Description: `Specifies the script ID. If you specify it, only objects containing that script ID will be listed. Otherwise, all objects are listed`,
				Args: []model.Arg{{
					Name: "script id",
				}},
			}, {
				Name:        []string{"--type"},
				Description: `Specifies the type of custom objects to list. To specify multiple types, enter the types separated by spaces. Otherwise, all types are listed. To see what custom objects are supported by SDF, see https://system.netsuite.com/app/help/helpcenter.nl?fid=sdfxml.html`,
				Args: []model.Arg{{
					Name:       "types",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"object:update"},
			Description: `Overwrites the custom objects in the project with the custom objects in an account`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--includeinstances"},
				Description: `Includes instances. This is only available for custom records`,
			}, {
				Name:        []string{"--scriptid"},
				Description: `Specifies the script ID of the objects you want to overwrite`,
				Args: []model.Arg{{
					Name: "script id",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"project:adddependencies"},
			Description: `Adds the missing dependencies to the manifest file`,
		}, {
			Name:        []string{"project:create"},
			Description: `Creates a SuiteCloud project, either a SuiteApp or an account customization project (ACP)`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--overwrite"},
				Description: `Overwrites the existing project`,
			}, {
				Name:        []string{"--projectid"},
				Description: `Specifies the project ID. It is mandatory for SuiteApps`,
				Args: []model.Arg{{
					Name: "project id",
				}},
			}, {
				Name:        []string{"--projectname"},
				Description: `Specifies the project name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"--projectversion"},
				Description: `Specifies the project version. It is mandatory for SuiteApps`,
				Args: []model.Arg{{
					Name: "project version",
				}},
			}, {
				Name:        []string{"--publisherid"},
				Description: `Specifies the publisher ID. It is mandatory for SuiteApps`,
				Args: []model.Arg{{
					Name: "publisher id",
				}},
			}, {
				Name:        []string{"--type"},
				Description: `Specifies the project type. Enter one of the following options: ACCOUNTCUSTOMIZATION or SUITEAPP`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`ACCOUNTCUSTOMIZATION`}}, {Name: []string{`SUITEAPP`}}},
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"project:deploy"},
			Description: `Deploys the folder containing the project. The project folder is zipped before deployment, only including the files and folders referenced in the deploy.xml file`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--dryrun"},
				Description: `Runs a preview of your deploy process. Your project is not deployed`,
			}, {
				Name:        []string{"--log"},
				Description: `Sets the deployment log file location, as either a directory or a file name. If it is a directory, a default log file is generated in the specified location. If a log file already exists in the specified location, deployment log details are appended to that existing file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders, model.TemplateFilepaths},
					Name:      "location",
				}},
			}, {
				Name:        []string{"--validate"},
				Description: `Validates the project before deploying. If an error ocurrs during the deployment, the process is stopped`,
			}, {
				Name:        []string{"--accountspecificvalues"},
				Description: `Indicates how to handle the presence of account-specific values in an account customization project. If there are account-specific values in the project, enter WARNING to continue with the deployment process, or ERROR to stop it. If the option is not specified, the default value is ERROR. It only applies to account customization projects`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`ERROR`}}, {Name: []string{`WARNING`}}},
				}},
			}, {
				Name:        []string{"-a", "--applyinstallprefs"},
				Description: `Applies the settings from the hiding.xml, locking.xml, and overwriting.xml files. It only applies to SuiteApps`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"project:package"},
			Description: `Generates a ZIP file from your project, respecting the structure specified in the deploy.xml file`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}, {
			Name:        []string{"project:validate"},
			Description: `Validates the folder containing the SuiteCloud project`,
			Options: []model.Option{{
				Name:        []string{"-i", "--interactive"},
				Description: `Runs the command in interactive mode`,
			}, {
				Name:        []string{"--log"},
				Description: `Sets the validation log file location, as either a directory or a file name. If it is a directory, a default log file is generated in the specified location. If a log file already exists in the specified location, validation log details are appended to that existing file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders, model.TemplateFilepaths},
					Name:      "location",
				}},
			}, {
				Name:        []string{"--server"},
				Description: `Indicates that the server will perform the validation`,
			}, {
				Name:        []string{"--accountspecificvalues"},
				Description: `Indicates how to handle the presence of account-specific values in an account customization project. If there are account-specific values in the project, enter WARNING to continue with the deployment process, or ERROR to stop it. If the option is not specified, the default value is ERROR. It only applies to account customization projects`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`ERROR`}}, {Name: []string{`WARNING`}}},
				}},
			}, {
				Name:        []string{"-a", "--applyinstallprefs"},
				Description: `Applies the settings from the hiding.xml, locking.xml, and overwriting.xml files. It only applies to SuiteApps`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Displays help for the command`,
			}},
		}},
	}
}
