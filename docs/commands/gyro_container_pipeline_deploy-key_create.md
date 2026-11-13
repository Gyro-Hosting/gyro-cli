## gyro container pipeline deploy-key create

Create the SSH deploy key

### Synopsis

Create the SSH deploy key

Creates an ed25519 key pair for an SSH repository and returns only the public key (add it to the repository as a read-only deploy key). The private key is stored encrypted and never returned. A pipeline has one key; `"rotate": true` replaces it (the old key stops working).

API: PUT /cloud/containers/{id}/pipelines/{pipeline}/deploy-key (operationId pipeline_deploy_key_create, scope services_write)

```
gyro container pipeline deploy-key create <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --rotate                   rotate
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro container pipeline deploy-key](gyro_container_pipeline_deploy-key.md)	 - Create the SSH deploy key

