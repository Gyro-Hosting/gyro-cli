## gyro auth login

Sign in with the browser (creates a time-limited API key for this profile)

### Synopsis

Opens the Gyro panel in your browser; you approve the request there (permissions and expiry are chosen on that page)
and the new key is stored in the profile, exactly like 'gyro configure'. No key is copied by hand.
Over SSH or without a display the address is only printed: open it on any device where you are signed in.
In CI keep using a key: echo "$KEY" | gyro configure --token-stdin

```
gyro auth login [flags]
```

### Examples

```
  gyro auth login
  gyro auth login --profile work --no-browser
  gyro auth login --template readonly
```

### Options

```
  -h, --help              help for login
      --keyring           store the key in the OS keychain
      --no-browser        only print the address, do not open a browser
      --scopes string     preselect scopes (comma separated, e.g. cloud_read,services_read); you can still change them in the browser
      --template string   preselect a permission template: readonly, dns, cloud, containers, storage, billing, full
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

* [gyro auth](gyro_auth.md)	 - Sign in with the browser, show or end the CLI session

