# gyro — Gyro Hosting komut satırı aracı

[English](README.en.md)

`gyro`, Gyro Hosting hizmetlerinizi terminalden yönetmenizi sağlayan resmi komut satırı aracıdır
(`aws`, `doctl`, `hcloud` gibi). Gyro Public API'nin sunduğu her işlem bir komuttur: bulut sunucular, snapshot'lar,
SSH anahtarları, alan adları ve DNS, kurumsal mail, S3 depolama, izleme, lisanslar, faturalar, bakiye, destek talepleri,
siparişler, webhook'lar, alt hesap anahtarları ve bayi (reseller) işlemleri.

Tek dosyalık, bağımlılıksız bir programdır (Linux, macOS, Windows; amd64 ve arm64).

## Kurulum

| Yöntem | Komut |
|---|---|
| Betik (Linux/macOS) | `curl -fsSL https://www.gyrohosting.com/cli/install.sh \| sh` |
| Homebrew | `brew install gyro-hosting/tap/gyro` |
| npm (Node.js ≥16) | `npm install -g @gyro-hosting/cli` ya da `npx @gyro-hosting/cli --version` |
| Scoop (Windows) | `scoop bucket add gyro https://github.com/Gyro-Hosting/scoop-bucket && scoop install gyro` |
| Debian/Ubuntu | sürüm sayfasındaki `.deb`: `sudo apt install ./gyro_<sürüm>_linux_amd64.deb` |
| RHEL/Alma/Rocky | `sudo dnf install ./gyro_<sürüm>_linux_amd64.rpm` |
| Docker | `docker run --rm -e GYRO_TOKEN ghcr.io/gyro-hosting/gyro server list` |
| Elle | <https://github.com/Gyro-Hosting/gyro-cli/releases> → arşivi açın, `gyro`'yu PATH'e koyun |

Sürüm dosyaları SHA-256 ile listelenir ve liste Gyro sürüm anahtarıyla imzalanır
(`E96B E963 5E77 A0D9 119A C8C3 A798 6589 A154 B990`, `packaging/gyro-release-signing-public.asc`).
Kurulum betiği gpg varsa imzayı, her durumda sağlama toplamını doğrular.

## Başlarken

```sh
gyro configure                    # API anahtarı: Müşteri paneli → sağ üstteki hesap menüsü → API Kimlik Bilgileri (wck_...)
gyro whoami                       # anahtar ve yetkileri
gyro plan list --location falkenstein
gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 \
     --ssh-key ~/.ssh/id_ed25519.pub --wait
gyro server ssh web-1
```

## Örnek oturum

```text
$ gyro server create web-1 --plan standart-1-1 --location falkenstein --image debian-12 --ssh-key laptop --wait
Sipariş özeti
  Plan:        standart-1-1 (eu-de-1, 1 vCPU / 1 GB / 20 GB)
  İmaj:        debian-12
  Ad:          web-1
  Fiyat:       719.39 TRY / monthly
  Vergi:       119.90 TRY
  Bakiye:      4280.61 TRY
Bu sunucu sipariş edilip bakiyenizden 719.39 TRY ödensin mi? [e/H] e
Sipariş #69140 verildi, sunucu kimliği 47485 (tekrar anahtarı 3f1c…).
⠼ sunucu 47485 starting, meşgul (41s)
Sunucu 47485 hazır.
Bağlanın: gyro server ssh 47485   (ssh root@203.0.113.10)

$ gyro server list
ID      HOSTNAME   STATUS   PLAN           REGION    IPV4           PRICE        DUE
47485   web-1      active   standart-1-1   eu-de-1   203.0.113.10   599.49 TRY   2026-12-09

$ gyro server power shutdown web-1 --wait && gyro server power on web-1 --wait
$ gyro server delete web-1
Sunucu 47485 (web-1, 203.0.113.10) silinsin mi? Geri alınamaz.
Devam etmek için "web-1" yazın: web-1
```

