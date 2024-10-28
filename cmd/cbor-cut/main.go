package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	cc "github.com/takanoriyanagitani/go-cbor-cut/cut"

	ci "github.com/takanoriyanagitani/go-cbor-cut/iter/cbor2iter"
	ia "github.com/takanoriyanagitani/go-cbor-cut/iter/cbor2iter/amacker"

	ac "github.com/takanoriyanagitani/go-cbor-cut/iter/iter2cbor"
	aa "github.com/takanoriyanagitani/go-cbor-cut/iter/iter2cbor/amacker"

	ss "github.com/takanoriyanagitani/go-cbor-cut/field/select/str2sel"
)

func sub(ctx context.Context, fields string, rdr io.Reader, wtr io.Writer) error {
	selector, e := ss.StringToSelectDefault(fields)
	if nil != e {
		return e
	}

	var arr2sel cc.ArrayToSelected = selector.ToArrayToSelected()

	var c2a ci.CborToArrays = ia.CborToArraysNew(rdr).AsCborToArrays()
	var a2c ac.ArrayToCbor = aa.ArrToCborNew(wtr).ToArrayToCbor()

	osel := cc.OutputSelected{
		ArrayToSelected: arr2sel,
		ArraySource:     c2a.ToArraySource(),
		ArrayOutput:     a2c.AsArrayOutput(),
	}
	return osel.ConvertAll(ctx)
}

func stdin2stdout(ctx context.Context, fields string) error {
	var r io.Reader = os.Stdin
	var br io.Reader = bufio.NewReader(r)

	var w io.Writer = os.Stdout
	var bw *bufio.Writer = bufio.NewWriter(w)
	defer bw.Flush()

	return sub(ctx, fields, br, bw)
}

func fromEnv(ctx context.Context, envKey string) error {
	var fields string = os.Getenv(envKey)
	return stdin2stdout(ctx, fields)
}

func main() {
	e := fromEnv(context.Background(), "ENV_FIELDS")
	if nil != e {
		log.Printf("%v\n", e)
	}
}
