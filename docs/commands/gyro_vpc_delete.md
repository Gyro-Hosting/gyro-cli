## gyro vpc delete

Delete a private network

### Synopsis

Delete a private network

Permanent. Detach the server cards and delete the port forwards first (409 `vpc_in_use`); an open gateway is turned off. Needs `"confirm": true` (body or query).

API: DELETE /cloud/vpcs/{vpc_id} (operationId deleteVpc, scope cloud_write)

```
gyro vpc delete <VPC-ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro vpc](gyro_vpc.md)	 - List private networks

