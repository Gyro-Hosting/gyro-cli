## gyro template list

List your templates

### Synopsis

List your templates

Billed 0.05 EUR per GB-month of the space the template really uses, hourly from the prepaid balance (through the account's template storage service, opened with the first template). Deleting a template stops its charge. At most 10 templates; visible only in your account. `meta` carries count, max, total_gb and pricing.

API: GET /cloud/templates (operationId listTemplates, scope cloud_read)

```
gyro template list [flags]
```

### Options

```
      --all         fetch every page
  -h, --help        help for list
      --limit int   items per page (max 100)
      --page int    page number (default 1)
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

* [gyro template](gyro_template.md)	 - List your templates

