// package main

// import (
//     "log"
// )

// func main() {
//     log.Print("Hey, I'm a log!")
// }

// Desta forma, ao executar o log.fatal, sai da aplicação
// com se fosse o comando os.Exit(1)
// package main

// import (
//     "fmt"
//     "log"
// )

// func main() {
//     log.Fatal("Hey, I'm an error log!")
//     fmt.Print("Can you see me?")
// }

// package main

// import (
//     "fmt"
//     "log"
// )

// func main() {
//     log.Panic("Hey, I'm an error log!")
//     fmt.Print("Can you see me?")
// }

// package main

// import (
//     "log"
// )

// func main() {
//     log.SetPrefix("main(): ")
//     log.Print("Hey, I'm a log!")
//     log.Fatal("Hey, I'm an error log!")
// }

// registando log em um arquivo
// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()

// 	log.SetOutput(file)
// 	log.Print("Hey, I'm a log!")
// }

// package main

// import (
// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"
// )

// func main() {
// 	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
// 	log.Print("Hey! I'm a log message!")
// }

package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
}
