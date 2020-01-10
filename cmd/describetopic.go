package cmd

import (
	"fmt"

	"github.com/gojekfarm/kat/logger"
	"github.com/gojekfarm/kat/pkg"
	"github.com/gojekfarm/kat/util"

	"github.com/spf13/cobra"
)

type describeTopic struct {
	BaseCmd
	topics []string
}

var describeTopicCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describes the given topic",
	Run: func(command *cobra.Command, args []string) {
		cobraUtil := util.NewCobraUtil(command)
		baseCmd := Init(cobraUtil)
		d := describeTopic{BaseCmd: baseCmd, topics: cobraUtil.GetTopicNames()}
		d.describeTopic()
	},
}

func init() {
	describeTopicCmd.PersistentFlags().StringP("topics", "t", "", "Comma separated list of topic names to describe")
	describeTopicCmd.MarkPersistentFlagRequired("topics")
}

func (d *describeTopic) describeTopic() {
	metadata, err := d.TopicCli.Describe(d.topics)
	if err != nil {
		logger.Fatalf("Error while retrieving topic metadata - %v\n", err)
	}
	printConfigs(metadata)
}

func printConfigs(metadata []*pkg.TopicMetadata) {
	for _, topicMetadata := range metadata {
		fmt.Printf("Topic Name: %v,\nIsInternal: %v,\nPartitions:\n", (*topicMetadata).Name, (*topicMetadata).IsInternal)

		partitions := (*topicMetadata).Partitions
		for _, partitionMetadata := range partitions {
			fmt.Printf("Id: %v, Leader: %v, Replicas: %v, ISR: %v, OfflineReplicas: %v\n", (*partitionMetadata).ID, (*partitionMetadata).Leader, (*partitionMetadata).Replicas, (*partitionMetadata).Isr, (*partitionMetadata).OfflineReplicas)
		}
		fmt.Println()
	}
}
