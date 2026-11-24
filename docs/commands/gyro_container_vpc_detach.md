## gyro container vpc detach

Detach the environment from its VPC (the /28 source range is released)

### Synopsis

Connects a container environment to one subnet of your private network (VPC) in tr-ist-1. The environment
gets a fixed VPC source range (/28) inside the subnet: connections from your containers to the VPC servers come
from it and it does not change when nodes are replaced, scaled or renewed. The first address is the VPC endpoint:
VPC servers reach the services you publish (--expose service:port) at endpoint:port. Never reachable from the
internet. The subnet must be /26 or larger.
Price: node environments (Mini, Small, Starter, Standard, Pro) free. Serverless environments: a connector dedicated
to the environment runs, so the connection is charged (2.99 EUR/month, billed hourly from the prepaid balance while
attached); attach shows the price in your currency and asks for confirmation (--yes accepts it).

```
gyro container vpc detach <ENVIRONMENT_ID> [flags]
```

### Examples

```
  gyro container vpc detach 17000 --wait
  gyro container vpc detach 17000 --yes
```

### Options

```
      --dry-run                 show the request without changing anything
  -h, --help                    help for detach
      --wait                    wait until it is detached (state off)
      --wait-timeout duration   maximum time for --wait (default 10m0s)
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

* [gyro container vpc](gyro_container_vpc.md)	 - Private network (VPC) of a container environment: fixed source range, VPC endpoint

