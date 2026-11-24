## gyro account auto-topup get

Automatic balance top-up setting

### Synopsis

Automatic balance top-up setting

When the balance drops below `threshold`, `amount` (TRY, excl. VAT) is charged to the saved card and credited as a balance top-up invoice incl. VAT. Checked hourly. Limits per day/month apply.

API: GET /cloud/account/auto-topup (operationId getAutoTopup, scope billing_read)

```
gyro account auto-topup get [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro account auto-topup](gyro_account_auto-topup.md)	 - Automatic balance top-up setting

