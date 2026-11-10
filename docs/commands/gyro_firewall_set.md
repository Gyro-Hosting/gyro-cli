## gyro firewall set

Replace the network firewall policy

### Synopsis

Replace the network firewall policy

Applies as a 5 minute trial unless `permanent` is true; confirm with POST /firewall/policy/confirm. Warnings return 422 `confirm_warnings_required` with `details.warnings`.

API: PUT /cloud/servers/{id}/firewall/policy (operationId putFirewallPolicy, scope cloud_write)

```
gyro firewall set <ID> [flags]
```

### Options

```
      --confirm-warnings         confirm_warnings
      --data string              request body as JSON, @file or @- (stdin)
      --default-action string    default_action (allow|deny)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --permanent                Skip the 5 minute trial (otherwise call POST /firewall/policy/confirm within 5 minutes or the change is reverted)
      --ping                     ping
      --rules strings            rules
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

* [gyro firewall](gyro_firewall.md)	 - Network firewall policy (rules, default action, status)

