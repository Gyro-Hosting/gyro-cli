## gyro dns rrset delete

Delete a record set

### Synopsis

Delete a record set

Deletes the record set given by `name` and `type` (query). Needs `"confirm": true`.

API: DELETE /cloud/dns/{id}/zones/{zone}/records (operationId typed_delete_dns__id__zones__zone__records, scope services_write)

```
gyro dns rrset delete <ID> <ZONE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              Record name
      --set stringArray          extra body field key=value (repeatable)
      --type string              Record type
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

