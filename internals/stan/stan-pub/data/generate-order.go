package data

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
	"tools/internals/models"

	"github.com/google/uuid"
	"github.com/haydenwoodhead/burner.kiwi/emailgenerator"
	"github.com/sirupsen/logrus"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

var currency = []string{
	"USD",
	"EUR",
	"GBP",
	"RUB",
}
var providers = []string{
	"wbpay",
	"sberpay",
	"sbp",
}

var banks = []string{
	"alpha",
	"sberbank",
	"vtb",
	"gasprom",
}

func GenerateOrder() *models.Orders {
	order := models.Orders{
		OrderID:           GenerateRandomUUID(),
		TrackNumber:       GenerateRandomString(64),
		Entry:             GenerateRandomString(32),
		Delivery:          GenerateDelivery(),
		Payment:           GeneratePayment(),
		Locale:            GenerateRandomString(32),
		IntersanSignature: GenerateRandomString(16),
		CustomerID:        GenerateRandomString(64),
		DeliveryService:   GenerateRandomString(32),
		Shardkey:          GenerateRandomString(64),
		SmID:              GenerateRandomBigInt(),
		DateCreated:       GenerateRandomTime(),
		OofShared:         GenerateRandomString(32),
	}
	return &order
}

func GenerateDelivery() models.Delivery {
	delivery := models.Delivery{
		Name:    GenerateRandomString(16),
		Phone:   GenerateRandomPhoneNumber(),
		Zip:     GenerateRandomZIP(),
		City:    GenerateRandomString(16),
		Address: GenerateRandomString(32),
		Region:  GenerateRandomString(32),
		Email:   GenerateRandomEmail(),
	}
	return delivery
}

func GeneratePayment() models.Payment {
	payment := models.Payment{
		Transaction:  GenerateRandomString(64),
		RequestID:    GenerateRandomString(5),
		Currency:     PickRandomElement(currency),
		Provider:     PickRandomElement(providers),
		Amount:       GenerateRandomInt(0),
		PaymentDT:    GenerateRandomBigInt(),
		Bank:         PickRandomElement(banks),
		DeliveryCost: GenerateRandomInt(0),
		GoodsTotal:   GenerateRandomInt(0),
		CustomFee:    GenerateRandomInt(0),
	}
	return payment
}

func GenerateItem() *models.Items {
	item := models.Items{
		ChrtID:      GenerateRandomBigInt(),
		TrackNumber: GenerateRandomString(64),
		Price:       GenerateRandomInt(0),
		Rid:         GenerateRandomString(64),
		Name:        GenerateRandomString(64),
		Sale:        GenerateRandomInt(100),
		Size:        GenerateRandomString(64),
		NmID:        GenerateRandomBigInt(),
		Brand:       GenerateRandomString(16),
		Status:      GenerateRandomInt(0),
	}
	item.TotalPrice = item.Price - (item.Price * item.Sale / 100)
	return &item
}

func GenerateRandomString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
func GenerateRandomPhoneNumber() string {

	areaCode := rand.Intn(900) + 100
	prefix := rand.Intn(900) + 100
	lineNumber := rand.Intn(9000) + 1000

	return fmt.Sprintf("(%d) %d-%04d", areaCode, prefix, lineNumber)
}

func GenerateRandomZIP() string {
	return fmt.Sprintf("%05d", rand.Intn(90000)+10000)
}

func GenerateRandomEmail() string {
	hosts := []string{"gmail.com", "yahoo.com", "hotmail.com"}
	minLength := 5

	eg := emailgenerator.New(hosts, minLength)

	randomEmail := eg.NewRandom()
	return randomEmail
}

func GenerateRandomBigInt() big.Int {
	var max *big.Int = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	n, err := crand.Int(crand.Reader, max)
	if err != nil {
		logrus.Errorf("Could not generate bigint: %v\n", err)
	}
	return *n
}
func GenerateRandomInt(max int) int {
	if max == 0 {
		return rand.Int()
	}
	return rand.Intn(max)
}

func GenerateRandomUUID() uuid.UUID {
	return uuid.New()
}

func GenerateRandomTime() time.Time {
	return time.Now()
}

func PickRandomElement(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}
