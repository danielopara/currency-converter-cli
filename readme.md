# Currency Converter

A simple command-line currency converter built with Go. Convert between different currencies using real-time exchange rates.

## Features

- Interactive terminal UI with dropdown menus
- Real-time exchange rates from OpenExchangeRates API
- Support for 9 major currencies including Nigerian Naira
- Input validation to prevent errors
- Multiple conversions in one session
- Clean, formatted output with exchange rate details

## Supported Currencies

- USD (US Dollar)
- EUR (Euro)
- GBP (British Pound)
- JPY (Japanese Yen)
- CAD (Canadian Dollar)
- AUD (Australian Dollar)
- CHF (Swiss Franc)
- CNY (Chinese Yuan)
- NGN (Nigerian Naira)

## Prerequisites

- Go 1.19 or higher
- OpenExchangeRates API key (free tier available)

## Installation

1. Clone this repository:

```bash
git clone <your-repo-url>
cd currency-converter
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create a `.env` file in the project root:

```
KEY=your_openexchangerates_api_key_here
```

4. Get your free API key from [OpenExchangeRates](https://openexchangerates.org/signup/free)

## Usage

Run the program:

```bash
go run main.go
```

The app will guide you through:

1. Selecting source currency
2. Selecting target currency
3. Entering amount to convert
4. Viewing the conversion result
5. Option to do another conversion

## Example Output

```
Convert from: US Dollar
Convert to: Nigerian Naira
Amount: 100

100.00 USD = 164750.00 NGN
Exchange rate: 1 USD = 1647.5000 NGN

Do you want to perform another conversion? Yes
```

## Dependencies

- [huh](https://github.com/charmbracelet/huh) - Terminal forms and prompts
- [godotenv](https://github.com/joho/godotenv) - Load environment variables from .env file

## Error Handling

The app handles common errors like:

- Missing or invalid API key
- Network connectivity issues
- Invalid number input
- Same currency conversions
- Unsupported currencies
