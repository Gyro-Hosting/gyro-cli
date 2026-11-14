## gyro container pipeline update

Change a pipeline

### Synopsis

Change a pipeline

Send only the fields to change. `token`: new value, or "-" to remove.

API: PATCH /cloud/containers/{id}/pipelines/{pipeline} (operationId pipeline_update, scope services_write)

```
gyro container pipeline update <ID> <PIPELINE> [flags]
```

### Options

```
      --approval                                  Wait for approval before deploying (24 h)
      --auto                                      Webhook pushes start runs
      --auto-rollback                             Roll back when the new version is unhealthy
      --branch string                             Branches whose pushes build (comma separated globs, empty = none)
      --build-args string                         Dockerfile ARG values (visible in the image history) (JSON)
      --build-env commands                        Variables for commands (JSON)
      --build-image commands                      Image that runs commands (e.g. node:22-alpine)
      --build-target FROM node:20 AS production   Stage of a multi-stage Dockerfile to build (FROM node:20 AS production -> `production`, same as `docker build --target`). Empty = the last stage. Stage names are case-insensitive (stored lowercase). A stage the Dockerfile does not have fails the run with `build_target_not_found`. gyro-pipeline.yml: `build_target:` overrides it per run. CLI: `--target` (or `--build-target`).
      --commands strings                          Shell commands run in the repository root before the Dockerfile (tests, npm ci ...); build secrets are environment variables here
      --context string                            Build folder (empty = repository root)
      --cpu-m int                                 Build CPU (millicores, plan limit)
      --data string                               request body as JSON, @file or @- (stdin)
      --dockerfile string                         dockerfile
      --dry-run                                   show the price quote / request without changing anything
      --git-user string                           User name for the token (default: x-access-token / oauth2 / git)
  -h, --help                                      help for update
      --idempotency-key string                    idempotency key (automatic for money operations)
      --image string                              Image name in the environment registry (<registry>/<code>/<image>:<sha7> + :latest)
      --memory-mi int                             Build memory (MiB, plan limit)
      --name string                               Display name (1-48)
      --notify strings                            Notifications (channels: the environment's alert settings)
      --provider string                           provider (github|gitlab|gitea|git|archive)
      --repo string                               https clone URL (no credentials) or SSH address git@host:org/repo.git / ssh://user@host[:port]/path (needs a deploy key); empty for provider archive
      --repo-file                                 Read gyro-pipeline.yml from the repository (GitHub, GitLab, Gitea)
      --set stringArray                           extra body field key=value (repeatable)
      --tags string                               Tags that build (globs, e.g. v*)
      --targets strings                           Services that get the new image (empty = build only)
      --timeout-min int                           Build time limit (minutes, plan limit)
      --token string                              Read-only access token for a private repository (stored encrypted, never returned; "-" on update removes it)
      --verify-min int                            Health wait after deploy (1-30 minutes)
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

