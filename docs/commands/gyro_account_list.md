## gyro account list

Accounts you can act in with --account (your own + accounts you are a panel user of)

### Synopsis

Lists your own account and every account whose owner added you under Account > Users, with your permissions there.
Act in one of them with --account <id>, GYRO_ACCOUNT=<id> or make it the profile default with 'gyro account use <id>'.
Permissions are exactly the panel permissions; owner-only operations (API keys, webhooks, users, balance, cards) are refused.

```
gyro account list [flags]
```

### Examples

```
  gyro account list
  gyro server list --account 12345
  gyro account use 12345
```

### Options

```
  -h, --help   help for list
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

* [gyro account](gyro_account.md)	 - Accounts you can act in, account settings

