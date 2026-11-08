## gyro configure

Store an API key and preferences in a named profile

### Synopsis

Interactive setup, like 'aws configure'. Profiles live in ~/.config/gyro/config.toml (mode 0600).
Use --profile NAME for more than one account; --keyring keeps the key in the OS keychain instead of the file.
Non-interactive: echo "$KEY" | gyro configure --token-stdin --profile ci

```
gyro configure [flags]
```

### Options

```
  -h, --help                 help for configure
      --keyring              store the key in the OS keychain
      --no-verify            do not call /whoami to check the key
      --set-api-url string   API URL to store in the profile
      --set-lang string      language to store (tr/en)
      --set-output string    default output to store
      --token-stdin          read the key from stdin (non-interactive)
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
* [gyro configure delete](gyro_configure_delete.md)	 - Remove a profile
* [gyro configure get](gyro_configure_get.md)	 - Print a profile setting (api_url, output, lang, ssh_user, ssh_identity)
* [gyro configure list](gyro_configure_list.md)	 - List profiles
* [gyro configure set](gyro_configure_set.md)	 - Change a profile setting (the API key: use 'gyro configure --token-stdin')

