## gyro api request

Send a raw request to a path below the API URL

```
gyro api request <METHOD> <PATH> [flags]
```

### Examples

```
  gyro api request GET /cloud/servers -p status=active
  gyro api request POST /cloud/ssh-keys --set name=ci --set public_key=@~/.ssh/id_ed25519.pub
```

### Options

```
      --all                      follow pagination (GET)
      --data string              JSON body, @file or @-
  -h, --help                     help for request
      --idempotency-key string   send an Idempotency-Key header
  -P, --param stringArray        query parameter key=value
      --set stringArray          body field key=value (value @file reads a file)
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

* [gyro api](gyro_api.md)	 - Call any API endpoint directly (like 'gh api')

