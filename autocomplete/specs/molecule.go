// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["molecule"] = model.Subcommand{
		Name:        []string{"molecule"},
		Description: `Molecule aids in the development and testing of Ansible roles`,
		Options: []model.Option{{
			Name:         []string{"--help"},
			Description:  `Show help`,
			IsPersistent: true,
		}, {
			Name:        []string{"--version"},
			Description: `Show molecule version`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"check"},
			Description: `Use the provisioner to perform a Dry-Run (destroy, dependency, create, prepare, converge)`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--parallel"},
				Description: `Enable parallel mode`,
				ExclusiveOn: []string{"--no-parallel"},
			}, {
				Name:        []string{"--no-parallel"},
				Description: `Disable parallel mode`,
				ExclusiveOn: []string{"--parallel"},
			}},
		}, {
			Name:        []string{"cleanup"},
			Description: `Use the provisioner to cleanup any changes made to external systems during the stages of testing`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"converge"},
			Description: `Use the provisioner to configure instances (dependency, create, prepare converge)`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"create"},
			Description: `Use the provisioner to start the instances`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--driver-name", "-d"},
				Description: `Name of driver to use. (delegated)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
				}},
			}},
		}, {
			Name:        []string{"dependency"},
			Description: `Manage the role's dependencies`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"destroy"},
			Description: `Use the provisioner to destroy the instances`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--driver-name", "-d"},
				Description: `Name of driver to use. (delegated)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
				}},
			}, {
				Name:        []string{"--all"},
				Description: `Destroy all scenarios. Default is False`,
				ExclusiveOn: []string{"--no-all"},
			}, {
				Name:        []string{"--no-all"},
				Description: `Do not destroy all scenarios`,
				ExclusiveOn: []string{"--all"},
			}, {
				Name:        []string{"--parallel"},
				Description: `Enable parallel mode`,
				ExclusiveOn: []string{"--no-parallel"},
			}, {
				Name:        []string{"--no-parallel"},
				Description: `Disable parallel mode`,
				ExclusiveOn: []string{"--parallel"},
			}},
		}, {
			Name:        []string{"drivers"},
			Description: `List drivers`,
			Options: []model.Option{{
				Name:        []string{"--format", "-f"},
				Description: `Change output format. (simple)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`simple`}}, {Name: []string{`plain`}}},
				}},
			}},
		}, {
			Name:        []string{"idempotence"},
			Description: `Use the provisioner to configure the instances and parse the output to determine idempotence`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"init"},
			Description: `Initialize a new role or scenario`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"role"},
				Description: `Initialize a new role for use with Molecule, namespace is required outside collections, like acme.myrole`,
				Args: []model.Arg{{
					Name: "ROLE_NAME",
				}},
				Options: []model.Option{{
					Name:        []string{"--dependency-name"},
					Description: `Name of dependency to initialize. (galaxy)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`galaxy`}}},
					}},
				}, {
					Name:        []string{"--driver-name", "-d"},
					Description: `Name of driver to use. (delegated)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
					}},
				}, {
					Name:        []string{"--lint-name"},
					Description: `Name of lint to initialize. (yamllint)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`yamllint`}}},
					}},
				}, {
					Name:        []string{"--provisioner-name"},
					Description: `Name of provisioner to initialize. (ansible)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`ansible`}}},
					}},
				}, {
					Name:        []string{"--verifier-name"},
					Description: `Name of verifier to initialize. (ansible)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`ansible`}}, {Name: []string{`testinfra`}}},
					}},
				}},
			}, {
				Name:        []string{"scenario"},
				Description: `Initialize a new scenario for use with Molecule`,
				Args: []model.Arg{{
					Name: "SCENARIO_NAME",
				}},
				Options: []model.Option{{
					Name:        []string{"--dependency-name"},
					Description: `Name of dependency to initialize. (galaxy)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`galaxy`}}},
					}},
				}, {
					Name:        []string{"--driver-name", "-d"},
					Description: `Name of driver to use. (delegated)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
					}},
				}, {
					Name:        []string{"--lint-name"},
					Description: `Name of lint to initialize. (yamllint)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`yamllint`}}},
					}},
				}, {
					Name:        []string{"--provisioner-name"},
					Description: `Name of provisioner to initialize. (ansible)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`ansible`}}},
					}},
				}, {
					Name:        []string{"--role-name", "-r"},
					Description: `Name of the role to create`,
					Args: []model.Arg{{
						Name: "TEXT",
					}},
				}, {
					Name:        []string{"--verifier-name"},
					Description: `Name of verifier to initialize. (ansible)`,
					Args: []model.Arg{{
						Name:        "TEXT",
						Suggestions: []model.Suggestion{{Name: []string{`ansible`}}, {Name: []string{`testinfra`}}},
					}},
				}},
			}},
		}, {
			Name:        []string{"lint"},
			Description: `Lint the role (dependency, lint)`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"list"},
			Description: `List status of instances`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--format", "-f"},
				Description: `Change output format. (simple)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`simple`}}, {Name: []string{`plain`}}, {Name: []string{`yaml`}}},
				}},
			}},
		}, {
			Name:        []string{"login"},
			Description: `Log in to one instance`,
			Options: []model.Option{{
				Name:        []string{"--host", "-h"},
				Description: `Host to access`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"matrix"},
			Description: `List matrix of steps used to test instances`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"prepare"},
			Description: `Use the provisioner to prepare the instances into a particular starting state`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--driver-name", "-d"},
				Description: `Name of driver to use. (delegated)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
				}},
			}, {
				Name:        []string{"--force"},
				Description: `Enable force mode. Default is disabled`,
				ExclusiveOn: []string{"--no-force"},
			}, {
				Name:        []string{"--no-force"},
				Description: `Disable force mode`,
				ExclusiveOn: []string{"--force"},
			}},
		}, {
			Name:        []string{"reset"},
			Description: `Reset molecule temporary folders`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"side-effect"},
			Description: `Use the provisioner to perform side-effects to the instances`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"syntax"},
			Description: `Use the provisioner to syntax check the role`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}},
		}, {
			Name:        []string{"test"},
			Description: `Test (dependency, lint, cleanup, destroy, syntax, create, prepare, converge, idempotence, side_effect, verify, cleanup, destroy)`,
			Options: []model.Option{{
				Name:        []string{"--scenario-name", "-s"},
				Description: `Name of the scenario to target. (default)`,
				Args: []model.Arg{{
					Name: "TEXT",
				}},
			}, {
				Name:        []string{"--driver-name", "-d"},
				Description: `Name of driver to use. (delegated)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`delegated`}}, {Name: []string{`docker`}}, {Name: []string{`podman`}}, {Name: []string{`vagrant`}}, {Name: []string{`azure`}}, {Name: []string{`hetzner`}}},
				}},
			}, {
				Name:        []string{"--all"},
				Description: `Test all scenarios. Default is False`,
				ExclusiveOn: []string{"--no-all"},
			}, {
				Name:        []string{"--no-all"},
				Description: `Do not test all scenarios`,
				ExclusiveOn: []string{"--all"},
			}, {
				Name:        []string{"--destroy"},
				Description: `The destroy strategy used at the conclusion of a Molecule run (always)`,
				Args: []model.Arg{{
					Name:        "TEXT",
					Suggestions: []model.Suggestion{{Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--parallel"},
				Description: `Enable parallel mode`,
				ExclusiveOn: []string{"--no-parallel"},
			}, {
				Name:        []string{"--no-parallel"},
				Description: `Disable parallel mode`,
				ExclusiveOn: []string{"--parallel"},
			}},
		}},
	}
}
