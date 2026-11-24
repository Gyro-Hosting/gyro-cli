## gyro activity get

Account activity (sign-ins and security events)

### Synopsis

Account activity (sign-ins and security events)

The records of **Account > Activity** in the client area (`/myaccount/info#activity`): sign-ins with IP, location, browser and platform, and account events (password, two-factor, passkey, session and trusted-device changes, profile and address changes ...). Newest first. Only your own account: the key decides the account, there is no account parameter, and sub-account keys get 403. Same privacy as the panel: full IP of your own sign-ins; the raw user agent and stored event data are not returned. Needs the `account_read` permission (Account / read on the key screen; the read-only template includes it). The write log of your API keys stays at `GET /cloud/activity`.

API: GET /activity (operationId getAccountActivity, scope account_read)

```
gyro activity get [flags]
```

### Options

```
      --all                fetch every page
      --from string        From this day (YYYY-MM-DD, whole day, Türkiye time) or time (ISO 8601), inclusive
  -h, --help               help for get
      --limit int          items per page (max 100)
      --page int           page number (default 1)
      --param-lang title   Language of title (default: the account language) (tr|en)
      --to string          Until this day (YYYY-MM-DD, whole day) or time (ISO 8601), inclusive
      --type string        login = sign-ins only, action = account events only (all|login|action)
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

* [gyro activity](gyro_activity.md)	 - Account activity (sign-ins and security events)

