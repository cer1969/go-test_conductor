// Copyright Cristian Echeverría Rabí

package main

import (
    "fmt"
    "time"
    cx "bitbucket.org/tormundo/go.conductor"
)

var cu300 = cx.Conductor{cx.CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}

func main() {
	//cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
	cc, _ := cx.NewCurrentCalc(cu300)
	//cc := CurrentCalc{Conductor: cu300}	// No verifica ni inicializa valores

	tas := [5]float64{10, 15, 20, 25, 30}
	tcs := [5]float64{30, 35, 40, 45, 50}

	fmt.Println(cc.Name)
	fmt.Println("----------------")

	t1 := time.Now()

	for i, va := range tas {
		for j, vc := range tcs {
			current, _ := cc.Current(va, vc)
			fmt.Printf("i=%d, i=%d, Ta=%.2f, Tc=%.2f, I=%.2f\n", i, j, va, vc, current)
		}
	}

	fmt.Println(time.Now().Sub(t1))

	//r, err := cc.Resistance(2001)
}