## Komut ağacı (özet)

```text
gyro
├── server        list · get · create · delete · update · rename · ip · ssh · wait · metrics
│                 power on|off|shutdown|reboot|reset · rebuild · reset-password · rescue enable|disable
│                 snapshot list|create|restore|delete · backup list|create · rdns list|set
│                 firewall list · firewall rule add|delete · console · activity list
├── plan list     image list   location list   ssh-key list|get|create|delete
├── service       list · get · tool list|get|run · method call     (tüm ürün hatları, genel köprü)
├── mail          list · tools · tool · mailbox|alias|catch-all|autoresponder|migration|dns|device|spam list|create|update|delete
├── s3            list · tools · bucket … · key …
├── monitor       list · check … · incident … · contact …
├── license       list · status · change-ip      app · vds · dedicated (list, tools, tool, call)
├── domain        list · get · check · register · transfer · renew · nameserver · dns-record · whois · dnssec …
├── dns record    list · add · update · delete
├── bulk          (en çok 50 sunucuda tek işlem)
├── invoice       list [--unpaid] · get · pay · coupon        balance (+ transaction list, setting update)
├── ticket        list · get · create · reply · close · reopen · rating · department list
├── order         list · get · preview · create
├── webhook       list · get · create · update · delete · test · delivery list
├── key           list · create · delete                       (alt hesap API anahtarları)
├── reseller      customer list|get|create|update|delete · customer key … · assign · report · setting
├── me · address · affiliate · reference · sms · ping · activity list · whoami
└── configure · context · api request|operations · completion · version
```

Tam liste: `gyro api operations` (her API ucu ve karşılığı olan komut) ve [docs/commands-tr](docs/commands-tr/gyro.md).

## Profiller ve kimlik

- `gyro configure [--profile ad]` — anahtarı doğrular (`/whoami`) ve `~/.config/gyro/config.toml` dosyasına (izin 0600) yazar.
  `--keyring` ile anahtar dosya yerine işletim sistemi anahtar zincirinde (macOS Keychain, Windows Credential Manager,
  Linux Secret Service) tutulur. Betikler için: `echo "$ANAHTAR" | gyro configure --token-stdin --profile ci`.
- Öncelik: bayrak (`--profile`, `--api-url`) > ortam (`GYRO_TOKEN`, `GYRO_PROFILE`, `GYRO_API_URL`, `GYRO_OUTPUT`, `GYRO_LANG`) > profil > varsayılan.
- `gyro context` etkin profili, adresi, maskeli anahtarı ve yetkileri gösterir; `gyro context use <profil>` varsayılanı değiştirir.
- Varsayılan API adresi: `https://www.gyrohosting.com/api/v1/client`.

## Çıktı

- `-o table` (varsayılan), `-o json`, `-o yaml`, `-o text` (sekme ayrılmış)
- `--query` JMESPath: `gyro server list -o json --query "[?status=='active'].ipv4"`
- `-q` yalnız kimlikler (`gyro server list -q | xargs -n1 gyro server get`), `--no-header`, `--wide`
- Renkler yalnız terminalde; `NO_COLOR=1` ya da `--no-color` kapatır. Mesaj dili yerel ayardan (`tr_*` → Türkçe) ya da `--lang tr|en`.

## Güvenlik

- Para harcayan (sipariş, fatura ödeme, yenileme, alan adı kaydı …) ve geri alınamayan (silme, yeniden kurulum, geri yükleme,
  elektrik kesme …) işlemler onay sorar. Betiklerde `--yes`; terminal yoksa ve `--yes` yoksa komut **çalışmaz** (çıkış kodu 7).
