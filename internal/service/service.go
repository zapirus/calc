package service

import (
	romans "github.com/summed/goromans"
	"log"
	"strconv"
)

func Validate(slc []string) bool {
	var (
		rm, arab int
		bl       bool
		operand  string
	)

	rims := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabic := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	operation := []string{"+", "-", "/", "*"}

	if len(slc) < 3 {
		log.Fatalln("Вывод ошибки, так как строка не является математической операцией.")
		return false
	} else if len(slc) > 3 {
		log.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return false
	}
	for _, s := range operation {
		if slc[1] == s {
			for i, _ := range slc {
				for i2, _ := range rims {
					if slc[i] == rims[i2] {
						rm += 1
					}
				}
			}
			for i, _ := range slc {
				for i2, _ := range arabic {
					if slc[i] == arabic[i2] {
						arab += 1
					}
				}
			}
			operand += s
		}
	}

	if rm == 2 || arab == 2 && operand != "" {
		bl = true
	} else {
		log.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return false
	}
	return bl

}

func Arithmetic(slcInput []string) string {
	rims := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabic := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for _, rim := range rims {
		if slcInput[0] == rim {
			if slcInput[1] == "+" {
				num1, err := romans.RtoA(slcInput[0])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				num2, err := romans.RtoA(slcInput[2])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				sum := num1 + num2
				return romans.AtoR(sum)
			} else if slcInput[1] == "-" {
				num1, err := romans.RtoA(slcInput[0])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				num2, err := romans.RtoA(slcInput[2])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}

				sum := int(num1) - int(num2)
				if sum < 0 {
					return "Рим не может быть < 0"
				}

				return romans.AtoR(uint(sum))

			} else if slcInput[1] == "*" {
				num1, err := romans.RtoA(slcInput[0])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				num2, err := romans.RtoA(slcInput[2])

				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				sum := num1 * num2
				return romans.AtoR(sum)
			} else if slcInput[1] == "/" {
				num1, err := romans.RtoA(slcInput[0])
				if err != nil {
					log.Fatalln("Ошибка конвертации")
				}
				num2, err := romans.RtoA(slcInput[2])
				sum := num1 / num2
				return romans.AtoR(sum)
			}
		} else {
			for _, arab := range arabic {
				if slcInput[0] == arab {
					if slcInput[1] == "-" {
						res, err := strconv.Atoi(slcInput[0])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						res2, err := strconv.Atoi(slcInput[2])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						sum := res - res2
						return strconv.Itoa(sum)
					} else if slcInput[1] == "+" {
						res, err := strconv.Atoi(slcInput[0])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						res2, err := strconv.Atoi(slcInput[2])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						sum := res + res2
						return strconv.Itoa(sum)
					} else if slcInput[1] == "*" {
						res, err := strconv.Atoi(slcInput[0])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						res2, err := strconv.Atoi(slcInput[2])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						sum := res * res2
						return strconv.Itoa(sum)
					} else if slcInput[1] == "/" {
						res, err := strconv.Atoi(slcInput[0])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						res2, err := strconv.Atoi(slcInput[2])
						if err != nil {
							log.Fatalln("Ошибка конвертации")
						}
						sum := res / res2
						return strconv.Itoa(sum)
					}

				}

			}
		}

	}
	return "Что-то пошло не так"

}
