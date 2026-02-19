package generators

import (
    "fmt"
    "math/rand"
    "time"
)

func GenerateCreditCard(holderName string) (string, string, string, string) {
    expiry := time.Now().AddDate(10, 0, 0)
    expiryDate := fmt.Sprintf("%02d/%02d", expiry.Month(), expiry.Year()%100)

    cardNumber := fmt.Sprintf("%04d-%04d-%04d-%04d",
        rand.Intn(9000)+1000,
        rand.Intn(9000)+1000,
        rand.Intn(9000)+1000,
        rand.Intn(9000)+1000,
    )

    cvv := fmt.Sprintf("%03d", rand.Intn(900)+100)
    return holderName, cardNumber, cvv, expiryDate
}