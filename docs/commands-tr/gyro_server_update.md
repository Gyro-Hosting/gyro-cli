## gyro server update

Ad ya da otomatik yenilemeyi değiştir

```
gyro server update <SERVER> [flags]
```

### Examples

```
  gyro server update web-1 --hostname web-01.example.com
  gyro server update 47466 --auto-renew=false
```

### Options

```
      --auto-renew        bakiyeden otomatik yenile
  -h, --help              help for update
      --hostname string   yeni sunucu adı
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...

