## gyro vpc port-forward

List port forwards

```
gyro vpc port-forward [flags]
```

### Options

```
  -h, --help   help for port-forward
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

* [gyro vpc](gyro_vpc.md)	 - List private networks
* [gyro vpc port-forward add](gyro_vpc_port-forward_add.md)	 - Forward a public port of the gateway
* [gyro vpc port-forward list](gyro_vpc_port-forward_list.md)	 - List port forwards
* [gyro vpc port-forward remove](gyro_vpc_port-forward_remove.md)	 - Delete a port forward
* [gyro vpc port-forward update](gyro_vpc_port-forward_update.md)	 - Change who may use a port forward (source allow list)

