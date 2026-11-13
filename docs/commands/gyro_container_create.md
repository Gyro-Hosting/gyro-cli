## gyro container create

Create container environments (batch)

### Synopsis

Create container environments (batch)

Orders 1..N environments of a plan in ONE order (`quantity`, default 1, max 25 per request). Prepaid: the order is 0 and usage is charged hourly from the balance; the balance must cover the first 24 hours of every environment and new accounts are limited in their first days. `name` (with `{n}`) and `tags` label every environment; `project_id` puts them in a project. Send `"dry_run": true` for a quote; a real order needs `"confirm": true` and an `Idempotency-Key` header.

API: POST /cloud/containers (operationId createContainerEnvironments, scope cloud_order)

```
gyro container create [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name (nickname); {n} = number in a batch
      --plan string              plan from GET /cloud/containers/plans (mini-1-1, small-1-2, starter, standart, pro, serverless)
      --product-id int           product id instead of plan
      --project-id int           project (0 = default project)
      --quantity int             environments to create in one order (--count), default 1, max 25
      --set stringArray          extra body field key=value (repeatable)
      --tags strings             tags (max 10)
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

