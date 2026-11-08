## gyro order create

Sipariş Verme

### Synopsis

Sipariş Verme

**Para harcar / fatura oluşturur.** Para kuralları (AK-40): gövdede `"confirm": true`, `Idempotency-Key` başlığı, yalnız bakiye (kayıtlı kart yok), hesap başına günlük sınır. Siparişi verir ve aynı çağrıda tahsil eder. — bakiye ya da kayıtlı kart

API: POST /orders (operationId coreCreateOrder, scope Orders/CreateOrder)

```
gyro order create [flags]
```

### Options

```
      --addons string              Seçilen ek hizmetler. Ek hizmet numarasını seçenek numarasına eşler. Ürünün ilk dönemi ücretsiz dahil ettiği ek hizmet gönderilmese de siparişe eklenir. İlk dönem tutarı 0 olur, yenilemeler ek hizmet fiyatından faturalanır ve siparişten çıkarılamaz. (JSON)
      --addons-qty string          Adetli ek hizmetlerin sayısı. (JSON)
      --billing-profile-id int     Fatura profilinin numarası.
      --card-id int                Çekilecek kartın numarası. Verilmezse hesabın varsayılan kartı çekilir.
      --coupons string             Kupon kodları. Operatörün kendiliğinden uygulanan kuponları da katılır.
      --cycle string               Ödeme döngüsü. Ürünün fiyatlı döngülerinden biri olmalıdır.
      --data string                istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string              Satırın alan adı ekseni. Seçim, ad, alt alan adı kökü ve lisans alanlarını taşır. (JSON)
      --dry-run                    hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                       help for create
      --idempotency-key string     tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --items string               items[]Sipariş satırları. Birden yirmiye kadar.
      --method string              Ödeme yolu: hesap bakiyesi ya da kayıtlı kart.
      --metrics string             Açılacak ölçümlerin numaraları.
      --notes string               Sipariş notu. Bin karakterden uzunu kırpılır.
      --payment string             paymentÖdeme kaynağı. Toplam sıfırın üstündeyse zorunludur. (JSON)
      --product-id int             Sipariş edilecek ürünün numarası.
      --quantity int               Adet. Yalnız ürün çoklu alıma izin veriyorsa dikkate alınır.
      --requirement-files string   Dosya isteyen alanların içeriği. Kodlanmış içerik olarak gönderilir; adres kabul edilmez. (JSON)
      --requirements string        Ürünün sorduğu alanların yanıtları. Zorunlu işaretlenen her alan gönderilmelidir. (JSON)
      --set stringArray            ek gövde alanı anahtar=değer (tekrarlanabilir)
      --terms                      Şartların kabulü. Doğru gönderilmelidir.
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

