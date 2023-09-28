// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["wscat"] = model.Subcommand{
		Name:        []string{"wscat"},
		Description: `Communicate over websocket`,
		Options: []model.Option{{
			Name:        []string{"-c", "--connect"},
			Description: `Connect to a WebSocket server`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateHistory},
				Name:      "url",
			}},
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Output the version number`,
		}, {
			Name:        []string{"--auth"},
			Description: `Add basic HTTP authentication header (--connect only)`,
			Args: []model.Arg{{
				Name: "username:password",
			}},
		}, {
			Name:        []string{"--ca"},
			Description: `Specify a Certificate Authority (--connect only)`,
			Args: []model.Arg{{
				Name: "ca",
			}},
		}, {
			Name:        []string{"--cert"},
			Description: `Specify a Client SSL Certificate (--connect only)`,
			Args: []model.Arg{{
				Name: "cert",
			}},
		}, {
			Name:        []string{"--host"},
			Description: `Optional host`,
			Args: []model.Arg{{
				Name: "host",
			}},
		}, {
			Name:        []string{"--key"},
			Description: `Specify a Client SSL Certificate's key (--connect only)`,
			Args: []model.Arg{{
				Name: "key",
			}},
		}, {
			Name:        []string{"--max-redirects"},
			Description: `Maximum number of redirects allowed (--connect only) (default: 10)`,
			Args: []model.Arg{{
				Name: "num",
			}},
		}, {
			Name:        []string{"--no-color"},
			Description: `Run without color`,
		}, {
			Name:        []string{"--passphrase"},
			Description: `Specify a Client SSL Certificate Key's passphrase (--connect only). If you don't provide a value, it will be prompted for`,
			Args: []model.Arg{{
				Name: "passphrase",
			}},
		}, {
			Name:        []string{"--proxy"},
			Description: `Connect via a proxy. Proxy must support CONNECT method`,
			Args: []model.Arg{{
				Name: "[protocol://]host[:port]",
			}},
		}, {
			Name:        []string{"--slash"},
			Description: `Enable slash commands for control frames (/ping, /pong, /close [code [, reason]])`,
		}, {
			Name:        []string{"-H", "--header"},
			Description: `Set an HTTP header. Repeat to set multiple (--connect only) (default: [])`,
			Args: []model.Arg{{
				Name: "header:value",
			}},
		}, {
			Name:        []string{"-L", "--location"},
			Description: `Follow redirects (--connect only)`,
		}, {
			Name:        []string{"-l", "--listen"},
			Description: `Listen on port`,
			Args: []model.Arg{{
				Name: "port",
			}},
		}, {
			Name:        []string{"-n", "--no-check"},
			Description: `Do not check for unauthorized certificates`,
		}, {
			Name:        []string{"-o", "--origin"},
			Description: `Optional origin`,
			Args: []model.Arg{{
				Name: "origin",
			}},
		}, {
			Name:        []string{"-p", "--protocol"},
			Description: `Optional protocol version`,
			Args: []model.Arg{{
				Name: "protocol",
			}},
		}, {
			Name:        []string{"-P", "--show-ping-pong"},
			Description: `Print a notification when a ping or pong is received`,
		}, {
			Name:        []string{"-s", "--subprotocol"},
			Description: `Optional subprotocol (default: [])`,
			Args: []model.Arg{{
				Name:        "protocol",
				Suggestions: []model.Suggestion{{Name: []string{`MBWS.huawei.com`}}, {Name: []string{`MBLWS.huawei.com`}}, {Name: []string{`soap`}}, {Name: []string{`wamp`}}, {Name: []string{`v10.stomp`}}, {Name: []string{`v11.stomp`}}, {Name: []string{`v12.stomp`}}, {Name: []string{`ocpp1.2`}}, {Name: []string{`ocpp1.5`}}, {Name: []string{`ocpp1.6`}}, {Name: []string{`ocpp2.0`}}, {Name: []string{`ocpp2.0.1`}}, {Name: []string{`rfb`}}, {Name: []string{`sip`}}, {Name: []string{`notificationchannel-netapi-rest.openmobilealliance.org`}}, {Name: []string{`wpcp`}}, {Name: []string{`amqp	`}}, {Name: []string{`mqtt`}}, {Name: []string{`jsflow`}}, {Name: []string{`rwpcp`}}, {Name: []string{`xmpp`}}, {Name: []string{`ship`}}, {Name: []string{`mielecloudconnect`}}, {Name: []string{`v10.pcp.sap.com`}}, {Name: []string{`msrp`}}, {Name: []string{`v1.saltyrtc.org`}}, {Name: []string{`TLCP-2.0.0.lightstreamer.com`}}, {Name: []string{`bfcp`}}, {Name: []string{`sldp.softvelum.com`}}, {Name: []string{`opcua+uacp`}}, {Name: []string{`opcua+uajson`}}, {Name: []string{`v1.swindon-lattice+json`}}, {Name: []string{`v1.usp`}}, {Name: []string{`mles-websocket`}}, {Name: []string{`coap`}}, {Name: []string{`TLCP-2.1.0.lightstreamer.com`}}, {Name: []string{`sqlnet.oracle.com`}}, {Name: []string{`oneM2M.R2.0.json`}}, {Name: []string{`oneM2M.R2.0.xml`}}, {Name: []string{`oneM2M.R2.0.cbor`}}, {Name: []string{`transit`}}, {Name: []string{`2016.serverpush.dash.mpeg.org`}}, {Name: []string{`2018.mmt.mpeg.org`}}, {Name: []string{`clue`}}, {Name: []string{`webrtc.softvelum.com`}}, {Name: []string{`cobra.v2.json`}}, {Name: []string{`drp`}}, {Name: []string{`hub.bsc.bacnet.org`}}, {Name: []string{`dc.bsc.bacnet.org`}}, {Name: []string{`jmap`}}, {Name: []string{`t140`}}, {Name: []string{`done`}}, {Name: []string{`TLCP-2.2.0.lightstreamer.com`}}, {Name: []string{`collection-update`}}, {Name: []string{`TLCP-2.3.0.lightstreamer.com`}}, {Name: []string{`text.ircv3.net`}}, {Name: []string{`binary.ircv3.net`}}, {Name: []string{`v3.penguin-stats.live+proto`}}},
			}},
		}, {
			Name:        []string{"-w", "--wait"},
			Description: `Wait given seconds after executing command`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-x", "--execute"},
			Description: `Execute command after connecting`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Display help for command`,
		}},
	}
}
