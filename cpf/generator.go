package cpf

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brazilian-utils/go-brazilian-utils/helpers"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomState() state {
	return stateCodes[rand.Intn(len(stateCodes))]
}

// Generate create a random valid CPF
func Generate() string {
	var result strings.Builder

	for i := 0; i < size-3; i++ {
		result.WriteString(strconv.Itoa(rand.Intn(10)))
	}

	result.WriteString(strconv.Itoa(rand.Intn(10)))
	result.WriteString(strconv.Itoa(helpers.Mod11(helpers.GenerateChecksum(result.String(), []int{10}))))
	result.WriteString(strconv.Itoa(helpers.Mod11(helpers.GenerateChecksum(result.String(), []int{11}))))

	return result.String()
}
