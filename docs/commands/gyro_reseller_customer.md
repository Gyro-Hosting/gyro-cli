## gyro reseller customer

List end-customers

```
gyro reseller customer [flags]
```

### Options

```
  -h, --help   help for customer
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

* [gyro reseller](gyro_reseller.md)	 - Reseller: end-customers, assignments, keys, report
* [gyro reseller customer create](gyro_reseller_customer_create.md)	 - Create an end-customer
* [gyro reseller customer delete](gyro_reseller_customer_delete.md)	 - Delete an end-customer (revokes keys and webhooks, unassigns services)
* [gyro reseller customer get](gyro_reseller_customer_get.md)	 - Get an end-customer
* [gyro reseller customer key](gyro_reseller_customer_key.md)	 - List an end-customer's keys
* [gyro reseller customer list](gyro_reseller_customer_list.md)	 - List end-customers
* [gyro reseller customer update](gyro_reseller_customer_update.md)	 - Update an end-customer (suspend = its keys stop working)

