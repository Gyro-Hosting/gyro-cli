## gyro service addon delete

Remove a prepaid add-on at once

### Synopsis

Remove a prepaid add-on at once

Prepaid add-ons only (invoiced add-ons are cancelled in the panel at the end of their term). Billing of the item stops at once.

API: DELETE /cloud/services/{id}/addons/{aid} (operationId removeServiceAddon, scope cloud_write)

```
gyro service addon delete <ID> <AID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
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

* [gyro service addon](gyro_service_addon.md)	 - Add-ons the service holds

