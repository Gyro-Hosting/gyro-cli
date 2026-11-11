## gyro

Gyro Hosting command-line interface

### Synopsis

gyro manages your Gyro Hosting services from the terminal: cloud servers, SSH keys, invoices,
balance and everything else the Gyro Public API offers (API spec 2.11.0).

Get started:
  gyro configure                 # store an API key (Client area → account menu (top right) → API Credentials)
  gyro server list
  gyro server create --plan standart-2-4 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server ssh web-1

Money and destructive actions ask for confirmation (use --yes in scripts, --dry-run for a quote).

### Options

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
  -h, --help               help for gyro
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

* [gyro activity](gyro_activity.md)	 - Write log of your API keys
* [gyro address](gyro_address.md)	 - Billing addresses
* [gyro affiliate](gyro_affiliate.md)	 - Get affiliate
* [gyro api](gyro_api.md)	 - Call any API endpoint directly (like 'gh api')
* [gyro app](gyro_app.md)	 - One-click applications
* [gyro balance](gyro_balance.md)	 - Account balance (bare command) and transactions
* [gyro billing](gyro_billing.md)	 - Prepaid hourly usage (Istanbul cloud, containers): balance, usage, estimate, statement
* [gyro bulk](gyro_bulk.md)	 - Run one action on up to 50 servers
* [gyro completion](gyro_completion.md)	 - Generate a shell completion script
* [gyro configure](gyro_configure.md)	 - Store an API key and preferences in a named profile
* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains
* [gyro context](gyro_context.md)	 - Show the active profile, API URL and key (and check the key)
* [gyro dedicated](gyro_dedicated.md)	 - Dedicated servers
* [gyro dns](gyro_dns.md)	 - DNS records of your domains
* [gyro domain](gyro_domain.md)	 - Domains: search, register, transfer, nameservers, DNS, WHOIS
* [gyro firewall](gyro_firewall.md)	 - Network firewall policy (rules, default action, status)
* [gyro image](gyro_image.md)	 - Operating system images
* [gyro invoice](gyro_invoice.md)	 - Invoices
* [gyro key](gyro_key.md)	 - Sub-account API keys
* [gyro lb](gyro_lb.md)	 - List load balancers
* [gyro license](gyro_license.md)	 - Software licenses (cPanel, Plesk, LiteSpeed ...)
* [gyro location](gyro_location.md)	 - Locations
* [gyro mail](gyro_mail.md)	 - Corporate mail: mailboxes, aliases, forwarding ...
* [gyro me](gyro_me.md)	 - Get me
* [gyro monitor](gyro_monitor.md)	 - Uptime monitors
* [gyro order](gyro_order.md)	 - Orders (preview and create)
* [gyro ping](gyro_ping.md)	 - Ping
* [gyro plan](gyro_plan.md)	 - Server plans and prices (catalog)
* [gyro product](gyro_product.md)	 - Full product catalog (all product lines)
* [gyro reference](gyro_reference.md)	 - Countries, states and cities
* [gyro reseller](gyro_reseller.md)	 - Reseller: end-customers, assignments, keys, report
* [gyro s3](gyro_s3.md)	 - S3 object storage: buckets and access keys
* [gyro security-group](gyro_security-group.md)	 - List security groups
* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...
* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)
* [gyro sms](gyro_sms.md)	 - SMS messages
* [gyro ssh-key](gyro_ssh-key.md)	 - SSH public keys of your account
* [gyro ticket](gyro_ticket.md)	 - Support tickets
* [gyro user](gyro_user.md)	 - List sub-users
* [gyro vds](gyro_vds.md)	 - VDS servers
* [gyro version](gyro_version.md)	 - Print the version (and check for a newer release with --check)
* [gyro webhook](gyro_webhook.md)	 - Webhooks (signed event callbacks)
* [gyro whoami](gyro_whoami.md)	 - Show the API key in use and its permissions

