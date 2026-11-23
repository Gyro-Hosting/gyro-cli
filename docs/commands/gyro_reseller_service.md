## gyro reseller service

Reseller view of a service (status, due date, open invoices)

```
gyro reseller service [flags]
```

### Options

```
  -h, --help   help for service
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

* [gyro reseller](gyro_reseller.md)	 - Reseller: end-customers, assignments, keys, report
* [gyro reseller service get](gyro_reseller_service_get.md)	 - Reseller view of a service (status, due date, open invoices)
* [gyro reseller service suspend](gyro_reseller_service_suspend.md)	 - Suspend a service (reseller)
* [gyro reseller service terminate](gyro_reseller_service_terminate.md)	 - Terminate a service (reseller)
* [gyro reseller service unsuspend](gyro_reseller_service_unsuspend.md)	 - Unsuspend a service (reseller)

