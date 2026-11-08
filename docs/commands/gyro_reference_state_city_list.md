## gyro reference state city list

List cities (path: {id})

### Synopsis

List cities (path: {id})

List cities (path: {id}). Core WISECP client endpoint (Turkish reference: ).

API: GET /reference/states/{id}/cities (operationId coreGetCities, scope Reference/GetCities)

```
gyro reference state city list <ID> [flags]
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

* [gyro reference state city](gyro_reference_state_city.md)	 - List cities (path: {id})

