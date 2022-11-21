# iso3166

> 3166-1 alpha-2, 3166-1 alpha-3, 3166-1 numeric

![Tests](https://github.com/verify-lab/iso3166/actions/workflows/tests.yml/badge.svg)

```sh
go get github.com/verify-lab/iso3166
```

## example

```go

if IsAlpha2CountryCode("pl") {
    ...
}

if IsAlpha3CountryCode("UKR") {
    ...
}

num, name, alpha2, alpha3 := CountryByCode("UKR")
```
