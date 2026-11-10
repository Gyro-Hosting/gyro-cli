## gyro lb target

List targets with health

```
gyro lb target [flags]
```

### Options

```
  -h, --help   help for target
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

* [gyro lb](gyro_lb.md)	 - List load balancers
* [gyro lb target add](gyro_lb_target_add.md)	 - Add a target
* [gyro lb target list](gyro_lb_target_list.md)	 - List targets with health
* [gyro lb target remove](gyro_lb_target_remove.md)	 - Remove a target

