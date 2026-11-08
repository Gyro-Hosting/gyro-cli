## gyro license tool

Bir aracın verisini oku ya da eylemini çalıştır ('tools' ile bakın)

```
gyro license tool <SERVICE> <TOOL> [ACTION] [flags]
```

### Examples

```
  gyro mail tool example.com mailboxes
  gyro mail tool example.com mailboxes create --set local_part=info --set password=@-
```

### Options

```
      --data string       JSON gövde, @dosya ya da @-
  -h, --help              help for tool
      --set stringArray   alan anahtar=değer (tekrarlanabilir)
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

* [gyro license](gyro_license.md)	 - Yazılım lisansları (cPanel, Plesk, LiteSpeed ...)

