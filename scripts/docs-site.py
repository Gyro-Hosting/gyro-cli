#!/usr/bin/env python3
"""Build the "CLI" pages of the developer docs site: docs-site/{tr,en}/cli.html (+ cli.md).

    python3 -I scripts/docs-site.py            (needs `go`; runs `go run ./cmd/gyro api operations -o json`)

The pages use the docs site's own ../assets/style.css (no inline script/style, CSP friendly) and are meant to be
served at /developers/{tr,en}/cli.html.
"""
import html, json, os, subprocess, sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
OUT = os.path.join(ROOT, "docs-site")

T = {
 "tr": {
  "title": "Gyro CLI — komut satırı aracı", "desc": "gyro: Gyro Hosting hizmetlerini terminalden yönetin.",
  "other": "en", "other_label": "English", "nav_api": "API belgeleri", "nav_title": "CLI",
  "sections": [
   ("install", "Kurulum", """<table><tr><th>Yöntem</th><th>Komut</th></tr>
<tr><td>Linux / macOS</td><td><code>curl -fsSL https://www.gyrohosting.com/cli/install.sh | sh</code></td></tr>
<tr><td>Homebrew</td><td><code>brew install gyro-hosting/tap/gyro</code></td></tr>
<tr><td>npm (Node.js ≥16)</td><td><code>npm install -g @gyro-hosting/cli</code><br><code>npx @gyro-hosting/cli --version</code></td></tr>
<tr><td>Windows (Scoop)</td><td><code>scoop bucket add gyro https://github.com/Gyro-Hosting/scoop-bucket</code><br><code>scoop install gyro</code></td></tr>
<tr><td>Debian / Ubuntu</td><td><code>sudo apt install ./gyro_&lt;sürüm&gt;_linux_amd64.deb</code></td></tr>
<tr><td>RHEL / Alma / Rocky</td><td><code>sudo dnf install ./gyro_&lt;sürüm&gt;_linux_amd64.rpm</code></td></tr>
<tr><td>Docker</td><td><code>docker run --rm -e GYRO_TOKEN ghcr.io/gyro-hosting/gyro server list</code></td></tr>
</table>
<p>Tek dosyalık program (Linux, macOS, Windows; amd64, arm64). Sürüm dosyalarının SHA-256 listesi Gyro sürüm anahtarıyla imzalıdır
(<code>E96B E963 5E77 A0D9 119A C8C3 A798 6589 A154 B990</code>).</p>"""),
   ("start", "Başlarken", """<pre><code>gyro configure                     # API anahtarı: Müşteri paneli → sağ üstteki hesap menüsü → API Kimlik Bilgileri
gyro whoami
gyro plan list --location falkenstein
gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 \\
     --ssh-key ~/.ssh/id_ed25519.pub --wait
gyro server ssh web-1</code></pre>
<p>Sipariş komutu önce API'den fiyat teklifini gösterir, onayınızı ister ve bakiyeden öder. <code>--wait</code> sunucu çalışana kadar bekler.</p>"""),
   ("session", "Örnek oturum", """<pre><code>$ gyro server create web-1 --plan standart-1-1 --location falkenstein --image debian-12 --ssh-key laptop --wait
Sipariş özeti
  Plan:        standart-1-1 (eu-de-1, 1 vCPU / 1 GB / 20 GB)
  İmaj:        debian-12
  Fiyat:       719.39 TRY / monthly
  Bakiye:      4280.61 TRY
Bu sunucu sipariş edilip bakiyenizden 719.39 TRY ödensin mi? [e/H] e
Sunucu 47485 hazır.
Bağlanın: gyro server ssh 47485   (ssh root@203.0.113.10)

$ gyro server power shutdown web-1 --wait
$ gyro server list -o json --query "[?status=='active'].ipv4"
$ gyro invoice list --unpaid
$ gyro dns record add example.com --type A --name www --value 203.0.113.10
$ gyro mail mailbox list example.com</code></pre>"""),
   ("profiles", "Profiller ve kimlik", """<ul>
<li><code>gyro configure [--profile ad]</code> anahtarı doğrular ve <code>~/.config/gyro/config.toml</code> dosyasına (izin 0600) yazar; <code>--keyring</code> işletim sistemi anahtar zincirini kullanır.</li>
<li>Betikler: <code>echo "$ANAHTAR" | gyro configure --token-stdin --profile ci</code> ya da yalnız <code>GYRO_TOKEN</code>.</li>
<li>Öncelik: bayrak &gt; ortam (<code>GYRO_TOKEN</code>, <code>GYRO_PROFILE</code>, <code>GYRO_API_URL</code>, <code>GYRO_OUTPUT</code>, <code>GYRO_LANG</code>) &gt; profil.</li>
<li><code>gyro context</code> etkin profili ve anahtar yetkilerini gösterir; <code>gyro context use ad</code> varsayılanı değiştirir.</li></ul>"""),
   ("output", "Çıktı ve betikler", """<ul>
<li><code>-o table|json|yaml|text</code>, <code>--query</code> (JMESPath), <code>-q</code> yalnız kimlikler, <code>--no-header</code>, <code>--wide</code>.</li>
<li>Renkler yalnız terminalde; <code>NO_COLOR</code> kapatır. Dil yerel ayardan ya da <code>--lang tr|en</code>.</li>
<li>Ad ya da kimlik: <code>gyro server get web-1</code> ve <code>gyro server get 47485</code> aynıdır.</li>
<li>Kabuk tamamlama: <code>gyro completion bash|zsh|fish|powershell</code> (sunucu ve anahtar adları API'den tamamlanır).</li></ul>"""),
   ("safety", "Güvenlik: onay, deneme, idempotency", """<ul>
<li>Para harcayan ve geri alınamayan işlemler onay sorar; betiklerde <code>--yes</code>. Terminal yoksa ve <code>--yes</code> verilmemişse komut çalışmaz (çıkış kodu 7).</li>
<li><code>--dry-run</code> fiyat teklifini (destekleyen uçlarda) ya da gönderilecek isteği gösterir; hiçbir şey değişmez.</li>
<li>Para işlemlerine otomatik <code>Idempotency-Key</code> eklenir: tekrar çalıştırmak ikinci sipariş vermez.</li>
<li>429/5xx yanıtlarında güvenli istekler artan beklemeyle tekrarlanır. Hatalar API kodu, HTTP durumu ve istek numarasıyla yazılır.</li></ul>
<table><tr><th>Çıkış kodu</th><th>Anlam</th></tr><tr><td>0</td><td>başarılı</td></tr><tr><td>2</td><td>kullanım hatası</td></tr><tr><td>3</td><td>kimlik / yetki (401, 403)</td></tr>
<tr><td>4</td><td>bulunamadı</td></tr><tr><td>5</td><td>çakışma (409)</td></tr><tr><td>6</td><td>ödeme</td></tr><tr><td>7</td><td>onay verilmedi</td></tr><tr><td>8</td><td>hız sınırı</td></tr><tr><td>9</td><td><code>--wait</code> zaman aşımı</td></tr></table>"""),
   ("limits", "Sınırlar ve sınır artırma", """<p><b>Hesap sınırları destek talebiyle artırılır.</b> <code>gyro limits</code> tüm sınırları ve kullanımı gösterir (yeni hesap ilk 7 gün: 5 İstanbul sunucusu, 1 kapsayıcı ortamı). Sınır yüzünden reddedilen komut sınırı, kullanımı ve hazır talep komutunu yazar:</p>
<pre><code>gyro limits
gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "İstenen sınır: 10. Gerekçe: ..."</code></pre>"""),
   ("lines", "Tüm ürünler", """<p>Mail, S3, izleme, lisans, VDS, fiziksel sunucu ve uygulama hizmetleri aynı köprüden yönetilir. Araçlar ve eylemler hizmetten okunur, yeni özellikler CLI güncellemesi beklemeden çalışır:</p>
<pre><code>gyro mail list
gyro mail tools example.com                       # araçlar, eylemler, onay/sır politikası
gyro mail mailbox create example.com --set local_part=info --set password=@-
gyro s3 bucket list depom
gyro license status 48001
gyro service tool run 48001 &lt;araç&gt; &lt;eylem&gt; --set alan=değer</code></pre>"""),
  ],
  "ref": "Komut başvurusu", "ref_p": "Her API ucu bir komuttur. Bu tablo CLI'nin kendisinden üretilmiştir (<code>gyro api operations</code>). Ayrıntı: <code>gyro &lt;komut&gt; --help</code>.",
  "cols": ("Komut", "Açıklama", "API"),
 },
 "en": {
  "title": "Gyro CLI — command-line interface", "desc": "gyro: manage Gyro Hosting services from your terminal.",
  "other": "tr", "other_label": "Türkçe", "nav_api": "API reference", "nav_title": "CLI",
  "sections": [
   ("install", "Install", """<table><tr><th>Method</th><th>Command</th></tr>
<tr><td>Linux / macOS</td><td><code>curl -fsSL https://www.gyrohosting.com/cli/install.sh | sh</code></td></tr>
<tr><td>Homebrew</td><td><code>brew install gyro-hosting/tap/gyro</code></td></tr>
<tr><td>npm (Node.js ≥16)</td><td><code>npm install -g @gyro-hosting/cli</code><br><code>npx @gyro-hosting/cli --version</code></td></tr>
<tr><td>Windows (Scoop)</td><td><code>scoop bucket add gyro https://github.com/Gyro-Hosting/scoop-bucket</code><br><code>scoop install gyro</code></td></tr>
<tr><td>Debian / Ubuntu</td><td><code>sudo apt install ./gyro_&lt;version&gt;_linux_amd64.deb</code></td></tr>
<tr><td>RHEL / Alma / Rocky</td><td><code>sudo dnf install ./gyro_&lt;version&gt;_linux_amd64.rpm</code></td></tr>
<tr><td>Docker</td><td><code>docker run --rm -e GYRO_TOKEN ghcr.io/gyro-hosting/gyro server list</code></td></tr>
</table>
<p>A single binary for Linux, macOS and Windows (amd64, arm64). Release checksums are signed with the Gyro release key
(<code>E96B E963 5E77 A0D9 119A C8C3 A798 6589 A154 B990</code>).</p>"""),
   ("start", "Getting started", """<pre><code>gyro configure                     # API key: Client area → account menu (top right) → API Credentials
gyro whoami
gyro plan list --location falkenstein
gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 \\
     --ssh-key ~/.ssh/id_ed25519.pub --wait
gyro server ssh web-1</code></pre>
<p>Ordering shows the API price quote first, asks for confirmation and pays from your balance. <code>--wait</code> blocks until the server runs.</p>"""),
   ("session", "Sample session", """<pre><code>$ gyro server create web-1 --plan standart-1-1 --location falkenstein --image debian-12 --ssh-key laptop --wait
Order summary
  Plan:        standart-1-1 (eu-de-1, 1 vCPU / 1 GB / 20 GB)
  Image:       debian-12
  Price:       719.39 TRY / monthly
  Balance:     4280.61 TRY
Order this server and pay 719.39 TRY from your balance? [y/N] y
Server 47485 is ready.
Connect: gyro server ssh 47485   (ssh root@203.0.113.10)

$ gyro server power shutdown web-1 --wait
$ gyro server list -o json --query "[?status=='active'].ipv4"
$ gyro invoice list --unpaid
$ gyro dns record add example.com --type A --name www --value 203.0.113.10
$ gyro mail mailbox list example.com</code></pre>"""),
   ("profiles", "Profiles and credentials", """<ul>
<li><code>gyro configure [--profile name]</code> checks the key and writes <code>~/.config/gyro/config.toml</code> (mode 0600); <code>--keyring</code> uses the OS keychain.</li>
<li>Scripts: <code>echo "$KEY" | gyro configure --token-stdin --profile ci</code> or just <code>GYRO_TOKEN</code>.</li>
<li>Precedence: flag &gt; environment (<code>GYRO_TOKEN</code>, <code>GYRO_PROFILE</code>, <code>GYRO_API_URL</code>, <code>GYRO_OUTPUT</code>, <code>GYRO_LANG</code>) &gt; profile.</li>
<li><code>gyro context</code> shows the active profile and key permissions; <code>gyro context use name</code> switches the default.</li></ul>"""),
   ("output", "Output and scripting", """<ul>
<li><code>-o table|json|yaml|text</code>, <code>--query</code> (JMESPath), <code>-q</code> ids only, <code>--no-header</code>, <code>--wide</code>.</li>
<li>Colours only on a terminal; <code>NO_COLOR</code> disables them. Language from the locale or <code>--lang tr|en</code>.</li>
<li>Names or ids: <code>gyro server get web-1</code> equals <code>gyro server get 47485</code>.</li>
<li>Shell completion: <code>gyro completion bash|zsh|fish|powershell</code> (server and key names complete from the API).</li></ul>"""),
   ("safety", "Safety: confirmation, dry run, idempotency", """<ul>
<li>Money and destructive actions ask for confirmation; use <code>--yes</code> in scripts. Without a terminal and without <code>--yes</code> the command refuses to run (exit code 7).</li>
<li><code>--dry-run</code> shows the price quote (where supported) or the request that would be sent; nothing changes.</li>
<li>Money operations carry an automatic <code>Idempotency-Key</code>: re-running never orders twice.</li>
<li>Safe requests are retried with backoff on 429/5xx. Errors show the API code, HTTP status and request id.</li></ul>
<table><tr><th>Exit code</th><th>Meaning</th></tr><tr><td>0</td><td>success</td></tr><tr><td>2</td><td>usage error</td></tr><tr><td>3</td><td>authentication / permission (401, 403)</td></tr>
<tr><td>4</td><td>not found</td></tr><tr><td>5</td><td>conflict (409)</td></tr><tr><td>6</td><td>payment</td></tr><tr><td>7</td><td>not confirmed</td></tr><tr><td>8</td><td>rate limited</td></tr><tr><td>9</td><td><code>--wait</code> timeout</td></tr></table>"""),
   ("limits", "Limits and raising them", """<p><b>Account limits are raised with a support ticket.</b> <code>gyro limits</code> lists every limit with its usage (new accounts, first 7 days: 5 Istanbul servers, 1 container environment). A command refused because of a limit prints the limit, the usage and the ready ticket command:</p>
<pre><code>gyro limits
gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "Requested limit: 10. Reason: ..."</code></pre>"""),
   ("lines", "Every product", """<p>Mail, S3, monitors, licences, VDS, dedicated servers and apps use the same bridge. Tools and actions are read from the service, so new features work without a CLI update:</p>
<pre><code>gyro mail list
gyro mail tools example.com                       # tools, actions, confirm/secret policy
gyro mail mailbox create example.com --set local_part=info --set password=@-
gyro s3 bucket list mystorage
gyro license status 48001
gyro service tool run 48001 &lt;tool&gt; &lt;action&gt; --set field=value</code></pre>"""),
  ],
  "ref": "Command reference", "ref_p": "Every API endpoint is a command. This table is generated from the CLI itself (<code>gyro api operations</code>). Details: <code>gyro &lt;command&gt; --help</code>.",
  "cols": ("Command", "Description", "API"),
 },
}


