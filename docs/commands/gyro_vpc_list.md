## gyro vpc list

List private networks

### Synopsis

List private networks

tr-ist-1 only. Limits (GET /cloud/vpcs `limits`, raised with a support ticket): 3 VPCs per account, 4 subnets per VPC, 3 private cards per server, 20 port forwards per VPC; a server is attached to one VPC at a time (to several of its subnets). Another account's object answers 404. Also returns `limits` and `pricing`.

API: GET /cloud/vpcs (operationId listVpcs, scope cloud_read)

```
gyro vpc list [flags]
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

* [gyro vpc](gyro_vpc.md)	 - List private networks

