## gyro balance

Account balance (bare command) and transactions

```
gyro balance [flags]
```

### Options

```
  -h, --help   help for balance
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro balance setting](gyro_balance_setting.md)	 - Update balance settings
* [gyro balance top-up](gyro_balance_top-up.md)	 - Create a balance top-up invoice and get its payment link
* [gyro balance transaction](gyro_balance_transaction.md)	 - Balance ledger

