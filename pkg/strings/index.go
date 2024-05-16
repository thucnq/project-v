package strings

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// Generate random bytes
func RandomBytes(size int) (rb []byte, err error) {
	rb = make([]byte, size)
	_, err = rand.Read(rb)
	return
}

// Generate random string
func RandomString(length int) string {
	result := ""
	for len(result) < length {
		size := length - len(result)
		randBytes, _ := RandomBytes(size)
		encoded := base64.StdEncoding.EncodeToString(randBytes)
		encoded = strings.ReplaceAll(encoded, "/", "")
		encoded = strings.ReplaceAll(encoded, "+", "")
		encoded = strings.ReplaceAll(encoded, "=", "")
		encoded = encoded[0:size]
		result = result + encoded
	}

	return result
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func FormatMoney(money int64) string {
	if money <= 0 {
		return fmt.Sprintf("%d", money)
	}

	var stack []string

	for money > 0 {
		stack = append(stack, fmt.Sprintf("%3d", money%1000))
		money /= 1000
	}

	reverse(stack)
	t := 3
	if len(stack) > 0 {
		for i, char := range stack[0] {
			if char != ' ' {
				t = i
				break
			}
		}
	}
	stack[0] = stack[0][t:3]
	res := strings.Replace(strings.Join(stack, "."), " ", "0", -1)
	if len(res) == 5 {
		res = strings.Replace(res, ".", "", -1)
	}
	return res
}

func FormatShortMoney(money int64, currency string) string {
	thousandStr := "K"
	millionStr := "M"
	if money >= 1000000 {
		return fmt.Sprintf("%s%s", strconv.FormatFloat(float64(money)/1000000, 'f', -1, 64), millionStr)
	}
	if money >= 1000 {
		return fmt.Sprintf("%d%s", money/1000, thousandStr)
	}
	return fmt.Sprintf("%d", money)
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}

func Capitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}
