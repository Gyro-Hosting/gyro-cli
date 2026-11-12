## gyro server console

Open a one-time VNC console ticket

### Synopsis

Open a one-time VNC console ticket

Returns a websocket URL and password valid for about a minute. Never send an Idempotency-Key here (400). In tr-ist-1 the ticket is a WebMKS websocket (`protocol: webmks`, empty password: the URL is the one-time secret). tr-ist-1: the WebSocket URL is on konsol.gyrohosting.com (single-use ticket). 2.14.0 (Gyro Compute): `protocol: web`, `url` is a one-time browser page (90 s) with keyboard + screen; `wss_url` carries the same URL for older clients.

API: POST /cloud/servers/{id}/console (operationId openConsole, scope cloud_write)

```
gyro server console <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for console
      --idempotency-key string   idempotency key (automatic for money operations)
      --wait                     wait until the server is no longer busy
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

