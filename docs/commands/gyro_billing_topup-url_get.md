## gyro billing topup-url get

Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

### Synopsis

Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

Prepaid (hourly, from the account balance) services: Gyro Containers (ECS) and Gyro Cloud Istanbul (tr-ist-1). Times are Europe/Istanbul. Usage is never invoiced: invoices are issued for balance top-ups.

API: GET /cloud/billing/topup-url (operationId getPrepaidTopupUrl, scope billing_read)

```
gyro billing topup-url get [flags]
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

* [gyro billing topup-url](gyro_billing_topup-url.md)	 - Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

