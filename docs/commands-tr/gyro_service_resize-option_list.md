## gyro service resize-option list

Hizmetin yükseltilebileceği ürünler ve fiyatları

### Synopsis

Hizmetin yükseltilebileceği ürünler ve fiyatları

**Veri güvencesi (Gyro Bulut İstanbul, tr-ist-1)**: plan değişikliği diski ve içindeki tüm veriyi, IP adreslerini, ek diskleri ve snapshot'ları korur. Kök disk yerinde büyütülür ve dosya sistemi kendiliğinden genişler; disk asla küçülmez, daha küçük diskli plan 409 `disk_shrink` ile reddedilir (`dry_run`'da da; resize-options bu planları `allowed: false` gösterir). Artan CPU/RAM hipervizör izin veriyorsa canlı uygulanır, yoksa sunucu düzgünce kapatılıp yeniden açılır (ölçüm: 1-2 -> 2-4 yeniden başladı, SSH yaklaşık 12 sn yok; 2-4 -> 4-8 canlı, yeniden başlatma yok). Sunucuda snapshot varken disk büyütülemez: 409 `resize_refused` (önce snapshot'ları silin). Yanıtta `data_preserved: true`, `disk: grows_only`, `restart: if_needed`. Veriyi yalnız yeniden kurulum siler (`POST /cloud/servers/{id}/actions/rebuild`, `confirm: true` ister): disk, ek diskler ve snapshot'lar silinir (IP adresleri kalır). **`recreate_required`** (422, yalnız faturalı ürünler): hedef plan hizmetin başka platformda yeniden kurulmasını gerektirir, bu da veriyi siler; API bunu asla yapmaz, hiçbir şey değişmez ve ücret alınmaz. resize-options'tan başka plan seçin ya da yeni sunucu açıp veriyi taşıyın.

API: GET /cloud/services/{id}/resize-options (operationId getResizeOptions, scope services_read)

```
gyro service resize-option list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro service resize-option](gyro_service_resize-option.md)	 - Hizmetin yükseltilebileceği ürünler ve fiyatları

