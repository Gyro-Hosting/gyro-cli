## gyro server firewall rule add

Gelen trafik kuralı ekle

### Synopsis

Gelen trafik kuralı ekle

Network firewall add-on: inserts at `priority` (default last) and applies as a trial unless `permanent`. Otherwise the cloud server's simple inbound allow list.

API: POST /cloud/servers/{id}/firewall/rules (operationId addFirewallRule, scope cloud_write)

```
gyro server firewall rule add <ID> [flags]
```

### Options

```
      --action string            action (allow|deny)
      --confirm-warnings         Apply even when the change has warnings (lock-out risk, shadowed or duplicate rules, private sources)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --destination string       any (all IPs of the server) or one IP address of the server
      --direction string         in = inbound to the server, out = new outbound connections (network firewall only) (in|out)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --permanent                Network firewall - skip the 5 minute trial and make the change permanent at once
      --port string              any | 22 | 80,443 | 1000-2000 (not for icmp/all)
      --priority int             Position (1 = evaluated first). Default = last. First matching rule wins.
      --protocol string          protocol (tcp|udp|tcp_udp|icmp|all)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --source string            any | IPv4/CIDR list (comma separated) | country:TR | not-country:TR (country add-on). For direction out: the remote address.
      --source-ips strings       Legacy - CIDRs (use source)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
```

### Options inherited from parent commands

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string        mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color           renkleri kapat (NO_COLOR da olur)
      --no-header          tablo başlığını yazma
  -o, --output string      çıktı biçimi: table, json, yaml, text
  -p, --profile string     kullanılacak profil (ortam GYRO_PROFILE)
      --query string       sonuca uygulanacak JMESPath sorgusu
  -q, --quiet              yalnız kimlikleri yaz
      --timeout duration   toplam zaman aşımı, ör. 30s, 10m
      --wide               tüm tablo sütunlarını göster
  -y, --yes                onay sorma
```

### SEE ALSO

* [gyro server firewall rule](gyro_server_firewall_rule.md)	 - Gelen trafik kuralı ekle

