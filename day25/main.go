package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
)

type snafu string

func parseInput(input []string) []snafu {
    var result []snafu


    for _, line := range input {
        result = append(result, snafu(line))
    }
    return result
}

func pow(n int, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}

func snafuToDecimal(input snafu) int {
    result := 0

    for i := len(input) - 1; i >= 0; i-- {
        power := len(input) - 1 - i

        switch input[i] {
        case '2':
            result += 2 * pow(5, power)
        case '1':
            result += pow(5, power)
        case '0':
            continue
        case '-':
            result -= pow(5, power)
        case '=':
            result -= 2 * pow(5, power)
        default:
            panic(input[i])
        }
    }
    return result
}

func decimalToSnafu(input int) snafu {

    var result snafu

    for input > 0 {
        mod := input % 5

        switch mod {
        case 0:
            result = "0" + result
            input /= 5
        case 1:
            result = "1" + result
            input /= 5
        case 2:
            result = "2" + result
            input /= 5
        case 3:
            result = "=" + result
            input = input / 5 + 1
        case 4:
            result = "-" + result
            input = input / 5 + 1
        default:
            panic(mod)
        }
    }

    return result
}
func solve(input []snafu) snafu {


    sum := 0

    for _, s := range input {
        sum += snafuToDecimal(s)
    }


    return decimalToSnafu(sum)
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

    result := solve(parsedInput)

	fmt.Printf("Result: %s\r\n", result)
}
