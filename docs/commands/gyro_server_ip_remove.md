## gyro server ip remove

Release an additional IPv4 address

### Synopsis

Release an additional IPv4 address

The address is removed from the server and returns to the pool (getting the same address back is not guaranteed); billing stops after the current hour. Needs `confirm=true`.

API: DELETE /cloud/servers/{id}/ips/{ip} (operationId removeServerIp, scope cloud_write)

```
gyro server ip remove <ID> <IP> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for remove
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro server ip](gyro_server_ip.md)	 - Print the public IPv4 (or --ipv6) of a server

