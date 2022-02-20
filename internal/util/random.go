package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SajjadManafi/simple-uber/models"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random int in the range [min, max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(length int) string {
	k := len(alphabet)
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(k)]
	}
	return string(b)
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomGender generates a random gender
func RandomGender() models.Gender {
	Genders := []models.Gender{models.Male, models.Female}
	Gender := Genders[rand.Intn(len(Genders))]
	return Gender
}

// RandomCabBrand generates a random cab brand
func RandomCabBrand() string {
	Brands := []string{"Audi", "BMW", "Dodge", "Ford", "Honda", "Mazda", "Maserati", "Mercedes-Benz", "Nissan", "Suzuki", "Tesla", "Toyota"}
	brand := Brands[rand.Intn(len(Brands))]
	return brand
}

// RandomCabModel generates a random cab model
func RandomCabModel() string {
	return fmt.Sprintf("%s-%d", RandomString(3), RandomInt(1, 9))
}

// RandomCabColor generates a random cab colorس
func RandomCabColor() string {
	Colors := []string{"Black", "Blue", "Brown", "Green", "Grey", "Orange", "Pink", "Purple", "Red", "Silver", "White", "Yellow"}
	color := Colors[rand.Intn(len(Colors))]
	return color
}

// RandomCabPlate generates a random cab plate
func RandomCabPlate() string {
	return fmt.Sprintf("%d%s%d-%d", RandomInt(10, 99), RandomString(1), RandomInt(100, 999), RandomInt(10, 99))
}

// RandomDriverRating generates a random driver rating
func RandomDriverRating() int32 {
	return int32(RandomInt(1, 5))
}
