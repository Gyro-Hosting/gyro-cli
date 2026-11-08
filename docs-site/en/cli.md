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
| `gyro address default` | Set default address | `POST /addresses/{id}/default` |
| `gyro address delete` | Delete address | `DELETE /addresses/{id}` |
| `gyro address get` | Get address | `GET /addresses/{id}` |
| `gyro address list` | Get addresses | `GET /addresses` |
| `gyro address set` | Update address | `PUT /addresses/{id}` |
| `gyro affiliate commission list` | Get affiliate commissions | `GET /affiliate/commissions` |
| `gyro affiliate enroll` | Enroll affiliate | `POST /affiliate/enroll` |
| `gyro affiliate get` | Get affiliate | `GET /affiliate` |
| `gyro affiliate payment-method set` | Save affiliate payout method | `PUT /affiliate/payment-method` |
| `gyro affiliate withdraw` | Withdraw affiliate | `POST /affiliate/withdraw` |
| `gyro balance` | Account balance | `GET /balance` |
| `gyro balance setting update` | Update balance settings | `PATCH /balance/settings` |
| `gyro balance transaction list` | Balance ledger | `GET /balance/transactions` |
| `gyro bulk` | Run one action on up to 50 servers | `POST /cloud/bulk` |
| `gyro core invoice pay` | Pay invoice | `POST /invoices/{id}/pay` |
| `gyro core product list` | Get products | `GET /products` |
| `gyro core service get` | Get service | `GET /services/{id}` |
| `gyro core service list` | Get services | `GET /services` |
| `gyro core service tool get` | Get service tool data | `GET /services/{id}/tools/{tool}` |
| `gyro core service tool list` | Get service tools | `GET /services/{id}/tools` |
| `gyro core service tool run` | Run service tool action | `POST /services/{id}/tools/{tool}/{action}` |
| `gyro domain addon buy` | Buy domain addon | `POST /domains/{domain}/addons/{key}` |
| `gyro domain auth-code send` | Send auth code | `POST /domains/{domain}/auth-code` |
| `gyro domain auth-code set` | Save auth code | `PUT /domains/{domain}/auth-code` |
| `gyro domain check` | Check domain | `POST /domains/check` |
| `gyro domain child-nameserver create` | Create child nameserver | `POST /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver delete` | Delete child nameserver | `DELETE /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver list` | Get child nameservers | `GET /domains/{domain}/child-nameservers` |
| `gyro domain default-nameserver list` | Get default nameservers | `GET /domains/default-nameservers` |
| `gyro domain default-nameserver set` | Update default nameservers | `PUT /domains/default-nameservers` |
| `gyro domain dns-record create` | Create dns record | `POST /domains/{domain}/dns-records` |
| `gyro domain dns-record delete` | Delete dns record | `DELETE /domains/{domain}/dns-records` |
| `gyro domain dns-record list` | Get dns records | `GET /domains/{domain}/dns-records` |
| `gyro domain dns-record set` | Update dns record | `PUT /domains/{domain}/dns-records` |
| `gyro domain dnssec create` | Create dnssec record | `POST /domains/{domain}/dnssec` |
| `gyro domain dnssec delete` | Delete dnssec record | `DELETE /domains/{domain}/dnssec` |
| `gyro domain dnssec get` | Get dnssec records | `GET /domains/{domain}/dnssec` |
| `gyro domain email-forward create` | Create email forward | `POST /domains/{domain}/email-forwards` |
| `gyro domain email-forward delete` | Delete email forward | `DELETE /domains/{domain}/email-forwards` |
| `gyro domain email-forward list` | Get email forwards | `GET /domains/{domain}/email-forwards` |
| `gyro domain forwarding delete` | Delete url forwarding | `DELETE /domains/{domain}/forwarding` |
| `gyro domain forwarding get` | Get url forwarding | `GET /domains/{domain}/forwarding` |
| `gyro domain forwarding set` | Update url forwarding | `PUT /domains/{domain}/forwarding` |
| `gyro domain get` | Get domain | `GET /domains/{domain}` |
| `gyro domain list` | Get domains | `GET /domains` |
| `gyro domain nameserver list` | Get nameservers | `GET /domains/{domain}/nameservers` |
| `gyro domain nameserver set` | Update nameservers | `PUT /domains/{domain}/nameservers` |
| `gyro domain register` | Register domain | `POST /domains/register` |
| `gyro domain renew` | Renew domain | `POST /domains/{domain}/renew` |
| `gyro domain tld list` | Get tlds | `GET /domains/tlds` |
| `gyro domain transfer` | Transfer domain | `POST /domains/transfer` |
| `gyro domain transfer-lock get` | Get transfer lock | `GET /domains/{domain}/transfer-lock` |
| `gyro domain transfer-lock set` | Update transfer lock | `PUT /domains/{domain}/transfer-lock` |
| `gyro domain transfer-status get` | Get transfer status | `GET /domains/{domain}/transfer-status` |
| `gyro domain update` | Update domain | `PATCH /domains/{domain}` |
| `gyro domain whois apply` | Apply whois profile | `POST /domains/{domain}/whois/apply` |
| `gyro domain whois get` | Get whois contacts | `GET /domains/{domain}/whois` |
| `gyro domain whois set` | Update whois contacts | `PUT /domains/{domain}/whois` |
| `gyro domain whois-privacy set` | Update whois privacy | `PUT /domains/{domain}/whois-privacy` |
| `gyro domain whois-profile create` | Create whois profile | `POST /domains/whois-profiles` |
| `gyro domain whois-profile default` | Set default whois profile | `POST /domains/whois-profiles/{pid}/default` |
| `gyro domain whois-profile delete` | Delete whois profile | `DELETE /domains/whois-profiles/{pid}` |
| `gyro domain whois-profile list` | Get whois profiles | `GET /domains/whois-profiles` |
| `gyro domain whois-profile set` | Update whois profile | `PUT /domains/whois-profiles/{pid}` |
| `gyro image list` | List operating system images | `GET /cloud/images` |
| `gyro invoice coupon` | Apply invoice coupon | `POST /invoices/{id}/coupon` |
| `gyro invoice get` | Get an invoice | `GET /invoices/{id}` |
| `gyro invoice list` | List invoices | `GET /invoices` |
| `gyro invoice pay` | Pay an invoice from the account balance | `POST /cloud/invoices/{id}/pay` |
| `gyro key create` | Create a sub-account key (token returned once) | `POST /cloud/keys` |
| `gyro key delete` | Delete a sub-account key | `DELETE /cloud/keys/{id}` |
| `gyro key list` | List sub-account keys (never shows tokens) | `GET /cloud/keys` |
| `gyro location list` | List locations with orderable products | `GET /cloud/locations` |
| `gyro me get` | Get me | `GET /me` |
| `gyro me notification list` | Get notifications | `GET /me/notifications` |
| `gyro me notification set` | Update notifications | `PUT /me/notifications` |
| `gyro me security get` | Get security | `GET /me/security` |
| `gyro me update` | Update me | `PATCH /me` |
| `gyro order create` | Create order | `POST /orders` |
| `gyro order get` | Get order | `GET /orders/{id}` |
| `gyro order list` | Get orders | `GET /orders` |
| `gyro order preview` | Preview order | `POST /orders/preview` |
| `gyro ping get` | Ping | `GET /ping` |
| `gyro plan list` | List orderable cloud server products | `GET /cloud/products` |
| `gyro product category list` | Get product categories | `GET /products/categories` |
| `gyro product get` | Get product | `GET /products/{id}` |
| `gyro reference country list` | Get countries | `GET /reference/countries` |
| `gyro reference country state list` | Get states | `GET /reference/countries/{code}/states` |
| `gyro reference state city list` | Get cities | `GET /reference/states/{id}/cities` |
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
| `gyro service addon list` | Get service addons | `GET /services/{id}/addons` |
| `gyro service cancellation create` | Create cancellation | `POST /services/{id}/cancellation` |
| `gyro service cancellation delete` | Revoke cancellation | `DELETE /services/{id}/cancellation` |
| `gyro service dashboard get` | Get service dashboard | `GET /services/{id}/dashboard` |
| `gyro service get` | Get a service (renewal, cancellation, access details) | `GET /cloud/services/{id}` |
| `gyro service invoice list` | Get service invoices | `GET /services/{id}/invoices` |
| `gyro service list` | List services of every product line | `GET /cloud/services` |
| `gyro service method call` | Call a client-callable module method (e.g. license change_ip) | `POST /cloud/services/{id}/methods/{method}` |
| `gyro service metrics billing get` | Get service metric billing | `GET /services/{id}/metrics/billing` |
| `gyro service metrics list` | Get service metrics | `GET /services/{id}/metrics` |
| `gyro service metrics update` | Update service metric | `PATCH /services/{id}/metrics/{key}` |
| `gyro service module-method` | Use service module method | `POST /services/{id}/module-method` |
| `gyro service password` | Change service password | `POST /services/{id}/password` |
| `gyro service renew` | Renew service | `POST /services/{id}/renew` |
| `gyro service sso` | Get service sso | `POST /services/{id}/sso` |
| `gyro service tool get` | Read a tool (list mailboxes, snapshots, DNS records, ...) | `GET /cloud/services/{id}/tools/{tool}` |
| `gyro service tool list` | Discover tools, actions and safety policy of a service | `GET /cloud/services/{id}/tools` |
| `gyro service tool run` | Run a tool action (create a mailbox, restore a snapshot, ...) | `POST /cloud/services/{id}/tools/{tool}/{action}` |
| `gyro service update` | Update service | `PATCH /services/{id}` |
| `gyro sms message get` | Get sms message | `GET /sms/messages/{id}` |
| `gyro sms message list` | Get sms messages | `GET /sms/messages` |
| `gyro sms quote` | Quote sms | `POST /sms/quote` |
| `gyro sms send` | Send sms | `POST /sms/send` |
| `gyro sms sender list` | Get sms senders | `GET /sms/senders` |
| `gyro ssh-key create` | Add an SSH key | `POST /cloud/ssh-keys` |
| `gyro ssh-key delete` | Delete an SSH key | `DELETE /cloud/ssh-keys/{id}` |
| `gyro ssh-key get` | Get an SSH key | `GET /cloud/ssh-keys/{id}` |
| `gyro ssh-key list` | List SSH keys | `GET /cloud/ssh-keys` |
| `gyro ticket access-group list` | Get ticket access groups | `GET /tickets/access-groups` |
| `gyro ticket attachment get` | Get ticket attachment | `GET /tickets/{id}/attachments/{aid}` |
| `gyro ticket close` | Close ticket | `POST /tickets/{id}/close` |
| `gyro ticket create` | Create ticket | `POST /tickets` |
| `gyro ticket department list` | Get ticket departments | `GET /tickets/departments` |
| `gyro ticket get` | Get ticket | `GET /tickets/{id}` |
| `gyro ticket list` | Get tickets | `GET /tickets` |
| `gyro ticket rating` | Rate ticket | `POST /tickets/{id}/rating` |
| `gyro ticket reopen` | Reopen ticket | `POST /tickets/{id}/reopen` |
| `gyro ticket reply` | Reply ticket | `POST /tickets/{id}/reply` |
| `gyro ticket reply rating` | Rate ticket reply | `POST /tickets/{id}/replies/{rid}/rating` |
| `gyro webhook create` | Create a webhook (secret returned once) | `POST /cloud/webhooks` |
| `gyro webhook delete` | Delete a webhook | `DELETE /cloud/webhooks/{id}` |
| `gyro webhook delivery list` | Last 100 deliveries | `GET /cloud/webhooks/{id}/deliveries` |
| `gyro webhook get` | Get a webhook | `GET /cloud/webhooks/{id}` |
| `gyro webhook list` | List webhooks | `GET /cloud/webhooks` |
| `gyro webhook test` | Send an api.test event now | `POST /cloud/webhooks/{id}/test` |
| `gyro webhook update` | Update url, events, description or status | `PATCH /cloud/webhooks/{id}` |
| `gyro whoami` | Describe the calling API key | `GET /whoami` |
