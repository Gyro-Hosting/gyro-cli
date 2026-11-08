## gyro mail

Corporate mail: mailboxes, aliases, forwarding ...

```
gyro mail [flags]
```

### Options

```
  -h, --help   help for mail
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro mail alias](gyro_mail_alias.md)	 - Aliases (forwarders)
* [gyro mail archive](gyro_mail_archive.md)	 - GET archive (GyroMail)
* [gyro mail autoresponder](gyro_mail_autoresponder.md)	 - Auto-responders
* [gyro mail call](gyro_mail_call.md)	 - Call a module method of a service
* [gyro mail catch-all](gyro_mail_catch-all.md)	 - Catch-all address
* [gyro mail deleted](gyro_mail_deleted.md)	 - GET deleted-mail (GyroMail)
* [gyro mail delivery-log](gyro_mail_delivery-log.md)	 - GET delivery-log (GyroMail)
* [gyro mail device](gyro_mail_device.md)	 - Device setup / app passwords
* [gyro mail dns](gyro_mail_dns.md)	 - Required DNS records and their status
* [gyro mail list](gyro_mail_list.md)	 - List your services of this product line
* [gyro mail mailbox](gyro_mail_mailbox.md)	 - Mailboxes
* [gyro mail migration](gyro_mail_migration.md)	 - IMAP migration jobs
* [gyro mail plan](gyro_mail_plan.md)	 - GET plan-size (GyroMail)
* [gyro mail quarantine](gyro_mail_quarantine.md)	 - GET quarantine (GyroMail)
* [gyro mail spam](gyro_mail_spam.md)	 - Spam filter
* [gyro mail spam-filter](gyro_mail_spam-filter.md)	 - GET spam-filter (GyroMail)
* [gyro mail tool](gyro_mail_tool.md)	 - Read a tool's data or run one of its actions (see 'tools')
* [gyro mail tools](gyro_mail_tools.md)	 - Show the tools, actions and safety policy of a service
* [gyro mail two-factor](gyro_mail_two-factor.md)	 - GET two-factor (GyroMail)

