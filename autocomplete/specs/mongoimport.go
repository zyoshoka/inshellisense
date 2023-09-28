// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["mongoimport"] = model.Subcommand{
		Name:        []string{"mongoimport"},
		Description: `Import data from a JSON, CSV, or TSV file into a MongoDB instance`,
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Returns information on the options and use of mongoimport`,
		}, {
			Name:        []string{"--verbose", "-v"},
			Description: `Increases the amount of internal reporting returned on standard output or in log files`,
		}, {
			Name:        []string{"--quiet"},
			Description: `Runs mongoimport in a quiet mode that attempts to limit the amount of output`,
		}, {
			Name:        []string{"--version"},
			Description: `Returns the mongoimport release number`,
		}, {
			Name:        []string{"--config"},
			Description: `Specifies the full path to a YAML configuration file containing sensitive values for the following options to mongoimport`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--uri"},
			Description: `Specifies the resolvable URI connection string of the MongoDB deployment`,
			Args: []model.Arg{{
				Name:        "connectionString",
				Description: `MongoDB connection string`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`mongodb://localhost:27017`},
					Description: `Default Connection String`,
				}, {
					Name:        []string{`mongodb+srv://username:password@cluster0.example.mongodb.net/database`},
					Description: `Atlas Connection String Example`,
				}},
			}},
		}, {
			Name:        []string{"--host", "-h"},
			Description: `Specifies the resolvable hostname of the MongoDB deployment`,
			Args: []model.Arg{{
				Name: "hostname:port",
				Suggestions: []model.Suggestion{{
					Name:        []string{`localhost:27017`},
					Description: `Default host`,
				}},
			}},
		}, {
			Name:        []string{"--port"},
			Description: `Specifies the TCP port on which the MongoDB instance listens for client connections`,
			Args: []model.Arg{{
				Name: "port",
				Suggestions: []model.Suggestion{{
					Name:        []string{`27017`},
					Description: `Default port`,
				}},
			}},
		}, {
			Name:        []string{"--ssl"},
			Description: `Enables connection to a mongod or mongos that has TLS/SSL support enabled`,
		}, {
			Name:        []string{"--sslCAFile"},
			Description: `Specifies the .pem file that contains the root certificate chain from the Certificate Authority`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--sslPEMKeyFile"},
			Description: `Specifies the .pem file that contains both the TLS/SSL certificate and key`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--sslPEMKeyPassword"},
			Description: `Specifies the password to de-crypt the certificate-key file (i.e. --sslPEMKeyFile)`,
			Args: []model.Arg{{
				Name:        "value",
				Description: `Password`,
			}},
		}, {
			Name:        []string{"--sslCRLFile"},
			Description: `Specifies the .pem file that contains the Certificate Revocation List`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--sslAllowInvalidCertificates"},
			Description: `Bypasses the validation checks for server certificates and allows the use of invalid certificates`,
		}, {
			Name:        []string{"--sslAllowInvalidHostnames"},
			Description: `Disables the validation of the hostnames in TLS/SSL certificates`,
		}, {
			Name:        []string{"--username", "-u"},
			Description: `Specifies a username with which to authenticate to a MongoDB database that uses authentication`,
			Args: []model.Arg{{
				Name: "username",
			}},
		}, {
			Name:        []string{"--password", "-p"},
			Description: `Specifies a password with which to authenticate to a MongoDB database that uses authentication`,
			Args: []model.Arg{{
				Name: "password",
			}},
		}, {
			Name:        []string{"--awsSessionToken"},
			Description: `Specifies the session token for MONGODB-AWS authentication mechanism`,
			Args: []model.Arg{{
				Name: "AWS Session Token",
			}},
		}, {
			Name:        []string{"--authenticationDatabase"},
			Description: `Specifies the authentication database where the specified --username has been created`,
			Args: []model.Arg{{
				Name:        "dbname",
				Description: `Database name`,
			}},
		}, {
			Name:        []string{"--authenticationMechanism"},
			Description: `Specifies the authentication mechanism the mongoimport instance uses to authenticate to the mongod or mongos`,
			Args: []model.Arg{{
				Name: "name",
				Suggestions: []model.Suggestion{{
					Name:        []string{`SCRAM-SHA-1`},
					Description: `Default`,
				}, {
					Name: []string{`SCRAM-SHA-256`},
				}, {
					Name: []string{`MONGODB-X509`},
				}, {
					Name: []string{`MONGODB-AWS`},
				}, {
					Name: []string{`GSSAPI`},
				}, {
					Name: []string{`PLAIN`},
				}},
			}},
		}, {
			Name:        []string{"--gssapiServiceName"},
			Description: `Specify the name of the service using GSSAPI/Kerberos. This option is available only in MongoDB Enterprise`,
			Args: []model.Arg{{
				Name: "serviceName",
			}},
		}, {
			Name:        []string{"--gssapiHostName"},
			Description: `Specify the hostname of the service using GSSAPI/Kerberos. This option is available only in MongoDB Enterprise`,
			Args: []model.Arg{{
				Name: "hostname",
			}},
		}, {
			Name:        []string{"--db", "-d"},
			Description: `Specifies the name of the database on which to run the mongoimport`,
			Args: []model.Arg{{
				Name: "database",
			}},
		}, {
			Name:        []string{"--collection", "-c"},
			Description: `Specifies the name of the collection to import`,
			Args: []model.Arg{{
				Name: "collection",
			}},
		}, {
			Name:        []string{"--fields", "-f"},
			Description: `Specify a comma separated list of field names when importing CSV or TSV files that do not have field names in the first (i.e. header) line of the file`,
			Args: []model.Arg{{
				Name:        "field1,field2,...",
				Description: `Comma separated list of fields`,
			}},
			ExclusiveOn: []string{"--fieldFile"},
		}, {
			Name:        []string{"--fieldFile"},
			Description: `Specify a file containing a comma separated list of field names when importing CSV or TSV files that do not have field names in the first (i.e. header) line of the file`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
			ExclusiveOn: []string{"--fields", "-f"},
		}, {
			Name:        []string{"--ignoreBlanks"},
			Description: `Ignores empty fields in CSV and TSV exports`,
		}, {
			Name:        []string{"--type"},
			Description: `Specifies the file type to import`,
			Args: []model.Arg{{
				Name: "type",
				Suggestions: []model.Suggestion{{
					Name:        []string{`json`},
					Description: `JSON`,
				}, {
					Name:        []string{`csv`},
					Description: `Comma-separated values`,
				}, {
					Name:        []string{`tsv`},
					Description: `Tab-separated values`,
				}},
			}},
		}, {
			Name:        []string{"--file"},
			Description: `Specifies the location and name of a file containing the data to import`,
			Args: []model.Arg{{
				Name:      "filename",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--drop"},
			Description: `Modifies the import process so that the target instance drops the collection before importing the data from the input`,
		}, {
			Name:        []string{"--headerline"},
			Description: `Treats the first line of the input file as a header line`,
		}, {
			Name:        []string{"--useArrayIndexFields"},
			Description: `Interpret natural numbers in fields as array indexes when importing CSV or TSV files`,
		}, {
			Name:        []string{"--mode"},
			Description: `Specifies how the import process should handle existing documents in the database that match documents in the import file`,
			Args: []model.Arg{{
				Name: "mode",
				Suggestions: []model.Suggestion{{
					Name:        []string{`insert`},
					Description: `Insert the documents in the import file`,
				}, {
					Name:        []string{`upsert`},
					Description: `Replace existing documents in the database with matching documents from the import file`,
				}, {
					Name:        []string{`merge`},
					Description: `Merge existing documents that match a document in the import file with the new document`,
				}, {
					Name:        []string{`delete`},
					Description: `Delete existing documents in the database that match a document in the import file`,
				}},
			}},
		}, {
			Name:        []string{"--upsertFields"},
			Description: `Specifies a list of fields for the query portion of the import process`,
			Args: []model.Arg{{
				Name:        "field1,field2,...",
				Description: `Comma separated list of fields`,
			}},
		}, {
			Name:        []string{"--stopOnError"},
			Description: `Forces mongoimport to halt the insert operation at the first error rather than continuing the operation despite errors`,
		}, {
			Name:        []string{"jsonArray"},
			Description: `Accepts the import of data expressed with multiple MongoDB documents within a single JSON array`,
		}, {
			Name:        []string{"--legacy"},
			Description: `Indicates that the import data is in Extended JSON v1 format`,
		}, {
			Name:        []string{"--maintainInsertionOrder"},
			Description: `Maintains the order of documents in the import file when inserting documents into the database`,
		}, {
			Name:        []string{"--numInsertionWorkers"},
			Description: `Specifies the number of insertion workers to run concurrently`,
			Args: []model.Arg{{
				Name:        "number",
				Description: `Number of workers`,
			}},
		}, {
			Name:        []string{"--writeConcern"},
			Description: `Specifies the write concern for each write operation that mongoimport performs`,
			Args: []model.Arg{{
				Name: "document",
			}},
		}, {
			Name:        []string{"--bypassDocumentValidation"},
			Description: `Enables mongoimport to bypass document validation during the operation`,
		}, {
			Name:        []string{"--columnsHaveTypes"},
			Description: `Instructs mongoimport that the field list specified in --fields, --fieldFile, or --headerline specifies the types of each field`,
		}, {
			Name:        []string{"--parseGrace"},
			Description: `Specifies how mongoimport handles type coercion failures when importing CSV or TSV files with --columnsHaveTypes`,
			Args: []model.Arg{{
				Name: "grace",
				Suggestions: []model.Suggestion{{
					Name: []string{`autoCast`},
				}, {
					Name: []string{`skipField`},
				}, {
					Name: []string{`skipRow`},
				}, {
					Name: []string{`stop`},
				}},
			}},
		}},
	}
}
