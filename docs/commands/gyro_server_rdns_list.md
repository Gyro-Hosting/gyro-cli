## gyro server rdns list

List reverse DNS entries

### Synopsis

List reverse DNS entries

**tr-ist-1 (Gyro Bulut İstanbul):** every address (primary and additional IPv4) gets an automatic PTR `<a-b-c-d>.ip.gyrohosting.com` with a matching A record. Set `hostname` to your own FQDN to replace it (published within seconds); an empty value restores the default. Names under Gyro domains and protected brand names are refused (HTTP 422 with an explanatory message); at most 30 changes per server per hour. The A record of your name should point to the address (forward-confirmed reverse DNS) for e-mail delivery; it is not enforced. When the server is deleted or an additional IP is released, its PTR is removed.

API: GET /cloud/servers/{id}/rdns (operationId listRdns, scope cloud_read)

```
gyro server rdns list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

