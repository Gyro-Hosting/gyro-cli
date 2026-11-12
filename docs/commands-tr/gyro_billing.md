## gyro billing

Ön ödemeli saatlik kullanım (İstanbul bulut, kapsayıcı): bakiye, kullanım, tahmin, döküm

```
gyro billing [flags]
```

### Options

```
  -h, --help   help for billing
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro billing balance](gyro_billing_balance.md)	 - Bakiye, saatlik ücret ve tahmini kalan süre
* [gyro billing estimate](gyro_billing_estimate.md)	 - Planların ve sunucusuz kapsayıcıların saatlik, günlük ve aylık maliyeti; bakiyenizin ne kadar yeteceği
* [gyro billing statement](gyro_billing_statement.md)	 - Aylık kullanım dökümü (fatura değildir); indirmek için --pdf / --csv
* [gyro billing topup-url](gyro_billing_topup-url.md)	 - Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)
* [gyro billing usage](gyro_billing_usage.md)	 - Saat saat kullanım (ya da günlük/aylık toplam)

