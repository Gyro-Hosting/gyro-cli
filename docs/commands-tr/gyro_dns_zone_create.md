## gyro dns zone create

Zone ekle

### Synopsis

Zone ekle

Alan adını talep eder. Sahiplik doğrulanınca ns1/ns2.gyrodns.net'te yayınlanır: ad sunucularımıza devir (zaten devredilmişse hemen) ya da mevcut DNS'inizde `_gyrodns-dogrulama.<zone>` = `txt_value` TXT kaydı. İsteğe bağlı `bind` (zone dosyası) etkinleşince içe aktarılır (SOA ve kök NS yok sayılır).

API: POST /cloud/dns/{id}/zones (operationId typed_post_dns__id__zones, scope services_write)

```
gyro dns zone create <ID> [flags]
```

### Options

```
      --bind string              Optional BIND zone file (max 2 MB)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              Zone name (IDN allowed)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro dns zone](gyro_dns_zone.md)	 - DNS zone'larını listele

