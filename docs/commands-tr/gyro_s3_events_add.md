## gyro s3 events add

Kovaya olay webhook'u ekle

### Synopsis

Kovaya olay webhook'u ekle

**Kova olayları**, bir nesne oluşturulunca (yükleme, kopyalama, çok parçalı yükleme), silinince ya da toplu silme yapılınca adresinize imzalı HTTPS POST gönderir. Gövde: `{"id", "event", "created_at", "data": {service_id, bucket, key, size, etag, content_type, version_id (Korumalı Depolama), method (put / multipart / delete / bulk), source (api / panel), access_key_id, event_time}}` - nesne içeriği yoktur. Başlıklar: `X-Gyro-Event`, `X-Gyro-Delivery` (tekil kimlik; tekrarı ayıklamak için), `X-Gyro-Timestamp`, `X-Gyro-Signature: sha256=HMAC-SHA256(gizli, timestamp + "." + ham gövde)` (hesap webhook'larıyla aynı biçim). Teslim en az bir kez yapılır, sıra garantisi yoktur; 2xx dışı yanıt ya da 6 sn içinde yanıt yoksa yeniden denenir (8 deneme, yaklaşık 2 saat); art arda 20 hata webhook'u kapatır (`disabled_reason` = `auto_failures`). İstekler 45.87.174.151, 45.87.174.153, 45.87.174.154 adreslerinden gelir. Hedef: yalnız https, port 443 ya da 8443, genel internet (özel, yerel ve Gyro platform adresleri reddedilir; ad her teslimde yeniden çözülür). Süzgeç: `prefix` (anahtar şununla başlar) ve `suffix` (anahtar bunlardan biriyle biter, en çok 10). Sınır: kova başına 5, hizmet başına 20 webhook (`s3.max_event_hooks`, destek talebiyle artırılır). Fiyat: kova başına aylık ücretsiz hak (`GET /cloud/storage/{id}/events/usage`) + satın alındığı ay geçerli ön ödemeli olay paketi; aşım faturası yoktur: hak dolunca yeni olaylar gönderilmez. Bilinen sınırlar: toplu silmede anahtar listesi gelmez; yaşam döngüsü silmeleri olay üretmez. Yanıtta `secret` bir kez gelir. Ret, hesap dilinde mesajla 422 `tool_action_failed` döner (geçersiz / özel / çözülemeyen adres, port, olay türü, sınırlar).

API: POST /cloud/storage/{id}/buckets/{bucket}/events (operationId createBucketEventWebhook, scope services_write)

```
gyro s3 events add <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --events strings           default: all three. Short names created / deleted / deleted_bulk are accepted
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --prefix string            prefix
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --suffix strings           a comma separated text is accepted too
      --url string               url
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

* [gyro s3 events](gyro_s3_events.md)	 - Kovanın olay webhook'larını listele

