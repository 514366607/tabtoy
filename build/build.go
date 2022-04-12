package build

import "fmt"

var (
	Version   = "3.1.1"
	GitCommit string
	BuildTime string
)

func Print() {
	fmt.Println("Version: ", Version)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println("BuildTime: ", BuildTime)
}
