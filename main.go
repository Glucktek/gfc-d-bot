package main

import (
	"context"
	"fmt"
	"log"

	"github.com/glucktek/gfc-d-bot/pkgs/lightsail"
)

func main() {

	fmt.Println("Starting bot service...")

	// Initialize the Lightsail client
	client, err := lightsail.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// HACK: Using hardcaded instance name for now
	instanceName := "GreaterFaithChurchSite"

	// Get status
	state, err := client.GetInstanceState(ctx, instanceName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Instance state: %s\n", state)

	// Start instance
	// if err := client.StartInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

	// Stop instance
	// if err := client.StopInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

	// // Reboot instance
	// if err := client.RebootInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

}
