## gyro service resize-option list

Products this service can be upgraded to, with prices

### Synopsis

Products this service can be upgraded to, with prices

Prepaid services answer `billing: prepaid`; each option has `plan`, `current` and one `hourly` price with `monthly_cap`.

API: GET /cloud/services/{id}/resize-options (operationId getResizeOptions, scope services_read)

```
gyro service resize-option list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro service resize-option](gyro_service_resize-option.md)	 - Products this service can be upgraded to, with prices

