package class1

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

//ControlProgramWithArgs Controls Programs flag inputs
//Note: flaga -h automatycznie wypisze wszystkie
func ControlProgramWithArgs() {
	helperMsg := "Sums all passed args"
	sumMode := flag.Bool("sum", false, helperMsg)
	powerMode := flag.Int("power", -1, "Podnosi do potęgi podany arg")
	flag.Parse()

	//Jezeli zmienna jest pointerem to aby pobrać wartośc dodajemy *przy nazwie zmiennej
	fmt.Printf("Power ma wartosc %v\n", *powerMode)

	if *sumMode {
		fmt.Println("Działam")
		w := sumElements(flag.Args())
		fmt.Println(w)
	}

	if *powerMode > 1 {
		a := powerArg(flag.Args(), *powerMode)
		fmt.Println(a)
	}
}

func sumElements(elements []string) int {
	fmt.Println("Summing elements")
	sum := 0

	for _, el := range elements {
		number, err := strconv.Atoi(el)

		if err != nil {
			log.Fatalf("Unable to convert element '%v' to int. Ended with Error %v", el, err)
		}

		sum += number
	}
	return sum
}

func powerArg(args []string, power int) int {
	//Najpierw spr długość slica czy jest wiekszy od 0
	if len(args) < 1 {
		//log.Fatalf konczy działanie całego programu
		log.Fatalf("Nie podałeś argumentu")
	}
	argument := args[0]
	number, err := strconv.Atoi(argument)

	if err != nil {
		log.Fatalf("Nie udalo się zrzutować argumentu: %v na int", err)
	}

	return int(math.Pow(float64(number), float64(power)))
}
