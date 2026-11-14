## gyro server nic detach

Detach a private network card

### Synopsis

Detach a private network card

The address is held for 10 minutes. While the public IP is off the last card cannot be detached (409 `last_network`).

API: DELETE /cloud/servers/{id}/nics/{nic_id} (operationId detachServerNic, scope cloud_write)

```
gyro server nic detach <ID> <NIC-ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for detach
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

* [gyro server nic](gyro_server_nic.md)	 - List the server's private network cards

