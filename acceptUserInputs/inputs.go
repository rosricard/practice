package inputs


import (
	"fmt"
	"bufio"
	"os"
)

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter plate type(s): ")
plate, _ := reader.ReadString('\n')
fmt.Println("Input value:", plate)

fmt.Println(reader)