## gyro invoice list

Faturaları listele

```
gyro invoice list [flags]
```

### Examples

```
  gyro invoice list --unpaid
  gyro invoice list --all -o json --query "[?state=='paid'].total.amount"
```

### Options

```
      --all             tüm sayfaları getir
  -h, --help            help for list
      --status string   API durum süzgeci
      --unpaid          yalnız ödenmemiş ve gecikmiş faturalar
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

* [gyro invoice](gyro_invoice.md)	 - Faturalar

