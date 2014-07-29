package main

import (
	"log"
	"os"

	"github.com/musicbeat/stddata"
	"github.com/musicbeat/stddata/bank"
	"github.com/musicbeat/stddata/country"
	"github.com/musicbeat/stddata/currency"
	"github.com/musicbeat/stddata/language"
)

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
	if err := bankService.LoadProvider(new(bank.BankProvider), "bank"); err != nil {
		log.Printf("Error loading BankProvider: %v\n", err)
		return 1
	}
	log.Printf("Serving %d bank entities at /%s\n", bankService.Count, bankService.EntityName)

	countryService := new(stddata.Service)
	if err := countryService.LoadProvider(new(country.CountryProvider), "country"); err != nil {
		log.Printf("Error loading CountryProvider: %v\n", err)
		return 2
	}
	log.Printf("Serving %d country entities at /%s\n", countryService.Count, countryService.EntityName)

	currencyService := new(stddata.Service)
	if err := currencyService.LoadProvider(new(currency.CurrencyProvider), "currency"); err != nil {
		log.Printf("Error loading CurrencyProvider: %v\n", err)
		return 1
	}
	log.Printf("Serving %d currency entities at /%s\n", currencyService.Count, currencyService.EntityName)

	languageService := new(stddata.Service)
	if err := languageService.LoadProvider(new(language.LanguageProvider), "language"); err != nil {
		log.Printf("Error loading LanguageProvider: %v\n", err)
		return 1
	}
	log.Printf("Serving %d language entities at /%s\n", languageService.Count, languageService.EntityName)

	return 0
}
