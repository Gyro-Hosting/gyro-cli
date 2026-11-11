## gyro container env set

Ortam değişkenlerini ayarla

### Synopsis

Ortam değişkenlerini ayarla

`set` değişken ekler/günceller (`{"KEY":"değer"}` ya da KEY=VALUE satırları), `unset` siler, `replace: 1` tümünü değiştirir. Sürüm oluşturur, servis yeniden başlar.

API: PUT /cloud/containers/{id}/services/{name}/env (operationId typed_put_containers__id__services__name__env, scope services_write)

```
gyro container env set <ID> <NAME> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --param-set string         set (JSON)
      --replace int              replace (0|1)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --unset strings            unset
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

* [gyro container env](gyro_container_env.md)	 - Ortam değişkenlerini ayarla

