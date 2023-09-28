// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["gource"] = model.Subcommand{
		Name:        []string{"gource"},
		Description: `Gource is an OpenGL-based 3D visualisation tool for source control repositories`,
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Help ('-H' for extended help)`,
		}, {
			Name:        []string{"--viewport"},
			Description: `Set the viewport size. If -f is also supplied, will attempt to set the video mode to this also. Add "!" to make the window non-resizable`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `The size of the viewport (e.g. WIDTHxHEIGHT)`,
				Suggestions: []model.Suggestion{{Name: []string{`3840x2160`}}, {Name: []string{`2560x1440`}}, {Name: []string{`1920x1080`}}, {Name: []string{`1280x720`}}, {Name: []string{`800x600`}}},
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Fullscreen`,
		}, {
			Name:        []string{"--screen"},
			Description: `Set the number of the screen to display on`,
			Args: []model.Arg{{
				Name:        "screen",
				Description: `The number of the screen to display on`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--high-dpi"},
			Description: `Request a high DPI display when creating the window. On some platforms such as MacOS, the window resolution is specified in points instead of pixels. The --high-dpi flag may be required to access some higher resolutions. E.g. requesting a high DPI 800x600 window may produce a window that is 1600x1200 pixels`,
		}, {
			Name:        []string{"--window-position"},
			Description: `Initial window position on your desktop which may be made up of multiple monitors. This will override the screen setting so don't specify both`,
			Args: []model.Arg{{
				Name:        "position",
				Description: `The initial window position on your desktop (e.g. XxY)`,
			}},
		}, {
			Name:        []string{"--frameless"},
			Description: `Frameless window`,
		}, {
			Name:        []string{"--transparent"},
			Description: `Make the background transparent. Only really useful for screenshots`,
		}, {
			Name:        []string{"--start-date"},
			Description: `Start with the first entry after the supplied date and optional time. If a time zone offset isn't specified the local time zone is used. Example accepted formats: 'YYYY-MM-DD', 'YYYY-MM-DD hh:mm', 'YYYY-MM-DD hh:mm:ss +tz'`,
			Args: []model.Arg{{
				Name:        "date",
				Description: `The start date and optional time`,
			}},
		}, {
			Name:        []string{"--stop-date"},
			Description: `Stop at the last entry prior to the supplied date and optional time. Uses the same format as --start-date`,
			Args: []model.Arg{{
				Name:        "date",
				Description: `The stop date and optional time`,
			}},
		}, {
			Name:        []string{"-p", "--start-position"},
			Description: `Begin at some position in the log (between 0.0 and 1.0 or 'random')`,
			Args: []model.Arg{{
				Name:        "position",
				Description: `The start position in the log`,
				Suggestions: []model.Suggestion{{Name: []string{`random`}}, {Name: []string{`0.0`}}, {Name: []string{`0.1`}}, {Name: []string{`0.2`}}, {Name: []string{`0.3`}}, {Name: []string{`0.4`}}, {Name: []string{`0.5`}}, {Name: []string{`0.6`}}, {Name: []string{`0.7`}}, {Name: []string{`0.8`}}, {Name: []string{`0.9`}}, {Name: []string{`1.0`}}},
			}},
		}, {
			Name:        []string{"--stop-position"},
			Description: `Stop (exit) at some position in the log (does not work with STDIN)`,
			Args: []model.Arg{{
				Name:        "position",
				Description: `The stop position in the log`,
				Suggestions: []model.Suggestion{{Name: []string{`random`}}, {Name: []string{`0.0`}}, {Name: []string{`0.1`}}, {Name: []string{`0.2`}}, {Name: []string{`0.3`}}, {Name: []string{`0.4`}}, {Name: []string{`0.5`}}, {Name: []string{`0.6`}}, {Name: []string{`0.7`}}, {Name: []string{`0.8`}}, {Name: []string{`0.9`}}, {Name: []string{`1.0`}}},
			}},
		}, {
			Name:        []string{"-t", "--stop-at-time"},
			Description: `Stop (exit) after a specified number of seconds`,
			Args: []model.Arg{{
				Name:        "seconds",
				Description: `The number of seconds to wait before stopping`,
			}},
		}, {
			Name:        []string{"--stop-at-end"},
			Description: `Stop (exit) at the end of the log / stream`,
		}, {
			Name:        []string{"--loop"},
			Description: `Loop back to the start of the log when the end is reached`,
		}, {
			Name:        []string{"--loop-delay-seconds"},
			Description: `Seconds to delay before looping`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-a", "--auto-skip-seconds"},
			Description: `Automatically skip to next entry if nothing happens for a specified number of seconds`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-s", "--seconds-per-day"},
			Description: `Speed of simulation in seconds per day`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"--realtime"},
			Description: `Realtime playback speed`,
		}, {
			Name:        []string{"--no-time-travel"},
			Description: `Use the time of the last commit if the time of a commit is in the past`,
		}, {
			Name:        []string{"-c", "--time-scale"},
			Description: `Change simulation time scale`,
		}, {
			Name:        []string{"-i", "--file-idle-time"},
			Description: `Time in seconds files remain idle before they are removed or 0 for no limit`,
		}, {
			Name:        []string{"--file-idle-time-at-end"},
			Description: `Time in seconds files remain idle at the end before they are removed`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-e", "--elasticity"},
			Description: `Elasticity of nodes`,
			Args: []model.Arg{{
				Name:        "elasticity",
				Description: `Elasticity of nodes`,
			}},
		}, {
			Name:        []string{"-b", "--background-colour"},
			Description: `Background colour in hex`,
			Args: []model.Arg{{
				Name: "colour",
			}},
		}, {
			Name:        []string{"--background-image"},
			Description: `Set a background image`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "image file",
			}},
		}, {
			Name:        []string{"--title"},
			Description: `Set a title`,
			Args: []model.Arg{{
				Name: "title",
			}},
		}, {
			Name:        []string{"--font-file"},
			Description: `Specify the font. Should work with most font file formats supported by FreeType, such as TTF and OTF, among others`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "font file",
			}},
		}, {
			Name:        []string{"--font-scale"},
			Description: `Scale the size of all fonts`,
			Args: []model.Arg{{
				Name:        "scale",
				Description: `Font scale`,
			}},
		}, {
			Name:        []string{"--font-size"},
			Description: `Font size used by the date and title`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Font size`,
			}},
		}, {
			Name:        []string{"--file-font-size"},
			Description: `Font size of filenames`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Font size`,
			}},
		}, {
			Name:        []string{"--dir-font-size"},
			Description: `Font size of directory names`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Font size`,
			}},
		}, {
			Name:        []string{"--user-font-size"},
			Description: `Font size of user names`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Font size`,
			}},
		}, {
			Name:        []string{"--font-colour"},
			Description: `Font colour used by the date and title in hex`,
			Args: []model.Arg{{
				Name:        "colour",
				Description: `Font colour`,
			}},
		}, {
			Name:        []string{"--key"},
			Description: `Show file extension key`,
		}, {
			Name:        []string{"--logo"},
			Description: `Logo to display in the foreground`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "image file",
			}},
		}, {
			Name:        []string{"--logo-offset"},
			Description: `Offset position of the logo`,
			Args: []model.Arg{{
				Name:        "offset",
				Description: `Logo offset`,
			}},
		}, {
			Name:        []string{"--date-format"},
			Description: `Specify display date string (strftime format)`,
			Args: []model.Arg{{
				Name:        "format",
				Description: `Date format`,
			}},
		}, {
			Name:        []string{"--log-command"},
			Description: `Show the log command used by gource (git,svn,hg,bzr,cvs2cl)`,
			Args: []model.Arg{{
				Name:        "command",
				Description: `Log command`,
				Suggestions: []model.Suggestion{{Name: []string{`git`}}, {Name: []string{`svn`}}, {Name: []string{`hg`}}, {Name: []string{`bzr`}}, {Name: []string{`cvs2cl`}}},
			}},
		}, {
			Name:        []string{"--log-format"},
			Description: `Specify format of the log being read (git,svn,hg,bzr,cvs2cl,custom). Required when reading from STDIN`,
			Args: []model.Arg{{
				Name:        "format",
				Description: `Log format`,
				Suggestions: []model.Suggestion{{Name: []string{`git`}}, {Name: []string{`svn`}}, {Name: []string{`hg`}}, {Name: []string{`bzr`}}, {Name: []string{`cvs2cl`}}, {Name: []string{`custom`}}},
			}},
		}, {
			Name:        []string{"--git-branch"},
			Description: `Get the git log of a branch other than the current one`,
			Args: []model.Arg{{
				Name: "branch",
			}},
		}, {
			Name:        []string{"--follow-user"},
			Description: `Have the camera automatically follow a particular user`,
			Args: []model.Arg{{
				Name:        "username",
				Description: `Name of user to follow`,
			}},
		}, {
			Name:        []string{"--highlight-dirs"},
			Description: `Highlight the names of all directories`,
		}, {
			Name:        []string{"--highlight-user"},
			Description: `Highlight the names of a particular user`,
			Args: []model.Arg{{
				Name:        "username",
				Description: `Name of user to highlight`,
			}},
		}, {
			Name:        []string{"--highlight-users"},
			Description: `Highlight the names of all users`,
		}, {
			Name:        []string{"--highlight-colour"},
			Description: `Font colour for highlighted users in hex`,
			Args: []model.Arg{{
				Name:        "colour",
				Description: `Font colour (HEX)`,
			}},
		}, {
			Name:        []string{"--selection-colour"},
			Description: `Font colour for selected users and files`,
			Args: []model.Arg{{
				Name:        "colour",
				Description: `Font colour (HEX)`,
			}},
		}, {
			Name:        []string{"--filename-colour"},
			Description: `Font colour for filenames`,
			Args: []model.Arg{{
				Name:        "colour",
				Description: `Font colour (HEX)`,
			}},
		}, {
			Name:        []string{"--dir-colour"},
			Description: `Font colour for directories`,
			Args: []model.Arg{{
				Name:        "color",
				Description: `Dir colour (HEX)`,
			}},
		}, {
			Name:        []string{"--dir-name-depth"},
			Description: `Draw names of directories down to a specific depth in the tree`,
			Args: []model.Arg{{
				Name: "depth",
			}},
		}, {
			Name:        []string{"--dir-name-position"},
			Description: `Position along edge of the directory name (between 0.1 and 1.0, default is 0.5)`,
			Args: []model.Arg{{
				Name:        "position",
				Suggestions: []model.Suggestion{{Name: []string{`0.1`}}, {Name: []string{`0.2`}}, {Name: []string{`0.3`}}, {Name: []string{`0.4`}}, {Name: []string{`0.5`}}, {Name: []string{`0.6`}}, {Name: []string{`0.7`}}, {Name: []string{`0.8`}}, {Name: []string{`0.9`}}, {Name: []string{`1.0`}}},
			}},
		}, {
			Name:        []string{"--filename-time"},
			Description: `Duration to keep filenames on screen (>= 2.0)`,
			Args: []model.Arg{{
				Name: "time",
			}},
		}, {
			Name:        []string{"--file-extensions"},
			Description: `Show filename extensions only`,
		}, {
			Name:        []string{"--file-extension-fallback"},
			Description: `Use filename as extension if the extension is missing or empty`,
		}, {
			Name:        []string{"--file-filter"},
			Description: `Filter out file paths matching the specified regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"--file-show-filter"},
			Description: `Show only file paths matching the specified regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"--user-filter"},
			Description: `Filter usernames matching the specified regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"--user-show-filter"},
			Description: `Show only usernames matching the specified regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"--user-image-dir"},
			Description: `Directory containing .jpg or .png images of users (eg 'Full Name.png') to use as avatars`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "directory",
			}},
		}, {
			Name:        []string{"--default-user-image"},
			Description: `Path of .jpg to use as the default user image`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "User image",
			}},
		}, {
			Name:        []string{"--fixed-user-size"},
			Description: `Forces the size of the user image to remain fixed throughout`,
		}, {
			Name:        []string{"--colour-images"},
			Description: `Colourize user images`,
		}, {
			Name:        []string{"--crop"},
			Description: `Crop view on an axis (vertical,horizontal)`,
			Args: []model.Arg{{
				Name: "axis",
			}},
		}, {
			Name:        []string{"--padding"},
			Description: `Camera view padding`,
			Args: []model.Arg{{
				Name: "padding",
			}},
		}, {
			Name:        []string{"--multi-sampling"},
			Description: `Enable multi-sampling`,
		}, {
			Name:        []string{"--no-vsync"},
			Description: `Disable vsync`,
		}, {
			Name:        []string{"--bloom-multiplier"},
			Description: `Adjust the amount of bloom`,
			Args: []model.Arg{{
				Name: "multiplier",
			}},
		}, {
			Name:        []string{"--bloom-intensity"},
			Description: `Adjust the intensity of the bloom`,
			Args: []model.Arg{{
				Name: "intensity",
			}},
		}, {
			Name:        []string{"--max-files"},
			Description: `Set the maximum number of files or 0 for no limit. Excess files will be discarded`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"--max-file-lag"},
			Description: `Max time files of a commit can take to appear. Use -1 for no limit`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"--max-user-speed"},
			Description: `Max speed users can travel per second`,
			Args: []model.Arg{{
				Name: "units",
			}},
		}, {
			Name:        []string{"--user-friction"},
			Description: `Time users take to come to a halt`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"--user-scale"},
			Description: `Change scale of user avatars`,
			Args: []model.Arg{{
				Name: "scale",
			}},
		}, {
			Name:        []string{"--camera-mode"},
			Description: `Camera mode (overview,track)`,
			Args: []model.Arg{{
				Name:        "mode",
				Suggestions: []model.Suggestion{{Name: []string{`overview`}}, {Name: []string{`track`}}},
			}},
		}, {
			Name:        []string{"--disable-auto-rotate"},
			Description: `Disable automatic camera rotation`,
		}, {
			Name:        []string{"--disable-input"},
			Description: `Disable keyboard and mouse input`,
		}, {
			Name:        []string{"--hide"},
			Description: `Hide one or more display elements from the list below: bloom - bloom effect date - current date dirnames - names of directories files - file icons filenames - names of files mouse - mouse cursor progress - progress bar widget root - root directory of the tree tree - animated tree structure users - user avatars usernames - names of users`,
			Args: []model.Arg{{
				Name:        "element",
				Description: `Element to hide (Separate multiple elements with commas (eg "mouse,progress"))`,
				Suggestions: []model.Suggestion{{Name: []string{`bloom`}}, {Name: []string{`date`}}, {Name: []string{`dirnames`}}, {Name: []string{`files`}}, {Name: []string{`filenames`}}, {Name: []string{`mouse`}}, {Name: []string{`progress`}}, {Name: []string{`root`}}, {Name: []string{`tree`}}, {Name: []string{`users`}}, {Name: []string{`usernames`}}},
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"--hash-seed"},
			Description: `Change the seed of hash function`,
			Args: []model.Arg{{
				Name: "seed",
			}},
		}, {
			Name:        []string{"--caption-file"},
			Description: `Caption file (see Caption Log Format)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"--caption-size"},
			Description: `Caption size`,
			Args: []model.Arg{{
				Name: "size",
			}},
		}, {
			Name:        []string{"--caption-colour"},
			Description: `Caption colour in hex`,
			Args: []model.Arg{{
				Name:        "colour",
				Description: `Colour in hex (eg #ff0000)`,
			}},
		}, {
			Name:        []string{"--caption-duration"},
			Description: `Caption duration`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"--caption-offset"},
			Description: `Caption horizontal offset (0 to centre captions)`,
			Args: []model.Arg{{
				Name: "offset",
			}},
		}, {
			Name:        []string{"-o", "--output-ppm-stream"},
			Description: `Output a PPM image stream to a file ('-' for STDOUT). This will automatically hide the progress bar initially and enable 'stop-at-end' unless other behaviour is specified`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"-r", "--output-framerate"},
			Description: `Framerate of output (25,30,60). Used with --output-ppm-stream`,
			Args: []model.Arg{{
				Name:        "fps",
				Suggestions: []model.Suggestion{{Name: []string{`25`}}, {Name: []string{`30`}}, {Name: []string{`60`}}, {Name: []string{`120`}}},
			}},
		}, {
			Name:        []string{"--output-custom-log"},
			Description: `Output a custom format log file ('-' for STDOUT)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"--load-config"},
			Description: `Load a config file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "config file",
			}},
		}, {
			Name:        []string{"--save-config"},
			Description: `Save a config file with the current options`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "config file",
			}},
		}, {
			Name:        []string{"--path"},
			Description: `Either a supported version control directory, a pre-generated log file (see log commands or the custom log format), a Gource conf file or '-' to read STDIN`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}},
	}
}
