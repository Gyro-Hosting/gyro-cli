## gyro domain dns-record create

DNS Kaydı Ekleme

### Synopsis

DNS Kaydı Ekleme

Alan adına yeni bir DNS kaydı ekler. — ek hizmet gerekir

API: POST /domains/{domain}/dns-records (operationId coreCreateDnsRecord, scope Domains/CreateDnsRecord)

```
gyro domain dns-record create <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --identity string          identity
      --name string              Kaydın yanıtlayacağı ad. Kabul edilen biçim sağlayıcıya göre değişir.
      --priority int             Öncelik. Posta ve hizmet kayıtları dışında yok sayılır.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ttl int                  Yaşam süresi. Verilmezse sağlayıcının varsayılanı kullanılır.
      --type string              Kayıt türü.
      --value string             Kaydın içeriği.
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

* [gyro domain dns-record](gyro_domain_dns-record.md)	 - DNS Kayıtlarını Listeleme

