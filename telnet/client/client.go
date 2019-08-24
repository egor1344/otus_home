package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Connect to telnet server
func Connect(addr, port string, timeout int) error {
	log.Println(addr, port)
	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	addrPort := addr + ":" + port
	conn, err := dialer.DialContext(ctx, "tcp", addrPort)
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(ctx, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn)
		wg.Done()
	}()

	wg.Wait()
	conn.Close()
	return nil
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		case s := <-sigs:
			log.Println("CTRL c", s)
			conn.Close()
			return
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
}

func readRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}
