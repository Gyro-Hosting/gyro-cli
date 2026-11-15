## gyro container vpc

Private network (VPC) of a node environment: fixed source range, VPC endpoint

### Synopsis

Connects a node container environment (Mini, Small, Starter, Standard, Pro) to one subnet of your private network
(VPC) in tr-ist-1. The environment gets a fixed VPC source range (/28) inside the subnet: connections from your
containers to the VPC servers come from it and it does not change when nodes are replaced, scaled or renewed.
The first address is the VPC endpoint: VPC servers reach the services you publish (--expose service:port) at
endpoint:port. Never reachable from the internet. Free. The subnet must be /26 or larger. Serverless: not yet.

```
gyro container vpc [flags]
```

### Options

```
  -h, --help   help for vpc
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
* [gyro container vpc attach](gyro_container_vpc_attach.md)	 - Attach the environment to a VPC subnet
* [gyro container vpc detach](gyro_container_vpc_detach.md)	 - Detach the environment from its VPC (the /28 source range is released)
* [gyro container vpc expose](gyro_container_vpc_expose.md)	 - Replace the services published at the VPC endpoint
* [gyro container vpc get](gyro_container_vpc_get.md)	 - Show the environment's VPC attachment (source range, endpoint, published services)

