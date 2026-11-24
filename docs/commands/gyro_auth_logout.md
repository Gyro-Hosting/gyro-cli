## gyro auth logout

Remove the key from this profile and revoke it on the server

### Synopsis

Keys created by 'gyro auth login' are revoked on the server and removed from the profile.
A key you pasted with 'gyro configure' is only removed locally (it may be used elsewhere); add --revoke to revoke it too.

```
gyro auth logout [flags]
```

### Options

```
  -h, --help       help for logout
      --keep-key   only remove the key locally, never revoke it
      --revoke     also revoke a key that was not created by gyro auth login
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

* [gyro auth](gyro_auth.md)	 - Sign in with the browser, show or end the CLI session

