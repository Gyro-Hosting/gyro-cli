## gyro server create

Order a new cloud server (paid from your balance)

### Synopsis

Order a cloud server. The price is shown first (a quote from the API) and you are asked to confirm;
the order is paid from your account balance, never from a card. An idempotency key is attached
automatically, so a retried command never orders twice.

--plan accepts a plan name (standart-2-4, pro-2-8 ...) or a product id (see 'gyro plan list').
--region (or --location) accepts a region (eu-de-1, tr-ist-1, see 'gyro location list') or a city (falkenstein, istanbul).
In tr-ist-1 (Istanbul) up to 10 --ssh-key values and --user-data go with cloud-init at creation (SSH in about 25 s).

```
gyro server create [HOSTNAME] [flags]
```

### Examples

```
  gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server create --plan standart-2-4 --location eu-de-1 --image debian-12 --ssh-key ~/.ssh/id_ed25519.pub --dry-run
  gyro server create db-1 --plan standart-2-4 --location eu-de-2 --image ubuntu-24.04 --user-data @cloud-init.yaml --yes -o json
  gyro server create ist-1 --plan standart-2-4 --region tr-ist-1 --image ubuntu-24.04 --ssh-key mykey --wait
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
      --location string          region (eu-de-1) or city
      --plan string              plan name or product id (required)
      --region string            region (eu-de-1, tr-ist-1); same as --location
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

