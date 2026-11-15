## gyro vpc port-forward

Port yönlendirmeleri listele

```
gyro vpc port-forward [flags]
```

### Options

```
  -h, --help   help for port-forward
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

* [gyro vpc](gyro_vpc.md)	 - Özel ağları listele
* [gyro vpc port-forward add](gyro_vpc_port-forward_add.md)	 - Kapının dış portunu yönlendir
* [gyro vpc port-forward list](gyro_vpc_port-forward_list.md)	 - Port yönlendirmeleri listele
* [gyro vpc port-forward remove](gyro_vpc_port-forward_remove.md)	 - Port yönlendirmeyi sil
* [gyro vpc port-forward update](gyro_vpc_port-forward_update.md)	 - Port yönlendirmeyi kimin kullanabileceğini değiştir (kaynak izin listesi)

