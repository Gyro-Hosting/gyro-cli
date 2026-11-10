# Gyro CLI — command-line interface

gyro: manage Gyro Hosting services from your terminal.

## Install

(HTML: cli.html#install)

## Getting started

(HTML: cli.html#start)

## Sample session

(HTML: cli.html#session)

## Profiles and credentials

(HTML: cli.html#profiles)

## Output and scripting

(HTML: cli.html#output)

## Safety: confirmation, dry run, idempotency

(HTML: cli.html#safety)

## Every product

(HTML: cli.html#lines)

## Command reference

| Command | Description | API |
|---|---|---|
| `gyro activity list` | List write calls made with your API keys | `GET /cloud/activity` |
| `gyro address create` | Create address | `POST /addresses` |
| `gyro address default` | Set default address (path: {id}) | `POST /addresses/{id}/default` |
| `gyro address delete` | Delete address (path: {id}) | `DELETE /addresses/{id}` |
| `gyro address get` | Get address (path: {id}) | `GET /addresses/{id}` |
| `gyro address list` | List addresses | `GET /addresses` |
| `gyro address set` | Update address (path: {id}) | `PUT /addresses/{id}` |
| `gyro affiliate commission list` | List affiliate commissions | `GET /affiliate/commissions` |
| `gyro affiliate enroll` | Enroll affiliate | `POST /affiliate/enroll` |
| `gyro affiliate get` | Get affiliate | `GET /affiliate` |
| `gyro affiliate payment-method set` | Save affiliate payout method | `PUT /affiliate/payment-method` |
| `gyro affiliate withdraw` | Withdraw affiliate | `POST /affiliate/withdraw` |
| `gyro api request DELETE /cloud/mail/{id}/aliases/{aid}` | DELETE aliases:delete (GyroMail) | `DELETE /cloud/mail/{id}/aliases/{aid}` |
| `gyro api request DELETE /cloud/mail/{id}/autoresponders/{local}` | DELETE autoresponder:delete (GyroMail) | `DELETE /cloud/mail/{id}/autoresponders/{local}` |
| `gyro api request DELETE /cloud/mail/{id}/mailboxes/{local}` | DELETE mailboxes:delete (GyroMail) | `DELETE /cloud/mail/{id}/mailboxes/{local}` |
| `gyro api request DELETE /cloud/mail/{id}/migrations/{jid}` | DELETE sync-jobs:delete (GyroMail) | `DELETE /cloud/mail/{id}/migrations/{jid}` |
| `gyro api request DELETE /cloud/storage/{id}/buckets/{bucket}` | DELETE buckets:delete (GyroDepolama) | `DELETE /cloud/storage/{id}/buckets/{bucket}` |
| `gyro api request DELETE /cloud/storage/{id}/keys/{key}` | DELETE access-keys:delete (GyroDepolama) | `DELETE /cloud/storage/{id}/keys/{key}` |
| `gyro api request GET /cloud/apps` | List installable apps | `GET /cloud/apps` |
| `gyro api request GET /cloud/mail/{id}/aliases` | GET aliases (GyroMail) | `GET /cloud/mail/{id}/aliases` |
| `gyro api request GET /cloud/mail/{id}/autoresponders` | GET autoresponder (GyroMail) | `GET /cloud/mail/{id}/autoresponders` |
| `gyro api request GET /cloud/mail/{id}/catch-all` | GET catch-all (GyroMail) | `GET /cloud/mail/{id}/catch-all` |
| `gyro api request GET /cloud/mail/{id}/dns` | GET dns-setup (GyroMail) | `GET /cloud/mail/{id}/dns` |
| `gyro api request GET /cloud/mail/{id}/mailboxes` | GET mailboxes (GyroMail) | `GET /cloud/mail/{id}/mailboxes` |
| `gyro api request GET /cloud/mail/{id}/migrations` | GET sync-jobs (GyroMail) | `GET /cloud/mail/{id}/migrations` |
| `gyro api request GET /cloud/storage/{id}/buckets` | GET buckets (GyroDepolama) | `GET /cloud/storage/{id}/buckets` |
| `gyro api request GET /cloud/storage/{id}/keys` | GET access-keys (GyroDepolama) | `GET /cloud/storage/{id}/keys` |
| `gyro api request PATCH /cloud/mail/{id}/aliases/{aid}` | PATCH aliases:edit (GyroMail) | `PATCH /cloud/mail/{id}/aliases/{aid}` |
| `gyro api request PATCH /cloud/mail/{id}/mailboxes/{local}` | PATCH mailboxes:edit (GyroMail) | `PATCH /cloud/mail/{id}/mailboxes/{local}` |
| `gyro api request POST /cloud/licenses/{id}/change-ip` | POST method change_ip (LisansAPP) | `POST /cloud/licenses/{id}/change-ip` |
| `gyro api request POST /cloud/mail/{id}/aliases` | POST aliases:create (GyroMail) | `POST /cloud/mail/{id}/aliases` |
| `gyro api request POST /cloud/mail/{id}/mailboxes` | POST mailboxes:create (GyroMail) | `POST /cloud/mail/{id}/mailboxes` |
| `gyro api request POST /cloud/mail/{id}/migrations` | POST sync-jobs:create (GyroMail) | `POST /cloud/mail/{id}/migrations` |
| `gyro api request POST /cloud/storage/{id}/buckets` | POST buckets:create (GyroDepolama) | `POST /cloud/storage/{id}/buckets` |
| `gyro api request POST /cloud/storage/{id}/keys` | POST access-keys:create (GyroDepolama) | `POST /cloud/storage/{id}/keys` |
| `gyro balance` | Account balance | `GET /balance` |
| `gyro balance setting update` | Update balance settings | `PATCH /balance/settings` |
| `gyro balance top-up` | Create a balance top-up invoice and get its payment link | `POST /cloud/balance/top-up` |
| `gyro balance transaction list` | Balance ledger | `GET /balance/transactions` |
| `gyro bulk` | Run one action on up to 50 servers | `POST /cloud/bulk` |
| `gyro core invoice pay` | Pay invoice (path: {id}) | `POST /invoices/{id}/pay` |
| `gyro core product list` | List products | `GET /products` |
| `gyro core service get` | Get service (path: {id}) | `GET /services/{id}` |
| `gyro core service list` | List services | `GET /services` |
| `gyro core service tool get` | Get service tool data (path: {id}, {tool}) | `GET /services/{id}/tools/{tool}` |
| `gyro core service tool list` | Get service tools (path: {id}) | `GET /services/{id}/tools` |
| `gyro core service tool run` | Run service tool action (path: {id}, {tool}, {action}) | `POST /services/{id}/tools/{tool}/{action}` |
| `gyro dns dnssec disable` | Disable DNSSEC | `DELETE /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns dnssec enable` | Enable DNSSEC | `POST /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns dnssec status` | DNSSEC status and DS records | `GET /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns health create` | Create a health-checked record | `POST /cloud/dns/{id}/health-checks` |
| `gyro dns health delete` | Delete a health check | `DELETE /cloud/dns/{id}/health-checks/{hid}` |
| `gyro dns health list` | List health-checked records | `GET /cloud/dns/{id}/health-checks` |
| `gyro dns rrset delete` | Delete a record set | `DELETE /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns rrset list` | List record sets | `GET /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns rrset set` | Create or replace a record set | `PUT /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns zone create` | Add a zone | `POST /cloud/dns/{id}/zones` |
| `gyro dns zone delete` | Delete a zone | `DELETE /cloud/dns/{id}/zones/{zone}` |
| `gyro dns zone export` | Export the zone file | `GET /cloud/dns/{id}/zones/{zone}/export` |
| `gyro dns zone get` | Get a zone with its records | `GET /cloud/dns/{id}/zones/{zone}` |
| `gyro dns zone import` | Import a BIND zone file | `POST /cloud/dns/{id}/zones/{zone}/import` |
| `gyro dns zone list` | List DNS zones | `GET /cloud/dns/{id}/zones` |
| `gyro dns zone verify` | Verify ownership now | `POST /cloud/dns/{id}/zones/{zone}/verify` |
| `gyro domain addon buy` | Buy domain addon (path: {domain}, {key}) | `POST /domains/{domain}/addons/{key}` |
| `gyro domain auth-code send` | Send auth code (path: {domain}) | `POST /domains/{domain}/auth-code` |
| `gyro domain auth-code set` | Save auth code (path: {domain}) | `PUT /domains/{domain}/auth-code` |
| `gyro domain check` | Check domain | `POST /domains/check` |
| `gyro domain child-nameserver create` | Create child nameserver (path: {domain}) | `POST /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver delete` | Delete child nameserver (path: {domain}) | `DELETE /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver list` | Get child nameservers (path: {domain}) | `GET /domains/{domain}/child-nameservers` |
| `gyro domain default-nameserver list` | Get default nameservers | `GET /domains/default-nameservers` |
| `gyro domain default-nameserver set` | Update default nameservers | `PUT /domains/default-nameservers` |
| `gyro domain dns-record create` | Create dns record (path: {domain}) | `POST /domains/{domain}/dns-records` |
| `gyro domain dns-record delete` | Delete dns record (path: {domain}) | `DELETE /domains/{domain}/dns-records` |
| `gyro domain dns-record list` | List dns records (path: {domain}) | `GET /domains/{domain}/dns-records` |
| `gyro domain dns-record set` | Update dns record (path: {domain}) | `PUT /domains/{domain}/dns-records` |
| `gyro domain dnssec create` | Create dnssec record (path: {domain}) | `POST /domains/{domain}/dnssec` |
| `gyro domain dnssec delete` | Delete dnssec record (path: {domain}) | `DELETE /domains/{domain}/dnssec` |
| `gyro domain dnssec list` | List dnssec records (path: {domain}) | `GET /domains/{domain}/dnssec` |
| `gyro domain email-forward create` | Create email forward (path: {domain}) | `POST /domains/{domain}/email-forwards` |
| `gyro domain email-forward delete` | Delete email forward (path: {domain}) | `DELETE /domains/{domain}/email-forwards` |
| `gyro domain email-forward list` | List email forwards (path: {domain}) | `GET /domains/{domain}/email-forwards` |
| `gyro domain forwarding delete` | Delete url forwarding (path: {domain}) | `DELETE /domains/{domain}/forwarding` |
| `gyro domain forwarding get` | Get url forwarding (path: {domain}) | `GET /domains/{domain}/forwarding` |
| `gyro domain forwarding set` | Update url forwarding (path: {domain}) | `PUT /domains/{domain}/forwarding` |
| `gyro domain get` | Get domain (path: {domain}) | `GET /domains/{domain}` |
| `gyro domain list` | List domains | `GET /domains` |
| `gyro domain nameserver list` | Get nameservers (path: {domain}) | `GET /domains/{domain}/nameservers` |
| `gyro domain nameserver set` | Update nameservers (path: {domain}) | `PUT /domains/{domain}/nameservers` |
| `gyro domain register` | Register domain | `POST /domains/register` |
| `gyro domain renew` | Renew domain (path: {domain}) | `POST /domains/{domain}/renew` |
| `gyro domain tld list` | List tlds | `GET /domains/tlds` |
| `gyro domain transfer` | Transfer domain | `POST /domains/transfer` |
| `gyro domain transfer-lock get` | Get transfer lock (path: {domain}) | `GET /domains/{domain}/transfer-lock` |
| `gyro domain transfer-lock set` | Update transfer lock (path: {domain}) | `PUT /domains/{domain}/transfer-lock` |
| `gyro domain transfer-status get` | Get transfer status (path: {domain}) | `GET /domains/{domain}/transfer-status` |
| `gyro domain update` | Update domain (path: {domain}) | `PATCH /domains/{domain}` |
| `gyro domain whois apply` | Apply whois profile (path: {domain}) | `POST /domains/{domain}/whois/apply` |
| `gyro domain whois get` | Get whois contacts (path: {domain}) | `GET /domains/{domain}/whois` |
| `gyro domain whois set` | Update whois contacts (path: {domain}) | `PUT /domains/{domain}/whois` |
| `gyro domain whois-privacy set` | Update whois privacy (path: {domain}) | `PUT /domains/{domain}/whois-privacy` |
| `gyro domain whois-profile create` | Create whois profile | `POST /domains/whois-profiles` |
| `gyro domain whois-profile default` | Set default whois profile (path: {pid}) | `POST /domains/whois-profiles/{pid}/default` |
| `gyro domain whois-profile delete` | Delete whois profile (path: {pid}) | `DELETE /domains/whois-profiles/{pid}` |
| `gyro domain whois-profile list` | List whois profiles | `GET /domains/whois-profiles` |
| `gyro domain whois-profile set` | Update whois profile (path: {pid}) | `PUT /domains/whois-profiles/{pid}` |
| `gyro image list` | List operating system images | `GET /cloud/images` |
| `gyro invoice coupon` | Apply invoice coupon (path: {id}) | `POST /invoices/{id}/coupon` |
| `gyro invoice get` | Get an invoice | `GET /invoices/{id}` |
| `gyro invoice list` | List invoices | `GET /invoices` |
| `gyro invoice pay` | Pay an invoice from the account balance | `POST /cloud/invoices/{id}/pay` |
| `gyro invoice pdf get` | Invoice PDF (base64 in JSON) | `GET /cloud/invoices/{id}/pdf` |
| `gyro key create` | Create a sub-account key (token returned once) | `POST /cloud/keys` |
| `gyro key delete` | Delete a sub-account key | `DELETE /cloud/keys/{id}` |
| `gyro key list` | List sub-account keys (never shows tokens) | `GET /cloud/keys` |
| `gyro lb algorithm set` | Change the balancing algorithm | `PUT /cloud/load-balancers/{id}/algorithm` |
| `gyro lb cert cancel` | Cancel a pending Let's Encrypt request | `DELETE /cloud/load-balancers/{id}/certificates/letsencrypt` |
| `gyro lb cert delete` | Delete a certificate | `DELETE /cloud/load-balancers/{id}/certificates/{cid}` |
| `gyro lb cert letsencrypt` | Request a free Let's Encrypt certificate | `POST /cloud/load-balancers/{id}/certificates/letsencrypt` |
| `gyro lb cert list` | List certificates | `GET /cloud/load-balancers/{id}/certificates` |
| `gyro lb cert upload` | Upload a certificate | `POST /cloud/load-balancers/{id}/certificates` |
| `gyro lb cert verify` | Continue a pending Let's Encrypt request | `POST /cloud/load-balancers/{id}/certificates/letsencrypt/verify` |
| `gyro lb create` | Order a load balancer | `POST /cloud/load-balancers` |
| `gyro lb delete` | Delete (cancel) a load balancer | `DELETE /cloud/load-balancers/{id}` |
| `gyro lb get` | Get a load balancer | `GET /cloud/load-balancers/{id}` |
| `gyro lb list` | List load balancers | `GET /cloud/load-balancers` |
| `gyro lb metrics` | Load balancer metrics | `GET /cloud/load-balancers/{id}/metrics` |
| `gyro lb network attach` | Attach to a private network | `POST /cloud/load-balancers/{id}/network` |
| `gyro lb network detach` | Detach from a private network | `DELETE /cloud/load-balancers/{id}/network/{nid}` |
| `gyro lb rdns set` | Set reverse DNS | `PUT /cloud/load-balancers/{id}/rdns` |
| `gyro lb service add` | Add a service | `POST /cloud/load-balancers/{id}/services` |
| `gyro lb service delete` | Delete a service | `DELETE /cloud/load-balancers/{id}/services/{port}` |
| `gyro lb service list` | List services (ports) of a load balancer | `GET /cloud/load-balancers/{id}/services` |
| `gyro lb service update` | Update a service | `PATCH /cloud/load-balancers/{id}/services/{port}` |
| `gyro lb settings` | Algorithm, reverse DNS, private network | `GET /cloud/load-balancers/{id}/settings` |
| `gyro lb target add` | Add a target | `POST /cloud/load-balancers/{id}/targets` |
| `gyro lb target list` | List targets with health | `GET /cloud/load-balancers/{id}/targets` |
| `gyro lb target remove` | Remove a target | `DELETE /cloud/load-balancers/{id}/targets` |
| `gyro lb types` | List load balancer sizes and prices | `GET /cloud/load-balancer-types` |
| `gyro license get` | GET method license_status (LisansAPP) | `GET /cloud/licenses/{id}` |
| `gyro location list` | List locations with orderable products | `GET /cloud/locations` |
| `gyro mail archive get` | GET archive (GyroMail) | `GET /cloud/mail/{id}/archive` |
| `gyro mail autoresponder set` | PUT autoresponder:edit (GyroMail) | `PUT /cloud/mail/{id}/autoresponders/{local}` |
| `gyro mail catch-all set` | PUT catch-all:edit (GyroMail) | `PUT /cloud/mail/{id}/catch-all` |
| `gyro mail deleted get` | GET deleted-mail (GyroMail) | `GET /cloud/mail/{id}/deleted` |
| `gyro mail delivery-log get` | GET delivery-log (GyroMail) | `GET /cloud/mail/{id}/delivery-log` |
| `gyro mail dns verify` | POST dns-setup:verify (GyroMail) | `POST /cloud/mail/{id}/dns/verify` |
| `gyro mail mailbox password` | POST mailboxes:password (GyroMail) | `POST /cloud/mail/{id}/mailboxes/{local}/password` |
| `gyro mail migration run` | POST sync-jobs:run (GyroMail) | `POST /cloud/mail/{id}/migrations/{jid}/run` |
| `gyro mail plan get` | GET plan-size (GyroMail) | `GET /cloud/mail/{id}/plan` |
| `gyro mail quarantine get` | GET quarantine (GyroMail) | `GET /cloud/mail/{id}/quarantine` |
| `gyro mail spam-filter get` | GET spam-filter (GyroMail) | `GET /cloud/mail/{id}/spam-filter` |
| `gyro mail two-factor get` | GET two-factor (GyroMail) | `GET /cloud/mail/{id}/two-factor` |
| `gyro me get` | Get me | `GET /me` |
| `gyro me notification list` | Get notifications | `GET /me/notifications` |
| `gyro me notification set` | Update notifications | `PUT /me/notifications` |
| `gyro me security get` | Get security | `GET /me/security` |
| `gyro me update` | Update me | `PATCH /me` |
| `gyro monitor channel create` | POST channels:create (GyroMonitor) | `POST /cloud/monitors/{id}/channels` |
| `gyro monitor channel delete` | DELETE channels:delete (GyroMonitor) | `DELETE /cloud/monitors/{id}/channels/{cid}` |
| `gyro monitor channel list` | GET channels (GyroMonitor) | `GET /cloud/monitors/{id}/channels` |
| `gyro monitor heartbeat create` | POST heartbeats:create (GyroMonitor) | `POST /cloud/monitors/{id}/heartbeats` |
| `gyro monitor heartbeat get` | GET heartbeats (GyroMonitor) | `GET /cloud/monitors/{id}/heartbeats` |
| `gyro monitor monitor create` | POST monitors:create (GyroMonitor) | `POST /cloud/monitors/{id}/monitors` |
| `gyro monitor monitor delete` | DELETE monitors:delete (GyroMonitor) | `DELETE /cloud/monitors/{id}/monitors/{mid}` |
| `gyro monitor monitor list` | GET monitors (GyroMonitor) | `GET /cloud/monitors/{id}/monitors` |
| `gyro monitor monitor pause` | POST monitors:pause (GyroMonitor) | `POST /cloud/monitors/{id}/monitors/{mid}/pause` |
| `gyro monitor monitor resume` | POST monitors:resume (GyroMonitor) | `POST /cloud/monitors/{id}/monitors/{mid}/resume` |
| `gyro monitor monitor update` | PATCH monitors:edit (GyroMonitor) | `PATCH /cloud/monitors/{id}/monitors/{mid}` |
| `gyro monitor status-page get` | GET status-pages (GyroMonitor) | `GET /cloud/monitors/{id}/status-pages` |
| `gyro order create` | Create order | `POST /orders` |
| `gyro order get` | Get order (path: {id}) | `GET /orders/{id}` |
| `gyro order list` | List orders | `GET /orders` |
| `gyro order preview` | Preview order | `POST /orders/preview` |
| `gyro ping get` | Ping | `GET /ping` |
| `gyro plan list` | List orderable cloud server products | `GET /cloud/products` |
| `gyro product category list` | List product categories | `GET /products/categories` |
| `gyro product get` | Get product (path: {id}) | `GET /products/{id}` |
| `gyro reference country list` | List countries | `GET /reference/countries` |
| `gyro reference country state list` | List states (path: {code}) | `GET /reference/countries/{code}/states` |
| `gyro reference state city list` | List cities (path: {id}) | `GET /reference/states/{id}/cities` |
| `gyro reseller assign` | Assign a service to an end-customer (customer_id 0 = unassign) | `POST /cloud/reseller/assign` |
| `gyro reseller customer create` | Create an end-customer | `POST /cloud/reseller/customers` |
| `gyro reseller customer delete` | Delete an end-customer (revokes keys and webhooks, unassigns services) | `DELETE /cloud/reseller/customers/{id}` |
| `gyro reseller customer get` | Get an end-customer | `GET /cloud/reseller/customers/{id}` |
| `gyro reseller customer key create` | Mint a key for an end-customer (token returned once) | `POST /cloud/reseller/customers/{id}/keys` |
| `gyro reseller customer key delete` | Revoke an end-customer key | `DELETE /cloud/reseller/customers/{id}/keys/{kid}` |
| `gyro reseller customer key list` | List an end-customer's keys | `GET /cloud/reseller/customers/{id}/keys` |
| `gyro reseller customer list` | List end-customers | `GET /cloud/reseller/customers` |
| `gyro reseller customer update` | Update an end-customer (suspend = its keys stop working) | `PATCH /cloud/reseller/customers/{id}` |
| `gyro reseller report get` | Usage and billing report per end-customer | `GET /cloud/reseller/report` |
| `gyro reseller setting get` | Reseller settings | `GET /cloud/reseller/settings` |
| `gyro reseller setting set` | Save reseller settings (white label) | `PUT /cloud/reseller/settings` |
| `gyro server activity list` | Recent platform actions on the server | `GET /cloud/servers/{id}/activity` |
| `gyro server app create` | Install an app (rebuilds the disk; confirm required; 3 per server per day) | `POST /cloud/servers/{id}/app` |
| `gyro server app get` | App install status (credentials shown once when ready) | `GET /cloud/servers/{id}/app` |
| `gyro server backup create` | Take a manual backup (backup add-on required) | `POST /cloud/servers/{id}/backups` |
| `gyro server backup list` | List automatic backups | `GET /cloud/servers/{id}/backups` |
| `gyro server console` | Open a one-time VNC console ticket | `POST /cloud/servers/{id}/console` |
| `gyro server create` | Create (order) a cloud server | `POST /cloud/servers` |
| `gyro server delete` | Delete (cancel) a server | `DELETE /cloud/servers/{id}` |
| `gyro server firewall list` | Inbound firewall rules | `GET /cloud/servers/{id}/firewall` |
| `gyro server firewall rule add` | Add an inbound firewall rule | `POST /cloud/servers/{id}/firewall/rules` |
| `gyro server firewall rule delete` | Remove a firewall rule (the last one needs confirm) | `DELETE /cloud/servers/{id}/firewall/rules/{n}` |
| `gyro server get` | Get a server with its live state | `GET /cloud/servers/{id}` |
| `gyro server list` | List cloud servers | `GET /cloud/servers` |
| `gyro server metrics` | CPU, disk and network time series | `GET /cloud/servers/{id}/metrics` |
| `gyro server power off` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power on` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power reboot` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power reset` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power shutdown` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rdns list` | List reverse DNS entries | `GET /cloud/servers/{id}/rdns` |
| `gyro server rdns set` | Set (or reset) the PTR record of one address | `PUT /cloud/servers/{id}/rdns` |
| `gyro server rebuild` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rescue disable` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rescue enable` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server reset-password` | Run a server action | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server snapshot create` | Take a snapshot | `POST /cloud/servers/{id}/snapshots` |
| `gyro server snapshot delete` | Delete a snapshot | `DELETE /cloud/servers/{id}/snapshots/{sid}` |
| `gyro server snapshot list` | List snapshots and the snapshot quota | `GET /cloud/servers/{id}/snapshots` |
| `gyro server snapshot restore` | Restore a snapshot onto the server (overwrites the disk) | `POST /cloud/servers/{id}/snapshots/{sid}/restore` |
| `gyro server update | rename` | Rename a server or switch auto-renew | `PATCH /cloud/servers/{id}` |
| `gyro service addon available get` | Add-ons offered for this service (with options and prices) | `GET /cloud/services/{id}/addons/available` |
| `gyro service addon list` | List service addons (path: {id}) | `GET /services/{id}/addons` |
| `gyro service addons` | Buy an add-on (paid from balance) | `POST /cloud/services/{id}/addons` |
| `gyro service cancellation create` | Create cancellation (path: {id}) | `POST /services/{id}/cancellation` |
| `gyro service cancellation delete` | Revoke cancellation (path: {id}) | `DELETE /services/{id}/cancellation` |
| `gyro service dashboard get` | Get service dashboard (path: {id}) | `GET /services/{id}/dashboard` |
| `gyro service get` | Get a service (renewal, cancellation, access details) | `GET /cloud/services/{id}` |
| `gyro service invoice list` | List service invoices (path: {id}) | `GET /services/{id}/invoices` |
| `gyro service label get` | Get the nickname and tags of a service | `GET /cloud/services/{id}/label` |
| `gyro service label set` | Set the nickname and/or tags of a service | `PUT /cloud/services/{id}/label` |
| `gyro service list` | List services of every product line | `GET /cloud/services` |
| `gyro service method call` | Call a client-callable module method (e.g. license change_ip) | `POST /cloud/services/{id}/methods/{method}` |
| `gyro service metrics billing list` | List service metric billing (path: {id}) | `GET /services/{id}/metrics/billing` |
| `gyro service metrics list` | Get service metrics (path: {id}) | `GET /services/{id}/metrics` |
| `gyro service metrics update` | Update service metric (path: {id}, {key}) | `PATCH /services/{id}/metrics/{key}` |
| `gyro service module-method` | Use service module method (path: {id}) | `POST /services/{id}/module-method` |
| `gyro service password` | Change service password (path: {id}) | `POST /services/{id}/password` |
| `gyro service renew` | Renew service (path: {id}) | `POST /services/{id}/renew` |
| `gyro service resize` | Upgrade the plan (paid from balance) | `POST /cloud/services/{id}/resize` |
| `gyro service resize-option list` | Products this service can be upgraded to, with prices | `GET /cloud/services/{id}/resize-options` |
| `gyro service sso` | Get service sso (path: {id}) | `POST /services/{id}/sso` |
| `gyro service tool get` | Read a tool (list mailboxes, snapshots, DNS records, ...) | `GET /cloud/services/{id}/tools/{tool}` |
| `gyro service tool list` | Discover tools, actions and safety policy of a service | `GET /cloud/services/{id}/tools` |
| `gyro service tool run` | Run a tool action (create a mailbox, restore a snapshot, ...) | `POST /cloud/services/{id}/tools/{tool}/{action}` |
| `gyro service update` | Update service (path: {id}) | `PATCH /services/{id}` |
| `gyro sms message get` | Get sms message (path: {id}) | `GET /sms/messages/{id}` |
| `gyro sms message list` | List sms messages | `GET /sms/messages` |
| `gyro sms quote` | Quote sms | `POST /sms/quote` |
| `gyro sms send` | Send sms | `POST /sms/send` |
| `gyro sms sender list` | List sms senders | `GET /sms/senders` |
| `gyro ssh-key create` | Add an SSH key | `POST /cloud/ssh-keys` |
| `gyro ssh-key delete` | Delete an SSH key | `DELETE /cloud/ssh-keys/{id}` |
| `gyro ssh-key get` | Get an SSH key | `GET /cloud/ssh-keys/{id}` |
| `gyro ssh-key list` | List SSH keys | `GET /cloud/ssh-keys` |
| `gyro storage bucket rename` | Rename a bucket | `POST /cloud/storage/{id}/buckets/{bucket}/rename` |
| `gyro storage connect get` | GET connect (GyroDepolama) | `GET /cloud/storage/{id}/connect` |
| `gyro storage folder count` | Count objects in a folder | `GET /cloud/storage/{id}/buckets/{bucket}/folders/count` |
| `gyro storage folder create` | Create a folder | `POST /cloud/storage/{id}/buckets/{bucket}/folders` |
| `gyro storage folder delete` | Delete a folder with its contents | `DELETE /cloud/storage/{id}/buckets/{bucket}/folders` |
| `gyro storage key rotate` | POST access-keys:rotate (GyroDepolama) | `POST /cloud/storage/{id}/keys/{key}/rotate` |
| `gyro storage object delete` | Delete objects | `POST /cloud/storage/{id}/buckets/{bucket}/objects/delete` |
| `gyro storage object download-url` | Presigned download URL | `POST /cloud/storage/{id}/buckets/{bucket}/objects/download-url` |
| `gyro storage object list` | List objects in a bucket folder | `GET /cloud/storage/{id}/buckets/{bucket}/objects` |
| `gyro storage object upload-url` | Presigned upload URLs | `POST /cloud/storage/{id}/buckets/{bucket}/objects/upload-url` |
| `gyro storage usage get` | GET usage (GyroDepolama) | `GET /cloud/storage/{id}/usage` |
| `gyro ticket access-group list` | List ticket access groups | `GET /tickets/access-groups` |
| `gyro ticket attachment get` | Get ticket attachment (path: {id}, {aid}) | `GET /tickets/{id}/attachments/{aid}` |
| `gyro ticket close` | Close ticket (path: {id}) | `POST /tickets/{id}/close` |
| `gyro ticket create` | Create ticket | `POST /tickets` |
| `gyro ticket department list` | List ticket departments | `GET /tickets/departments` |
| `gyro ticket get` | Get ticket (path: {id}) | `GET /tickets/{id}` |
| `gyro ticket list` | List tickets | `GET /tickets` |
| `gyro ticket rating` | Rate ticket (path: {id}) | `POST /tickets/{id}/rating` |
| `gyro ticket reopen` | Reopen ticket (path: {id}) | `POST /tickets/{id}/reopen` |
| `gyro ticket reply` | Reply ticket (path: {id}) | `POST /tickets/{id}/reply` |
| `gyro ticket reply rating` | Rate ticket reply (path: {id}, {rid}) | `POST /tickets/{id}/replies/{rid}/rating` |
| `gyro user create` | Invite a sub-user | `POST /cloud/users` |
| `gyro user delete` | Remove a sub-user | `DELETE /cloud/users/{id}` |
| `gyro user list` | List sub-users | `GET /cloud/users` |
| `gyro user resend-invite` | Resend the invitation e-mail | `POST /cloud/users/{id}/resend-invite` |
| `gyro user update` | Update a sub-user (label, status, permissions) | `PATCH /cloud/users/{id}` |
| `gyro webhook create` | Create a webhook (secret returned once) | `POST /cloud/webhooks` |
| `gyro webhook delete` | Delete a webhook | `DELETE /cloud/webhooks/{id}` |
| `gyro webhook delivery list` | Last 100 deliveries | `GET /cloud/webhooks/{id}/deliveries` |
| `gyro webhook get` | Get a webhook | `GET /cloud/webhooks/{id}` |
| `gyro webhook list` | List webhooks | `GET /cloud/webhooks` |
| `gyro webhook test` | Send an api.test event now | `POST /cloud/webhooks/{id}/test` |
| `gyro webhook update` | Update url, events, description or status | `PATCH /cloud/webhooks/{id}` |
| `gyro whoami` | Describe the calling API key | `GET /whoami` |
