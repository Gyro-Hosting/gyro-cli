## gyro billing estimate

Hourly, daily and monthly cost of plans and serverless containers; how long your balance lasts

### Synopsis

Plans (Istanbul servers, container plans) are capped at their monthly price. Serverless containers (--cpu-m/--memory-mi, or --item <cpu_m>/<memory_mi>:<count>) are billed per running copy per hour without a cap; sizes: gyro container sizes.

```
gyro billing estimate [flags]
```

### Examples

```
  gyro billing estimate --plan standart-2-4 --region tr-ist-1 --count 3
  gyro billing estimate --item standart-2-4:2 --item standart-4-8:1
  gyro billing estimate --cpu-m 500 --memory-mi 1024 --count 2 --hours 730
  gyro billing estimate --item 250/512:3 --item standart-2-4:1
```

### Options

```
      --count int          how many (default 1)
      --cpu-m int          serverless container: CPU per copy in millicores (250, 500, 1000, 2000, 4000)
  -h, --help               help for estimate
      --hours int          serverless container: running hours for the total (default 730)
      --item stringArray   plan:count or serverless <cpu_m>/<memory_mi>:count, repeatable
      --memory-mi int      serverless container: memory per copy in MiB (see gyro container sizes)
      --plan string        plan name, e.g. standart-2-4
      --product-id int     or the product id
      --region string      region (prepaid: tr-ist-1) (default "tr-ist-1")
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

* [gyro billing](gyro_billing.md)	 - Prepaid hourly usage (Istanbul cloud, containers): balance, usage, estimate, statement

