## gyro affiliate enroll

Enroll affiliate

### Synopsis

Enroll affiliate

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/satis-ortakligi-kazanci

API: POST /affiliate/enroll (operationId coreEnrollAffiliate, scope Affiliate/EnrollAffiliate)

```
gyro affiliate enroll [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for enroll
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro affiliate](gyro_affiliate.md)	 - Get affiliate

