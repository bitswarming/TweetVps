package cmd

// rebootclient.go encode  --secret=xxx --targetdomain=quad.node.consul --message=reboot:robot.node.consul
import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"time"
)

var SecretPass string
var DomainThis string
var MessageMes string
var Userss string
var ExpireS int64
var Period int64


func init() {
	daemonCmd.Flags().StringVarP(&SecretPass, "secret", "", "", "Password for decoding")
	daemonCmd.Flags().StringVarP(&Userss, "user", "", "", "Twitter user name")
	daemonCmd.Flags().StringVarP(&DomainThis, "mydomain", "", "", "target domain for commands")
	daemonCmd.Flags().StringVarP(&MessageMes, "message", "", "", "Message")
	daemonCmd.Flags().Int64Var(&ExpireS, "expire", 0, "Message")
	daemonCmd.Flags().Int64Var(&Period, "period", 0, "Message")
	RootCmd.AddCommand(daemonCmd)
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Daemonize twitter fetcher",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Println(err)
		}
		tick := time.Tick(3000 * time.Millisecond)
		// Keep trying until we're timed out or got a result or got an error
		for {
			select {
			case <-tick:
				fmt.Println("tick")
				fetchtwitter(Userss, ExpireS, SecretPass, DomainThis)
			}
		}
	},
}
