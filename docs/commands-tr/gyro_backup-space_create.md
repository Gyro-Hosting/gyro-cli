## gyro backup-space create

Yedekleme alanı sipariş et

### Synopsis

Yedekleme alanı sipariş et

TEK bir yedekleme alanı sipariş eder ve hesap bakiyesinden öder (kayıtlı kart asla). Alan `server_id` ile hesabınızdaki etkin bir sunucuya bağlanır (Gyro Bulut tr-ist-1, TR VDS, yönetilen ya da fiziksel); hesapta tek sunucu varsa gönderilmeyebilir. Yalnız o sunucunun IPv4 adresleri bağlanabilir (SFTP, port 2201). `cycle` dönemine göre ücretlendirilir (sabit EUR fiyat hesap para biriminde, KDV hariç) ve diğer hizmetler gibi bakiyeden yenilenir. Teklif için `"dry_run": true`; gerçek sipariş `"confirm": true` ve `Idempotency-Key` başlığı ister. Alan yaklaşık bir dakikada kurulur (`backup_space.state` `pending` -> `active`); parola `GET /cloud/backup-spaces/{id}/credentials` ile alınır. Hatalar: 402 `insufficient_balance`; 403 `not_available` (bu hesap için satışta değil), `orders_disabled`; 409 `server_required` (etkin sunucu yok), `capacity_unavailable`; 422 `confirm_required` (`quote` ile), `product_not_orderable`, `cycle_invalid`, `server_id_required`, `invalid_server`, `invalid_quantity`; 429 `daily_limit`.

API: POST /cloud/backup-spaces (operationId createBackupSpace, scope cloud_order)

```
gyro backup-space create [flags]
```

### Options

```
      --cycle string             fatura dönemi (monthly|quarterly|semiannually|annually)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --end-customer-id int      reseller main key: assign to this end-customer
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              ad (takma ad)
      --plan string              GET /cloud/backup-spaces/plans paketi
      --product-id int           plan yerine ürün kimliği
      --project-id int           proje (0 = Genel)
      --server-id int            alanın bağlanacağı sunucu (hesapta tek etkin sunucu varsa isteğe bağlı)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tags strings             tags
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

* [gyro backup-space](gyro_backup-space.md)	 - Yedekleme Alanı: sunucunuza bağlı SFTP yedekleme alanı (tr-ist-1)

