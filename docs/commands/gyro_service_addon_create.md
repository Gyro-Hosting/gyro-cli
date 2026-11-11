## gyro service addon create

Buy an add-on (paid from balance)

### Synopsis

Buy an add-on (paid from balance)

Prepaid services + prepaid add-ons: no invoice, the add-on is active at once (`status: active`, `billing: prepaid`) and billed hourly from the balance.

API: POST /cloud/services/{id}/addons (operationId buyAddon, scope cloud_order)

```
gyro service addon create <ID> [flags]
```

### Options

```
      --addon-id int             addon_id
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --option-id int            option_id
      --quantity int             quantity
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

* [gyro service addon](gyro_service_addon.md)	 - Add-ons the service holds

