## gyro reseller service terminate

Hizmeti sonlandır (bayi)

### Synopsis

Hizmeti sonlandır (bayi)

`"confirm": true` zorunlu (422 `confirm_required`). `"immediate": true` hemen iptal eder: ücretlendirme durur, kaynak ve verisi silinir (durum `cancelling`, yaklaşık bir dakikada `cancelled`; geri alınamaz; operatör kapattıysa 403 `immediate_disabled`). `immediate` olmadan iptal talebi açılır (yalnız etkin hizmet; askıdaki hizmette 409 `terminate_needs_immediate`). Önce ürünün iptal kuralları denetlenir (422 `cancellation_rejected`). İptal edilmiş hizmette ya da açık talepte `changed: false` döner. Yalnız bayinin kendi hesabındaki hizmetler (aksi 404 `service_not_found`); yalnız bayi anahtarları (403 `not_reseller`, alt hesap anahtarı 403 `subkey_forbidden`). Paneldeki yaşam döngüsünü Gyro'ya yansıtan fatura paneli modülleri (WHMCS / WISECP Gyro bayi modülü) için.

API: POST /cloud/reseller/services/{id}/terminate (operationId terminateResellerService, scope reseller)

```
gyro reseller service terminate <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for terminate
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --immediate                şimdi sil (geri alınamaz)
      --reason string            reason
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro reseller service](gyro_reseller_service.md)	 - Hizmetin bayi görünümü (durum, bitiş, açık faturalar)

