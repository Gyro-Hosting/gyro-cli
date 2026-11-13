## gyro container pipeline

List pipelines

```
gyro container pipeline [flags]
```

### Options

```
  -h, --help   help for pipeline
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
* [gyro container pipeline approve](gyro_container_pipeline_approve.md)	 - Approve the deployment
* [gyro container pipeline cancel](gyro_container_pipeline_cancel.md)	 - Stop a run
* [gyro container pipeline create](gyro_container_pipeline_create.md)	 - Create a pipeline
* [gyro container pipeline delete](gyro_container_pipeline_delete.md)	 - Delete a pipeline
* [gyro container pipeline get](gyro_container_pipeline_get.md)	 - Get a pipeline
* [gyro container pipeline list](gyro_container_pipeline_list.md)	 - List pipelines
* [gyro container pipeline logs](gyro_container_pipeline_logs.md)	 - Run log
* [gyro container pipeline redeploy](gyro_container_pipeline_redeploy.md)	 - Deploy this run's image again (rollback)
* [gyro container pipeline reject](gyro_container_pipeline_reject.md)	 - Reject the deployment
* [gyro container pipeline rotate-secret](gyro_container_pipeline_rotate-secret.md)	 - Rotate the webhook secret
* [gyro container pipeline run](gyro_container_pipeline_run.md)	 - Start a run
* [gyro container pipeline run-get](gyro_container_pipeline_run-get.md)	 - Get a run
* [gyro container pipeline runs](gyro_container_pipeline_runs.md)	 - Runs of a pipeline
* [gyro container pipeline secret](gyro_container_pipeline_secret.md)	 - Delete a build secret
* [gyro container pipeline test](gyro_container_pipeline_test.md)	 - Test repository access
* [gyro container pipeline update](gyro_container_pipeline_update.md)	 - Change a pipeline
* [gyro container pipeline webhook](gyro_container_pipeline_webhook.md)	 - Webhook address and secret

