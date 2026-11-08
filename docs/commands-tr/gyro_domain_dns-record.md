## gyro domain dns-record

DNS Kayıtlarını Listeleme

```
gyro domain dns-record [flags]
```

### Options

```
  -h, --help   help for dns-record
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
* [gyro domain dns-record create](gyro_domain_dns-record_create.md)	 - DNS Kaydı Ekleme
* [gyro domain dns-record delete](gyro_domain_dns-record_delete.md)	 - DNS Kaydını Silme
* [gyro domain dns-record list](gyro_domain_dns-record_list.md)	 - DNS Kayıtlarını Listeleme
* [gyro domain dns-record set](gyro_domain_dns-record_set.md)	 - DNS Kaydını Düzenleme

