## gyro dns acme-challenge status

Is the challenge record live?

### Synopsis

Is the challenge record live?

Asks both authoritative name servers (ns1/ns2.gyrodns.net) directly. `ready` is true when both publish `value`. Poll every few seconds after adding (usually < 1 minute).

API: GET /cloud/dns/acme-challenge (operationId acmeChallengeStatus, scope dns_acme)

```
gyro dns acme-challenge status [flags]
```

### Options

```
  -h, --help           help for status
      --name string    _acme-challenge.<domain>
      --value string   Challenge value (key authorization digest)
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

* [gyro dns acme-challenge](gyro_dns_acme-challenge.md)	 - Add an ACME challenge TXT value

