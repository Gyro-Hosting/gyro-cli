## gyro dns acme-challenge remove

ACME doğrulama TXT değerini sil

### Synopsis

ACME doğrulama TXT değerini sil

Yalnız `value` değerini siler (sorgu ya da JSON gövde); küme boşalınca kayıt silinir.

API: DELETE /cloud/dns/acme-challenge (operationId acmeChallengeRemove, scope dns_acme)

```
gyro dns acme-challenge remove [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for remove
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              _acme-challenge.<domain>
      --value string             Challenge value (key authorization digest)
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

* [gyro dns acme-challenge](gyro_dns_acme-challenge.md)	 - ACME doğrulama TXT değeri ekle

