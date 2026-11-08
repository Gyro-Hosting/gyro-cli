## gyro service list

List services of every product line

### Synopsis

List services of every product line

API: GET /cloud/services (operationId listServices, scope services_read)

```
gyro service list [flags]
```

### Options

```
      --all                   fetch every page
      --customer-id int       Reseller main key only
  -h, --help                  help for list
      --limit int             items per page (max 100)
      --page int              page number (default 1)
      --product-line string    (cloud_server|vds|dedicated_server|mail|monitor|storage|license|app)
      --status string         
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)

