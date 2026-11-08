## gyro reseller customer create

Create an end-customer

### Synopsis

Create an end-customer

API: POST /cloud/reseller/customers (operationId createResellerCustomer, scope reseller)

```
gyro reseller customer create [flags]
```

### Options

```
      --allow-order              May this customer's keys order servers (paid from YOUR balance)?
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             email
      --external-ref string      Your own customer id (unique per reseller)
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --note string              note
      --set stringArray          extra body field key=value (repeatable)
      --status string            status (active|suspended)
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

* [gyro reseller customer](gyro_reseller_customer.md)	 - List end-customers

