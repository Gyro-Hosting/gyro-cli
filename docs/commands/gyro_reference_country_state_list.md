## gyro reference country state list

List states (path: {code})

### Synopsis

List states (path: {code})

List states (path: {code}). Core WISECP client endpoint.

API: GET /reference/countries/{code}/states (operationId coreGetStates, scope Reference/GetStates)

```
gyro reference country state list <CODE> [flags]
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

* [gyro reference country state](gyro_reference_country_state.md)	 - List states (path: {code})

