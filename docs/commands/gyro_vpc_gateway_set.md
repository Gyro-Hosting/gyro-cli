## gyro vpc gateway set

Turn the VPC gateway on/off and change its settings

### Synopsis

Turn the VPC gateway on/off and change its settings

The gateway gives private-only servers internet access (NAT), port forwarding and the inter-subnet rule. While it is on it is billed hourly from the prepaid balance (`monthly_eur` / 672, at least 1 hour, never more than `monthly_eur` a month) and needs at least one prepaid tr-ist-1 server on the account (409 `no_billing_server`). Turning it off releases its public IP; refused while port forwards exist, or while private-only servers use it (409 `gateway_in_use`, repeat with `allow_no_egress: true`). Outgoing connections are logged as required by Turkish Law No. 5651 (2 years). Only the given fields change.

API: PUT /cloud/vpcs/{vpc_id}/gateway (operationId setVpcGateway, scope cloud_write)

```
gyro vpc gateway set <VPC-ID> [flags]
```

### Options

```
      --allow-no-egress          allow_no_egress
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  enabled
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --inter-subnet string      inter_subnet (allow|deny)
      --nat                      nat
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro vpc gateway](gyro_vpc_gateway.md)	 - Get the VPC gateway

