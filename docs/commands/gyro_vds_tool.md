## gyro vds tool

Read a tool's data or run one of its actions (see 'tools')

```
gyro vds tool <SERVICE> <TOOL> [ACTION] [flags]
```

### Examples

```
  gyro mail tool example.com mailboxes
  gyro mail tool example.com mailboxes create --set local_part=info --set password=@-
```

### Options

```
      --data string       JSON body, @file or @-
  -h, --help              help for tool
      --set stringArray   field key=value (repeatable)
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

* [gyro vds](gyro_vds.md)	 - VDS servers

