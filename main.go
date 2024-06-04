// Ivan Orshak, 1321, 13/05/2024
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"zipper/pkg/bwt"
	"zipper/pkg/huffman"
	"zipper/pkg/lz77"
	"zipper/pkg/mtf"
	"zipper/pkg/rle"
)

// Zipper is an interface for all compression/decompression algs
type Zipper interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

func main() {
	// handle input flags
	task, err := Handler()
	if err != nil {
		log.Fatal(err)
	}

	// Start archiver
	if err = Archiver(task); err != nil {
		log.Fatal(err)
	}
}

// Archiver opens the input file, parses the input data,
// calls the archiving and unarchiving algorithms in the
// specified order, writes the result to the output file
func Archiver(task map[string]string) error {
	// main implementation
	var o Zipper
	// queue of de/compressing algs
	q := strings.Split(task["alg"], "-")

	// Input file with data
	fin, err := os.Open(task["fin"])
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return err
	}
	defer fin.Close()

	data, err := io.ReadAll(fin)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return err
	}

	// The main cycle, making one by one de/compression algs
	for _, k := range q {
		switch k {
		case "RLE", "rle":
			o = rle.Rle{}
		case "HA", "ha":
			o = &huffman.Ha{}
		case "LZ77", "lz77":
			o = lz77.Lz77{}
		case "BWT", "bwt":
			o = bwt.Bwt{}
		case "MTF", "mtf":
			o = mtf.Mft{}
		default:
			return errors.New(fmt.Sprintf("wrong de/compressinng method: %s", k))
		}

		switch task["mode"] {
		case "zip":
			tmp := len(data)
			t1 := time.Now()
			data, err = o.Encode(data) // TODO: make compressing status line bar
			t2 := time.Since(t1)

			if err != nil {
				log.Println("Error compressing data")
				return err
			}

			fmt.Printf("Execution %s algorithm time: %d ms\n", k, t2.Milliseconds())
			fmt.Printf("Compression ratio: %f\n", float64(tmp)/float64(len(data)))
		case "unzip":
			t1 := time.Now()
			data, err = o.Decode(data)
			t2 := time.Since(t1)

			if err != nil {
				log.Printf("Error decompressing data: %v\n", err)
				return err
			}

			fmt.Printf("Execution %s algorithm time: %d ms\n", k, t2.Milliseconds())
		}
	}

	// Writing output data to file
	if err = os.WriteFile(task["fout"], data, 0644); err != nil {
		log.Printf("Error writing to file: %v\n", err)
		return err
	}

	return nil
}

// Handler handles program сmd flags and writes it to the task map[string]string
func Handler() (map[string]string, error) {
	task := make(map[string]string)
	// cmd flags
	zip := flag.Bool("zip", false, "Compress data")
	unzip := flag.Bool("unzip", false, "Decompress data")
	fin := flag.String("f", "", "Input data file")
	alg := flag.String("with", "", "Compressing algorithm")

	flag.Parse()

	// Parsing logic
	if *fin == "" {
		flag.Usage()
		return nil, errors.New("input data file is required")
	} else {
		task["fin"] = *fin
		task["fout"] = *fin
	}

	if *alg == "" {
		flag.Usage()
		return nil, errors.New("algorithm is required")
	} else {
		task["alg"] = *alg
	}

	if *zip {
		task["mode"] = "zip"
		task["fout"] += "." + *alg
	} else if *unzip {
		task["mode"] = "unzip"

		lastIndex := strings.LastIndex(task["fout"], ".")
		if lastIndex != -1 {
			task["fout"] = task["fout"][:lastIndex]
		} else {
			task["fout"] = "unzipped.txt"
		}

	} else {
		flag.Usage()
		return nil, errors.New("archiver mode is required")
	}

	return task, nil
}
