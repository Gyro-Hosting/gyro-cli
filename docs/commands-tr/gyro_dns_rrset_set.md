## gyro dns rrset set

Kayıt kümesi oluştur ya da değiştir

### Synopsis

Kayıt kümesi oluştur ya da değiştir

Bir kayıt kümesini (ad + tip) yazar. CNAME ve ALIAS tek değer alır; kökte CNAME olmaz (ALIAS kullanın); kök NS ve SOA bizde.

API: PUT /cloud/dns/{id}/zones/{zone}/records (operationId typed_put_dns__id__zones__zone__records, scope services_write)

```
gyro dns rrset set <ID> <ZONE> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              Record name relative to the zone (@ = apex, * = wildcard) or FQDN with trailing dot
      --records strings          Values in zone file syntax, one per item (TXT quotes optional)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ttl int                  60-86400 s
      --type string              type (A|AAAA|CNAME|MX|TXT|SRV|CAA|ALIAS|NS|PTR)
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

* [gyro dns rrset](gyro_dns_rrset.md)	 - Kayıt kümelerini listele

