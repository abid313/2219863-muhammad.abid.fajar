package main

import "fmt"

func main() {
	/*
		Given total payment and price, return the excess payment or -1 if payment is less than the final price
		Give 5% discount if total price is >= 100000
		Ex: payment: 5000, price 1000 -> result 4000
			payment: 2000, price 5000 -> result -1
			payment: 100000, price 100000 -> result 5000
	*/
	res := CalculateChange(5000, 1000)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := CalculateChangeCorrect(arr)
	// fmt.Println(resCorrect)
}

func CalculateChange(totalPayment, totalPrice float64) float64 {
	if totalPayment < totalPrice {
		return -1
	}

	minimumPriceForDiscount := float64(100000)
	if totalPrice >= minimumPriceForDiscount {
		// Give 5% Discount rate
		discountRate := 0.5
		totalPrice = totalPrice * discountRate
	}

	change := totalPayment - totalPrice

	return change
}

func CalculateChangeCorrect(totalPayment, totalPrice float64) float64 {
<<<<<<< HEAD
	if totalPayment < totalPrice && totalPrice != 100000 {
		return -1
	}

	minimumPriceForDiscount := float64(100000)
	if totalPrice >= minimumPriceForDiscount {
		// Give 5% Discount rate
		f := totalPrice * 0.05
		totalPrice -= f
	}
	change := totalPayment - totalPrice
	if change < 0 {
		return -1
	}
	return change
=======
	return 0 // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
