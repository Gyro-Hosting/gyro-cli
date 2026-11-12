## gyro container service update

Deploy a new version / change settings

### Synopsis

Deploy a new version / change settings

Send only the fields to change (e.g. `image`). Every change except `replicas` creates a new revision and rolls out instance by instance without downtime. With autoscaling on, `replicas` is ignored (the autoscaler owns it). `autoscaling` / `alerts` are settings, not revisions.

API: PATCH /cloud/containers/{id}/services/{name} (operationId typed_patch_containers__id__services__name, scope services_write)

```
gyro container service update <ID> <NAME> [flags]
```

### Options

```
      --alerts string               Alert threshold: notify when fewer replicas are ready (0 = off). Channels: PUT /cloud/containers/{id}/notifications. (JSON)
      --args strings                Arguments for the image's default command
      --autoscaling enabled:false   Autoscaling (horizontal): replicas follow the average CPU (and optionally memory) use of the reserved size. enabled:false = off (fixed replicas). (JSON)
      --command strings             Override the image entrypoint
      --cpu-m int                   CPU per instance in millicores (250 = 0.25 vCPU)
      --data string                 request body as JSON, @file or @- (stdin)
      --dry-run                     show the price quote / request without changing anything
      --env string                  Plain environment variables (replaces the whole set on update) (JSON)
      --health-path string          HTTP health check path (empty = TCP check)
  -h, --help                        help for update
      --idempotency-key string      idempotency key (automatic for money operations)
      --image string                Image reference, e.g. ghcr.io/acme/web:1.4 (public registries or your private registry)
      --memory-mi int               Memory per instance in MiB
      --port int                    Container port (0 = no network port). Containers run as a non-root user: use a port >= 1024
      --public                      HTTPS address <name>-<code>.gyroapp.net (and your domains)
      --reason string               Free text shown in the revision history (e.g. a git sha)
      --replicas int                Number of instances (default 1)
      --run-as-user int             User id inside the container (default 10001)
      --set stringArray             extra body field key=value (repeatable)
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

