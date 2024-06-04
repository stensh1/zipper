.PHONY:
.SILENT:


# Compiling
build:
	go build -o ./.bin/zipper main.go

# Emulating users requests
##########################
# Huffman only
zip_ha: build
	./.bin/zipper -with "HA" -f "./src/text/data.in" -zip

unzip_ha: build
	./.bin/zipper -with "HA" -f "./src/text/data.in.HA" -unzip

##########################
# RLE only
zip_rle: build
	./.bin/zipper -with "RLE" -f "./src/text/data.in" -zip

unzip_rle: build
	./.bin/zipper -with "RLE" -f "./src/text/data.in.RLE" -unzip

##########################
# BWT only
zip_bwt: build
	./.bin/zipper -with "BWT" -f "./src/text/data.in" -zip

unzip_bwt: build
	./.bin/zipper -with "BWT" -f "./src/text/data.in.BWT" -unzip

##########################
# MTF only
zip_mtf: build
	./.bin/zipper -with "MTF" -f "./src/text/data.in" -zip

unzip_mtf: build
	./.bin/zipper -with "MTF" -f "./src/text/data.in.MTF" -unzip

##########################
# LZ77 only
zip_lz77: build
	./.bin/zipper -with "LZ77" -f "./src/text/data.in" -zip

unzip_lz77: build
	./.bin/zipper -with "LZ77" -f "./src/text/data.in.LZ77" -unzip

##########################
# Combos
##########################
# BWT+RLE
zip_bwt_rle: build
	./.bin/zipper -with "BWT-RLE" -f "./src/text/data.in" -zip

unzip_rle_bwt: build
	./.bin/zipper -with "RLE-BWT" -f "./src/text/data.in.BWT-RLE" -unzip

##########################
# BWT+MTF+Huffman
zip_bwt_mtf_ha: build
	./.bin/zipper -with "BWT-MTF-HA" -f "./src/text/data.in" -zip

unzip_ha_mtf_bwt: build
	./.bin/zipper -with "HA-MTF-BWT" -f "./src/text/data.in.BWT-MTF-HA" -unzip

##########################
# BWT+MTF+RLE+Huffman
zip_bwt_mtf_rle_ha: build
	./.bin/zipper -with "BWT-MTF-RLE-HA" -f "./src/text/data.in" -zip

unzip_ha_rle_mtf_bwt: build
	./.bin/zipper -with "HA-RLE-MTF-BWT" -f "./src/text/data.in.BWT-MTF-RLE-HA" -unzip

##########################
# LZ77+Huffman
zip_lz77_ha: build
	./.bin/zipper -with "LZ77-HA" -f "./src/text/data.in" -zip

unzip_ha_lz77: build
	./.bin/zipper -with "HA-LZ77" -f "./src/text/data.in.LZ77-HA" -unzip

##########################
# Testing commands
test_rle:
	go test ./tests/rle_test.go ./tests/tests.go -v

test_bwt:
	go test ./tests/bwt_test.go ./tests/tests.go -v

test_mtf:
	go test ./tests/mtf_test.go ./tests/tests.go -v

test_ha:
	go test ./tests/ha_test.go ./tests/tests.go -v

test_lz77:
	go test ./tests/lz77_test.go ./tests/tests.go -v

test_all: test_rle test_bwt test_mtf test_ha test_lz77