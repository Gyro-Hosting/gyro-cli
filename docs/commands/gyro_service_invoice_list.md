## gyro service invoice list

List service invoices (path: {id})

### Synopsis

List service invoices (path: {id})

List service invoices (path: {id}). Service of the calling account (any product line). Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-hizmetler).

API: GET /services/{id}/invoices (operationId coreGetServiceInvoices, scope Services/GetServiceInvoices)

```
gyro service invoice list <ID> [flags]
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

* [gyro service invoice](gyro_service_invoice.md)	 - List service invoices (path: {id})

