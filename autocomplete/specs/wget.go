// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["wget"] = model.Subcommand{
		Name:        []string{"wget"},
		Description: `A non-interactive network retriever`,
		Args: []model.Arg{{
			Name:        "url",
			Description: `The url(s) to retrieve`,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-V", "--version"},
			Description: `Display the version of Wget and exit`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print this help`,
		}, {
			Name:        []string{"-b", "--background"},
			Description: `Go to background after startup`,
		}, {
			Name:        []string{"-e", "--execute=COMMAND"},
			Description: `Execute a ".wgetrc'-style command`,
		}, {
			Name:        []string{"-o", "--output-file=FILE"},
			Description: `Log messages to FILE`,
		}, {
			Name:        []string{"-a", "--append-output=FILE"},
			Description: `Append messages to FILE`,
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Quiet (no output)`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Be verbose (this is the default)`,
		}, {
			Name:        []string{"-nv", "--no-verbose"},
			Description: `Turn off verboseness, without being quiet`,
		}, {
			Name:        []string{"--report-speed=TYPE"},
			Description: `Output bandwidth as TYPE.  TYPE can be bits`,
		}, {
			Name:        []string{"-i", "--input-file=FILE"},
			Description: `Download URLs found in local or external FILE`,
		}, {
			Name:        []string{"-F", "--force-html"},
			Description: `Treat input file as HTML`,
		}, {
			Name:        []string{"-B", "--base=URL"},
			Description: `Resolves HTML input-file links (-i -F) relative to URL`,
		}, {
			Name:        []string{"--config=FILE"},
			Description: `Specify config file to use`,
		}, {
			Name:        []string{"--no-config"},
			Description: `Do not read any config file`,
		}, {
			Name:        []string{"--rejected-log=FILE"},
			Description: `Log reasons for URL rejection to FILE`,
		}, {
			Name:        []string{"-t", "--tries=NUMBER"},
			Description: `Set number of retries to NUMBER (0 unlimits)`,
		}, {
			Name:        []string{"--retry-connrefused"},
			Description: `Retry even if connection is refused`,
		}, {
			Name:        []string{"--retry-on-http-error"},
			Description: `Comma-separated list of HTTP errors to retry`,
		}, {
			Name:        []string{"-O", "--output-document=FILE"},
			Description: `Write documents to FILE`,
		}, {
			Name:        []string{"-nc", "--no-clobber"},
			Description: `Skip downloads that would download to existing files (overwriting them)`,
		}, {
			Name:        []string{"--no-netrc"},
			Description: `Don't try to obtain credentials from .netrc`,
		}, {
			Name:        []string{"-c", "--continue"},
			Description: `Resume getting a partially-downloaded file`,
		}, {
			Name:        []string{"--start-pos=OFFSET"},
			Description: `Start downloading from zero-based position OFFSET`,
		}, {
			Name:        []string{"--progress=TYPE"},
			Description: `Select progress gauge type`,
		}, {
			Name:        []string{"--show-progress"},
			Description: `Display the progress bar in any verbosity mode`,
		}, {
			Name:        []string{"-N", "--timestamping"},
			Description: `Don't re-retrieve files unless newer than local`,
		}, {
			Name:        []string{"-S", "--server-response"},
			Description: `Print server response`,
		}, {
			Name:        []string{"--spider"},
			Description: `Don't download anything`,
		}, {
			Name:        []string{"-T", "--timeout=SECONDS"},
			Description: `Set all timeout values to SECONDS`,
		}, {
			Name:        []string{"--dns-timeout=SECS"},
			Description: `Set the DNS lookup timeout to SECS`,
		}, {
			Name:        []string{"--connect-timeout=SECS"},
			Description: `Set the connect timeout to SECS`,
		}, {
			Name:        []string{"--read-timeout=SECS"},
			Description: `Set the read timeout to SECS`,
		}, {
			Name:        []string{"-w", "--wait=SECONDS"},
			Description: `Wait SECONDS between retrievals`,
		}, {
			Name:        []string{"--waitretry=SECONDS"},
			Description: `Wait 1..SECONDS between retries of a retrieval`,
		}, {
			Name:        []string{"--random-wait"},
			Description: `Wait from 0.5*WAIT...1.5*WAIT secs between retrievals`,
		}, {
			Name:        []string{"--no-proxy"},
			Description: `Explicitly turn off proxy`,
		}, {
			Name:        []string{"-Q", "--quota=NUMBER"},
			Description: `Set retrieval quota to NUMBER`,
		}, {
			Name:        []string{"--bind-address=ADDRESS"},
			Description: `Bind to ADDRESS (hostname or IP) on local host`,
		}, {
			Name:        []string{"--limit-rate=RATE"},
			Description: `Limit download rate to RATE`,
		}, {
			Name:        []string{"--no-dns-cache"},
			Description: `Disable caching DNS lookups`,
		}, {
			Name:        []string{"--restrict-file-names=OS"},
			Description: `Restrict chars in file names to ones OS allows`,
		}, {
			Name:        []string{"--ignore-case"},
			Description: `Ignore case when matching files/directories`,
		}, {
			Name:        []string{"-4", "--inet4-only"},
			Description: `Connect only to IPv4 addresses`,
		}, {
			Name:        []string{"-6", "--inet6-only"},
			Description: `Connect only to IPv6 addresses`,
		}, {
			Name:        []string{"--user=USER"},
			Description: `Set both ftp and http user to USER`,
		}, {
			Name:        []string{"--password=PASS"},
			Description: `Set both ftp and http password to PASS`,
		}, {
			Name:        []string{"--ask-password"},
			Description: `Prompt for passwords`,
		}, {
			Name:        []string{"--no-iri"},
			Description: `Turn off IRI support`,
		}, {
			Name:        []string{"--local-encoding=ENC"},
			Description: `Use ENC as the local encoding for IRIs`,
		}, {
			Name:        []string{"--remote-encoding=ENC"},
			Description: `Use ENC as the default remote encoding`,
		}, {
			Name:        []string{"--unlink"},
			Description: `Remove file before clobber`,
		}, {
			Name:        []string{"--xattr"},
			Description: `Turn on storage of metadata in extended file attributes`,
		}, {
			Name:        []string{"-nd", "--no-directories"},
			Description: `Don't create directories`,
		}, {
			Name:        []string{"-x", "--force-directories"},
			Description: `Force creation of directories`,
		}, {
			Name:        []string{"-nH", "--no-host-directories"},
			Description: `Don't create host directories`,
		}, {
			Name:        []string{"--protocol-directories"},
			Description: `Use protocol name in directories`,
		}, {
			Name:        []string{"-P", "--directory-prefix=PREFIX"},
			Description: `Save files to PREFIX/`,
		}, {
			Name:        []string{"--cut-dirs=NUMBER"},
			Description: `Ignore NUMBER remote directory components`,
		}, {
			Name:        []string{"--http-user=USER"},
			Description: `Set http user to USER`,
		}, {
			Name:        []string{"--http-password=PASS"},
			Description: `Set http password to PASS`,
		}, {
			Name:        []string{"--no-cache"},
			Description: `Disallow server-cached data`,
		}, {
			Name:        []string{"-E", "--adjust-extension"},
			Description: `Save HTML/CSS documents with proper extensions`,
		}, {
			Name:        []string{"--ignore-length"},
			Description: `Ignore 'Content-Length' header field`,
		}, {
			Name:        []string{"--header=STRING"},
			Description: `Insert STRING among the headers`,
		}, {
			Name:        []string{"--compression=TYPE"},
			Description: `Choose compression, one of auto, gzip and none. (default: none)`,
		}, {
			Name:        []string{"--max-redirect"},
			Description: `Maximum redirections allowed per page`,
		}, {
			Name:        []string{"--proxy-user=USER"},
			Description: `Set USER as proxy username`,
		}, {
			Name:        []string{"--proxy-password=PASS"},
			Description: `Set PASS as proxy password`,
		}, {
			Name:        []string{"--referer=URL"},
			Description: `Include 'Referer: URL' header in HTTP request`,
		}, {
			Name:        []string{"--save-headers"},
			Description: `Save the HTTP headers to file`,
		}, {
			Name:        []string{"-U", "--user-agent=AGENT"},
			Description: `Identify as AGENT instead of Wget/VERSION`,
		}, {
			Name:        []string{"--no-http-keep-alive"},
			Description: `Disable HTTP keep-alive (persistent connections)`,
		}, {
			Name:        []string{"--no-cookies"},
			Description: `Don't use cookies`,
		}, {
			Name:        []string{"--load-cookies=FILE"},
			Description: `Load cookies from FILE before session`,
		}, {
			Name:        []string{"--save-cookies=FILE"},
			Description: `Save cookies to FILE after session`,
		}, {
			Name:        []string{"--keep-session-cookies"},
			Description: `Load and save session (non-permanent) cookies`,
		}, {
			Name:        []string{"--post-data=STRING"},
			Description: `Use the POST method; send STRING as the data`,
		}, {
			Name:        []string{"--post-file=FILE"},
			Description: `Use the POST method; send contents of FILE`,
		}, {
			Name:        []string{"--method=HTTPMethod"},
			Description: `Use method "HTTPMethod" in the request`,
		}, {
			Name:        []string{"--body-data=STRING"},
			Description: `Send STRING as data. --method MUST be set`,
		}, {
			Name:        []string{"--body-file=FILE"},
			Description: `Send contents of FILE. --method MUST be set`,
		}, {
			Name:        []string{"--content-on-error"},
			Description: `Output the received content on server errors`,
		}, {
			Name:        []string{"--secure-protocol=PR"},
			Description: `Choose secure protocol, one of auto, SSLv2,`,
		}, {
			Name:        []string{"--https-only"},
			Description: `Only follow secure HTTPS links`,
		}, {
			Name:        []string{"--no-check-certificate"},
			Description: `Don't validate the server's certificate`,
		}, {
			Name:        []string{"--certificate=FILE"},
			Description: `Client certificate file`,
		}, {
			Name:        []string{"--certificate-type=TYPE"},
			Description: `Client certificate type, PEM or DER`,
		}, {
			Name:        []string{"--private-key=FILE"},
			Description: `Private key file`,
		}, {
			Name:        []string{"--private-key-type=TYPE"},
			Description: `Private key type, PEM or DER`,
		}, {
			Name:        []string{"--ca-certificate=FILE"},
			Description: `File with the bundle of CAs`,
		}, {
			Name:        []string{"--ca-directory=DIR"},
			Description: `Directory where hash list of CAs is stored`,
		}, {
			Name:        []string{"--crl-file=FILE"},
			Description: `File with bundle of CRLs`,
		}, {
			Name:        []string{"--ciphers=STR"},
			Description: `Set the priority string (GnuTLS) or cipher list string (OpenSSL) directly`,
		}, {
			Name:        []string{"-r", "--recursive"},
			Description: `Specify recursive download`,
		}, {
			Name:        []string{"-l", "--level=NUMBER"},
			Description: `Maximum recursion depth (inf or 0 for infinite)`,
		}, {
			Name:        []string{"--delete-after"},
			Description: `Delete files locally after downloading them`,
		}, {
			Name:        []string{"-k", "--convert-links"},
			Description: `Make links in downloaded HTML or CSS point to local files`,
		}, {
			Name:        []string{"-K", "--backup-converted"},
			Description: `Before converting file X, back up as X.orig`,
		}, {
			Name:        []string{"-m", "--mirror"},
			Description: `Shortcut for -N -r -l inf --no-remove-listing`,
		}, {
			Name:        []string{"-p", "--page-requisites"},
			Description: `Get all images, etc. needed to display HTML page`,
		}, {
			Name:        []string{"-A", "--accept=LIST"},
			Description: `Comma-separated list of accepted extensions`,
		}, {
			Name:        []string{"-R", "--reject=LIST"},
			Description: `Comma-separated list of rejected extensions`,
		}, {
			Name:        []string{"--accept-regex=REGEX"},
			Description: `Regex matching accepted URLs`,
		}, {
			Name:        []string{"--reject-regex=REGEX"},
			Description: `Regex matching rejected URLs`,
		}, {
			Name:        []string{"--regex-type=TYPE"},
			Description: `Regex type (posix)`,
		}, {
			Name:        []string{"-D", "--domains=LIST"},
			Description: `Comma-separated list of accepted domains`,
		}, {
			Name:        []string{"--exclude-domains=LIST"},
			Description: `Comma-separated list of rejected domains`,
		}, {
			Name:        []string{"--follow-ftp"},
			Description: `Follow FTP links from HTML documents`,
		}, {
			Name:        []string{"--follow-tags=LIST"},
			Description: `Comma-separated list of followed HTML tags`,
		}, {
			Name:        []string{"--ignore-tags=LIST"},
			Description: `Comma-separated list of ignored HTML tags`,
		}, {
			Name:        []string{"-H", "--span-hosts"},
			Description: `Go to foreign hosts when recursive`,
		}, {
			Name:        []string{"-L", "--relative"},
			Description: `Follow relative links only`,
		}, {
			Name:        []string{"-I", "--include-directories=LIST"},
			Description: `List of allowed directories`,
		}, {
			Name:        []string{"-X", "--exclude-directories=LIST"},
			Description: `List of excluded directories`,
		}, {
			Name:        []string{"-np", "--no-parent"},
			Description: `Don't ascend to the parent directory`,
		}},
	}
}
