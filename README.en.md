# gyro — the Gyro Hosting command-line interface

[Türkçe](README.md)

`gyro` is the official CLI for Gyro Hosting (in the spirit of `aws`, `doctl` and `hcloud`). Every operation of the
Gyro Public API is a command: cloud servers, snapshots, SSH keys, domains and DNS, corporate mail, S3 storage,
monitors, licences, invoices, balance, support tickets, orders, webhooks, sub-account keys and reseller endpoints.

A single static binary for Linux, macOS and Windows (amd64, arm64).

## Install

| Method | Command |
|---|---|
| Script (Linux/macOS) | `curl -fsSL https://www.gyrohosting.com/cli/install.sh \| sh` |
| Homebrew | `brew install gyro-hosting/tap/gyro` |
| npm (Node.js ≥16) | `npm install -g @gyro-hosting/cli` or `npx @gyro-hosting/cli --version` |
| Scoop (Windows) | `scoop bucket add gyro https://github.com/Gyro-Hosting/scoop-bucket && scoop install gyro` |
| Debian/Ubuntu | `.deb` from the release page: `sudo apt install ./gyro_<version>_linux_amd64.deb` |
| RHEL/Alma/Rocky | `sudo dnf install ./gyro_<version>_linux_amd64.rpm` |
| Docker | `docker run --rm -e GYRO_TOKEN ghcr.io/gyro-hosting/gyro server list` |
| Manual | <https://github.com/Gyro-Hosting/gyro-cli/releases> — unpack, put `gyro` on your PATH |

Release files are listed with SHA-256 checksums, signed with the Gyro release key
(`E96B E963 5E77 A0D9 119A C8C3 A798 6589 A154 B990`, `packaging/gyro-release-signing-public.asc`).
The install script verifies the checksum and, when gpg is available, the signature.

## Getting started

```sh
gyro auth login                   # sign in with the browser (no key to copy)
# or: gyro configure              # paste an API key: Client area → account menu (top right) → API Credentials (wck_...)
gyro whoami
gyro plan list --location falkenstein
gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 \
     --ssh-key ~/.ssh/id_ed25519.pub --wait
gyro server ssh web-1
```

## Sample session

```text
$ gyro server create web-1 --plan standart-1-1 --location falkenstein --image debian-12 --ssh-key laptop --wait
Order summary
  Plan:        standart-1-1 (eu-de-1, 1 vCPU / 1 GB / 20 GB)
  Image:       debian-12
  Hostname:    web-1
  Price:       719.39 TRY / monthly
  Tax:         119.90 TRY
  Balance:     4280.61 TRY
Order this server and pay 719.39 TRY from your balance? [y/N] y
Order #69140 placed, server id 47485 (idempotency key 3f1c…).
Server 47485 is ready.
Connect: gyro server ssh 47485   (ssh root@203.0.113.10)

$ gyro server list -o json --query "[?status=='active'].{id:id,ip:ipv4}"
[ { "id": 47485, "ip": "203.0.113.10" } ]
```

## Command tree (summary)

```text
gyro
├── server        list · get · create · delete · update · rename · ip · ssh · wait · metrics
│                 power on|off|shutdown|reboot|reset · rebuild · reset-password · rescue enable|disable
│                 snapshot list|create|restore|delete · backup list|create · rdns list|set
│                 firewall list · firewall rule add|delete · console · activity list
├── plan list     image list   location list   ssh-key list|get|create|delete
├── service       list · get · tool list|get|run · method call     (every product line, generic bridge)
├── mail          list · tools · tool · mailbox|alias|catch-all|autoresponder|migration|dns|device|spam list|create|update|delete
├── s3            list · tools · bucket … · key …
├── backup-space plans · list · get · create · credentials · reset-password · attach-server · delete   (SFTP backup space, tr-ist-1)
├── monitor       list · check … · incident … · contact …
├── license       list · status · change-ip      app · vds · dedicated (list, tools, tool, call)
├── domain        list · get · check · register · transfer · renew · nameserver · dns-record · whois · dnssec …
├── dns record    list · add · update · delete
├── bulk          (one action on up to 50 servers)
├── invoice       list [--unpaid] · get · pay · coupon        balance (+ transaction list, setting update)
├── ticket        list · get · create · reply · close · reopen · rating · department list
├── order         list · get · preview · create
├── webhook       list · get · create · update · delete · test · delivery list
├── key           list · create · delete                       (sub-account API keys)
├── reseller      customer list|get|create|update|delete · customer key … · assign · report · setting
├── me · address · affiliate · reference · sms · ping · activity list · whoami
└── auth login|status|logout · configure · context · api request|operations · completion · version
```

