## gyro template

List your templates

```
gyro template [flags]
```

### Options

```
  -h, --help   help for template
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
* [gyro template create](gyro_template_create.md)	 - Capture a server as a template
* [gyro template delete](gyro_template_delete.md)	 - Delete a template
* [gyro template deploy](gyro_template_deploy.md)	 - New server from a template
* [gyro template get](gyro_template_get.md)	 - Get a template
* [gyro template list](gyro_template_list.md)	 - List your templates
* [gyro template rebuild](gyro_template_rebuild.md)	 - Reinstall a server from a template
* [gyro template update](gyro_template_update.md)	 - Rename a template

