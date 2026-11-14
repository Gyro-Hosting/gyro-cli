## gyro vpc create

Create a private network

### Synopsis

Create a private network

tr-ist-1 only. Limits (GET /cloud/vpcs `limits`, raised with a support ticket): 3 VPCs per account, 4 subnets per VPC, 3 private cards per server, 20 port forwards per VPC; a server is attached to one VPC at a time (to several of its subnets). Another account's object answers 404. Private range /16 to /24 inside 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16; the network address must be given (10.10.0.0/16, not 10.10.0.5/16). Reserved by the platform: 10.42.0.0/16, 10.43.0.0/16, 10.200.0.0/16, 10.250.0.0/16, 10.255.0.0/16. The first subnet is created with it: `subnet.cidr` or, when left out, the first /24 of the range (name `default`). Free.

API: POST /cloud/vpcs (operationId createVpc, scope cloud_write)

```
gyro vpc create [flags]
```

### Options

```
      --cidr string              cidr
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --project-id int           project_id
      --set stringArray          extra body field key=value (repeatable)
      --subnet string            subnet (JSON)
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

