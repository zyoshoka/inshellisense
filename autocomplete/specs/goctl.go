// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["goctl"] = model.Subcommand{
		Name:        []string{"goctl"},
		Description: `A cli tool to generate go-zero code`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"api"},
			Description: `Generate api related files`,
			Options: []model.Option{{
				Name:        []string{"--branch"},
				Description: `The branch of the remote repo, it does work with --remote`,
				Args: []model.Arg{{
					Name: "branch",
				}},
			}, {
				Name:        []string{"--home"},
				Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
				Args: []model.Arg{{
					Name: "home",
				}},
			}, {
				Name:        []string{"--o"},
				Description: `Output a sample api file`,
				Args: []model.Arg{{
					Name: "o",
				}},
			}, {
				Name:        []string{"--remote"},
				Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
				Args: []model.Arg{{
					Name: "remote",
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"dart"},
				Description: `Generate dart files for provided api in api file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--hostname"},
					Description: `Hostname of the server`,
					Args: []model.Arg{{
						Name: "hostname",
					}},
				}, {
					Name:        []string{"--legacy"},
					Description: `Legacy generator for flutter v1`,
				}},
			}, {
				Name:        []string{"doc"},
				Description: `Generate doc files`,
				Options: []model.Option{{
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--o"},
					Description: `The output markdown directory`,
					Args: []model.Arg{{
						Name: "o",
					}},
				}},
			}, {
				Name:        []string{"format"},
				Description: `Format api files`,
				Options: []model.Option{{
					Name:        []string{"--declare"},
					Description: `Use to skip check api types already declare`,
				}, {
					Name:        []string{"--dir"},
					Description: `The format target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--iu"},
					Description: `Ignore update`,
				}, {
					Name:        []string{"--stdin"},
					Description: `Use stdin to input api doc content, press "ctrl + d" to send EOF`,
				}},
			}, {
				Name:        []string{"go"},
				Description: `Generate go files for provided api in yaml file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}},
			}, {
				Name:        []string{"java"},
				Description: `Generate java files for provided api in api file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}},
			}, {
				Name:        []string{"kt"},
				Description: `Generate kotlin code for provided api file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--pkg"},
					Description: `Define package name for kotlin file`,
					Args: []model.Arg{{
						Name: "pkg",
					}},
				}},
			}, {
				Name:        []string{"new"},
				Description: `Fast create api service`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}},
			}, {
				Name:        []string{"plugin"},
				Description: `Custom file generator`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--plugin", "-p"},
					Description: `The plugin file`,
					Args: []model.Arg{{
						Name: "plugin",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}},
			}, {
				Name:        []string{"ts"},
				Description: `Generate ts files for provided api in api file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `The api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}, {
					Name:        []string{"--caller"},
					Description: `The web api caller`,
					Args: []model.Arg{{
						Name: "caller",
					}},
				}, {
					Name:        []string{"--dir"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--unwrap"},
					Description: `Unwrap the webapi caller for import`,
				}, {
					Name:        []string{"--webapi"},
					Description: `The web api file path`,
					Args: []model.Arg{{
						Name: "webapi",
					}},
				}},
			}, {
				Name:        []string{"validate"},
				Description: `Validate api file`,
				Options: []model.Option{{
					Name:        []string{"--api"},
					Description: `Validate target api file`,
					Args: []model.Arg{{
						Name: "api",
					}},
				}},
			}},
		}, {
			Name:        []string{"bug"},
			Description: `Report a bug`,
		}, {
			Name:        []string{"completion"},
			Description: `Generate the autocompletion script for the specified shell`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate the autocompletion script for bash`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"fish"},
				Description: `Generate the autocompletion script for fish`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"powershell"},
				Description: `Generate the autocompletion script for powershell`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"zsh"},
				Description: `Generate the autocompletion script for zsh`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}},
		}, {
			Name:        []string{"docker"},
			Description: `Generate Dockerfile`,
			Options: []model.Option{{
				Name:        []string{"--base"},
				Description: `The base image to build the docker image, default scratch`,
				Args: []model.Arg{{
					Name: "base",
				}},
			}, {
				Name:        []string{"--branch"},
				Description: `The branch of the remote repo, it does work with --remote`,
				Args: []model.Arg{{
					Name: "branch",
				}},
			}, {
				Name:        []string{"--go"},
				Description: `The directory that contains main function`,
				Args: []model.Arg{{
					Name: "go",
				}},
			}, {
				Name:        []string{"--home"},
				Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
				Args: []model.Arg{{
					Name: "home",
				}},
			}, {
				Name:        []string{"--port"},
				Description: `The port to expose, default none`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--remote"},
				Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
				Args: []model.Arg{{
					Name: "remote",
				}},
			}, {
				Name:        []string{"--tz"},
				Description: `The timezone of the container`,
				Args: []model.Arg{{
					Name: "tz",
				}},
			}, {
				Name:        []string{"--version"},
				Description: `The goctl builder golang image version`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}},
		}, {
			Name:        []string{"env"},
			Description: `Check or edit goctl environment`,
			Options: []model.Option{{
				Name:        []string{"--write", "-w"},
				Description: `Edit goctl environment`,
				Args: []model.Arg{{
					Name: "write",
				}},
			}, {
				Name:         []string{"--force", "-f"},
				Description:  `Silent installation of non-existent dependencies`,
				IsPersistent: true,
			}, {
				Name:         []string{"--verbose", "-v"},
				Description:  `Enable log output`,
				IsPersistent: true,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"check"},
				Description: `Detect goctl env and dependency tools`,
				Options: []model.Option{{
					Name:        []string{"--install", "-i"},
					Description: `Install dependencies if not found`,
				}},
			}, {
				Name:        []string{"install"},
				Description: `Goctl env installation`,
			}},
		}, {
			Name:        []string{"kube"},
			Description: `Generate kubernetes files`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"deploy"},
				Description: `Generate deployment yaml file`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--image"},
					Description: `The docker image of deployment (required)`,
					Args: []model.Arg{{
						Name: "image",
					}},
				}, {
					Name:        []string{"--limitCpu"},
					Description: `The limit cpu to deploy`,
					Args: []model.Arg{{
						Name: "limitCpu",
					}},
				}, {
					Name:        []string{"--limitMem"},
					Description: `The limit memory to deploy`,
					Args: []model.Arg{{
						Name: "limitMem",
					}},
				}, {
					Name:        []string{"--maxReplicas"},
					Description: `The max replicas to deploy`,
					Args: []model.Arg{{
						Name: "maxReplicas",
					}},
				}, {
					Name:        []string{"--minReplicas"},
					Description: `The min replicas to deploy`,
					Args: []model.Arg{{
						Name: "minReplicas",
					}},
				}, {
					Name:        []string{"--name"},
					Description: `The name of deployment (required)`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--namespace"},
					Description: `The namespace of deployment (required)`,
					Args: []model.Arg{{
						Name: "namespace",
					}},
				}, {
					Name:        []string{"--nodePort"},
					Description: `The nodePort of the deployment to expose`,
					Args: []model.Arg{{
						Name: "nodePort",
					}},
				}, {
					Name:        []string{"--o"},
					Description: `The output yaml file (required)`,
					Args: []model.Arg{{
						Name: "o",
					}},
				}, {
					Name:        []string{"--port"},
					Description: `The port of the deployment to listen on pod (required)`,
					Args: []model.Arg{{
						Name: "port",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--replicas"},
					Description: `The number of replicas to deploy`,
					Args: []model.Arg{{
						Name: "replicas",
					}},
				}, {
					Name:        []string{"--requestCpu"},
					Description: `The request cpu to deploy`,
					Args: []model.Arg{{
						Name: "requestCpu",
					}},
				}, {
					Name:        []string{"--requestMem"},
					Description: `The request memory to deploy`,
					Args: []model.Arg{{
						Name: "requestMem",
					}},
				}, {
					Name:        []string{"--revisions"},
					Description: `The number of revision history to limit`,
					Args: []model.Arg{{
						Name: "revisions",
					}},
				}, {
					Name:        []string{"--secret"},
					Description: `The secret to image pull from registry`,
					Args: []model.Arg{{
						Name: "secret",
					}},
				}, {
					Name:        []string{"--serviceAccount"},
					Description: `The ServiceAccount for the deployment`,
					Args: []model.Arg{{
						Name: "serviceAccount",
					}},
				}},
			}},
		}, {
			Name:        []string{"migrate"},
			Description: `Migrate from tal-tech to zeromicro`,
			Options: []model.Option{{
				Name:        []string{"--verbose", "-v"},
				Description: `Verbose enables extra logging`,
			}, {
				Name:        []string{"--version"},
				Description: `The target release version of github.com/zeromicro/go-zero to migrate`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}},
		}, {
			Name:        []string{"model"},
			Description: `Generate model code`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"mongo"},
				Description: `Generate mongo model`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--cache", "-c"},
					Description: `Generate code with cache [optional]`,
				}, {
					Name:        []string{"--dir", "-d"},
					Description: `The target dir`,
					Args: []model.Arg{{
						Name: "dir",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}, {
					Name:        []string{"--type", "-t"},
					Description: `Specified model type name`,
					Args: []model.Arg{{
						Name: "type",
					}},
				}},
			}, {
				Name:        []string{"mysql"},
				Description: `Generate mysql model`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"datasource"},
					Description: `Generate model from datasource`,
					Options: []model.Option{{
						Name:        []string{"--branch"},
						Description: `The branch of the remote repo, it does work with --remote`,
						Args: []model.Arg{{
							Name: "branch",
						}},
					}, {
						Name:        []string{"--cache", "-c"},
						Description: `Generate code with cache [optional]`,
					}, {
						Name:        []string{"--dir", "-d"},
						Description: `The target dir`,
						Args: []model.Arg{{
							Name: "dir",
						}},
					}, {
						Name:        []string{"--home"},
						Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
						Args: []model.Arg{{
							Name: "home",
						}},
					}, {
						Name:        []string{"--idea"},
						Description: `For idea plugin [optional]`,
					}, {
						Name:        []string{"--remote"},
						Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
						Args: []model.Arg{{
							Name: "remote",
						}},
					}, {
						Name:        []string{"--style"},
						Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
						Args: []model.Arg{{
							Name: "style",
						}},
					}, {
						Name:        []string{"--table", "-t"},
						Description: `The table or table globbing patterns in the database`,
						Args: []model.Arg{{
							Name: "table",
						}},
					}, {
						Name:        []string{"--url"},
						Description: `The data source of database,like "root:password@tcp(127.0.0.1:3306)/database"`,
						Args: []model.Arg{{
							Name: "url",
						}},
					}},
				}, {
					Name:        []string{"ddl"},
					Description: `Generate mysql model from ddl`,
					Options: []model.Option{{
						Name:        []string{"--branch"},
						Description: `The branch of the remote repo, it does work with --remote`,
						Args: []model.Arg{{
							Name: "branch",
						}},
					}, {
						Name:        []string{"--cache", "-c"},
						Description: `Generate code with cache [optional]`,
					}, {
						Name:        []string{"--database"},
						Description: `The name of database [optional]`,
						Args: []model.Arg{{
							Name: "database",
						}},
					}, {
						Name:        []string{"--dir", "-d"},
						Description: `The target dir`,
						Args: []model.Arg{{
							Name: "dir",
						}},
					}, {
						Name:        []string{"--home"},
						Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
						Args: []model.Arg{{
							Name: "home",
						}},
					}, {
						Name:        []string{"--idea"},
						Description: `For idea plugin [optional]`,
					}, {
						Name:        []string{"--remote"},
						Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
						Args: []model.Arg{{
							Name: "remote",
						}},
					}, {
						Name:        []string{"--src", "-s"},
						Description: `The path or path globbing patterns of the ddl`,
						Args: []model.Arg{{
							Name: "src",
						}},
					}, {
						Name:        []string{"--style"},
						Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
						Args: []model.Arg{{
							Name: "style",
						}},
					}},
				}, {
					Name:        []string{"pg"},
					Description: `Generate postgresql model`,
					Options: []model.Option{{
						Name:        []string{"--branch"},
						Description: `The branch of the remote repo, it does work with --remote`,
						Args: []model.Arg{{
							Name: "branch",
						}},
					}, {
						Name:        []string{"--cache", "-c"},
						Description: `Generate code with cache [optional]`,
					}, {
						Name:        []string{"--dir", "-d"},
						Description: `The target dir`,
						Args: []model.Arg{{
							Name: "dir",
						}},
					}, {
						Name:        []string{"--home"},
						Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
						Args: []model.Arg{{
							Name: "home",
						}},
					}, {
						Name:        []string{"--idea"},
						Description: `For idea plugin [optional]`,
					}, {
						Name:        []string{"--remote"},
						Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
						Args: []model.Arg{{
							Name: "remote",
						}},
					}, {
						Name:        []string{"--schema", "-s"},
						Description: `The table schema`,
						Args: []model.Arg{{
							Name: "schema",
						}},
					}, {
						Name:        []string{"--style"},
						Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
						Args: []model.Arg{{
							Name: "style",
						}},
					}, {
						Name:        []string{"--table", "-t"},
						Description: `The table or table globbing patterns in the database`,
						Args: []model.Arg{{
							Name: "table",
						}},
					}, {
						Name:        []string{"--url"},
						Description: `The data source of database,like "root:password@tcp(127.0.0.1:3306)/database"`,
						Args: []model.Arg{{
							Name: "url",
						}},
					}},
				}},
			}},
		}, {
			Name:        []string{"quickstart"},
			Description: `Quickly start a project`,
			Options: []model.Option{{
				Name:        []string{"--service-type", "-t"},
				Description: `Specify the service type, supported values: [mono, micro]`,
				Args: []model.Arg{{
					Name: "service-type",
				}},
			}},
		}, {
			Name:        []string{"rpc"},
			Description: `Generate rpc code`,
			Options: []model.Option{{
				Name:        []string{"--branch"},
				Description: `The branch of the remote repo, it does work with --remote`,
				Args: []model.Arg{{
					Name: "branch",
				}},
			}, {
				Name:        []string{"--home"},
				Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
				Args: []model.Arg{{
					Name: "home",
				}},
			}, {
				Name:        []string{"--o"},
				Description: `Output a sample proto file`,
				Args: []model.Arg{{
					Name: "o",
				}},
			}, {
				Name:        []string{"--remote"},
				Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
				Args: []model.Arg{{
					Name: "remote",
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"new"},
				Description: `Generate rpc demo service`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--go-grpc_opt"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go-grpc_opt",
					}},
				}, {
					Name:        []string{"--go_opt"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go_opt",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--idea"},
					Description: `Whether the command execution environment is from idea plugin`,
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}, {
					Name:        []string{"--verbose", "-v"},
					Description: `Enable log output`,
				}},
			}, {
				Name:        []string{"protoc"},
				Description: `Generate grpc code`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--go-grpc_opt"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go-grpc_opt",
					}},
				}, {
					Name:        []string{"--go-grpc_out"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go-grpc_out",
					}},
				}, {
					Name:        []string{"--go_opt"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go_opt",
					}},
				}, {
					Name:        []string{"--go_out"},
					Description: ``,
					Args: []model.Arg{{
						Name: "go_out",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--multiple", "-m"},
					Description: `Generated in multiple rpc service mode`,
				}, {
					Name:        []string{"--plugin"},
					Description: ``,
					Args: []model.Arg{{
						Name: "plugin",
					}},
				}, {
					Name:        []string{"--proto_path", "-I"},
					Description: ``,
					Args: []model.Arg{{
						Name: "proto_path",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}, {
					Name:        []string{"--style"},
					Description: `The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]`,
					Args: []model.Arg{{
						Name: "style",
					}},
				}, {
					Name:        []string{"--verbose", "-v"},
					Description: `Enable log output`,
				}, {
					Name:        []string{"--zrpc_out"},
					Description: `The zrpc output directory`,
					Args: []model.Arg{{
						Name: "zrpc_out",
					}},
				}},
			}, {
				Name:        []string{"template"},
				Description: `Generate proto template`,
				Options: []model.Option{{
					Name:        []string{"--branch"},
					Description: `The branch of the remote repo, it does work with --remote`,
					Args: []model.Arg{{
						Name: "branch",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--o"},
					Description: `Output a sample proto file`,
					Args: []model.Arg{{
						Name: "o",
					}},
				}, {
					Name:        []string{"--remote"},
					Description: `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority 	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`,
					Args: []model.Arg{{
						Name: "remote",
					}},
				}},
			}},
		}, {
			Name:        []string{"template"},
			Description: `Template operation`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"clean"},
				Description: `Clean the all cache templates`,
				Options: []model.Option{{
					Name:        []string{"--home"},
					Description: `The goctl home path of the template`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}},
			}, {
				Name:        []string{"init"},
				Description: `Initialize the all templates(force update)`,
				Options: []model.Option{{
					Name:        []string{"--home"},
					Description: `The goctl home path of the template`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}},
			}, {
				Name:        []string{"revert"},
				Description: `Revert the target template to the latest`,
				Options: []model.Option{{
					Name:        []string{"--category", "-c"},
					Description: `The category of template, enum [api,rpc,model,docker,kube]`,
					Args: []model.Arg{{
						Name: "category",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}, {
					Name:        []string{"--name", "-n"},
					Description: `The target file name of template`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}},
			}, {
				Name:        []string{"update"},
				Description: `Update template of the target category to the latest`,
				Options: []model.Option{{
					Name:        []string{"--category", "-c"},
					Description: `The category of template, enum [api,rpc,model,docker,kube]`,
					Args: []model.Arg{{
						Name: "category",
					}},
				}, {
					Name:        []string{"--home"},
					Description: `The goctl home path of the template`,
					Args: []model.Arg{{
						Name: "home",
					}},
				}},
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Upgrade goctl to latest version`,
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"api"},
				Description: `Generate api related files`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"dart"},
					Description: `Generate dart files for provided api in api file`,
				}, {
					Name:        []string{"doc"},
					Description: `Generate doc files`,
				}, {
					Name:        []string{"format"},
					Description: `Format api files`,
				}, {
					Name:        []string{"go"},
					Description: `Generate go files for provided api in yaml file`,
				}, {
					Name:        []string{"java"},
					Description: `Generate java files for provided api in api file`,
				}, {
					Name:        []string{"kt"},
					Description: `Generate kotlin code for provided api file`,
				}, {
					Name:        []string{"new"},
					Description: `Fast create api service`,
				}, {
					Name:        []string{"plugin"},
					Description: `Custom file generator`,
				}, {
					Name:        []string{"ts"},
					Description: `Generate ts files for provided api in api file`,
				}, {
					Name:        []string{"validate"},
					Description: `Validate api file`,
				}},
			}, {
				Name:        []string{"bug"},
				Description: `Report a bug`,
			}, {
				Name:        []string{"completion"},
				Description: `Generate the autocompletion script for the specified shell`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"bash"},
					Description: `Generate the autocompletion script for bash`,
				}, {
					Name:        []string{"fish"},
					Description: `Generate the autocompletion script for fish`,
				}, {
					Name:        []string{"powershell"},
					Description: `Generate the autocompletion script for powershell`,
				}, {
					Name:        []string{"zsh"},
					Description: `Generate the autocompletion script for zsh`,
				}},
			}, {
				Name:        []string{"docker"},
				Description: `Generate Dockerfile`,
			}, {
				Name:        []string{"env"},
				Description: `Check or edit goctl environment`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"check"},
					Description: `Detect goctl env and dependency tools`,
				}, {
					Name:        []string{"install"},
					Description: `Goctl env installation`,
				}},
			}, {
				Name:        []string{"kube"},
				Description: `Generate kubernetes files`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"deploy"},
					Description: `Generate deployment yaml file`,
				}},
			}, {
				Name:        []string{"migrate"},
				Description: `Migrate from tal-tech to zeromicro`,
			}, {
				Name:        []string{"model"},
				Description: `Generate model code`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"mongo"},
					Description: `Generate mongo model`,
				}, {
					Name:        []string{"mysql"},
					Description: `Generate mysql model`,
					Subcommands: []model.Subcommand{{
						Name:        []string{"datasource"},
						Description: `Generate model from datasource`,
					}, {
						Name:        []string{"ddl"},
						Description: `Generate mysql model from ddl`,
					}, {
						Name:        []string{"pg"},
						Description: `Generate postgresql model`,
					}},
				}},
			}, {
				Name:        []string{"quickstart"},
				Description: `Quickly start a project`,
			}, {
				Name:        []string{"rpc"},
				Description: `Generate rpc code`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"new"},
					Description: `Generate rpc demo service`,
				}, {
					Name:        []string{"protoc"},
					Description: `Generate grpc code`,
				}, {
					Name:        []string{"template"},
					Description: `Generate proto template`,
				}},
			}, {
				Name:        []string{"template"},
				Description: `Template operation`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"clean"},
					Description: `Clean the all cache templates`,
				}, {
					Name:        []string{"init"},
					Description: `Initialize the all templates(force update)`,
				}, {
					Name:        []string{"revert"},
					Description: `Revert the target template to the latest`,
				}, {
					Name:        []string{"update"},
					Description: `Update template of the target category to the latest`,
				}},
			}, {
				Name:        []string{"upgrade"},
				Description: `Upgrade goctl to latest version`,
			}},
		}},
	}
}
