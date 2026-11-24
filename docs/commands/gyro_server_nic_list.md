## gyro server nic list

List the server's private network cards

### Synopsis

List the server's private network cards

Private network cards of a tr-ist-1 server. Attaching is hot (no restart): the card is ready in about 10-60 s on a running server (`state` attaching -> ready); on a stopped server the settings apply at boot. A server can join one VPC at a time (409 `other_vpc`), at most 3 cards. `ip` empty = automatic.

API: GET /cloud/servers/{id}/nics (operationId listServerNics, scope cloud_read)

```
gyro server nic list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

