## gyro container service get

Get a container service

### Synopsis

Get a container service

Spec, state, problem code (ImagePullBackOff, CrashLoopBackOff, OOMKilled, Unschedulable ...), instances, revision, domains and the names of secrets (values are never returned). `autoscaling` = settings + live status (current/desired replicas, current CPU/memory %, at_max), `capacity` = {blocked, fits}. The answer carries `namespace` and, for services with a port, `internal` = `{host, short, port: 80, target_port, address, url}`: other services of the same environment reach it at `internal.address` (`<name>.<namespace>.svc.cluster.local:80`) or `<name>:80`.

API: GET /cloud/containers/{id}/services/{name} (operationId typed_get_containers__id__services__name, scope services_read)

```
gyro container service get <ID> <NAME> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro container service](gyro_container_service.md)	 - List container services

