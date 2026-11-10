## gyro dns zone import

BIND zon dosyası içe aktar

### Synopsis

BIND zon dosyası içe aktar

`merge` dosyadaki kayıt kümelerini değiştirir, diğerlerini korur; `replace` dosyada olmayanları da siler. SOA ve kök NS yok sayılır; desteklenmeyen tipler `skipped` içinde. `"confirm": true` gerekir.

API: POST /cloud/dns/{id}/zones/{zone}/import (operationId typed_post_dns__id__zones__zone__import, scope services_write)

```
gyro dns zone import <ID> <ZONE> [flags]
```

### Options

```
      --bind string              bind
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for import
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --mode string              mode (merge|replace)
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

* [gyro dns zone](gyro_dns_zone.md)	 - DNS zonlarını listele

