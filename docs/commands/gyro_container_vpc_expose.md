## gyro container vpc expose

Replace the services published at the VPC endpoint

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
gyro container vpc expose <ENVIRONMENT_ID> --expose service:port ... | --clear [flags]
```

### Examples

```
  gyro container vpc expose 17000 --expose api:8080 --expose web:80
  gyro container vpc expose 17000 --clear
```

### Options

```
      --clear                remove every published service
      --dry-run              show the request without changing anything
      --expose stringArray   service:port (repeatable; the list replaces the old one)
  -h, --help                 help for expose
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

