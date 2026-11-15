## gyro auth

Sign in with the browser, show or end the CLI session

### Options

```
  -h, --help   help for auth
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
* [gyro auth login](gyro_auth_login.md)	 - Sign in with the browser (creates a time-limited API key for this profile)
* [gyro auth logout](gyro_auth_logout.md)	 - Remove the key from this profile and revoke it on the server
* [gyro auth status](gyro_auth_status.md)	 - Show the account, key, permissions and expiry of the current profile

