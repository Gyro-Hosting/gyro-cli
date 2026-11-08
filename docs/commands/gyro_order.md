## gyro order

Orders (preview and create)

```
gyro order [flags]
```

### Options

```
  -h, --help   help for order
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
* [gyro order create](gyro_order_create.md)	 - Create order
* [gyro order get](gyro_order_get.md)	 - Get order (path: {id})
* [gyro order list](gyro_order_list.md)	 - List orders
* [gyro order preview](gyro_order_preview.md)	 - Preview order

