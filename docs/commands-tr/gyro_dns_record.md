## gyro dns record

DNS kayıtlarını listele, ekle, değiştir, sil

```
gyro dns record [flags]
```

### Options

```
  -h, --help   help for record
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

* [gyro dns](gyro_dns.md)	 - Alan adlarınızın DNS kayıtları
* [gyro dns record add](gyro_dns_record_add.md)	 - DNS Kaydı Ekleme
* [gyro dns record delete](gyro_dns_record_delete.md)	 - DNS Kaydını Silme
* [gyro dns record list](gyro_dns_record_list.md)	 - DNS Kayıtlarını Listeleme
* [gyro dns record update](gyro_dns_record_update.md)	 - DNS Kaydını Düzenleme

