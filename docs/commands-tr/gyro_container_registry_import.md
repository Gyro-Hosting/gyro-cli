## gyro container registry import

Başka bir depodan imaj içe aktar

### Synopsis

Başka bir depodan imaj içe aktar

Bir imajı (ör. `nginx:1.27`, `ghcr.io/gyro-ornek/uygulama:1.0`) `docker push` gerekmeden ortamın Gyro imaj deposuna kopyalar. Yalnız https depolar; çok mimarili imajlarda yalnız linux/amd64 kopyalanır. `repo` (varsayılan: kaynak adının son parçası) ve `tag` (varsayılan: kaynak etiketi) kopyanın adıdır; var olan ad:etiket için `overwrite: true` gerekir. Özel imaj: `username` + `password` (token) yalnız bu kopya için kullanılır, saklanmaz, geri dönmez, API günlüğünde karartılır. Arka planda çalışır: GET .../registry/imports ile izleyin. Sınırlar: ortam başına aynı anda 1 (409 `import_running`), saatte 10 (429 `import_rate_limited`), 15 dakika, 128 katman, boyut deponuzdaki boş alanı aşamaz. İç/özel/Gyro altyapı adreslerine çözülen kaynaklar reddedilir (`source_blocked`).

API: POST /cloud/containers/{id}/registry/imports (operationId importContainerImage, scope services_write)

```
gyro container registry import <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for import
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --overwrite                var olan ad:etiketin üzerine yaz
      --password string          özel kaynak: token / parola (saklanmaz)
      --repo string              hedef ad (küçük harf, en çok 3 seviye)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --source string            kaynak imaj (host/depo:etiket ya da @sha256:...); host yoksa docker.io
      --tag string               hedef etiket
      --username string          özel kaynak: kullanıcı adı (saklanmaz)
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

* [gyro container registry](gyro_container_registry.md)	 - Ortamın Gyro imaj deposu

