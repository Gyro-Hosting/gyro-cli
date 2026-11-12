## gyro container service

List container services

```
gyro container service [flags]
```

### Options

```
  -h, --help   help for service
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains
* [gyro container service autoscale](gyro_container_service_autoscale.md)	 - Autoscaling / alert threshold of a service
* [gyro container service create](gyro_container_service_create.md)	 - Create a container service
* [gyro container service delete](gyro_container_service_delete.md)	 - Delete a container service
* [gyro container service get](gyro_container_service_get.md)	 - Get a container service
* [gyro container service list](gyro_container_service_list.md)	 - List container services
* [gyro container service restart](gyro_container_service_restart.md)	 - Restart a service
* [gyro container service revisions](gyro_container_service_revisions.md)	 - Revision history
* [gyro container service rollback](gyro_container_service_rollback.md)	 - Roll back to an earlier revision
* [gyro container service scale](gyro_container_service_scale.md)	 - Scale a service
* [gyro container service update](gyro_container_service_update.md)	 - Deploy a new version / change settings

