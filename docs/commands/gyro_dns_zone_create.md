## gyro dns zone create

Add a zone

### Synopsis

Add a zone

Claims a domain. It is published on ns1/ns2.gyrodns.net after ownership is verified: delegation to our name servers (immediate when already delegated) or the TXT record `_gyrodns-dogrulama.<zone>` = `txt_value` at your current DNS. Optional `bind` (zone file) is imported on activation (SOA and apex NS are ignored).

API: POST /cloud/dns/{id}/zones (operationId typed_post_dns__id__zones, scope services_write)

```
gyro dns zone create <ID> [flags]
```

### Options

```
      --bind string              Optional BIND zone file (max 2 MB)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              Zone name (IDN allowed)
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

