## gyro container registry import

Import an image from another registry

### Synopsis

Import an image from another registry

Copies an image (e.g. `nginx:1.27`, `docker.io/library/redis:7`, `ghcr.io/gyro-ornek/uygulama:1.0`) into the environment's Gyro registry without `docker push`. HTTPS registries only; only the linux/amd64 variant of multi-platform images is copied. `repo` (default: last part of the source name) and `tag` (default: source tag) name the copy under `<host>/<code>/`; an existing name:tag needs `overwrite: true`. Private images: `username` + `password` (token) are used only for this copy, never stored or returned, and redacted in the API log. Runs in the background: poll GET .../registry/imports. Limits: 1 running import per environment (409 `import_running`), 10 per hour (429 `import_rate_limited`), 15 minutes, 128 layers, size within the free registry space (50 GB cap, 2 GB free, then the usual €0.05/GB-month). Sources that resolve to private, internal or Gyro infrastructure addresses are refused (`source_blocked`).

API: POST /cloud/containers/{id}/registry/imports (operationId importContainerImage, scope services_write)

```
gyro container registry import <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for import
      --idempotency-key string   idempotency key (automatic for money operations)
      --overwrite                replace an existing name:tag
      --password string          private source: token / password (not stored, redacted in logs)
      --repo string              target name under your prefix (lowercase, max 3 levels)
      --set stringArray          extra body field key=value (repeatable)
      --source string            source image (host/repo:tag or @sha256:...); docker.io when no host
      --tag latest               target tag (default: source tag, latest for a digest)
      --username string          private source: username (not stored)
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

* [gyro container registry](gyro_container_registry.md)	 - Gyro image registry of the environment