- Sunucu silmede sunucunun adını yazmanız istenir.
- `--dry-run`: fiyat teklifi destekleyen uçlarda API'den teklif alır, diğerlerinde gönderilecek isteği gösterir; hiçbir şey değişmez.
- Para işlemlerine otomatik `Idempotency-Key` eklenir: ağ hatasında komutu tekrar çalıştırmak ikinci sipariş vermez.
  Tek seferlik sır dönen işlemlerde (root parolası, konsol bileti, anahtar) bu anahtar gönderilmez, sır yanıtta gelir ve bir kez gösterilir.
- 429 ve 5xx yanıtlarında (yalnız güvenli isteklerde) artan beklemeyle tekrar dener; `Retry-After`'a uyar.
- Hata mesajı API hata kodunu, HTTP durumunu ve istek numarasını içerir (destek talebinde iletin).
- `--debug` HTTP trafiğini stderr'e yazar; anahtar ve sırlar maskelenir.

## Çıkış kodları

| Kod | Anlam |
|---|---|
| 0 | başarılı |
| 1 | genel hata |
| 2 | kullanım hatası (bayrak, belirsiz ad) |
| 3 | kimlik/yetki (401, 403) |
| 4 | bulunamadı (404) |
| 5 | çakışma (409: sunucu meşgul, aynı anahtar …) |
| 6 | ödeme (402, yetersiz bakiye, ödeme tamamlanmadı) |
| 7 | onay verilmedi / `--yes` gerekli |
| 8 | hız sınırı (429, tekrarlardan sonra) |
| 9 | `--wait` zaman aşımı |

## Kabuk tamamlama

```sh
source <(gyro completion bash)          # bash
gyro completion zsh > "${fpath[1]}/_gyro" # zsh
gyro completion fish > ~/.config/fish/completions/gyro.fish
gyro completion powershell | Out-String | Invoke-Expression
```
Sunucu, SSH anahtarı, hizmet vb. kimlikler API'den tamamlanır. Paketlerle (deb/rpm/brew) tamamlama dosyaları otomatik kurulur.

## Mimari (geliştiriciler için)

```text
gyro-api/openapi.yaml ──scripts/sync-spec.sh──▶ internal/spec/openapi.yaml
                                   │ go generate (cmd/gyro-gen)
                                   ▼
                     internal/spec/zz_generated_ops.go   (her işlem: yol, parametreler, gövde alanları,
                                   │                      kapsam, onay/para bayrağı, komut adı, tablo sütunları)
                                   ▼
internal/cli/generated.go  → her işlem için otomatik komut (bayraklar, --data/--set, --all, --dry-run, onay, ad→kimlik çözümü)
internal/cli/server.go …   → elle yazılmış kullanıcı deneyimi (sipariş özeti, --wait, ssh, güç, ürün hattı kısayolları)
```

- Yeni bir API ucu: `scripts/sync-spec.sh` → komut kendiliğinden gelir. Adlandırma yoldan türetilir
  (`GET /cloud/servers/{id}/snapshots` → `gyro server snapshot list <ID>`); spec'te `x-cli: {command: "a b c", confirm: true, money: true, hidden: true}` ile değiştirilebilir.
- Elle yazılmış komut bir işlemi kapsıyorsa `markHandled("operationId")` ile üretileni gizler; `gyro api operations` her işlemin hangi komutla çalıştığını gösterir.
- Çekirdek (WISECP) uçları Gyro bulut uçlarıyla aynı adı alırsa `gyro core …` altında kalır.

## Geliştirme ve test

```sh
go generate ./... && go vet ./... && go test ./...           # birim + golden testler (sahte API)
go test ./internal/cli -update                                # golden dosyalarını yenile
GYRO_E2E_URL=<api-url> GYRO_E2E_KEYS=<anahtarlar.json> go test -tags e2e ./e2e -v   # gerçek API'ye karşı
GYRO_E2E_REAL=1 GYRO_E2E_PURGE="<temizleme komutu>" go test -tags e2e ./e2e -run TestRealServer -v
goreleaser release --snapshot --clean --skip=sign,publish,docker
```

Yayın adımları: [RELEASE.md](RELEASE.md). Lisans: MPL-2.0.
