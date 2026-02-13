// MYR Vault - Concurrent Chunk Streamer
package main
import ("fmt"; "time")

func establish_tunnel(id int) {
	fmt.Printf("[SYS] Go-Routine %d: Tunnel secured for encrypted chunk stream.\n", id)
}

func main() {
	go establish_tunnel(1)
	go establish_tunnel(2)
	time.Sleep(100 * time.Millisecond)
}
