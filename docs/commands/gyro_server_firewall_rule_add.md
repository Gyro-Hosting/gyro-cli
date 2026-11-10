## gyro server firewall rule add

Add a firewall rule

### Synopsis

Add a firewall rule

Network firewall add-on: inserts at `priority` (default last) and applies as a trial unless `permanent`. Otherwise the cloud server's simple inbound allow list.

API: POST /cloud/servers/{id}/firewall/rules (operationId addFirewallRule, scope cloud_write)

```
gyro server firewall rule add <ID> [flags]
```

### Options

```
      --action string            action (allow|deny)
      --confirm-warnings         Apply even when the change has warnings (lock-out risk, shadowed or duplicate rules, private sources)
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --destination string       any (all IPs of the server) or one IP address of the server
      --direction string         in = inbound to the server, out = new outbound connections (network firewall only) (in|out)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --permanent                Network firewall - skip the 5 minute trial and make the change permanent at once
      --port string              any | 22 | 80,443 | 1000-2000 (not for icmp/all)
      --priority int             Position (1 = evaluated first). Default = last. First matching rule wins.
      --protocol string          protocol (tcp|udp|tcp_udp|icmp|all)
      --set stringArray          extra body field key=value (repeatable)
      --source string            any | IPv4/CIDR list (comma separated, max 10) | country:TR,DE | not-country:TR (country add-on; any ISO 3166-1 alpha-2 code with RIPE address space, max 10 per rule). For direction out: the remote address.
      --source-ips strings       Legacy - CIDRs (use source)
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

* [gyro server firewall rule](gyro_server_firewall_rule.md)	 - Add a firewall rule

