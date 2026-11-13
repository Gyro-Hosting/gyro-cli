## gyro container

Gyro Containers (ECS), Istanbul: services, versions, logs, domains

```
gyro container [flags]
```

### Options

```
  -h, --help   help for container
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
* [gyro container create](gyro_container_create.md)	 - Create container environments (batch)
* [gyro container domain](gyro_container_domain.md)	 - Domains of a service
* [gyro container env](gyro_container_env.md)	 - Set environment variables
* [gyro container events](gyro_container_events.md)	 - Recent warnings
* [gyro container external-registry](gyro_container_external-registry.md)	 - Remove external registry credentials
* [gyro container logs](gyro_container_logs.md)	 - Service logs
* [gyro container metrics](gyro_container_metrics.md)	 - Service CPU / memory
* [gyro container nodes](gyro_container_nodes.md)	 - Node pool
* [gyro container notifications](gyro_container_notifications.md)	 - Notification settings and current alerts
* [gyro container pipeline](gyro_container_pipeline.md)	 - List pipelines
* [gyro container plans](gyro_container_plans.md)	 - List container environment plans
* [gyro container registry](gyro_container_registry.md)	 - Gyro image registry of the environment
* [gyro container secret](gyro_container_secret.md)	 - Delete secret values
* [gyro container service](gyro_container_service.md)	 - List container services
* [gyro container sizes](gyro_container_sizes.md)	 - Serverless container sizes (vCPU / memory pairs) and prices

