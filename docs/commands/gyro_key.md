## gyro key

Sub-account API keys

```
gyro key [flags]
```

### Options

```
  -h, --help   help for key
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
* [gyro key create](gyro_key_create.md)	 - Create a sub-account key (token returned once)
* [gyro key delete](gyro_key_delete.md)	 - Delete a sub-account key
* [gyro key list](gyro_key_list.md)	 - List sub-account keys (never shows tokens)

