package services

import (
    "encoding/json"
    "errors"
    "net/http"
   // "paypal-clone/models"
    "time"
)

func GetExchangeRate(baseCurrency string, targetCurrency string) (float64, error) {
    cacheKey := "exchange_rate_" + baseCurrency + "_" + targetCurrency
    var cachedRate float64

    if err := GetCache(cacheKey, &cachedRate); err == nil {
        return cachedRate, nil
    }

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

    // Кэширование обменного курса на 1 час
    if err := SetCache(cacheKey, rate, time.Hour); err != nil {
        return 0, err
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
