## gyro service addon

Add-ons the service holds

```
gyro service addon [flags]
```

### Options

```
  -h, --help   help for addon
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)
* [gyro service addon available](gyro_service_addon_available.md)	 - Add-ons offered for this service (with options and prices)
* [gyro service addon create](gyro_service_addon_create.md)	 - Buy an add-on (paid from balance)
* [gyro service addon delete](gyro_service_addon_delete.md)	 - Remove a prepaid add-on at once
* [gyro service addon list](gyro_service_addon_list.md)	 - Add-ons the service holds

