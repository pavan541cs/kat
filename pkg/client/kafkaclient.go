package client

type TopicDetail struct {
	NumPartitions     int32
	ReplicationFactor int16
	ReplicaAssignment map[int32][]int32
	Config            map[string]*string
}

type TopicMetadata struct {
	Err        error
	Name       string
	IsInternal bool
	Partitions []*PartitionMetadata
}

type PartitionMetadata struct {
	Err             error
	ID              int32
	Leader          int32
	Replicas        []int32
	Isr             []int32
	OfflineReplicas []int32
}

type ConfigResource struct {
	Type        int
	Name        string
	ConfigNames []string
}

type ConfigEntry struct {
	Name      string
	Value     string
	ReadOnly  bool
	Default   bool
	Source    string
	Sensitive bool
	Synonyms  []*ConfigSynonym
}

type ConfigSynonym struct {
	ConfigName  string
	ConfigValue string
	Source      string
}

type ListTopicsRequest struct {
	LastWritten int64
	DataDir     string
}

type KafkaAPIClient interface {
	CreateTopic(topic string, detail TopicDetail, validateOnly bool) error
	CreatePartitions(topic string, count int32, assignment [][]int32, validateOnly bool) error
	ListBrokers() map[int]string
	ListTopicDetails() (map[string]TopicDetail, error)
	DeleteTopic(topics []string) error
	DescribeTopicMetadata(topics []string) ([]*TopicMetadata, error)
	UpdateConfig(resourceType int, name string, entries map[string]*string, validateOnly bool) error
	GetTopicResourceType() int
	GetConfig(resource ConfigResource) ([]ConfigEntry, error)
}

type KafkaSSHClient interface {
	ListTopics(ListTopicsRequest) ([]string, error)
}
