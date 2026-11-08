## gyro domain whois-profile list

Get whois profiles

### Synopsis

Get whois profiles

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik

API: GET /domains/whois-profiles (operationId coreGetWhoisProfiles, scope Domains/GetWhoisProfiles)

```
gyro domain whois-profile list [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro domain whois-profile](gyro_domain_whois-profile.md)	 - Get whois profiles

