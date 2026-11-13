## gyro server rebuild

İşletim sistemini yeniden kur - DİSK SİLİNİR

```
gyro server rebuild <SERVER> [flags]
```

### Options

```
      --app string              yeni sisteme kurulacak tek tıkla uygulama ('gyro app list' kimliği; API 2.19.1)
      --app-domain string       uygulama alan adı (A kaydı sunucuyu göstermeli)
  -h, --help                    help for rebuild
      --image string            imaj (zorunlu)
      --user-data string        yeni sistem için cloud-init: metin, @dosya ya da @- (tr-ist-1)
      --wait                    sunucu tekrar çalışana kadar bekle
      --wait-timeout duration   --wait için en uzun süre (default 20m0s)
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

