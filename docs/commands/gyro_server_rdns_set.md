## gyro server rdns set

Set (or reset) the PTR record of one address

### Synopsis

Set (or reset) the PTR record of one address

**tr-ist-1 (Gyro Bulut İstanbul):** every address (primary and additional IPv4) gets an automatic PTR `<a-b-c-d>.ip.gyrohosting.com` with a matching A record. Set `hostname` to your own FQDN to replace it (published within seconds); an empty value restores the default. Names under Gyro domains and protected brand names are refused (HTTP 422 with an explanatory message); at most 30 changes per server per hour. The A record of your name should point to the address (forward-confirmed reverse DNS) for e-mail delivery; it is not enforced. When the server is deleted or an additional IP is released, its PTR is removed. `dns_ptr` and `ptr` are accepted as aliases of `hostname`. Legacy Istanbul VDS servers keep the request flow (`status: pending` until applied).

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

