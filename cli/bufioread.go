package cli

import (
	"bufio"
	"os"
)

func ScanInputString() string {
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	scanned := reader.Text()

	return scanned
}