## gyro s3 alerts set

Kovanın doluluk uyarısını ayarla

### Synopsis

Kovanın doluluk uyarısını ayarla

Eşiği ayarlar ya da değiştirir (uyarı yeniden kurulur). Kova zaten eşiğin üzerindeyse e-posta ~5 dakika içinde gider. `type` percent (1-100, paketin yüzdesi) ya da gb (1 ile paket boyutu arası). Hata: hesap dilinde mesajla 422 `tool_action_failed`. **Doluluk uyarısı**, kova kullanımı eşiğe ulaşınca hesap e-posta adresine bir e-posta gönderir: `percent` = paket boyutunun yüzdesi (hizmetin tüm kovaları paketi ortak kullanır) ya da `gb` = mutlak boyut. 5 dakikada bir kontrol edilir. E-postadan sonra uyarı `triggered` olur; kullanım eşiğin %95'inin altına inince yeniden kurulur ve eşik tekrar aşılırsa yine gönderilir. Uyarı yalnız bilgilendirir, yüklemeyi durdurmaz. Paket genelindeki %80 / %100 e-postaları aynen sürer. Nesne Depolama ve Korumalı Depolama (eski sürümler ve silme işaretleri kullanıma dahildir).

API: PUT /cloud/storage/{id}/buckets/{bucket}/alert (operationId setStorageBucketAlert, scope services_write)

```
gyro s3 alerts set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --type string              type (percent|gb)
      --value int                value
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

* [gyro s3 alerts](gyro_s3_alerts.md)	 - Tüm kovaların doluluk uyarıları

