## gyro service label get

Get the nickname and tags of a service

### Synopsis

Get the nickname and tags of a service

Customer-defined nickname and colour tags of one of your services (any product line). A service without a label answers `"nickname": null` and `"tags": []`. Labels are private to the service owner; after a transfer the old owner's label is not shown.

API: GET /cloud/services/{id}/label (operationId getServiceLabel, scope services_read)

```
gyro service label get <ID> [flags]
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

* [gyro service label](gyro_service_label.md)	 - Get the nickname and tags of a service

