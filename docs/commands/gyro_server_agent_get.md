## gyro server agent get

Guest agent version

### Synopsis

Guest agent version

The Gyro guest agent configures private network cards and the public-IP-off route inside the server. Servers installed before 14 November 2026 run an older agent: a card attaches but `applied` stays false and turning the public IP off answers 409 `agent_update_required`. `update_command` is a one-line command to run on the server as root (downloads Gyro's versioned updater and checks its SHA-256); it only replaces the agent, needs no restart and does not touch your data. If you gave the card an address yourself (netplan), delete that file and run `netplan apply` afterwards. The version is read from the server when it is running.

API: GET /cloud/servers/{id}/agent (operationId getServerAgent, scope cloud_read)

```
gyro server agent get <ID> [flags]
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

* [gyro server agent](gyro_server_agent.md)	 - Guest agent version

