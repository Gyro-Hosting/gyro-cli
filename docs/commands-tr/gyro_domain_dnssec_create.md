## gyro domain dnssec create

İmza Kaydı Ekleme

### Synopsis

İmza Kaydı Ekleme

Alan adına bir imza doğrulama kaydı ekler. — algoritma denetlenir

API: POST /domains/{domain}/dnssec (operationId coreCreateDnssecRecord, scope Domains/CreateDnssecRecord)

```
gyro domain dnssec create <DOMAIN> [flags]
```

### Options

```
      --algorithm int                    Anahtar algoritmasının numarası. Modülün bildirdiği kümede olmalıdır.
      --data string                      istek gövdesi JSON, @dosya ya da @- (stdin)
      --digest string                    İmza özeti. DNSKEY isteyen servis sağlayıcı burada açık anahtarı (base64) alır.
      --digest-type meta.dnssec_fields   Özet algoritmasının numarası ya da meta.dnssec_fields'in adlandırdığı anahtar bayrağı veya protokol. Modülün bildirdiği kümede olmalıdır.
      --dry-run                          hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                             help for create
      --idempotency-key string           tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --key-tag dnskey-keytag-flags      Anahtarın etiketi. dnskey-keytag-flags'te anahtar bayrağını taşır.
      --set stringArray                  ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro domain dnssec](gyro_domain_dnssec.md)	 - İmza Kayıtlarını Listeleme

