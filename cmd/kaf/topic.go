package main

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"

	"strings"

	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var (
	partitionsFlag           int32
	partitionAssignmentsFlag string
	replicasFlag             int16
	noHeaderFlag             bool
	compactFlag              bool
	zookeepersFlag           []string
	rootPathFlag             string
)

type PartitionState struct {
	ControllerEpoch int   `json:"controller_epoch"`
	Leader          int   `json:"leader"`
	Version         int   `json:"version"`
	LeaderEpoch     int   `json:"leader_epoch"`
	Isr             []int `json:"isr"`
}

type TopicInfo struct {
	Name     string
	ParState []PartitionState
}

func init() {
	rootCmd.AddCommand(topicCmd)
	rootCmd.AddCommand(topicsCmd)
	topicCmd.AddCommand(createTopicCmd)
	topicCmd.AddCommand(deleteTopicCmd)
	topicCmd.AddCommand(lsTopicsCmd)
	topicCmd.AddCommand(describeTopicCmd)
	topicCmd.AddCommand(addConfigCmd)
	topicCmd.AddCommand(topicSetConfig)
	topicCmd.AddCommand(updateTopicCmd)

	createTopicCmd.Flags().Int32VarP(&partitionsFlag, "partitions", "p", int32(1), "Number of partitions")
	createTopicCmd.Flags().Int16VarP(&replicasFlag, "replicas", "r", int16(1), "Number of replicas")
	createTopicCmd.Flags().BoolVar(&compactFlag, "compact", false, "Enable topic compaction")

	lsTopicsCmd.Flags().BoolVar(&noHeaderFlag, "no-headers", false, "Hide table headers")
	lsTopicsCmd.Flags().StringSliceVarP(&zookeepersFlag, "zookeepers", "z", nil, "Comma separated list of zookeeper server ip:port pairs")
	lsTopicsCmd.Flags().StringVar(&rootPathFlag, "root-path", "/", "kafka root path in zookeeper")

	topicsCmd.Flags().BoolVar(&noHeaderFlag, "no-headers", false, "Hide table headers")
	topicsCmd.Flags().StringSliceVarP(&zookeepersFlag, "zookeepers", "z", nil, "Comma separated list of zookeeper server ip:port pairs")
	topicsCmd.Flags().StringVar(&rootPathFlag, "root-path", "/", "kafka root path in zookeeper")

	describeTopicCmd.Flags().StringSliceVarP(&zookeepersFlag, "zookeepers", "z", nil, "Comma separated list of zookeeper server ip:port pairs")
	describeTopicCmd.Flags().StringVar(&rootPathFlag, "root-path", "/", "kafka root path in zookeeper")

	updateTopicCmd.Flags().Int32VarP(&partitionsFlag, "partitions", "p", int32(-1), "Number of partitions")
	updateTopicCmd.Flags().StringVar(&partitionAssignmentsFlag, "partition-assignments", "", "Partition Assignments. Optional. If set in combination with -p, an assignment must be provided for each new partition. Example: '[[1,2,3],[1,2,3]]' (JSON Array syntax) assigns two new partitions to brokers 1,2,3. If used by itself, a reassignment must be provided for all partitions.")
}

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Create and describe topics.",
}

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "List topics",
	Run:   lsTopicsCmd.Run,
}

var topicSetConfig = &cobra.Command{
	Use:     "set-config",
	Short:   "set topic config",
	Example: "kaf topic set-config topic.name \"cleanup.policy=delete\"",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		topic := args[0]

		splt := strings.Split(args[1], ",")
		configs := make(map[string]*string)

		for _, kv := range splt {
			s := strings.Split(kv, "=")

			if len(s) != 2 {
				continue
			}

			key := s[0]
			value := s[1]
			configs[key] = &value
		}

		if len(configs) < 1 {
			errorExit("No valid configs found")
		}

		err := admin.AlterConfig(sarama.TopicResource, topic, configs, false)
		if err != nil {
			errorExit("Unable to alter topic config: %v\n", err)
		}
		fmt.Printf("\xE2\x9C\x85 Updated config.")
	},
}

var updateTopicCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update topic",
	Example: "kaf topic update -p 5 --partition-assignments '[[1,2,3],[1,2,3]]'",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		if partitionsFlag == -1 && partitionAssignmentsFlag == "" {
			errorExit("Number of partitions and/or partition assigments must be given")
		}

		var assignments [][]int32
		if partitionAssignmentsFlag != "" {
			if err := json.Unmarshal([]byte(partitionAssignmentsFlag), &assignments); err != nil {
				errorExit("Invalid partition assignments: %v", err)
			}
		}

		if partitionsFlag != int32(-1) {
			err := admin.CreatePartitions(args[0], partitionsFlag, assignments, false)
			if err != nil {
				errorExit("Failed to create partitions: %v", err)
			}
		} else {
			// Needs at least Kafka version 2.4.0.
			err := admin.AlterPartitionReassignments(args[0], assignments)
			if err != nil {
				errorExit("Failed to reassign the partition assigments: %v", err)
			}
		}
		fmt.Printf("\xE2\x9C\x85 Updated topic!\n")
	},
}

var lsTopicsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List topics",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)

		if len(zookeepersFlag) > 0 {

			zc := NewZkConfig(zookeepersFlag, rootPathFlag, time.Duration(3)*time.Second, verbose)
			err := zc.Connect()
			if err != nil {
				errorExit("Unable to connect zookeeper ensemble %v to get topics: %v\n", zookeepersFlag, err)
			}
			rootChildren, err := zc.LsDir(rootPathFlag)
			if err != nil {
				errorExit("Unable to ls %s in zookeeper ensemble %v: %v\n", rootPathFlag, zookeepersFlag, err)
			}
			isRightRootPath := false
			for _, dir := range rootChildren {
				if dir == "brokers" {
					isRightRootPath = true
				}
			}
			if isRightRootPath {
				topicPath := path.Join(rootPathFlag, "brokers", "topics")
				topicChildren, err := zc.LsDir(topicPath)
				if err != nil {
					errorExit("Unable to ls %s in zookeeper ensemble %v: %v\n", topicPath, zookeepersFlag, err)
				}

				topics := make([]TopicInfo, len(topicChildren))

				for idx, tName := range topicChildren {
					partitionPath := path.Join(rootPathFlag, "brokers", "topics", tName, "partitions")
					partitionChildren, err := zc.LsDir(partitionPath)

					if err != nil {
						errorExit("Unable to ls %s in zookeeper ensemble %v: %v\n", partitionChildren, zookeepersFlag, err)
					}

					sort.Slice(partitionChildren, func(i int, j int) bool {
						iInt, _ := strconv.Atoi(partitionChildren[i])
						jInt, _ := strconv.Atoi(partitionChildren[j])
						return iInt < jInt
					})

					for _, num := range partitionChildren {
						partitionNumStatePath := path.Join(rootPathFlag, "brokers", "topics", tName, "partitions", num, "state")

						stateData, err := zc.GetData(partitionNumStatePath)

						if err != nil {
							errorExit("Unable to get %s of %s in zookeeper ensemble %v: %v\n", stateData, partitionNumStatePath, zookeepersFlag, err)
						}
						var state PartitionState
						err = json.Unmarshal(stateData, &state)
						if err != nil {
							errorExit("Unable to decode %s of %s zookeeper ensemble %v: %v\n", string(stateData), partitionNumStatePath, zookeepersFlag, err)
						}
						topics[idx].Name = tName
						topics[idx].ParState = append(topics[idx].ParState, state)
					}
				}
				if len(topics) > 0 {
					if !noHeaderFlag {
						_, _ = fmt.Fprintf(w, "NAME\tPARTITIONS\tLEADER\tISR\t\n")
					}
					sort.Slice(topics, func(i int, j int) bool {
						return strings.ToLower(topics[i].Name) < strings.ToLower(topics[j].Name)
					})
					for _, topic := range topics {
						for idx, par := range topic.ParState {
							_, _ = fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n", topic.Name, idx, par.Leader, par.Isr)
						}
					}
				}
			} else {
				errorExit("Unable to %s in zookeeper ensemble %v\n", rootPathFlag, zookeepersFlag)
			}
		} else {
			admin := getClusterAdmin()

			topics, err := admin.ListTopics()
			if err != nil {
				errorExit("Unable to list topics: %v\n", err)
			}

			sortedTopics := make(
				[]struct {
					name string
					sarama.TopicDetail
				}, len(topics))

			i := 0
			for name, topic := range topics {
				sortedTopics[i].name = name
				sortedTopics[i].TopicDetail = topic
				i++
			}

			sort.Slice(sortedTopics, func(i int, j int) bool {
				return sortedTopics[i].name < sortedTopics[j].name
			})

			if !noHeaderFlag {
				fmt.Fprintf(w, "NAME\tPARTITIONS\tREPLICAS\t\n")
			}

			for _, topic := range sortedTopics {
				fmt.Fprintf(w, "%v\t%v\t%v\t\n", topic.name, topic.NumPartitions, topic.ReplicationFactor)
			}
		}
		w.Flush()
	},
}

