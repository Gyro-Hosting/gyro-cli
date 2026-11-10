## gyro s3 website set

Statik web sitesini aç / güncelle

### Synopsis

Statik web sitesini aç / güncelle

Tam değiştirme: verilmeyen alanlar varsayılana döner (index.html, 404 sayfası yok, SPA kapalı). `index_document` bir dosya adıdır (`/` yok); kökte ve her klasörde açılır. `spa: true` bulunmayan her yolu giriş dosyasıyla ve 200 koduyla yanıtlar (React/Vue/Angular yönlendirme, sayfa yenileme); gerçek dosyalar olduğu gibi gelir, SPA açıkken 404 sayfası kullanılmaz. Kapalı siteyi açmak `accept: true` ister (kötüye kullanım kuralları: kovadaki her şey herkese açılır).

API: PUT /cloud/storage/{id}/buckets/{bucket}/website (operationId typed_put_storage__id__buckets__bucket__website, scope services_write)

```
gyro s3 website set <ID> <BUCKET> [flags]
```

### Options

```
      --accept                   Kötüye kullanım kurallarını kabul et (site kapalıyken gerekli)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --error-document string    404 sayfasının nesne anahtarı (ör. 404.html), boş = yok
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --index-document string    Giriş dosyası adı, varsayılan index.html (ör. app.html)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --spa                      Bulunmayan alt yolları giriş dosyasına 200 ile yönlendir
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

* [gyro s3 website](gyro_s3_website.md)	 - Kovanın statik web sitesi ayarları

