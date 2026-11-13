## gyro container plans

List container environment plans

### Synopsis

List container environment plans

Orderable Gyro Container (ECS) environment plans with prepaid prices. A plan is ONE environment (a dedicated node pool); `max_services` is the most apps you can define, `max_replicas` the most replicas per service; all of them share the pool.

API: GET /cloud/containers/plans (operationId listContainerPlans, scope cloud_read)

```
gyro container plans [flags]
```

### Options

```
  -h, --help   help for plans
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

