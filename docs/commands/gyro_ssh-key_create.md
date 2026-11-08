## gyro ssh-key create

Upload an SSH public key

```
gyro ssh-key create [PUBLIC-KEY-FILE] [flags]
```

### Examples

```
  gyro ssh-key create                      # ~/.ssh/id_ed25519.pub
  gyro ssh-key create ~/.ssh/work.pub --name work
  cat key.pub | gyro ssh-key create - --name ci
```

### Options

```
  -h, --help          help for create
      --name string   key name (default: the key comment)
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

* [gyro ssh-key](gyro_ssh-key.md)	 - SSH public keys of your account

