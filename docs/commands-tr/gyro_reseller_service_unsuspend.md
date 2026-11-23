## gyro reseller service unsuspend

Hizmeti askıdan çıkar (bayi)

### Synopsis

Hizmeti askıdan çıkar (bayi)

Bayinin koyduğu askıyı kaldırır. Gyro'nun koyduğu askı buradan kaldırılamaz: 409 `suspended_by_provider`; hizmetin vadesi geçmiş Gyro faturası varken 409 `overdue_invoice` (önce ödeyin). Etkin hizmette tekrar `changed: false` döner. Yalnız bayinin kendi hesabındaki hizmetler (aksi 404 `service_not_found`); yalnız bayi anahtarları (403 `not_reseller`, alt hesap anahtarı 403 `subkey_forbidden`). Paneldeki yaşam döngüsünü Gyro'ya yansıtan fatura paneli modülleri (WHMCS / WISECP Gyro bayi modülü) için.

API: POST /cloud/reseller/services/{id}/unsuspend (operationId unsuspendResellerService, scope reseller)

```
gyro reseller service unsuspend <ID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for unsuspend
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

