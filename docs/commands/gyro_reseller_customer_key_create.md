## gyro reseller customer key create

Mint a key for an end-customer (token returned once)

### Synopsis

Mint a key for an end-customer (token returned once)

API: POST /cloud/reseller/customers/{id}/keys (operationId createResellerCustomerKey, scope reseller)

```
gyro reseller customer key create <ID> [flags]
```

### Options

```
      --actions strings          Optional action allow list: server actions (reboot) or tool:action / tool:* / method:name
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --ips strings              ips
      --name string              name
      --rate-limit int           Requests per minute (<= your own limit; 0 = default)
      --scopes strings           scopes
      --services int64Slice      Optional service allow list (default [])
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

* [gyro reseller customer key](gyro_reseller_customer_key.md)	 - List an end-customer's keys