def ops(lang):
    env = dict(os.environ, GYRO_LANG=lang, GYRO_CONFIG_FILE="/nonexistent", NO_COLOR="1")
    out = subprocess.run(["go", "run", "./cmd/gyro", "api", "operations", "-o", "json"], cwd=ROOT, env=env,
                         check=True, capture_output=True, text=True).stdout
    rows = json.loads(out)
    rows.sort(key=lambda r: r["command"])
    return rows


def page(lang):
    t = T[lang]
    e = html.escape
    nav = "".join(f'<a href="#{sid}">{e(title)}</a>' for sid, title, _ in t["sections"])
    body = "".join(f'<section id="{sid}"><h2>{e(title)}</h2>{content}</section>' for sid, title, content in t["sections"])
    rows = ops(lang)
    trs = "".join(
        f'<tr><td><code>{e(r["command"])}</code></td><td>{e(r.get("summary") or "")}</td>'
        f'<td><span class="m {e(r["method"].lower())}">{e(r["method"])}</span> <code>{e(r["path"])}</code></td></tr>' for r in rows)
    c = t["cols"]
    ref = (f'<section id="reference"><h2>{e(t["ref"])}</h2><p>{t["ref_p"]}</p>'
           f'<table><tr><th>{c[0]}</th><th>{c[1]}</th><th>{c[2]}</th></tr>{trs}</table></section>')
    return f"""<!doctype html><html lang="{lang}"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>{e(t["title"])}</title>
<meta name="description" content="{e(t["desc"])}">
<link rel="alternate" hreflang="{t["other"]}" href="../{t["other"]}/cli.html"><link rel="stylesheet" href="../assets/style.css"><link rel="icon" href="../assets/logo.svg"></head>
<body><header><a class="logo" href="./"><img src="../assets/logo.svg" width="26" height="26" alt="">Gyro Developers</a><span class="sp"></span>
<a href="./">{e(t["nav_api"])}</a><a href="../{t["other"]}/cli.html" hreflang="{t["other"]}">{e(t["other_label"])}</a></header>
<div class="wrap"><nav aria-label="cli"><h4>{e(t["nav_title"])}</h4>{nav}<a href="#reference">{e(t["ref"])}</a></nav>
<main><section id="top"><h1>gyro <span class="badge">CLI</span></h1><p class="mut">{e(t["desc"])}</p></section>{body}{ref}</main></div>
</body></html>
""", rows


def markdown(lang, rows):
    t = T[lang]
    lines = [f"# {t['title']}", "", t["desc"], ""]
    for sid, title, content in t["sections"]:
        lines += [f"## {title}", "", "(HTML: cli.html#" + sid + ")", ""]
    lines += [f"## {t['ref']}", "", "| " + " | ".join(t["cols"]) + " |", "|---|---|---|"]
    lines += [f"| `{r['command']}` | {r.get('summary') or ''} | `{r['method']} {r['path']}` |" for r in rows]
    return "\n".join(lines) + "\n"


def main():
    for lang in ("tr", "en"):
        os.makedirs(os.path.join(OUT, lang), exist_ok=True)
        h, rows = page(lang)
        open(os.path.join(OUT, lang, "cli.html"), "w").write(h)
        open(os.path.join(OUT, lang, "cli.md"), "w").write(markdown(lang, rows))
        print(f"{lang}: {len(rows)} operations", file=sys.stderr)


if __name__ == "__main__":
    main()
