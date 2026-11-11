## gyro dns acme-challenge status

Doğrulama kaydı yayında mı?

### Synopsis

Doğrulama kaydı yayında mı?

İki yetkili ad sunucusuna (ns1/ns2.gyrodns.net) doğrudan sorar. İkisi de `value` değerini yayınlıyorsa `ready` true. Ekledikten sonra birkaç saniyede bir sorun (genelde < 1 dk).

API: GET /cloud/dns/acme-challenge (operationId acmeChallengeStatus, scope dns_acme)

```
gyro dns acme-challenge status [flags]
```

### Options

```
  -h, --help           help for status
      --name string    _acme-challenge.<domain>
      --value string   Challenge value (key authorization digest)
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

