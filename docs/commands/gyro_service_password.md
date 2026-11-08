## gyro service password

Change service password (path: {id})

### Synopsis

Change service password (path: {id})

Change service password (path: {id}). Service of the calling account (any product line). Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/hizmetin-icine-erisim).

API: POST /services/{id}/password (operationId coreChangeServicePassword, scope Services/ChangeServicePassword)

```
gyro service password <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for password
      --idempotency-key string   idempotency key (automatic for money operations)
      --password string          Yeni panel parolası. Operatörün asgari uzunluğunu karşılamalıdır.
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)

