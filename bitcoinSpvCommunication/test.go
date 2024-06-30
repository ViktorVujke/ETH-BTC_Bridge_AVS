package bitcoin

import (
	"fmt"
	"os/exec"
)

func SendTx(signatures []string) {
	cmdArgs := append([]string{"spremo/send.js"}, signatures...)
	fmt.Println("SACUUUU")
	fmt.Println(signatures[0], "!!!!!!!!!")
	cmd := exec.Command("node", cmdArgs...)

	opt, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Error executing command2:", string(opt))
		return
	}
}

func SignTx(destAddress string, amount string) string {
	// Command to run
	cmd := exec.Command("node", "spremo/app.js",
		"1000000000000000000000000000000000000000000000000000000000000001", // Privatni kljuc
		destAddress,                           // ovde saljemo bitcoin
		"2MyUHzyJuWVjXCNV2JML2Muh56vSJt5PZr2", // multisig
		amount,                                // amount
		"512102347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df74621037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e28552ae") // redeemSkripta

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {

		fmt.Println("Error executing command:", err)
		fmt.Println("Error executing command:", string(output))
		return ""
	}

	// Store output in a string variable
	outputString := string(output)

	// Print the output
	//fmt.Println("Command Output:")
	//fmt.Println(outputString)

	return outputString

	// You can use the outputString variable as needed
}
