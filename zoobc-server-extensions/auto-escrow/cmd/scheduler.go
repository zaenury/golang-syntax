package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zoobc-server-extensions/auto-escrow/core"
)

type (
	// RunCommand represent of output function from scheduler commands
	RunCommand func(ccmd *cobra.Command, args []string)
)

var (
	accountAddress, seedPhrase, nodeAddress string

	schedulerCmd = &cobra.Command{
		Use:   "scheduler",
		Short: "run scheduler for escrow server",
		Long:  ``,
	}

	runServerCmd = &cobra.Command{
		Use:   "server",
		Short: "scheduler for escrow server",
	}
)

func init() {
	schedulerCmd.PersistentFlags().StringVar(
		&accountAddress,
		"account-address", "", "Account Address for the escrow server",
	)

	schedulerCmd.PersistentFlags().StringVar(
		&seedPhrase,
		"seed-phrase", "", "a phrase used to sign the transactions",
	)

	schedulerCmd.PersistentFlags().StringVar(
		&nodeAddress,
		"node-address", "", "a node address",
	)
}

// Commands will return the main scheduler cmd
func Commands() *cobra.Command {
	runServerCmd.Run = CronJobBlockAndTransaction()
	schedulerCmd.AddCommand(runServerCmd)
	return schedulerCmd
}

// PostTransactionNodeRegistry scheduler for post transction related with node registry
func CronJobBlockAndTransaction() RunCommand {
	return func(cmd *cobra.Command, args []string) {

		var ticker = time.NewTicker(time.Duration(6000) * time.Millisecond)
		doneSignal := make(chan bool)

		//Set the AccAddress, seedPhrase & NodeAddress to environtment
		_ = os.Setenv("Account_Address", accountAddress)
		_ = os.Setenv("Seed_Phrase", seedPhrase)
		_ = os.Setenv("Node_Address", nodeAddress)

		for {
			go core.StartServer()
			fmt.Println("CRONJOB FUNCTION WORKS IN HERE FOR PULLING BLOCKS AND TRANSACTIONS")

			select {
			case <-doneSignal:
				return
			case <-ticker.C:
				continue
			}
		}
	}
}
