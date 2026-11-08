## gyro service cancellation create

Create cancellation (path: {id})

### Synopsis

Create cancellation (path: {id})

Create cancellation (path: {id}). Service of the calling account (any product line). Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-hizmetler).

API: POST /services/{id}/cancellation (operationId coreCreateCancellation, scope Services/CreateCancellation)

```
gyro service cancellation create <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --note string              Operatöre not.
      --reason string            Gerekçe. Beş değerden biri: gerek kalmadı, pahalı, başka yere geçiliyor, özellik eksik ya da diğer.
      --reason-detail string     Serbest açıklama. Gerekçe diğer ise zorunludur.
      --set stringArray          extra body field key=value (repeatable)
      --urgency string           Ne zaman iptal edilsin: hemen ya da dönem sonunda.
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

* [gyro service cancellation](gyro_service_cancellation.md)	 - Create cancellation (path: {id})

