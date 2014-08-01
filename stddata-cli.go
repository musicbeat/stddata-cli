// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package main implements the command line interface used to load and 
serve the data set providers. The server will listen on port 6060.

So far, no options or config file. But a future version may have
these.

Install golang for your development architecture (see golang.org).

Usage:
	go get github.com/musicbeat/stddata
	go get github.com/musicbeat/stddata-cli
	go install
	$GOPATH/bin/stddata-cli

Use an http client, say curl, to interact with the server:
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
	curl -v http://localhost:6060/language?alpha=_dump
	... et cetera ...

*/
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/musicbeat/stddata"
	"github.com/musicbeat/stddata/bank"
	"github.com/musicbeat/stddata/country"
	"github.com/musicbeat/stddata/currency"
	"github.com/musicbeat/stddata/language"
)

// TODO: config and command flag:
var port string = ":6060"

func main() {
	// os.Exit will terminate the program at the place of call without running
	// any deferred cleanup statements. It might cause unintended effects. To
	// be safe, we wrap the program in run() and only os.Exit() outside the
	// wrapper. Be careful not to indirectly trigger os.Exit() in the program,
	// notably via log.Fatal() and on flag.Parse() where the default behavior
	// is ExitOnError.
	os.Exit(run())
}

// Run the program and return exit code.
func run() int {
	bankService := new(stddata.Service)
	if err := bankService.LoadProvider(new(bank.BankProvider), "/bank"); err != nil {
		log.Printf("Error loading BankProvider: %v\n", err)
		return 1
	}
	http.HandleFunc(bankService.EntityName, bankService.ServeHTTP)
	log.Printf("Serving %d bank entities at %s\n", bankService.Count, bankService.EntityName)

	countryService := new(stddata.Service)
	if err := countryService.LoadProvider(new(country.CountryProvider), "/country"); err != nil {
		log.Printf("Error loading CountryProvider: %v\n", err)
		return 2
	}
	http.HandleFunc(countryService.EntityName, countryService.ServeHTTP)
	log.Printf("Serving %d country entities at %s\n", countryService.Count, countryService.EntityName)

	currencyService := new(stddata.Service)
	if err := currencyService.LoadProvider(new(currency.CurrencyProvider), "/currency"); err != nil {
		log.Printf("Error loading CurrencyProvider: %v\n", err)
		return 3
	}
	http.HandleFunc(currencyService.EntityName, currencyService.ServeHTTP)
	log.Printf("Serving %d currency entities at %s\n", currencyService.Count, currencyService.EntityName)

	languageService := new(stddata.Service)
	if err := languageService.LoadProvider(new(language.LanguageProvider), "/language"); err != nil {
		log.Printf("Error loading LanguageProvider: %v\n", err)
		return 4
	}
	http.HandleFunc(languageService.EntityName, languageService.ServeHTTP)
	log.Printf("Serving %d language entities at %s\n", languageService.Count, languageService.EntityName)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return 0
}
