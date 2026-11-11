## gyro dns acme-challenge add

ACME doğrulama TXT değeri ekle

### Synopsis

ACME doğrulama TXT değeri ekle

`value` değerini `name` adındaki TXT kümesine ekler (TTL 60); diğer değerler kalır, böylece bir sertifikanın kök ve joker adları birlikte doğrulanır (ad başına en çok 10 değer). Hizmet ve zone addan bulunur (adı kapsayan en uzun Gyro DNS zone'unuz; hosting/cPanel zone'ları desteklenmez). Tekrarlanabilir.

API: POST /cloud/dns/acme-challenge (operationId acmeChallengeAdd, scope dns_acme)

```
gyro dns acme-challenge add [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              _acme-challenge.<alan adı> (*.ornek.com için: _acme-challenge.ornek.com)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --value string             Doğrulama değeri (base64url, 16-128 karakter)
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

