## gyro server ip add

Add additional IPv4 addresses

### Synopsis

Add additional IPv4 addresses

**Minimum charge:** each added IP is billed for at least 6 hours (`pricing.minimum_hours`): releasing it earlier still bills the minimum (debited at release), and taking the same address again starts a new minimum. At most `adds_per_hour` addresses (default 10) can be added to a server per 60 minutes. Additional IPv4 addresses are billed per address hourly from the prepaid balance (default 2 EUR per address-month, `pricing` shows the current price), from the hour they are added until they are released. Up to `max` per server (default 5). Each address is configured on the server automatically (/32 on the primary interface, kept across reboots and reinstalls); reverse DNS: PUT /cloud/servers/{id}/rdns. Prepaid tr-ist-1 servers only; not while the server is suspended.

API: POST /cloud/servers/{id}/ips (operationId addServerIps, scope cloud_write)

```
gyro server ip add <ID> [flags]
```

### Options

```
      --count int                count
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro server ip](gyro_server_ip.md)	 - Print the public IPv4 (or --ipv6) of a server

