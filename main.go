package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	flagType   = flag.String("t", "svg", "출력 형식(png, svg, txt, pdf)")
	flagOutput = flag.String("o", "", "출력 파일 경로(기본값: 입력 파일 위치)")
	flagServer = flag.String("s", "http://127.0.0.1:8080", "PlantUML 서버 주소")
	flagQuiet  = flag.Bool("q", false, "조용한 모드(로그 출력 안함)")
)

const (
	timeoutDuration = 10 * time.Second
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	inputPath := flag.Arg(0)

	if *flagOutput == "" {
		dir := filepath.Dir(inputPath)
		base := filepath.Base(inputPath)
		ext := filepath.Ext(base)
		name := base[:len(base)-len(ext)]
		*flagOutput = filepath.Join(dir, name+"."+*flagType)
	}

	in, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v\n", err)
	}
	defer in.Close()

	serverUrl := fmt.Sprintf("%s/%s", *flagServer, *flagType)
	if !*flagQuiet {
		log.Println("PlantUML Server: ", serverUrl)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, serverUrl, in)
	if err != nil {
		log.Fatalf("failed to create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("request failed with status: %s\n", res.Status)
	}

	out, err := os.Create(*flagOutput)
	if err != nil {
		log.Fatalf("failed to create output file: %v\n", err)
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatalf("failed to write output file: %v\n", err)
	}

	if !*flagQuiet {
		log.Println("completed successfully: ", *flagOutput)
	}
}
