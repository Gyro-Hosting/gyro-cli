## gyro lb service add

Add a service

### Synopsis

Add a service

Adds a listener. HTTPS needs at least one certificate of this load balancer; TLS ends at the load balancer and targets get HTTP. Use TCP 443 for passthrough.

API: POST /cloud/load-balancers/{id}/services (operationId typed_post_load_balancers__id__services, scope services_write)

```
gyro lb service add <ID> [flags]
```

### Options

```
      --certificates string      HTTPS: comma separated certificate ids of this load balancer
      --cookie-lifetime int      cookie_lifetime
      --cookie-name string       cookie_name
      --data string              request body as JSON, @file or @- (stdin)
      --destination-port int     Port on the targets
      --dry-run                  show the price quote / request without changing anything
      --hc-domain string         HTTP check Host header
      --hc-interval int          3-60 s
      --hc-path string           HTTP check path (starts with /)
      --hc-port int              hc_port
      --hc-protocol string       Health check type (tcp|http)
      --hc-retries int           0-5
      --hc-status-codes string   Comma separated, e.g. 2??,3??
      --hc-timeout int           1-30 s, below interval
      --hc-tls                   HTTP check over TLS
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --idle-timeout int         Idle timeout (s)
      --listen-port int          Port the load balancer listens on
      --protocol string          protocol (tcp|http|https)
      --proxyprotocol            Send PROXY protocol header to targets
      --redirect-http            HTTPS: redirect port 80 to HTTPS (port 80 must be free)
      --set stringArray          extra body field key=value (repeatable)
      --sticky                   Sticky sessions (cookie)
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

* [gyro lb service](gyro_lb_service.md)	 - List services (ports) of a load balancer

