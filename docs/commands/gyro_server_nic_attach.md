## gyro server nic attach

Attach the server to a subnet

### Synopsis

Attach the server to a subnet

Private network cards of a tr-ist-1 server. Attaching is hot (no restart): the card is ready in about 10-60 s on a running server (`state` attaching -> ready); on a stopped server the settings apply at boot. A server can join one VPC at a time (409 `other_vpc`), at most 3 cards. `ip` empty = automatic.

API: POST /cloud/servers/{id}/nics (operationId attachServerNic, scope cloud_write)

```
gyro server nic attach <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for attach
      --idempotency-key string   idempotency key (automatic for money operations)
      --ip string                ip
      --set stringArray          extra body field key=value (repeatable)
      --subnet-id string         subnet_id
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro server nic](gyro_server_nic.md)	 - List the server's private network cards

