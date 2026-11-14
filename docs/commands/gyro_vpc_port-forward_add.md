## gyro vpc port-forward add

Forward a public port of the gateway

### Synopsis

Forward a public port of the gateway

public_ip:public_port of the gateway -> target_ip:target_port inside the VPC. The gateway must be on (409 `gateway_disabled`); one forward per protocol + public port; at most 20.

API: POST /cloud/vpcs/{vpc_id}/port-forwards (operationId createVpcPortForward, scope cloud_write)

```
gyro vpc port-forward add <VPC-ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --protocol string          protocol (tcp|udp)
      --public-port int          public_port
      --set stringArray          extra body field key=value (repeatable)
      --target-ip string         target_ip
      --target-port int          target_port
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

* [gyro vpc port-forward](gyro_vpc_port-forward.md)	 - List port forwards

