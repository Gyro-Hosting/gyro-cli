## gyro container run

Run a one-off job (migration, seed) with a service's image, variables and secrets

### Synopsis

Runs a command once with the image (pinned to the digest of the running version), variables and secrets of an
existing service, then stops. Everything after "--" is the command; it replaces the image ENTRYPOINT and is passed as
a list (no shell: write sh -c "a && b" for shell features). Without a command the image's own command runs.
--follow streams the output; --wait waits silently. With either, the CLI exits with the job's exit code
(124 = time limit reached, 130 = stopped). Free on node plans (uses your nodes); on Serverless the job is billed per
vCPU-second / GB-second like a task (billing line "<service>-job"). Parallel jobs per environment are limited
(2-3 by plan); see "gyro limits".

```
gyro container run <ENVIRONMENT_ID> --service <NAME> [flags] [-- COMMAND [ARG...]]
```

### Examples

```
  gyro container run 17000 --service api --follow -- npm run migrate
  gyro container run 17000 --service web --wait --time-limit 10m -- sh -c "php artisan migrate --force && php artisan db:seed --force"
  gyro container run 17000 --service api --cpu 1 --memory 1024 -- python manage.py migrate
```

### Options

```
      --cpu string            vCPU for the job (default: same as the service)
      --follow                stream the output until the job ends (implies --wait)
  -h, --help                  help for run
      --memory int            memory in MB for the job (default: same as the service)
      --service string        service whose image, variables and secrets are used (required)
      --time-limit duration   stop the job after this time (10s-6h, default 30m)
      --wait                  wait until the job ends; exit with its exit code
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

