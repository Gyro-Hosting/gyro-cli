## gyro container pipeline host-key scan

SSH server host key

### Synopsis

SSH server host key

Connects to the repository's SSH server (public addresses only) and returns the host key fingerprint. GitHub / GitLab / Bitbucket keys are pinned (`pinned: true`). Other servers: compare the fingerprint with your own record, then confirm it with PUT.

API: GET /cloud/containers/{id}/pipelines/{pipeline}/host-key (operationId pipeline_host_key_scan, scope services_read)

```
gyro container pipeline host-key scan <ID> <PIPELINE> [flags]
```

### Options

```
  -h, --help   help for scan
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

* [gyro container pipeline host-key](gyro_container_pipeline_host-key.md)	 - SSH server host key

