# [stddata-cli](https://github.com/musicbeat/stddata-cli)

The [stddata-cli](https://github.com/musicbeat/stddata-cli) is the command line "main" package for a set of standard data providers. An [stddata](https://github.com/musicbeat/stddata) data provider loads data from a standards provider, indexes it, and serves it over HTTP with REST interface, delivering json results to the client.

There are currently data providers for:

 * Banks - the Fed's ACH directory
 * Countries - ISO 3166-1
 * Currencies - ISO 4217
 * Languages - ISO 639.2

## Getting Started

So far, no options or config file. But a future version may have
these.

Install golang for your development architecture (see [golang.org](http://golang.org)).

## Installation and Launch
```
	go get github.com/musicbeat/stddata
	go get github.com/musicbeat/stddata-cli
	go install
	$GOPATH/bin/stddata-cli
```
## Usage
Use an http client, say curl, to interact with the server. Request URLs are in the form:
```
    http://server:port/entity?index=query
```
Here are the currently included entities and their indexes:
```
  Entity      | Indexes | Notes
  --------- | ----- | --------
  bank  | number | ACH - ABA Routing number
  bank  | name   | ACH Customer name
  country | name | English name
  country | alpha2 | Alpha2 code
  country | alpha3 | Alpha3 code
  country | number | Numeric code
  currency | country | Country name
  currency | name | Currency name
  currency | code | Currency code (alphabetic)
  currency | number | Currency code (numeric)
  language | alpha | Alpha code
  language | name | English name
```

### Examples
```
curl -v http://localhost:6060/bank?name=X
curl -v http://localhost:6060/bank?number=271972899
curl -v http://localhost:6060/country?name=argentina
curl -v http://localhost:6060/country?alpha2=AR
curl -v http://localhost:6060/country?alpha3=USA
curl -v http://localhost:6060/country?number=800
curl -v http://localhost:6060/currency?name=Euro
curl -v http://localhost:6060/currency?code=USD
curl -v http://localhost:6060/currency?number=840
curl -v http://localhost:6060/language?name=Eng
curl -v http://localhost:6060/language?alpha=fr
	... et cetera ...
```

Also note that the `reserved` word `_dump` can be used as a query to retrieve all of the entities in the order of the index.

```
curl -v http://localhost:6060/bank?name=_dump
curl -v http://localhost:6060/bank?number=_dump
curl -v http://localhost:6060/country?name=_dump
curl -v http://localhost:6060/country?alpha2=_dump
curl -v http://localhost:6060/country?alpha3=_dump
curl -v http://localhost:6060/country?number=_dump
curl -v http://localhost:6060/currency?name=_dump
curl -v http://localhost:6060/currency?code=_dump
curl -v http://localhost:6060/currency?number=_dump
curl -v http://localhost:6060/language?name=_dump
curl -v http://localhost:6060/language?alpha=_dump
```
