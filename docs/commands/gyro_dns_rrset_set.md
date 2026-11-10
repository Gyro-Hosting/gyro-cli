## gyro dns rrset set

Create or replace a record set

### Synopsis

Create or replace a record set

Upserts one record set (name + type). CNAME and ALIAS take one value; no CNAME at the apex (use ALIAS); apex NS and SOA are managed by us.

API: PUT /cloud/dns/{id}/zones/{zone}/records (operationId typed_put_dns__id__zones__zone__records, scope services_write)

```
gyro dns rrset set <ID> <ZONE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              Record name relative to the zone (@ = apex, * = wildcard) or FQDN with trailing dot
      --records strings          Values in zone file syntax, one per item (TXT quotes optional)
      --set stringArray          extra body field key=value (repeatable)
      --ttl int                  60-86400 s
      --type string              type (A|AAAA|CNAME|MX|TXT|SRV|CAA|ALIAS|NS|PTR)
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

* [gyro dns rrset](gyro_dns_rrset.md)	 - List record sets

