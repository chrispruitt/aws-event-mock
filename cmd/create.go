package cmd

import (
	lib "aws-event-mock/lib"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	event     lib.Event
	eventType string

	// Run Command go run main.go create
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "returns a valid aws event in json format",
		Run: func(cmd *cobra.Command, args []string) {
			if event.Message == "" {
				log.Fatal("--message cannot be empty")
			}

			// Validate EventTypeEnum
			var err error
			eventTypeEnum, err := lib.ParseEventTypeEnum(eventType)
			if err != nil {
				log.Fatal(fmt.Sprintf("--type flag input incorrect. %s", err.Error()))
			}

			resp, err := lib.GetEvent(&event, eventTypeEnum)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
		},
	}
)

func init() {
	createCmd.PersistentFlags().StringVarP(&event.Message, "message", "m", "", "Message of the event. (Required)")
	createCmd.PersistentFlags().StringVarP(&eventType, "type", "t", lib.EventTypeEnumCloudwatchLog.String(), "Event Type. (default: cloudwatch-log)")

	// runCmd.PersistentFlags().StringVarP(&task.Cluster, "cluster", "", "", "ECS cluster")
	// runCmd.PersistentFlags().StringVarP(&task.Name, "name", "n", "ephemeral-task-from-ecs-cli", "Assign a name to the task")
	// runCmd.PersistentFlags().StringVar(&task.Family, "family", "", "Family for ECS task")
	// runCmd.PersistentFlags().StringVar(&task.ExecutionRoleArn, "execution-role", "", "Execution role ARN (required for Fargate)")
	// runCmd.PersistentFlags().StringVar(&task.RoleArn, "role", "", "Task role ARN")
	// runCmd.PersistentFlags().BoolVarP(&task.Detach, "detach", "d", false, "Run the task in the background")
	// runCmd.PersistentFlags().Int64VarP(&task.Count, "count", "c", 1, "Spawn n tasks")
	// runCmd.PersistentFlags().Int64VarP(&task.Memory, "memory", "m", 0, "Memory limit")
	// runCmd.PersistentFlags().Int64Var(&task.CPUReservation, "cpu-reservation", 0, "CPU reservation")
	// runCmd.PersistentFlags().Int64Var(&task.MemoryReservation, "memory-reservation", 2048, "Memory reservation")
	// runCmd.PersistentFlags().StringArrayVarP(&task.Environment, "env", "e", nil, "Set environment variables")
	// runCmd.PersistentFlags().StringArrayVarP(&task.Publish, "publish", "p", nil, "Publish a container's port(s) to the host")
	// // TODO: attach a specific security group
	// runCmd.PersistentFlags().StringArrayVar(&task.SecurityGroups, "security-groups", nil, "[TODO] Attach security groups to task")
	// runCmd.PersistentFlags().StringArrayVar(&task.Subnets, "subnet", nil, "Subnet(s) where task should run")
	// runCmd.PersistentFlags().StringArrayVarP(&task.Volumes, "volume", "v", nil, "Map volume to ECS Container Instance")
	// runCmd.PersistentFlags().StringArrayVarP(&task.EfsVolumes, "efs-volume", "", nil, "Map EFS volume to ECS Container Instance (ex. fs-23kj2f:/efs/dir:/container/mnt/dir)")
	// // TODO: support assigning public ip address
	// runCmd.PersistentFlags().BoolVar(&task.Public, "public", false, "assign public ip")
	// runCmd.PersistentFlags().BoolVar(&task.Fargate, "fargate", false, "Launch in Fargate")
	// runCmd.PersistentFlags().BoolVar(&task.Deregister, "no-deregister", false, "do not deregister the task definition")
	// runCmd.Flags().SetInterspersed(false)
}
