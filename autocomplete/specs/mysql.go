// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["mysql"] = model.Subcommand{
		Name:        []string{"mysql"},
		Description: `Mysql is a terminal-based front-end to MySQL`,
		Args:        []model.Arg{{}},
		Options: []model.Option{{
			Name:        []string{"--auto-rehash"},
			Description: `Enable automatic rehashing`,
		}, {
			Name:        []string{"--auto-vertical-output"},
			Description: `Enable automatic vertical result set display`,
		}, {
			Name:        []string{"--batch", "-B"},
			Description: `Do not use history file`,
		}, {
			Name:        []string{"--binary-as-hex"},
			Description: `Display binary values in hexadecimal notation`,
		}, {
			Name:        []string{"--binary-mode"},
			Description: `Disable \\r\\n - to - \\n translation and treatment of \\0 as end-of-query`,
		}, {
			Name:        []string{"--bind-address"},
			Description: `Use specified network interface to connect to MySQL Server`,
			Args: []model.Arg{{
				Name: "ip_address",
			}},
		}, {
			Name:        []string{"--character-sets-dir"},
			Description: `Directory where character sets are installed`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir_name",
				Description: `Directory where character sets are installed`,
			}},
		}, {
			Name:        []string{"--column-names"},
			Description: `Write column names in results`,
		}, {
			Name:        []string{"--column-type-info"},
			Description: `Display result set metadata`,
		}, {
			Name:        []string{"--comments", "-c"},
			Description: `Whether to retain or strip comments in statements sent to the server`,
		}, {
			Name:        []string{"--compress", "-C"},
			Description: `Compress all information sent between client and server`,
		}, {
			Name:        []string{"--compression-algorithms"},
			Description: `Permitted compression algorithms for connections to server`,
			Args: []model.Arg{{
				Name: "value",
				Suggestions: []model.Suggestion{{
					Name: []string{`zlib`},
				}, {
					Name: []string{`zstd`},
				}, {
					Name: []string{`uncompressed`},
				}},
			}},
		}, {
			Name:        []string{"--connect-expired-password"},
			Description: `Indicate to server that client can handle expired-password sandbox mode`,
		}, {
			Name:        []string{"--connect-timeout"},
			Description: `Number of seconds before connection timeout`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"-D", "--database"},
			Description: `The database to use`,
			Args: []model.Arg{{
				Name: "db_name",
			}},
		}, {
			Name:        []string{"--debug"},
			Description: `Write debugging log; supported only if MySQL was built with debugging support`,
		}, {
			Name:        []string{"--debug-check"},
			Description: `Print debugging information when program exits`,
		}, {
			Name:        []string{"-T", "--debug-info"},
			Description: `Print debugging information, memory, and CPU statistics when program exits`,
		}, {
			Name:        []string{"--default-auth"},
			Description: `Authentication plugin to use`,
			Args: []model.Arg{{
				Name: "plugin",
			}},
		}, {
			Name:        []string{"--default-character-set"},
			Description: `Specify default character set`,
			Args: []model.Arg{{
				Name: "charset_name",
			}},
		}, {
			Name:        []string{"--defaults-extra-file"},
			Description: `Read named option file in addition to usual option files`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file_name",
			}},
		}, {
			Name:        []string{"--defaults-file"},
			Description: `Read only named option file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file_name",
			}},
		}, {
			Name:        []string{"--defaults-group-suffix"},
			Description: `Option group suffix value`,
			Args: []model.Arg{{
				Name: "str",
			}},
		}, {
			Name:        []string{"--delimiter"},
			Description: `Set the statement delimiter`,
			Args: []model.Arg{{
				Name: "str",
			}},
		}, {
			Name:        []string{"--disable-named-commands"},
			Description: `Disable named commands. Use the \* form only, or use named commands only at the beginning of a line ending with a semicolon (;). mysql starts with this option enabled by default. However, even with this option, long-format commands still work from the first line`,
		}, {
			Name:        []string{"--dns-srv-name"},
			Description: `Use DNS SRV lookup for host information`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--enable-cleartext-plugin"},
			Description: `Enable cleartext authentication plugin`,
		}, {
			Name:        []string{"-e", "--execute"},
			Description: `Execute the statement and quit`,
			Args: []model.Arg{{
				Name: "statement",
			}},
		}, {
			Name:        []string{"-f", "--force"},
			Description: `Continue even if an SQL error occurs`,
		}, {
			Name:        []string{"--get-server-public-key"},
			Description: `Request RSA public key from server`,
		}, {
			Name:        []string{"--help", "-?"},
			Description: `Display help message and exit`,
		}, {
			Name:        []string{"--histignore"},
			Description: `Patterns specifying which statements to ignore for logging`,
		}, {
			Name:        []string{"-h", "--host"},
			Description: `Host on which MySQL server is located`,
			Args: []model.Arg{{
				Name: "host_name",
			}},
		}, {
			Name:        []string{"-H", "--html"},
			Description: `Produce HTML output`,
		}, {
			Name:        []string{"-i", "--ignore-spaces"},
			Description: `Ignore spaces after function names`,
		}, {
			Name:        []string{"--init-command"},
			Description: `SQL statement to execute after connecting`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"--line-numbers"},
			Description: `Write line numbers for errors`,
		}, {
			Name:        []string{"--load-data-local-dir"},
			Description: `Directory for files named in LOAD DATA LOCAL statements`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "dir_name",
			}},
		}, {
			Name:        []string{"--local-infile"},
			Description: `Enable or disable for LOCAL capability for LOAD DATA`,
			Args: []model.Arg{{
				Name: "is-local-inflie",
				Suggestions: []model.Suggestion{{
					Name:        []string{`0`},
					Description: `Disable`,
				}, {
					Name:        []string{`1`},
					Description: `Enable`,
				}},
			}},
		}, {
			Name:        []string{"--login-path"},
			Description: `Read login path options from .mylogin.cnf`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--max-allowed-packet"},
			Description: `Maximum packet length to send to or receive from server`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--max-join-size"},
			Description: `The automatic limit for rows in a join when using --safe-updates`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"-G", "--named-commands"},
			Description: `Enable named mysql commands`,
		}, {
			Name:        []string{"--net-buffer-length"},
			Description: `Buffer size for TCP/IP and socket communication`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--network-namespace"},
			Description: `Specify network namespace`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"-A", "--no-auto-rehash"},
			Description: `Disable automatic rehashing`,
		}, {
			Name:        []string{"--no-beep", "-b"},
			Description: `Do not beep when errors occur`,
		}, {
			Name:        []string{"--no-defaults"},
			Description: `Read no option files`,
		}, {
			Name:        []string{"-o", "--one-database"},
			Description: `Ignore statements except those for the default database named on the command line`,
		}, {
			Name:        []string{"--pager"},
			Description: `Use the given command for paging query output`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"-p", "--password"},
			Description: `Password to use when connecting to server`,
			Args: []model.Arg{{
				Name: "password",
			}},
		}, {
			Name:        []string{"--pipe", "-W"},
			Description: `Connect to server using named pipe (Windows only)`,
		}, {
			Name:        []string{"--plugin-dir"},
			Description: `Directory where plugins are installed`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "dir_name",
			}},
		}, {
			Name:        []string{"-P", "--port"},
			Description: `TCP/IP port number for connection`,
			Args: []model.Arg{{
				Name: "port_num",
			}},
		}, {
			Name:        []string{"--print-defaults"},
			Description: `Print default options`,
		}, {
			Name:        []string{"--prompt"},
			Description: `Set the prompt to the specified format`,
			Args: []model.Arg{{
				Name: "format_str",
			}},
		}, {
			Name:        []string{"--protocol"},
			Description: `Transport protocol to use`,
			Args: []model.Arg{{
				Name: "format_str",
				Suggestions: []model.Suggestion{{
					Name: []string{`TCP`},
				}, {
					Name: []string{`SOCKET`},
				}, {
					Name: []string{`PIPE`},
				}, {
					Name: []string{`MEMORY`},
				}},
			}},
		}, {
			Name:        []string{"-q", "--quick"},
			Description: `Do not cache each query result`,
		}, {
			Name:        []string{"-r", "--raw"},
			Description: `Write column values without escape conversion`,
		}, {
			Name:        []string{"--reconnect"},
			Description: `If the connection to the server is lost, automatically try to reconnect`,
		}, {
			Name:        []string{"-U", "--safe-updates", "--i-am-a-dummy"},
			Description: `Allow only UPDATE and DELETE statements that specify key values`,
		}, {
			Name:        []string{"--select-limit"},
			Description: `The automatic limit for SELECT statements when using --safe-updates`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--server-public-key-path"},
			Description: `Path name to file containing RSA public key`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"--shared-memory-base-name"},
			Description: `Shared-memory name for shared-memory connections (Windows only)`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--show-warnings"},
			Description: `Show warnings after each statement if there are any`,
		}, {
			Name:        []string{"--sigint-ignore"},
			Description: `Ignore SIGINT signals (typically the result of typing Control+C)`,
		}, {
			Name:        []string{"-s", "--silent"},
			Description: `Silent mode.  Produce less output. This option can be given multiple times to produce less and less output`,
		}, {
			Name:        []string{"--skip-auto-rehash"},
			Description: `Disable automatic rehashing`,
		}, {
			Name:        []string{"-N", "--skip-column-names"},
			Description: `Do not write column names in results`,
		}, {
			Name:        []string{"-L", "--skip-line-numbers"},
			Description: `Skip line numbers for errors`,
		}, {
			Name:        []string{"--skip-named-commands"},
			Description: `Disable named mysql commands`,
		}, {
			Name:        []string{"--skip-pager"},
			Description: `Disable paging`,
		}, {
			Name:        []string{"--skip-reconnect"},
			Description: `Disable reconnecting`,
		}, {
			Name:        []string{"-S", "--socket"},
			Description: `Unix socket file or Windows named pipe to use"File that contains list of trusted SSL Certificate Authorities"`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"--ssl-ca"},
			Description: `File that contains list of trusted SSL Certificate Authorities`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"--ssl-capath"},
			Description: `Directory that contains trusted SSL Certificate Authority certificate files`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "dirname",
			}},
		}, {
			Name:        []string{"--ssl-cert"},
			Description: `File that contains X.509 certificate`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"--ssl-cipher"},
			Description: `Permissible ciphers for connection encryption`,
		}, {
			Name:        []string{"--ssl-crl"},
			Description: `File that contains certificate revocation lists`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"--ssl-crlpath"},
			Description: `Directory that contains certificate revocation-list files`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "dirname",
			}},
		}, {
			Name:        []string{"--ssl-fips-mode"},
			Description: `Whether to enable FIPS mode on client side`,
			Args: []model.Arg{{
				Name: "mode",
				Suggestions: []model.Suggestion{{
					Name: []string{`OFF`},
				}, {
					Name: []string{`ON`},
				}, {
					Name: []string{`STRICT`},
				}},
			}},
		}, {
			Name:        []string{"--ssl-key"},
			Description: `File that contains X.509 key`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"--ssl-mode"},
			Description: `Desired security state of connection to server`,
		}, {
			Name:        []string{"-j", "--syslog"},
			Description: `Log interactive statements to syslog`,
		}, {
			Name:        []string{"-t", "--table"},
			Description: `Display output in tabular format`,
		}, {
			Name:        []string{"--tee"},
			Description: `Append a copy of output to named file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file_name",
			}},
		}, {
			Name:        []string{"--tls-ciphersuites"},
			Description: `Permissible TLSv1.3 ciphersuites for encrypted connections`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "ciphersuite_list",
			}},
		}, {
			Name:        []string{"--tls-version"},
			Description: `Permissible TLS protocols for encrypted connections`,
			Args: []model.Arg{{
				Name: "protocol_list",
			}},
		}, {
			Name:        []string{"-n", "--unbuffered"},
			Description: `Flush the buffer after each query`,
		}, {
			Name:        []string{"-u", "--user"},
			Description: `MySQL user name to use when connecting to server`,
			Args: []model.Arg{{
				Name: "user_name",
			}},
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Verbose mode`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Display version information and exit`,
		}, {
			Name:        []string{"-E", "--vertical"},
			Description: `Print query output rows vertically (one line per column value)`,
		}, {
			Name:        []string{"-w", "--wait"},
			Description: `If the connection cannot be established, wait and retry instead of aborting`,
		}, {
			Name:        []string{"-X", "--xml"},
			Description: `Produce XML output`,
		}, {
			Name:        []string{"--zstd-compression-level"},
			Description: `Compression level for connections to server that use zstd compression`,
			Args: []model.Arg{{
				Name:        "level",
				Description: `Compression level between 1-22. The default zstd compression level is 3. The compression level setting has no effect on connections that do not use zstd compression`,
			}},
		}},
	}
}
