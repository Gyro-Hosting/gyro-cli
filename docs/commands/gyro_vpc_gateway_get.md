## gyro vpc gateway get

Get the VPC gateway

### Synopsis

Get the VPC gateway

The gateway gives private-only servers internet access (NAT), port forwarding and the inter-subnet rule. While it is on it is billed hourly from the prepaid balance (`monthly_eur` / 672, at least 1 hour, never more than `monthly_eur` a month) and needs at least one prepaid tr-ist-1 server on the account (409 `no_billing_server`). Turning it off releases its public IP; refused while port forwards exist, or while private-only servers use it (409 `gateway_in_use`, repeat with `allow_no_egress: true`). Outgoing connections are logged as required by Turkish Law No. 5651 (2 years).

API: GET /cloud/vpcs/{vpc_id}/gateway (operationId getVpcGateway, scope cloud_read)

```
gyro vpc gateway get <VPC-ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro vpc gateway](gyro_vpc_gateway.md)	 - Get the VPC gateway

