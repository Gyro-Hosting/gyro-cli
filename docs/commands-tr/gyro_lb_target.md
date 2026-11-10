## gyro lb target

Hedefleri sağlık durumuyla listele

```
gyro lb target [flags]
```

### Options

```
  -h, --help   help for target
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

* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele
* [gyro lb target add](gyro_lb_target_add.md)	 - Hedef ekle
* [gyro lb target list](gyro_lb_target_list.md)	 - Hedefleri sağlık durumuyla listele
* [gyro lb target remove](gyro_lb_target_remove.md)	 - Hedefi çıkar

