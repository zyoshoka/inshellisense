// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["pre-commit"] = model.Subcommand{
		Name:        []string{"pre-commit"},
		Description: `Pre-commit`,
		Args:        []model.Arg{{}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Show help message and exit`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Show program's version number and exit`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"autoupdate"},
			Description: `Auto-update pre-commit config to the latest repos' versions`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name:        []string{"--bleeding-edge"},
				Description: `Update to the bleeding edge of "master" instead of the latest tagged version (the default behavior)`,
			}, {
				Name:        []string{"--freeze"},
				Description: `Store 'frozen' hashes in "rev" instead of tag names`,
			}, {
				Name:        []string{"--repo"},
				Description: `Only update this repository -- may be specified multiple times`,
				Args: []model.Arg{{
					Name: "REPO",
				}},
			}},
		}, {
			Name:        []string{"clean"},
			Description: `Clean out pre-commit files`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}},
		}, {
			Name:        []string{"gc"},
			Description: `Clean unused cached repos`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}},
		}, {
			Name:        []string{"init-templatedir"},
			Description: `Install hook script in a directory intended for use with "git config init.templateDir"`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name:        []string{"-t", "--hook-type"},
				Description: `Type of hook to install`,
				Args: []model.Arg{{
					Name:        "HOOK_TYPE",
					Suggestions: []model.Suggestion{{Name: []string{`pre-commit`}}, {Name: []string{`pre-merge-commit`}}, {Name: []string{`pre-push`}}, {Name: []string{`prepare-commit-msg`}}, {Name: []string{`commit-msg`}}, {Name: []string{`post-commit`}}, {Name: []string{`post-checkout`}}, {Name: []string{`post-merge`}}, {Name: []string{`post-rewrite`}}},
				}},
			}, {
				Name:        []string{"--no-allow-missing-config"},
				Description: `Assume cloned repos should have a "pre-commit" config`,
			}},
		}, {
			Name:        []string{"install"},
			Description: `Install the pre-commit script`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name:        []string{"-f", "--overwrite"},
				Description: `Overwrite existing hooks / remove migration mode`,
			}, {
				Name:        []string{"install-hooks"},
				Description: `Whether to install hook environments for all environments in the config file`,
			}, {
				Name:        []string{"-t", "--hook-type"},
				Description: `Type of hook to install`,
				Args: []model.Arg{{
					Name:        "HOOK_TYPE",
					Suggestions: []model.Suggestion{{Name: []string{`pre-commit`}}, {Name: []string{`pre-merge-commit`}}, {Name: []string{`pre-push`}}, {Name: []string{`prepare-commit-msg`}}, {Name: []string{`commit-msg`}}, {Name: []string{`post-commit`}}, {Name: []string{`post-checkout`}}, {Name: []string{`post-merge`}}, {Name: []string{`post-rewrite`}}},
				}},
			}, {
				Name:        []string{"--allow-missing-config"},
				Description: `Whether to allow a missing "pre-commit" configuration file or exit with a failure code`,
			}},
		}, {
			Name:        []string{"migrate-config"},
			Description: `Migrate list configuration to new map configuration`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run hooks`,
			Args: []model.Arg{{
				Name:        "hook",
				Description: `A single hook-id to run`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name: []string{"--verbose", "-v"},
			}, {
				Name:        []string{"--all-files", "-a"},
				Description: `Run all files in the repo`,
			}, {
				Name:        []string{"--files"},
				Description: `Specific filenames to run hooks on`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
					Name:       "FILES",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--show-diff-on-failure"},
				Description: `When hooks fail, run "git diff" directly afterward`,
			}, {
				Name:        []string{"--hook-stage"},
				Description: `The stage during which the hook is fired`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`pre-commit`}}, {Name: []string{`pre-merge-commit`}}, {Name: []string{`pre-push`}}, {Name: []string{`prepare-commit-msg`}}, {Name: []string{`commit-msg`}}, {Name: []string{`post-commit`}}, {Name: []string{`post-checkout`}}, {Name: []string{`post-merge`}}, {Name: []string{`post-rewrite`}}},
				}},
			}, {
				Name:        []string{"--remote-branch"},
				Description: `Remote branch ref used by "git push"`,
				Args: []model.Arg{{
					Name:      "REMOTE_BRANCH",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--local-branch"},
				Description: `Local branch ref used by "git push"`,
				Args: []model.Arg{{
					Name:      "LOCAL_BRANCH",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--from-ref", "--source", "-s"},
				Description: `Represents the original ref in a "from_ref...to_ref" diff expression. For "pre-push" hooks, this represents the branch you are pushing to. For "post-checkout" hooks, this represents the branch that was previously checked out`,
				Args: []model.Arg{{
					Name:      "FROM_REF",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--to-ref", "--origin", "-o"},
				Description: `Represents the destination ref in a "from_ref...to_ref" diff expression. For "pre-push" hooks, this represents the branch being pushed. For "post-checkout" hooks, this represents the branch that is now checked out`,
				Args: []model.Arg{{
					Name:      "TO_REF",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--commit-msg-filename"},
				Description: `Filename to check when running during "commit-msg"`,
				Args: []model.Arg{{
					Name: "COMMIT_MSG_FILENAME",
				}},
			}, {
				Name:        []string{"--remote-name"},
				Description: `Remote name used by "git push"`,
				Args: []model.Arg{{
					Name:      "REMOTE_NAME",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--remote-url"},
				Description: `Remote URL used by "git push"`,
				Args: []model.Arg{{
					Name: "REMOTE_URL",
				}},
			}, {
				Name:        []string{"--checkout-type"},
				Description: `Indicates whether the checkout was a branch checkout (changing branches, flag=1) or a file checkout (retrieving a file from the index, flag=0)`,
				Args: []model.Arg{{
					Name: "CHECKOUT_TYPE",
				}},
			}, {
				Name:        []string{"--is-squash-merge"},
				Description: `During a post-merge hook, indicates whether the merge was a squash merge`,
				Args: []model.Arg{{
					Name: "IS_SQUASH_MERGE",
				}},
			}, {
				Name:        []string{"--rewrite-command"},
				Description: `During a post-rewrite hook, specifies the command that invoked the rewrite`,
				Args: []model.Arg{{
					Name: "REWRITE_COMMAND",
				}},
			}},
		}, {
			Name:        []string{"sample-config"},
			Description: `Produce a sample .pre-commit-config.yaml file`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}},
		}, {
			Name:        []string{"try-repo"},
			Description: `Try the hooks in a repository, useful for developing new hooks`,
			Args: []model.Arg{{
				Name:        "repo",
				Description: `Repository to source hooks from`,
			}, {
				Name:        "hook",
				Description: `A single hook-id to run`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name: []string{"--verbose", "-v"},
			}, {
				Name:        []string{"--all-files", "-a"},
				Description: `Run all files in the repo`,
			}, {
				Name:        []string{"--files"},
				Description: `Specific filenames to run hooks on`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
					Name:       "FILES",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--show-diff-on-failure"},
				Description: `When hooks fail, run "git diff" directly afterward`,
			}, {
				Name:        []string{"--hook-stage"},
				Description: `The stage during which the hook is fired`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`pre-commit`}}, {Name: []string{`pre-merge-commit`}}, {Name: []string{`pre-push`}}, {Name: []string{`prepare-commit-msg`}}, {Name: []string{`commit-msg`}}, {Name: []string{`post-commit`}}, {Name: []string{`post-checkout`}}, {Name: []string{`post-merge`}}, {Name: []string{`post-rewrite`}}},
				}},
			}, {
				Name:        []string{"--remote-branch"},
				Description: `Remote branch ref used by "git push"`,
				Args: []model.Arg{{
					Name:      "REMOTE_BRANCH",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--local-branch"},
				Description: `Local branch ref used by "git push"`,
				Args: []model.Arg{{
					Name:      "LOCAL_BRANCH",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--from-ref", "--source", "-s"},
				Description: `Represents the original ref in a "from_ref...to_ref" diff expression. For "pre-push" hooks, this represents the branch you are pushing to. For "post-checkout" hooks, this represents the branch that was previously checked out`,
				Args: []model.Arg{{
					Name:      "FROM_REF",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--to-ref", "--origin", "-o"},
				Description: `Represents the destination ref in a "from_ref...to_ref" diff expression. For "pre-push" hooks, this represents the branch being pushed. For "post-checkout" hooks, this represents the branch that is now checked out`,
				Args: []model.Arg{{
					Name:      "TO_REF",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--commit-msg-filename"},
				Description: `Filename to check when running during "commit-msg"`,
				Args: []model.Arg{{
					Name: "COMMIT_MSG_FILENAME",
				}},
			}, {
				Name:        []string{"--remote-name"},
				Description: `Remote name used by "git push"`,
				Args: []model.Arg{{
					Name:      "REMOTE_NAME",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--remote-url"},
				Description: `Remote URL used by "git push"`,
				Args: []model.Arg{{
					Name: "REMOTE_URL",
				}},
			}, {
				Name:        []string{"--checkout-type"},
				Description: `Indicates whether the checkout was a branch checkout (changing branches, flag=1) or a file checkout (retrieving a file from the index, flag=0)`,
				Args: []model.Arg{{
					Name: "CHECKOUT_TYPE",
				}},
			}, {
				Name:        []string{"--is-squash-merge"},
				Description: `During a post-merge hook, indicates whether the merge was a squash merge`,
				Args: []model.Arg{{
					Name: "IS_SQUASH_MERGE",
				}},
			}, {
				Name:        []string{"--rewrite-command"},
				Description: `During a post-rewrite hook, specifies the command that invoked the rewrite`,
				Args: []model.Arg{{
					Name: "REWRITE_COMMAND",
				}},
			}, {
				Name:        []string{"--ref", "--rev"},
				Description: `Manually select a rev to run against, otherwise the "HEAD" revision will be used`,
				Args: []model.Arg{{
					Name: "REV",
				}},
			}},
		}, {
			Name:        []string{"uninstall"},
			Description: `Uninstall the pre-commit script`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}, {
				Name:        []string{"-t", "--hook-type"},
				Description: `Type of hook to install`,
				Args: []model.Arg{{
					Name:        "HOOK_TYPE",
					Suggestions: []model.Suggestion{{Name: []string{`pre-commit`}}, {Name: []string{`pre-merge-commit`}}, {Name: []string{`pre-push`}}, {Name: []string{`prepare-commit-msg`}}, {Name: []string{`commit-msg`}}, {Name: []string{`post-commit`}}, {Name: []string{`post-checkout`}}, {Name: []string{`post-merge`}}, {Name: []string{`post-rewrite`}}},
				}},
			}},
		}, {
			Name:        []string{"help"},
			Description: `Show help for a specific command`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Show help message and exit`,
			}, {
				Name:        []string{"--color"},
				Description: `Whether to use color in output. Defaults to "auto"`,
				Args: []model.Arg{{
					Name:        "color",
					Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`always`}}, {Name: []string{`never`}}},
				}},
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Path to alternate config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "CONFIG",
				}},
			}},
		}},
	}
}
