## gyro security-group attach

Attach a security group to a server

### Synopsis

Attach a security group to a server

API: POST /cloud/servers/{id}/security-groups/{sg} (operationId attachSecurityGroup, scope cloud_write)

```
gyro security-group attach <ID> <SG> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for attach
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

* [gyro security-group](gyro_security-group.md)	 - List security groups

