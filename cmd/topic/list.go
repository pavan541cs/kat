package topic

import (
	"fmt"
	"github.com/spf13/cobra"
	"source.golabs.io/hermes/kafka-admin-tools/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the topics satisfying the passed criteria if any",
	Run:   listTopics,
}

func init() {
	listCmd.PersistentFlags().StringP("broker-list", "b", "", "Comma separated list of broker ips")
	listCmd.PersistentFlags().IntP("replication-factor", "r", 0, "Replication Factor of the topic")

	listCmd.MarkPersistentFlagRequired("broker-list")
}

func listTopics(cmd *cobra.Command, args []string) {
	admin := utils.GetAdminClient(cmd)
	replicationFactor := getReplicationFactor(cmd)

	topicDetails, err := admin.ListTopics()
	if err != nil {
		fmt.Printf("Err while retrieving topic details: %v\n", err)
	}
	for topicDetail := range topicDetails {
		if replicationFactor != 0 {
			if int(topicDetails[topicDetail].ReplicationFactor) == replicationFactor {
				fmt.Println(topicDetail)
			}
		} else {
			fmt.Println(topicDetail)
		}
	}
}


