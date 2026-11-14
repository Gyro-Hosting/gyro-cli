## gyro vpc

List private networks

```
gyro vpc [flags]
```

### Options

```
  -h, --help   help for vpc
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro vpc create](gyro_vpc_create.md)	 - Create a private network
* [gyro vpc delete](gyro_vpc_delete.md)	 - Delete a private network
* [gyro vpc gateway](gyro_vpc_gateway.md)	 - Get the VPC gateway
* [gyro vpc get](gyro_vpc_get.md)	 - Get a private network
* [gyro vpc list](gyro_vpc_list.md)	 - List private networks
* [gyro vpc port-forward](gyro_vpc_port-forward.md)	 - List port forwards
* [gyro vpc subnet](gyro_vpc_subnet.md)	 - List subnets
* [gyro vpc update](gyro_vpc_update.md)	 - Rename a private network

