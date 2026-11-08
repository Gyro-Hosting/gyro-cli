## gyro server firewall rule add

Add an inbound firewall rule

### Synopsis

Add an inbound firewall rule

API: POST /cloud/servers/{id}/firewall/rules (operationId addFirewallRule, scope cloud_write)

```
gyro server firewall rule add <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --port string              N or N-M (empty = all ports)
      --protocol string          protocol (tcp|udp|icmp)
      --set stringArray          extra body field key=value (repeatable)
      --source-ips strings       CIDRs (default any)
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

* [gyro server firewall rule](gyro_server_firewall_rule.md)	 - Add an inbound firewall rule

