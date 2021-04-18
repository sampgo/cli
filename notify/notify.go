package notify

import (
	"fmt"
	"github.com/ttacon/chalk"
)

func Info(msg string) {
	fmt.Println(chalk.White, chalk.Bold, "INFO:", chalk.Reset, chalk.White, " ", msg, chalk.Reset)
}

func Error(msg string) {
	fmt.Println(chalk.Red, chalk.Bold, "ERROR:", chalk.Reset, chalk.Red, " ", msg, chalk.Reset)
}

func Warning(msg string) {
	fmt.Println(chalk.Yellow, chalk.Bold, "WARNING:", chalk.Reset, chalk.Yellow, " ", msg, chalk.Reset)
}
