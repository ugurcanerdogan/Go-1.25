package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go 1.25+ GOMAXPROCS behavior example:")

	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS (default):", runtime.GOMAXPROCS(0))

	fmt.Println("\nSetting GOMAXPROCS=4 manually...")
	runtime.GOMAXPROCS(4)
	fmt.Println("GOMAXPROCS (manual):", runtime.GOMAXPROCS(0))

	fmt.Println("\nReverting to default value...")
	runtime.SetDefaultGOMAXPROCS()
	fmt.Println("GOMAXPROCS (default):", runtime.GOMAXPROCS(0))

	fmt.Println("\nNote: When running inside Docker or a cgroup environment with a CPU limit, GOMAXPROCS will automatically reflect the limit.")
}