var describeTopicCmd = &cobra.Command{
	Use:               "describe",
	Short:             "Describe topic",
	Long:              "Describe a topic. Default values of the configuration are omitted.",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validTopicArgs,
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		topicDetails, err := admin.DescribeTopics([]string{args[0]})
		if err != nil {
			errorExit("Unable to describe topics: %v\n", err)
		}

		if topicDetails[0].Err == sarama.ErrUnknownTopicOrPartition {
			fmt.Printf("Topic %v not found.\n", args[0])
			return
		}

		cfg, err := admin.DescribeConfig(sarama.ConfigResource{
			Type: sarama.TopicResource,
			Name: args[0],
		})
		if err != nil {
			errorExit("Unable to describe config: %v\n", err)
		}

		var compacted bool
		for _, e := range cfg {
			if e.Name == "cleanup.policy" && e.Value == "compact" {
				compacted = true
			}
		}

		detail := topicDetails[0]
		sort.Slice(detail.Partitions, func(i, j int) bool { return detail.Partitions[i].ID < detail.Partitions[j].ID })

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		fmt.Fprintf(w, "Name:\t%v\t\n", detail.Name)
		fmt.Fprintf(w, "Internal:\t%v\t\n", detail.IsInternal)
		fmt.Fprintf(w, "Compacted:\t%v\t\n", compacted)
		fmt.Fprintf(w, "Partitions:\n")

		w.Flush()
		w.Init(outWriter, tabwriterMinWidthNested, 4, 2, tabwriterPadChar, tabwriterFlags)

		fmt.Fprintf(w, "\tPartition\tHigh Watermark\tLeader\tReplicas\tISR\t\n")
		fmt.Fprintf(w, "\t---------\t--------------\t------\t--------\t---\t\n")

		partitions := make([]int32, 0, len(detail.Partitions))
		for _, partition := range detail.Partitions {
			partitions = append(partitions, partition.ID)
		}
		highWatermarks := getHighWatermarks(args[0], partitions)
		highWatermarksSum := 0

		for _, partition := range detail.Partitions {
			sortedReplicas := partition.Replicas
			sort.Slice(sortedReplicas, func(i, j int) bool { return sortedReplicas[i] < sortedReplicas[j] })

			sortedISR := partition.Isr
			sort.Slice(sortedISR, func(i, j int) bool { return sortedISR[i] < sortedISR[j] })

			highWatermarksSum += int(highWatermarks[partition.ID])

			fmt.Fprintf(w, "\t%v\t%v\t%v\t%v\t%v\t\n", partition.ID, highWatermarks[partition.ID], partition.Leader, sortedReplicas, sortedISR)
		}

		w.Flush()

		fmt.Fprintf(w, "Summed HighWatermark:\t%d\n", highWatermarksSum)
		w.Flush()

		fmt.Fprintf(w, "Config:\n")
		fmt.Fprintf(w, "\tName\tValue\tReadOnly\tSensitive\t\n")
		fmt.Fprintf(w, "\t----\t-----\t--------\t---------\t\n")

		for _, entry := range cfg {
			if entry.Default {
				continue
			}
			fmt.Fprintf(w, "\t%v\t%v\t%v\t%v\t\n", entry.Name, entry.Value, entry.ReadOnly, entry.Sensitive)
		}

		w.Flush()
	},
}

var createTopicCmd = &cobra.Command{
	Use:   "create TOPIC",
	Short: "Create a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		topicName := args[0]
		compact := "delete"
		if compactFlag {
			compact = "compact"
		}
		err := admin.CreateTopic(topicName, &sarama.TopicDetail{
			NumPartitions:     partitionsFlag,
			ReplicationFactor: replicasFlag,
			ConfigEntries: map[string]*string{
				"cleanup.policy": &compact,
			},
		}, false)
		if err != nil {
			fmt.Printf("Could not create topic %v: %v\n", topicName, err.Error())
		} else {
			w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
			fmt.Fprintf(w, "\xE2\x9C\x85 Created topic!\n")
			fmt.Fprintln(w, "\tTopic Name:\t", topicName)
			fmt.Fprintln(w, "\tPartitions:\t", partitionsFlag)
			fmt.Fprintln(w, "\tReplication Factor:\t", replicasFlag)
			fmt.Fprintln(w, "\tCleanup Policy:\t", compact)
			w.Flush()
		}
	},
}

var addConfigCmd = &cobra.Command{
	Use:   "add-config TOPIC KEY VALUE",
	Short: "Add config key/value pair to topic",
	Args:  cobra.ExactArgs(3), // TODO how to unset ? support empty VALUE ?
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		topic := args[0]
		key := args[1]
		value := args[2]

		err := admin.AlterConfig(sarama.TopicResource, topic, map[string]*string{
			key: &value,
		}, false)
		if err != nil {
			errorExit("failed to update topic config: %v", err)
		} else {
			fmt.Printf("Added config %v=%v to topic %v.\n", key, value, topic)
		}
	},
}

var deleteTopicCmd = &cobra.Command{
	Use:               "delete TOPIC",
	Short:             "Delete a topic",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validTopicArgs,
	Run: func(cmd *cobra.Command, args []string) {
		admin := getClusterAdmin()

		topicName := args[0]
		err := admin.DeleteTopic(topicName)
		if err != nil {
			fmt.Fprintf(outWriter, "Could not delete topic %v: %v\n", topicName, err.Error())
		} else {
			fmt.Fprintf(outWriter, "\xE2\x9C\x85 Deleted topic %v!\n", topicName)
		}
	},
}
