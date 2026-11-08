## gyro domain renew

Yenileme Faturası Kesme

### Synopsis

Yenileme Faturası Kesme

**Para harcar / fatura oluşturur.** Para kuralları (AK-40): gövdede `"confirm": true`, `Idempotency-Key` başlığı, yalnız bakiye (kayıtlı kart yok), hesap başına günlük sınır. Alan adı için istenen yıl sayısı kadar yenileme faturası keser. — fatura üretir

API: POST /domains/{domain}/renew (operationId coreRenewDomain, scope Domains/RenewDomain)

```
gyro domain renew <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for renew
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --years int                Yenileme süresi. Detaydaki fiyatlı süre satırlarından biri olmalıdır.
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

* [gyro domain](gyro_domain.md)	 - Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS

