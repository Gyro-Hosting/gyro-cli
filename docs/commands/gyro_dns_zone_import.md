## gyro dns zone import

Import a BIND zone file

### Synopsis

Import a BIND zone file

`merge` replaces record sets present in the file and keeps the others; `replace` also deletes record sets not in the file. SOA and apex NS are ignored; unsupported types are listed in `skipped`. Needs `"confirm": true`.

API: POST /cloud/dns/{id}/zones/{zone}/import (operationId typed_post_dns__id__zones__zone__import, scope services_write)

```
gyro dns zone import <ID> <ZONE> [flags]
```

### Options

```
      --bind string              bind
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for import
      --idempotency-key string   idempotency key (automatic for money operations)
      --mode string              mode (merge|replace)
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

* [gyro dns zone](gyro_dns_zone.md)	 - List DNS zones

