// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["passwd"] = model.Subcommand{
		Name:        []string{"passwd"},
		Description: `Modify a user's password`,
		Options: []model.Option{{
			Name:        []string{"-i"},
			Description: `Specify where the password update should be applied`,
			Args: []model.Arg{{
				Name:        "infosystem",
				Description: `The directory system`,
				Suggestions: []model.Suggestion{{Name: []string{`PAM`}}, {Name: []string{`opendirectory`}}, {Name: []string{`file`}}, {Name: []string{`nis`}}},
			}},
		}, {
			Name:        []string{"-l"},
			Description: `Causes the password to be updated in the given location of the chosen directory system`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths, model.TemplateFolders},
				Name:        "location",
				Description: `The location of the chosen directory system`,
			}},
		}, {
			Name:        []string{"-u"},
			Description: `Specify the user name to use when authenticating to the directory node`,
			Args: []model.Arg{{
				Name:        "authname",
				Description: `The user name`,
				Generator:   nil, // TODO: port over generator
			}},
		}},
	}
}