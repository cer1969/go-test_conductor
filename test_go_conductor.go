// Copyright Cristian Echeverría Rabí

package main

import (
	"fmt"
	"time"

	cx "bitbucket.org/tormundo/go.conductor"
)

//var cu300 = cx.Conductor{cx.CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}

func main() {

	cuArgs := cx.ConductorArgs{"CU 300 MCM", cx.CC_CU, 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
	fmt.Printf("%T\n", cuArgs)
	fmt.Printf("%v\n", cuArgs)
	fmt.Printf("%v\n", cuArgs.Category.Name())

	cu300 := cuArgs.Get()

	cc, _ := cx.NewCurrentCalc(cu300)
	//cc = &cx.CurrentCalc{0.12270, 15.95, 0.00374, 300.0, 2.0, 1.0, 0.5, cx.CF_IEEE, 0.01}
	//cc := CurrentCalc{Conductor: cu300}	// No verifica ni inicializa valores

	fmt.Printf("%T\n", cc)
	fmt.Printf("%v\n", cc.Conductor())

	tas := [5]float64{10, 15, 20, 25, 30}
	tcs := [5]float64{30, 35, 40, 45, 50}

	fmt.Println(cu300.Name)
	fmt.Println("----------------")

	t1 := time.Now()

	for i, va := range tas {
		for j, vc := range tcs {
			current, _ := cc.Current(va, vc)
			fmt.Printf("i=%d, i=%d, Ta=%.2f, Tc=%.2f, I=%.2f\n", i, j, va, vc, current)
		}
	}

	tc, _ := cc.Tc(30, 171.14)
	fmt.Printf("%.3f\n", tc)
	ta, _ := cc.Ta(40, 171.14)
	fmt.Printf("%.3f\n", ta)

	fmt.Printf("%v\n", time.Now().Sub(t1))

	//r, err := cc.Resistance(2001)
}
