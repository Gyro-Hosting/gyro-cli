## gyro billing statement

Aylık kullanım dökümü (fatura değildir); indirmek için --pdf / --csv

```
gyro billing statement YYYY-MM [flags]
```

### Examples

```
  gyro billing statement 2026-11
  gyro billing statement 2026-11 --pdf .
  gyro billing statement 2026-11 --csv kasim.csv
```

### Options

```
      --csv string   CSV kaydet (dosya, '.' = sunucunun adı)
  -h, --help         help for statement
      --pdf string   PDF kaydet (dosya, '.' = sunucunun adı)
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

* [gyro billing](gyro_billing.md)	 - Ön ödemeli saatlik kullanım (İstanbul bulut, kapsayıcı): bakiye, kullanım, tahmin, döküm

