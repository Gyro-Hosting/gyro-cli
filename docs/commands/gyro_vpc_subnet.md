## gyro vpc subnet

List subnets

```
gyro vpc subnet [flags]
```

### Options

```
  -h, --help   help for subnet
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

* [gyro vpc](gyro_vpc.md)	 - List private networks
* [gyro vpc subnet create](gyro_vpc_subnet_create.md)	 - Add a subnet
* [gyro vpc subnet delete](gyro_vpc_subnet_delete.md)	 - Delete a subnet
* [gyro vpc subnet list](gyro_vpc_subnet_list.md)	 - List subnets

