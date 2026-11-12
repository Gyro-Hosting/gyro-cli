## gyro server rdns set

Set (or reset) the PTR record of one address

### Synopsis

Set (or reset) the PTR record of one address

`dns_ptr` and `ptr` are accepted as aliases of `hostname`. tr-ist-1: a PTR on an address without a record is created (request; goes live after approval, `status: pending`); resetting an address without a record is a no-op.

API: PUT /cloud/servers/{id}/rdns (operationId setRdns, scope cloud_write)

```
gyro server rdns set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --hostname string          FQDN; empty string resets to the default
      --idempotency-key string   idempotency key (automatic for money operations)
      --ip string                An address of this server (IPv6 addresses inside the server's /64 are allowed)
      --set stringArray          extra body field key=value (repeatable)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro server rdns](gyro_server_rdns.md)	 - List reverse DNS entries

