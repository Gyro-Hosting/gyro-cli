## gyro s3 alerts get

Kovanın doluluk uyarısı

### Synopsis

Kovanın doluluk uyarısı

Bir kovanın eşiği, durumu ve güncel kullanımı. **Doluluk uyarısı**, kova kullanımı eşiğe ulaşınca hesap e-posta adresine bir e-posta gönderir: `percent` = paket boyutunun yüzdesi (hizmetin tüm kovaları paketi ortak kullanır) ya da `gb` = mutlak boyut. 5 dakikada bir kontrol edilir. E-postadan sonra uyarı `triggered` olur; kullanım eşiğin %95'inin altına inince yeniden kurulur ve eşik tekrar aşılırsa yine gönderilir. Uyarı yalnız bilgilendirir, yüklemeyi durdurmaz. Paket genelindeki %80 / %100 e-postaları aynen sürer. Nesne Depolama ve Korumalı Depolama (eski sürümler ve silme işaretleri kullanıma dahildir).

API: GET /cloud/storage/{id}/buckets/{bucket}/alert (operationId getStorageBucketAlert, scope services_read)

```
gyro s3 alerts get <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro s3 alerts](gyro_s3_alerts.md)	 - Tüm kovaların doluluk uyarıları

