package learn

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func Logging() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with fine/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	mySlog := slog.New(jsonHandler)
	mySlog.Info("hsdasd")
	mySlog.Info("heasdasd", "key", "val", "age", 25)
}
