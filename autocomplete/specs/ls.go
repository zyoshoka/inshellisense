// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["ls"] = model.Subcommand{
		Name:        []string{"ls"},
		Description: `List directory contents`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-@"},
			Description: `Display extended attribute keys and sizes in long (-l) output`,
		}, {
			Name:        []string{"-1"},
			Description: `(The numeric digit ""one''.)  Force output to be one entry per line.  This is the default when output is not to a terminal`,
		}, {
			Name:        []string{"-A"},
			Description: `List all entries except for . and ...  Always set for the super-user`,
		}, {
			Name:        []string{"-a"},
			Description: `Include directory entries whose names begin with a dot (.)`,
		}, {
			Name:        []string{"-B"},
			Description: `Force printing of non-printable characters (as defined by ctype(3) and current locale settings) in file names as xxx, where xxx is the numeric value of the character in octal`,
		}, {
			Name:        []string{"-b"},
			Description: `As -B, but use C escape codes whenever possible`,
		}, {
			Name:        []string{"-C"},
			Description: `Force multi-column output; this is the default when output is to a terminal`,
		}, {
			Name:        []string{"-c"},
			Description: `Use time when file status was last changed for sorting (-t) or long printing (-l)`,
		}, {
			Name:        []string{"-d"},
			Description: `Directories are listed as plain files (not searched recursively)`,
		}, {
			Name:        []string{"-e"},
			Description: `Print the Access Control List (ACL) associated with the file, if present, in long (-l) output`,
		}, {
			Name:        []string{"-F"},
			Description: `Display a slash (/) immediately after each pathname that is a directory, an asterisk (*) after each that is executable, an at sign (@) after each symbolic link, an equals sign (=) after each socket, a percent sign (%) after each whiteout, and a vertical bar (|) after each that is a FIFO`,
		}, {
			Name:        []string{"-f"},
			Description: `Output is not sorted.  This option turns on the -a option`,
		}, {
			Name:        []string{"-G"},
			Description: `Enable colorized output.  This option is equivalent to defining CLICOLOR in the environment.  (See below.)`,
		}, {
			Name:        []string{"-g"},
			Description: `This option is only available for compatibility with POSIX; it is used to display the group name in the long (-l) format output (the owner name is suppressed)`,
		}, {
			Name:        []string{"-H"},
			Description: `Symbolic links on the command line are followed.  This option is assumed if none of the -F, -d, or -l options are specified`,
		}, {
			Name:        []string{"-h"},
			Description: `When used with the -l option, use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte in order to reduce the number of digits to three or less using base 2 for sizes`,
		}, {
			Name:        []string{"-i"},
			Description: `For each file, print the file's file serial number (inode number)`,
		}, {
			Name:        []string{"-k"},
			Description: `If the -s option is specified, print the file size allocation in kilobytes, not blocks.  This option overrides the environment variable BLOCKSIZE`,
		}, {
			Name:        []string{"-L"},
			Description: `Follow all symbolic links to final target and list the file or directory the link references rather than the link itself.  This option cancels the -P option`,
		}, {
			Name:        []string{"-l"},
			Description: `(The lowercase letter ""ell''.)  List in long format.  (See below.)  A total sum for all the file sizes is output on a line before the long listing`,
		}, {
			Name:        []string{"-m"},
			Description: `Stream output format; list files across the page, separated by commas`,
		}, {
			Name:        []string{"-n"},
			Description: `Display user and group IDs numerically, rather than converting to a user or group name in a long (-l) output.  This option turns on the -l option`,
		}, {
			Name:        []string{"-O"},
			Description: `Include the file flags in a long (-l) output`,
		}, {
			Name:        []string{"-o"},
			Description: `List in long format, but omit the group id`,
		}, {
			Name:        []string{"-P"},
			Description: `If argument is a symbolic link, list the link itself rather than the object the link references.  This option cancels the -H and -L options`,
		}, {
			Name:        []string{"-p"},
			Description: `Write a slash ("/') after each filename if that file is a directory`,
		}, {
			Name:        []string{"-q"},
			Description: `Force printing of non-graphic characters in file names as the character "?'; this is the default when output is to a terminal`,
		}, {
			Name:        []string{"-R"},
			Description: `Recursively list subdirectories encountered`,
		}, {
			Name:        []string{"-r"},
			Description: `Reverse the order of the sort to get reverse lexicographical order or the oldest entries first (or largest files last, if combined with sort by size`,
		}, {
			Name:        []string{"-S"},
			Description: `Sort files by size`,
		}, {
			Name:        []string{"-s"},
			Description: `Display the number of file system blocks actually used by each file, in units of 512 bytes, where partial units are rounded up to the next integer value.  If the output is to a terminal, a total sum for all the file sizes is output on a line before the listing.  The environment variable BLOCKSIZE overrides the unit size of 512 bytes`,
		}, {
			Name:        []string{"-T"},
			Description: `When used with the -l (lowercase letter ""ell'') option, display complete time information for the file, including month, day, hour, minute, second, and year`,
		}, {
			Name:        []string{"-t"},
			Description: `Sort by time modified (most recently modified first) before sorting the operands by lexicographical order`,
		}, {
			Name:        []string{"-u"},
			Description: `Use time of last access, instead of last modification of the file for sorting (-t) or long printing (-l)`,
		}, {
			Name:        []string{"-U"},
			Description: `Use time of file creation, instead of last modification for sorting (-t) or long output (-l)`,
		}, {
			Name:        []string{"-v"},
			Description: `Force unedited printing of non-graphic characters; this is the default when output is not to a terminal`,
		}, {
			Name:        []string{"-W"},
			Description: `Display whiteouts when scanning directories.  (-S) flag)`,
		}, {
			Name:        []string{"-w"},
			Description: `Force raw printing of non-printable characters.  This is the default when output is not to a terminal`,
		}, {
			Name:        []string{"-x"},
			Description: `The same as -C, except that the multi-column output is produced with entries sorted across, rather than down, the columns`,
		}, {
			Name:        []string{"-%"},
			Description: `Distinguish dataless files and directories with a '%' character in long (-l) output, and don't materialize dataless directories when listing them`,
		}, {
			Name: []string{"-,"},
			Description: `When the -l option is set, print file sizes grouped and separated by thousands using the non-monetary separator returned
by localeconv(3), typically a comma or period.  If no locale is set, or the locale does not have a non-monetary separator, this
option has no effect.  This option is not defined in IEEE Std 1003.1-2001 (“POSIX.1”)`,
		}, {
			Name:        []string{"--color"},
			Description: `Output colored escape sequences based on when, which may be set to either always, auto, or never`,
			Args: []model.Arg{{
				Name: "when",
				Suggestions: []model.Suggestion{{
					Name:        []string{`always`, `yes`, `force`},
					Description: `Will make ls always output color`,
				}, {
					Name:        []string{`auto`},
					Description: `Will make ls output escape sequences based on termcap(5), but only if stdout is a tty and either the -G flag is specified or the COLORTERM environment variable is set and not empty`,
				}, {
					Name:        []string{`never`, `no`, `none`},
					Description: `Will disable color regardless of environment variables`,
				}},
			}},
		}},
	}
}
