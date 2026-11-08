## gyro server power off

Hard power off

```
gyro server power off <SERVER> [flags]
```

### Options

```
  -h, --help                    help for off
      --wait                    wait until the server is stopped
      --wait-timeout duration   maximum time for --wait (default 10m0s)
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

* [gyro server power](gyro_server_power.md)	 - Power on, off, shutdown, reboot or reset a server