Full list: `gyro api operations` (every endpoint and its command) and [docs/commands](docs/commands/gyro.md).

## Profiles and credentials

- `gyro auth login [--profile name] [--no-browser] [--template readonly]` opens the panel in your browser (device authorization, like `gh auth login`):
  check the code, pick permissions and the key lifetime, approve; the CLI stores the new time-limited key in the profile. Over SSH the address is only printed.
  `gyro auth status` shows account, key, permissions and expiry; `gyro auth logout` revokes the key and removes it locally.
- `gyro configure [--profile name]` (asks for the browser sign-in first) checks the key (`/whoami`) and writes `~/.config/gyro/config.toml` (mode 0600).
  `--keyring` keeps the key in the OS keychain instead. Scripts: `echo "$KEY" | gyro configure --token-stdin --profile ci`.
- Working in another account (API 2.46.0): when the owner of an account adds you under *Account > Users*, you act there with your own key.
  `gyro account list` shows the accounts you can reach and your permissions; `--account <id>` for one command, `GYRO_ACCOUNT=<id>` for a shell,
  `gyro account use <id>` for the profile default (`gyro account use self` resets it). `gyro auth login` stores the account picked on the approval page.
  Permissions are exactly the panel user permissions; a refused operation answers 403 (`account_permission_denied`, `account_owner_only`), exit code 3.
- Precedence: flags (`--profile`, `--api-url`, `--account`) > environment (`GYRO_TOKEN`, `GYRO_PROFILE`, `GYRO_API_URL`, `GYRO_OUTPUT`, `GYRO_LANG`, `GYRO_ACCOUNT`) > profile > default.
- `gyro context` shows the active profile, URL, masked key and permissions; `gyro context use <profile>` switches the default.

## Output

`-o table|json|yaml|text`, `--query` (JMESPath), `-q` (ids only), `--no-header`, `--wide`. Colours only on a terminal;
`NO_COLOR` disables them. Messages follow the locale (`tr_*` → Turkish) or `--lang tr|en`.

## Safety

- Money and destructive actions ask for confirmation; `--yes` in scripts. Without a terminal and without `--yes` the command
  refuses to run (exit code 7). Deleting a server asks you to type its name.
- `--dry-run` returns the API price quote where supported, otherwise prints the request; nothing changes.
- Money operations carry an automatic `Idempotency-Key`, so re-running after a network error never orders twice. Calls that
  return a one-time secret (root password, console ticket, keys) do not send one, so the secret is returned (shown once).
- Retries with exponential backoff on 429/5xx for safe requests, honouring `Retry-After`.
- Errors show the API error code, HTTP status and request id. `--debug` logs HTTP traffic with secrets masked.

## Limits

**Account limits are raised with a support ticket.** `gyro limits` lists every limit with its usage (new accounts, first 7 days: 5 Istanbul servers, 1 container environment). A command refused because of a limit prints the limit, the usage and the ready ticket command:

```sh
gyro limits
gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "Requested limit: 10. Reason: ..."
```

## Exit codes

0 ok · 1 error · 2 usage · 3 auth (401/403) · 4 not found · 5 conflict (409) · 6 payment · 7 not confirmed · 8 rate limited · 9 `--wait` timeout.

## Shell completion

`gyro completion bash|zsh|fish|powershell` — server, SSH key and service ids complete from the API.

## Architecture

The command layer is generated from the OpenAPI document (`cmd/gyro-gen` → `internal/spec/zz_generated_ops.go`); every
operation becomes a command automatically. Hand-written commands (server create/ssh/power, product-line shortcuts, billing)
add UX on top. New endpoints: run `scripts/sync-spec.sh`. Override names with the `x-cli` extension
(`command`, `confirm`, `money`, `hidden`).

## Development

```sh
go generate ./... && go vet ./... && go test ./...
GYRO_E2E_URL=... GYRO_E2E_KEYS=... go test -tags e2e ./e2e -v   # against a real API
goreleaser release --snapshot --clean --skip=sign,publish,docker
```

Release: [RELEASE.md](RELEASE.md). Licence: MPL-2.0.
