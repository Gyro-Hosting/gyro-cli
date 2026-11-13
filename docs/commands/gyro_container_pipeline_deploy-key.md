## gyro container pipeline deploy-key

Create the SSH deploy key

```
gyro container pipeline deploy-key [flags]
```

### Options

```
  -h, --help   help for deploy-key
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

* [gyro container pipeline](gyro_container_pipeline.md)	 - List pipelines
* [gyro container pipeline deploy-key create](gyro_container_pipeline_deploy-key_create.md)	 - Create the SSH deploy key
* [gyro container pipeline deploy-key delete](gyro_container_pipeline_deploy-key_delete.md)	 - Delete the SSH deploy key

