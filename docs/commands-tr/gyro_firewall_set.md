## gyro firewall set

Ağ güvenlik duvarı politikasını değiştir

### Synopsis

Ağ güvenlik duvarı politikasını değiştir

Applies as a 5 minute trial unless `permanent` is true; confirm with POST /firewall/policy/confirm. Warnings return 422 `confirm_warnings_required` with `details.warnings`.

API: PUT /cloud/servers/{id}/firewall/policy (operationId putFirewallPolicy, scope cloud_write)

```
gyro firewall set <ID> [flags]
```

### Options

```
      --confirm-warnings         confirm_warnings
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --default-action string    default_action (allow|deny)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --permanent                Skip the 5 minute trial (otherwise call POST /firewall/policy/confirm within 5 minutes or the change is reverted)
      --ping                     ping
      --rules strings            rules
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro firewall](gyro_firewall.md)	 - Ağ güvenlik duvarı politikası (kurallar, varsayılan kural, durum)

