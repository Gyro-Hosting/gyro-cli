## gyro billing estimate

Planların ve sunucusuz kapsayıcıların saatlik, günlük ve aylık maliyeti; bakiyenizin ne kadar yeteceği

### Synopsis

Planlar (İstanbul sunucuları, kapsayıcı planları) aylık fiyatla sınırlıdır. Sunucusuz kapsayıcılar (--cpu-m/--memory-mi ya da --item <cpu_m>/<memory_mi>:<adet>) çalışan kopya başına saatlik, tavansız ücretlenir; boyutlar: gyro container sizes.

```
gyro billing estimate [flags]
```

### Examples

```
  gyro billing estimate --plan standart-2-4 --region tr-ist-1 --count 3
  gyro billing estimate --item standart-2-4:2 --item standart-4-8:1
  gyro billing estimate --cpu-m 500 --memory-mi 1024 --count 2 --hours 730
  gyro billing estimate --item 250/512:3 --item standart-2-4:1
```

### Options

```
      --count int          adet (default 1)
      --cpu-m int          sunucusuz kapsayıcı: kopya başına işlemci, milicore (250, 500, 1000, 2000, 4000)
  -h, --help               help for estimate
      --hours int          sunucusuz kapsayıcı: toplam için çalışma saati (varsayılan 730)
      --item stringArray   plan:adet ya da sunucusuz <cpu_m>/<memory_mi>:adet, tekrarlanabilir
      --memory-mi int      sunucusuz kapsayıcı: kopya başına bellek, MiB (bkz. gyro container sizes)
      --plan string        plan adı, ör. standart-2-4
      --product-id int     ya da ürün no
      --region string      bölge (ön ödemeli: tr-ist-1) (default "tr-ist-1")
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

