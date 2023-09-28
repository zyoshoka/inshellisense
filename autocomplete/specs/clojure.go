// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["clojure"] = model.Subcommand{
		Name:        []string{"clojure"},
		Description: `Use the Clojure tools to run Clojure programs on the JVM, start a REPL, or invoke a specific function with data`,
		Options: []model.Option{{
			Name:        []string{"-A"},
			Description: `Use concatenated aliases to modify classpath`,
		}, {
			Name: []string{"-X"},
			Args: []model.Arg{{
				Name:        "a/fn",
				Description: `An alias to refer to its function or a qualified function`,
				IsOptional:  true,
			}, {
				Name:        "kvs",
				Description: `A list of key-value arguments to merge with the exec-args`,
				IsOptional:  true,
				IsVariadic:  true,
			}, {
				Name:        "kv-map",
				Description: `Map to merge with the :exec-args map, taking precedence over kvs`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"-T"},
			Description: `Invoke tool by name or via aliases ala -X`,
			Args: []model.Arg{{
				Name:        "a/fn",
				Description: `An alias to refer to its function or a qualified function`,
				IsOptional:  true,
			}, {
				Name:        "kvs",
				Description: `A list of key-value arguments to merge with the exec-args`,
				IsOptional:  true,
				IsVariadic:  true,
			}, {
				Name:        "kv-map",
				Description: `Map to merge with the :exec-args map, taking precedence over kvs`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"-M"},
			Description: `Use concatenated aliases to modify classpath or supply main opts`,
			Args: []model.Arg{{
				Name:       "args",
				IsOptional: true,
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"-P"},
			Description: `Prepare deps - download libs, cache classpath, but don't exec`,
		}, {
			Name:        []string{"-J"},
			Description: `Pass opt through in java_opts`,
		}, {
			Name:        []string{"-Sdeps"},
			Description: `Pass the deps data on the command line`,
			Args: []model.Arg{{
				Name:        "edn",
				Description: `The deps data in edn`,
			}},
		}, {
			Name:        []string{"-Spath"},
			Description: `Compute classpath and echo to stdout only`,
		}, {
			Name:        []string{"-Scp"},
			Description: `Use specified classpath instead of cached or computed one`,
			Args: []model.Arg{{
				Name:        "cp",
				Description: `The classpath to use`,
			}},
		}, {
			Name:        []string{"-Sdescribe"},
			Description: `Print environment and command parsing information as data`,
		}, {
			Name:        []string{"-Sforce"},
			Description: `Ignore classpath cache and force recomputation`,
		}, {
			Name:        []string{"-Spom"},
			Description: `Generate (or update) pom.xml with deps and paths`,
		}, {
			Name:        []string{"-Srepro"},
			Description: `Ignore the ~/.clojure/deps.edn config file`,
		}, {
			Name:        []string{"-Sthreads"},
			Description: `Set the number of threads to use when downloading dependencies`,
		}, {
			Name:        []string{"-Strace"},
			Description: `Write a trace.edn file that traces deps expansion`,
		}, {
			Name:        []string{"-Stree"},
			Description: `Print dependency tree`,
		}, {
			Name:        []string{"-Sverbose"},
			Description: `Print all path locations`,
		}, {
			Name:        []string{"-version", "--version"},
			Description: `Print the Clojure CLI version`,
		}, {
			Name:        []string{"-i", "--init"},
			Description: `Load a file or resource`,
			Args: []model.Arg{{
				Name: "path",
			}},
		}, {
			Name:        []string{"-e", "--eval"},
			Description: `Evaluate expressions in string; print non-nil values`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--report"},
			Description: `Report uncaught exceptions`,
			Args: []model.Arg{{
				Name:        "target",
				Description: `Where to report`,
				Suggestions: []model.Suggestion{{Name: []string{`file`}}, {Name: []string{`stderr`}}, {Name: []string{`none`}}},
			}},
		}, {
			Name:        []string{"-m", "--main"},
			Description: `Call the -main function from a namespace with args`,
			Args: []model.Arg{{
				Name:        "ns-name",
				Description: `The namespace of the -main function`,
			}, {
				Name:        "args",
				Description: `The arguments to pass to the -main function`,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"-r", "--repl"},
			Description: `Run a REPL`,
		}, {
			Name:        []string{"-h", "-?", "--help"},
			Description: `Show help for clojure`,
		}},
	}
}
