package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
)

// ExchangeRates represents the JSON structure returned by the exchange rate API
type ExchangeRates struct {
    Rates map[string]float64 `json:"rates"` // Map of currency codes to their exchange rates
    Base  string             `json:"base"`  // Base currency (usually USD for this API)
}

// supportBase defines the currencies we support (though not currently used in the original code)
var supportBase = []string{"USD", "EUR", "GBP", "JPY"}

// validateFloat checks if the input string can be converted to a valid float64
func validateFloat(input string) error {
    // Try to parse the string as a float64
    if _, err := strconv.ParseFloat(input, 64); err != nil {
        return fmt.Errorf("Please enter a valid number")
    }
    return nil // Return nil if parsing succeeds
}

func fetchAmount(apiURL string) (*ExchangeRates, error) {
    // Make HTTP GET request to the API
    resp, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close() 
    
    // Create a variable to hold the decoded JSON data
    var data ExchangeRates
    
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return nil, err
    }
    
    return &data, nil
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("env file does not exist")
        return
    }
    
    // Get API key from environment variable
    apiKey := os.Getenv("KEY")
    if apiKey == "" {
        fmt.Println("no key")
        return
    }
    
    // Construct API URL with the API key
    apiURL := "https://openexchangerates.org/api/latest.json?app_id=" + apiKey
    
    // Variables to store form input values
    var currencyFrom, currencyTo, amountStr string
    var shouldContinue bool = true
    
    // Main application loop
    for shouldContinue {
        // Create the form with improved options and validation
        form := huh.NewForm(
            huh.NewGroup(
                huh.NewSelect[string]().
                    Title("Convert from:").
                    Description("Select the currency you want to convert from").
                    Options(
                        huh.NewOption("US Dollar", "USD"),
                        huh.NewOption("Euro", "EUR"),
                        huh.NewOption("British Pound", "GBP"),
                        huh.NewOption("Japanese Yen", "JPY"),
                        huh.NewOption("Canadian Dollar", "CAD"),
                        huh.NewOption("Australian Dollar", "AUD"),
                        huh.NewOption("Swiss Franc", "CHF"),
                        huh.NewOption("Chinese Yuan", "CNY"),
						huh.NewOption("Nigeria Naira", "NGN"),
                    ).
                    Value(&currencyFrom),
                
                // Select dropdown for target currency
                huh.NewSelect[string]().
                    Title("Convert to:").
                    Description("Select the currency you want to convert to").
                    Options(
                        huh.NewOption("US Dollar", "USD"),
                        huh.NewOption("Euro", "EUR"),
                        huh.NewOption("British Pound", "GBP"),
                        huh.NewOption("Japanese Yen", "JPY"),
                        huh.NewOption("Canadian Dollar", "CAD"),
                        huh.NewOption("Australian Dollar", "AUD"),
                        huh.NewOption("Swiss Franc", "CHF"),
                        huh.NewOption("Chinese Yuan", "CNY"),
						huh.NewOption("Nigeria Naira", "NGN"),

                    ).
                    Value(&currencyTo),
                
                // Input field for amount with validation
                huh.NewInput().
                    Title("Amount:").
                    Description("Enter the amount you want to convert").
                    Placeholder("e.g., 100.50").
                    Validate(validateFloat).
                    Value(&amountStr),
            ),
        )
        
        // Run the form and handle errors
        if err := form.Run(); err != nil {
            fmt.Println("Error running form:", err)
            return
        }
        
        // Convert string amount to float64 (we know it's valid due to validation)
        amount, _ := strconv.ParseFloat(amountStr, 64)
        
        // Check if user is trying to convert same currency
        if currencyFrom == currencyTo {
            fmt.Printf("\n%.2f %s = %.2f %s (same currency)\n", amount, currencyFrom, amount, currencyTo)
        } else {
            // Fetch exchange rate data from API
            data, err := fetchAmount(apiURL)
            if err != nil {
                fmt.Println("Error fetching exchange rates:", err)
                return
            }
            
            // Get exchange rates for both currencies
            rateFrom := data.Rates[currencyFrom]
            rateTo := data.Rates[currencyTo]
            
            // Check if currencies exist in the API response
            if rateFrom == 0 || rateTo == 0 {
                fmt.Println("Error: One or both currencies not found in API response")
                continue
            }
            
            // Calculate conversion: amount * (target_rate / source_rate)
            converted := amount * (rateTo / rateFrom)
            
            // Display the result with proper formatting
            fmt.Printf("\n%.2f %s = %.2f %s\n", amount, currencyFrom, converted, currencyTo)
            fmt.Printf("Exchange rate: 1 %s = %.4f %s\n", currencyFrom, rateTo/rateFrom, currencyTo)
        }
        
        // Ask if user wants to perform another conversion
        continueForm := huh.NewForm(
            huh.NewGroup(
                huh.NewConfirm().
                    Title("Do you want to perform another conversion?").
                    Value(&shouldContinue),
            ),
        )
        
        if err := continueForm.Run(); err != nil {
            fmt.Println("Error running continue form:", err)
            return
        }
        
        fmt.Println()
    }
    
    fmt.Println("Thank you for using the currency converter!")
}