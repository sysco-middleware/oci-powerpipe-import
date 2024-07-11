# oci-powerpipe-import
                                 
Tool to import oci powerpipe snapshots into database. 

## Running imports
- build `env GOOS=linux GOARCH=amd64 go build -o importer  main.go ` from project root
- this will generate an executable binary `importer` from project root for linux 64-bit
- run the import
```bash
./importer ociCompliance -H <host-name> -s <schema-name> -u <user> -P <password> -p <port> -f <filepath>
```

### importer usage
```shell

./importer --help                                                                                                
Usage:
  importer [command]

Available Commands:
  help          Help about any command
  ociCompliance imports powerpipe oci compliance report

Flags:
  -f, --filepath string   * Import file path
  -h, --help              help for importer
  -H, --host string       * Database host (default "127.0.0.1")
  -P, --pass string       * Database password
  -p, --port int          * Database port (default 9133)
  -s, --schema string     * Database service (default "steampipe")
  -u, --user string       * Database user (default "steampipe")
```
### importer ociCompliance usage
```shell
./importer ociCompliance --help
Usage:
  importer ociCompliance [flags]

Flags:
  -h, --help   help for ociCompliance

Global Flags:
  -f, --filepath string   * Import file path
  -H, --host string       * Database host (default "127.0.0.1")
  -P, --pass string       * Database password
  -p, --port int          * Database port (default 9133)
  -s, --schema string     * Database service (default "steampipe")
  -u, --user string       * Database user (default "steampipe")
```

### example
```shell
 importer ociCompliance \
  -f <absolute-path>/oci_compliance.benchmark.cis_v200.20240711T095642.csv \
  -p 9193 \
  -P ******* \
  -H 127.0.0.1 \
  -s steampipe \
  -u steampipe
```
