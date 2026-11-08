## gyro server create

Order a new cloud server (paid from your balance)

### Synopsis

Order a cloud server. The price is shown first (a quote from the API) and you are asked to confirm;
the order is paid from your account balance, never from a card. An idempotency key is attached
automatically, so a retried command never orders twice.

--plan accepts a plan name (standart-2-4, pro-2-8 ...), a server type (cpx22, ccx13 ...) or a product id
(see 'gyro plan list'). --location accepts a code (fsn1) or a city (falkenstein).

```
gyro server create [HOSTNAME] [flags]
```

### Examples

```
  gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server create --plan cpx22 --location fsn1 --image debian-12 --ssh-key ~/.ssh/id_ed25519.pub --dry-run
  gyro server create db-1 --plan standart-2-4 --location nbg1 --image ubuntu-24.04 --user-data @cloud-init.yaml --yes -o json
```

### Options

```
      --app string               one-click app to install (needs API v2)
      --cycle string             monthly|quarterly|semiannually|annually (default "monthly")
      --data string              extra request body JSON / @file
      --dry-run                  only show the price quote
  -h, --help                     help for create
      --hostname string          host name (or give it as the argument)
      --idempotency-key string   idempotency key (default: random)
      --image string             OS image slug, e.g. ubuntu-24.04 (required)
      --location string          location code or city
      --plan string              plan name, server type or product id (required)
      --set stringArray          extra body field key=value (repeatable)
      --ssh-key strings          SSH key id, name or .pub file (uploaded if new)
      --user-data string         cloud-init user data, inline or @file
      --wait                     wait until the server is running
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...

