## gyro vpc subnet create

Add a subnet

### Synopsis

Add a subnet

/20 to /28 inside the VPC range, not overlapping another subnet. Give `cidr`, or `prefix` (default 24) for the next free range. In every subnet .1 is the gateway; .0-.4 and the last address are reserved.

API: POST /cloud/vpcs/{vpc_id}/subnets (operationId createVpcSubnet, scope cloud_write)

```
gyro vpc subnet create <VPC-ID> [flags]
```

### Options

```
      --cidr string              cidr
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --prefix int               prefix
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro vpc subnet](gyro_vpc_subnet.md)	 - List subnets

