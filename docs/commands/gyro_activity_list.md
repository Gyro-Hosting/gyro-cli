## gyro activity list

List write calls made with your API keys

### Synopsis

List write calls made with your API keys

Every non-GET call to `/cloud/...` is recorded for 365 days (method, path, scope, HTTP status, IP, key id). Request and response bodies are never stored.

API: GET /cloud/activity (operationId listActivity, scope cloud_read)

```
gyro activity list [flags]
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

* [gyro activity](gyro_activity.md)	 - Write log of your API keys

