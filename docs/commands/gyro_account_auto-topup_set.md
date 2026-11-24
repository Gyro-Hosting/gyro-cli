## gyro account auto-topup set

Turn automatic top-up on/off

### Synopsis

Turn automatic top-up on/off

Enabling needs `consent: true` (authorises the charges). `amount` at least the minimum (250 TRY), at most `limits.max_single`. Omitted fields keep their value. `card_id` 0 = the auto-payment card.

API: PUT /cloud/account/auto-topup (operationId setAutoTopup, scope billing_pay)

```
gyro account auto-topup set [flags]
```

### Options

```
      --amount float             amount
      --card-id int              card_id
      --consent                  consent
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  enabled
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --threshold float          threshold
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

