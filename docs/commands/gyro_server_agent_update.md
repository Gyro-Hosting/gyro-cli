## gyro server agent update

Update the guest agent

### Synopsis

Update the guest agent

Starts the automatic update through the old agent's fixed Gyro download channel (agent 3, or unknown version which is checked first); the server keeps running, about 30 s, then the private network settings are applied again. Refusals (409): `manual_update_required` (agent 1-2: run `update_command`), `already_latest`, `not_running`, `not_supported`, `busy`. The task fails when the server cannot reach the internet; then run `update_command` from the web console. The Gyro guest agent configures private network cards and the public-IP-off route inside the server. Servers installed before 15 November 2026 run an older agent: a card attaches but `applied` stays false and turning the public IP off answers 409 `agent_update_required`. `update_command` is a one-line command to run on the server as root (downloads Gyro's versioned updater and checks its SHA-256); it only replaces the agent, needs no restart and does not touch your data. If you gave the card an address yourself (netplan), delete that file and run `netplan apply` afterwards.

API: POST /cloud/servers/{id}/agent/update (operationId updateServerAgent, scope cloud_write)

```
gyro server agent update <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro server agent](gyro_server_agent.md)	 - Guest agent version

