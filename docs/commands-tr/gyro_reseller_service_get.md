## gyro reseller service get

Hizmetin bayi görünümü (durum, bitiş, açık faturalar)

### Synopsis

Hizmetin bayi görünümü (durum, bitiş, açık faturalar)

Hizmetin durumu, ürünü, dönemi, fiyatı, bitiş tarihi, kimin askıya aldığı, atanmış uç müşteri ve açık faturaları. Yalnız bayinin kendi hesabındaki hizmetler (aksi 404 `service_not_found`); yalnız bayi anahtarları (403 `not_reseller`, alt hesap anahtarı 403 `subkey_forbidden`). Paneldeki yaşam döngüsünü Gyro'ya yansıtan fatura paneli modülleri (WHMCS / WISECP Gyro bayi modülü) için.

API: GET /cloud/reseller/services/{id} (operationId getResellerService, scope reseller)

```
gyro reseller service get <ID> [flags]
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

* [gyro reseller service](gyro_reseller_service.md)	 - Hizmetin bayi görünümü (durum, bitiş, açık faturalar)

