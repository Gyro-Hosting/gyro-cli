## gyro order preview

Siparişi Fiyatlandırma

### Synopsis

Siparişi Fiyatlandırma

Sipariş gövdesini fiyatlandırır ve tek satır bile yazmaz. — hiçbir şey yazılmaz

API: POST /orders/preview (operationId corePreviewOrder, scope Orders/PreviewOrder)

```
gyro order preview [flags]
```

### Options

```
      --addons string              Seçilen ek hizmetler. Ek hizmet numarasını seçenek numarasına eşler. Ürünün ilk dönemi ücretsiz dahil ettiği ek hizmet gönderilmese de siparişe eklenir. İlk dönem tutarı 0 olur, yenilemeler ek hizmet fiyatından faturalanır ve siparişten çıkarılamaz. (JSON)
      --addons-qty string          Adetli ek hizmetlerin sayısı. (JSON)
      --billing-profile-id int     Fatura profilinin numarası. Vergi bağlamını belirler; verilmezse varsayılan profil, hesapta kayıtlı profil yoksa hesabın ülkesi kullanılır.
      --card-id string             card_id
      --coupons string             Denenecek kupon kodları.
      --cycle string               Ödeme döngüsü. Ürünün fiyatlı döngülerinden biri olmalıdır.
      --data string                istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string              Satırın alan adı ekseni. Seçim, ad, alt alan adı kökü ve lisans alanlarını taşır. (JSON)
      --dry-run                    hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                       help for preview
      --idempotency-key string     tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --items string               items[]Sipariş satırları. Birden yirmiye kadar; şekli sipariş ucuyla aynıdır.
      --method string              method
      --metrics string             Açılacak ölçümlerin numaraları.
      --payment string             Ödeme bloğu. Verilirse yöntemin komisyonu toplama eklenir; bakiye yeterliliği burada denetlenmez. (JSON)
      --product-id int             Sipariş edilecek ürünün numarası.
      --quantity int               Adet. Yalnız ürün çoklu alıma izin veriyorsa dikkate alınır.
      --requirement-files string   Dosya isteyen alanların içeriği. Kodlanmış içerik olarak gönderilir; adres kabul edilmez. (JSON)
      --requirements string        Ürünün sorduğu alanların yanıtları. Zorunlu işaretlenen her alan gönderilmelidir. (JSON)
      --set stringArray            ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro order](gyro_order.md)	 - Siparişler (önizleme ve oluşturma)

