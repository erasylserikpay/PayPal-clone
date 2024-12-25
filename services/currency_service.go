package services

import (
    "encoding/json"
    "net/http"
    "errors"
)

func GetExchangeRate(baseCurrency string, targetCurrency string) (float64, error) {
    // Пример запроса к публичному API обмена валют (например, ExchangeRatesAPI или Fixer)
    apiURL := "https://api.exchangerate-api.com/v4/latest/" + baseCurrency

    resp, err := http.Get(apiURL)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    rates, ok := result["rates"].(map[string]interface{})
    if !ok {
        return 0, errors.New("invalid response structure")
    }

    rate, ok := rates[targetCurrency].(float64)
    if !ok {
        return 0, errors.New("rate not found for target currency")
    }

    return rate, nil
}

func ConvertCurrency(amount float64, baseCurrency string, targetCurrency string) (float64, error) {
    rate, err := GetExchangeRate(baseCurrency, targetCurrency)
    if err != nil {
        return 0, err
    }

    return amount * rate, nil
}
