# trkit

[![CI](https://github.com/sakiphan/trkit/actions/workflows/ci.yml/badge.svg)](https://github.com/sakiphan/trkit/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sakiphan/trkit.svg)](https://pkg.go.dev/github.com/sakiphan/trkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/sakiphan/trkit)](https://goreportcard.com/report/github.com/sakiphan/trkit)

Türkiye'ye özgü günlük işler için derli toplu, **sıfır harici bağımlılıklı** (yalnız stdlib) Go kütüphanesi. TC Kimlik / VKN doğrulama, IBAN, telefon, Türkçe-doğru metin, para ve tarih işlemleri tek import altında.

```bash
go get github.com/sakiphan/trkit
```

> Go 1.24+ • yalnızca standart kütüphane

## Neden trkit?

Bu işlerin her biri için ayrı ayrı küçük paketler var; ama dağınıklar, bir kısmı ekstra bağımlılık getiriyor ve Türkiye'ye özgü günlük işleri tek çatı altında toplamıyor. `trkit` bunları küratörlü, tek import'la kullanılan, `go.sum`'ı şişirmeyen tek bir kütüphanede toplar.

Örnek olarak metin tarafı: plain `strings.ToUpper("istanbul")` Türkçe'de `"ISTANBUL"` verir — yanlış. Standart kütüphane doğrusunu `strings.ToUpperSpecial(unicode.TurkishCase, s)` ile yapabilir ama bunu akılda tutmak zor. `trkit` doğrusunu **tek satırda, ezber gerektirmeden** verir; üstüne `Slug`, `Title` gibi stdlib'de olmayan yardımcıları ekler.

## Paketler

| Paket | İşlev |
|-------|-------|
| [`identity`](identity) | TC Kimlik No (TCKN) ve Vergi Kimlik No (VKN) doğrulama |
| [`iban`](iban) | TR IBAN doğrula / formatla / maskele / banka kodu |
| [`phone`](phone) | Telefon normalize (E.164) / doğrula / formatla / operatör |
| [`text`](text) | Türkçe-doğru Upper/Lower/Title + Slug |
| [`money`](money) | TL formatı + yazıyla tutar |
| [`date`](date) | Türkçe ay/gün isimleri + resmi tatil + iş günü |

## Kullanım

```go
package main

import (
	"fmt"
	"time"

	"github.com/sakiphan/trkit/date"
	"github.com/sakiphan/trkit/iban"
	"github.com/sakiphan/trkit/identity"
	"github.com/sakiphan/trkit/money"
	"github.com/sakiphan/trkit/phone"
	"github.com/sakiphan/trkit/text"
)

func main() {
	// identity
	fmt.Println(identity.ValidTCKN("10000000832")) // true
	fmt.Println(identity.ValidVKN("1234567899"))   // true

	// iban
	fmt.Println(iban.Valid("TR33 0006 1005 1978 6457 8413 26"))  // true
	fmt.Println(iban.Mask("TR330006100519786457841326"))         // TR33 **** **** **** **** **13 26

	// phone
	n, _ := phone.Normalize("0532 123 45 67")
	fmt.Println(n)                  // +905321234567
	fmt.Println(phone.Operator(n))  // Turkcell

	// text
	fmt.Println(text.Upper("istanbul"))          // İSTANBUL
	fmt.Println(text.Slug("Çağrı'nın İşi"))      // cagrinin-isi

	// money
	fmt.Println(money.FormatKurus(123456789))    // 1.234.567,89 ₺
	fmt.Println(money.Yaziyla(1234))             // bin iki yüz otuz dört lira

	// date
	t := time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println(date.Format(t))                  // 2 Ocak 2026 Cuma
	fmt.Println(date.IsGunuEkle(t, 1).Format("2006-01-02")) // 2026-01-05
}
```

## Bilinen sınırlar

- **Dini bayramlar** (Ramazan / Kurban) kameri takvime dayanır. `date` paketi 2020–2035 aralığını gömülü tabloyla kapsar; 2020–2026 resmi tarihler, 2027 sonrası Diyanet'in önceden hesaplanmış takvimidir ve nadiren bir gün kayabilir. Tablo 2035'te biter.
- **Operatör tahmini** numara taşıma nedeniyle %100 kesin değildir; yapısal doğrulama (`^5\d{9}$`) kesin, operatör eşlemesi "best effort"tur.

## Katkı

Issue ve PR'lar açıktır. Bağımlılık eklenmez (yalnız stdlib). Her değişiklik table-driven testlerle ve `go test ./...` temiz geçecek şekilde gelmelidir.
