## gyro service addon available get

Add-ons offered for this service (with options and prices)

### Synopsis

Add-ons offered for this service (with options and prices)

Prepaid add-on options show `cycle: hourly`, `billing: prepaid` and `price {amount (EUR/hour), monthly_cap}`.

API: GET /cloud/services/{id}/addons/available (operationId listAvailableAddons, scope services_read)

```
gyro service addon available get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro service addon available](gyro_service_addon_available.md)	 - Add-ons offered for this service (with options and prices)

