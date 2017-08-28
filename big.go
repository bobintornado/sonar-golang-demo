package main

import (
    "math/big"
)

func main() {
    z := new(big.Int)
    x := new(big.Int)
    x = x.SetUint64(2)
    y := new(big.Int)
    y = y.SetUint64(4)
    m := new(big.Int)
    m = m.SetUint64(0)

    z = z.Exp(x, y, m)
}
