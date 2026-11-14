## gyro container get

Container environment detail (namespace, internal addressing, egress IPs)

### Synopsis

Container environment detail (namespace, internal addressing, egress IPs)

Environment code, Kubernetes **namespace**, in-cluster addressing (`cluster_dns`), plan limits, nodes, quota, ingress IPs and the current **egress IPs**. Egress IPs of a **dedicated** environment = the public IPs of its own nodes + the shared pool node(s) used while an update rolls out (the new copy starts there for about a minute, then moves back to your node). **Serverless** = the shared sandboxed pool nodes (the pool can grow: a new address can appear). Build machines of Gyro Pipeline are listed separately (`build_ips`). The list changes when the nodes change (plan change, node replacement, pool scaling); the owner is notified (Alerts channels, event `egress`). No IPv6 egress; outbound TCP 25 is closed. A fixed egress address is not offered yet.

API: GET /cloud/containers/{id} (operationId getContainerEnvironment, scope services_read)

```
gyro container get <ID> [flags]
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains

