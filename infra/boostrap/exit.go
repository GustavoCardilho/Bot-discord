package boostrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Exit() {
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	session.Close()
}
