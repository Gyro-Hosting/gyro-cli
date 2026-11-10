## gyro server list

List your servers

```
gyro server list [flags]
```

### Examples

```
  gyro server list
  gyro server list --status active -o json
  gyro server list --query "[?region=='eu-de-1'].ipv4" -o text
```

### Options

```
      --all             fetch every page
  -h, --help            help for list
      --status string   waiting|inprocess|active|suspended|cancelled
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

* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...

