package main

import (
	"fmt"

	"github.com/mflendrich/ci-experiment/mathutil"
	"github.com/mflendrich/ci-experiment/stringutil"
)

func main() {
	fmt.Println("mathutil.Add(2, 3) =", mathutil.Add(2, 3))
	fmt.Println("stringutil.Reverse(\"hello\") =", stringutil.Reverse("hello"))
}
