// Copyright © 2016 Absolute DevOps Ltd <info@absolutedevops.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/absolutedevops/civo/api"
	"github.com/absolutedevops/civo/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var quota api.QuotaParams

// quotaCmd represents the quota command
var quotaCmd = &cobra.Command{
	Use:     "quota",
	Aliases: []string{"quotas"},
	Short:   "List the current account token's quota",
	Long:    `Show all limits for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		if quota.Account == "" {
			quota.Account = api.AccountFindByToken(config.CurrentToken())
			if quota.Account == "" {
				fmt.Println("Couldn't find the default account by its token")
				return
			}
		}

		result, err := api.QuotaGet(quota.Account)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		changing := false
		instanceCountUsage := fmt.Sprintf("%.0f", result.S("instance_count_usage").Data().(float64))
		instanceCountLimit := fmt.Sprintf("%.0f", result.S("instance_count_limit").Data().(float64))
		if quota.InstanceCount == "" {
			quota.InstanceCount = instanceCountLimit
		} else {
			changing = true
		}
		cpuCoreUsage := fmt.Sprintf("%.0f", result.S("cpu_core_usage").Data().(float64))
		cpuCoreLimit := fmt.Sprintf("%.0f", result.S("cpu_core_limit").Data().(float64))
		if quota.CpuCore == "" {
			quota.CpuCore = cpuCoreLimit
		} else {
			changing = true
		}
		ramMbUsage := fmt.Sprintf("%.0f", result.S("ram_mb_usage").Data().(float64))
		ramMbLimit := fmt.Sprintf("%.0f", result.S("ram_mb_limit").Data().(float64))
		if quota.RamMB == "" {
			quota.RamMB = ramMbLimit
		} else {
			changing = true
		}
		diskGbUsage := fmt.Sprintf("%.0f", result.S("disk_gb_usage").Data().(float64))
		diskGbLimit := fmt.Sprintf("%.0f", result.S("disk_gb_limit").Data().(float64))
		if quota.DiskGB == "" {
			quota.DiskGB = diskGbLimit
		} else {
			changing = true
		}
		diskVolumeCountUsage := fmt.Sprintf("%.0f", result.S("disk_volume_count_usage").Data().(float64))
		diskVolumeCountLimit := fmt.Sprintf("%.0f", result.S("disk_volume_count_limit").Data().(float64))
		if quota.DiskVolumeCount == "" {
			quota.DiskVolumeCount = diskVolumeCountLimit
		} else {
			changing = true
		}
		diskSnapshotCountUsage := fmt.Sprintf("%.0f", result.S("disk_snapshot_count_usage").Data().(float64))
		diskSnapshotCountLimit := fmt.Sprintf("%.0f", result.S("disk_snapshot_count_limit").Data().(float64))
		if quota.DiskSnapshotCount == "" {
			quota.DiskSnapshotCount = diskSnapshotCountLimit
		} else {
			changing = true
		}
		publicIpAddressUsage := fmt.Sprintf("%.0f", result.S("public_ip_address_usage").Data().(float64))
		publicIpAddressLimit := fmt.Sprintf("%.0f", result.S("public_ip_address_limit").Data().(float64))
		if quota.PublicIPAddress == "" {
			quota.PublicIPAddress = publicIpAddressLimit
		} else {
			changing = true
		}
		subnetCountUsage := fmt.Sprintf("%.0f", result.S("subnet_count_usage").Data().(float64))
		subnetCountLimit := fmt.Sprintf("%.0f", result.S("subnet_count_limit").Data().(float64))
		if quota.SubnetCount == "" {
			quota.SubnetCount = subnetCountLimit
		} else {
			changing = true
		}
		networkCountUsage := fmt.Sprintf("%.0f", result.S("network_count_usage").Data().(float64))
		networkCountLimit := fmt.Sprintf("%.0f", result.S("network_count_limit").Data().(float64))
		if quota.NetworkCount == "" {
			quota.NetworkCount = networkCountLimit
		} else {
			changing = true
		}
		securityGroupUsage := fmt.Sprintf("%.0f", result.S("security_group_usage").Data().(float64))
		securityGroupLimit := fmt.Sprintf("%.0f", result.S("security_group_limit").Data().(float64))
		if quota.SecurityGroup == "" {
			quota.SecurityGroup = securityGroupLimit
		} else {
			changing = true
		}
		securityGroupRuleUsage := fmt.Sprintf("%.0f", result.S("security_group_rule_usage").Data().(float64))
		securityGroupRuleLimit := fmt.Sprintf("%.0f", result.S("security_group_rule_limit").Data().(float64))
		if quota.SecurityGroupRule == "" {
			quota.SecurityGroupRule = securityGroupRuleLimit
		} else {
			changing = true
		}
		portCountUsage := fmt.Sprintf("%.0f", result.S("port_count_usage").Data().(float64))
		portCountLimit := fmt.Sprintf("%.0f", result.S("port_count_limit").Data().(float64))
		if quota.PortCount == "" {
			quota.PortCount = portCountLimit
		} else {
			changing = true
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetAutoWrapText(false)
		if changing {
			table.SetHeader([]string{"Title", "Used", "Previous", "New"})
			table.Append([]string{"Number of instances", instanceCountUsage, instanceCountLimit, quota.InstanceCount})
			table.Append([]string{"Total CPU cores", cpuCoreUsage, cpuCoreLimit, quota.CpuCore})
			table.Append([]string{"Total RAM (MB)", ramMbUsage, ramMbLimit, quota.RamMB})
			table.Append([]string{"Total disk space (GB)", diskGbUsage, diskGbLimit, quota.DiskGB})
			table.Append([]string{"Disk volumes", diskVolumeCountUsage, diskVolumeCountLimit, quota.DiskVolumeCount})
			table.Append([]string{"Disk snapshots", diskSnapshotCountUsage, diskSnapshotCountLimit, quota.DiskSnapshotCount})
			table.Append([]string{"Public IP addresses", publicIpAddressUsage, publicIpAddressLimit, quota.PublicIPAddress})
			table.Append([]string{"Private subnets", subnetCountUsage, subnetCountLimit, quota.SubnetCount})
			table.Append([]string{"Private networks", networkCountUsage, networkCountLimit, quota.NetworkCount})
			table.Append([]string{"Security groups", securityGroupUsage, securityGroupLimit, quota.SecurityGroup})
			table.Append([]string{"Security group rules", securityGroupRuleUsage, securityGroupRuleLimit, quota.SecurityGroupRule})
			table.Append([]string{"Number of ports (network connections)", portCountUsage, portCountLimit, quota.PortCount})
		} else {
			table.SetHeader([]string{"Title", "Used", "Limit"})
			table.Append([]string{"Number of instances", instanceCountUsage, instanceCountLimit})
			table.Append([]string{"Total CPU cores", cpuCoreUsage, cpuCoreLimit})
			table.Append([]string{"Total RAM (MB)", ramMbUsage, ramMbLimit})
			table.Append([]string{"Total disk space (GB)", diskGbUsage, diskGbLimit})
			table.Append([]string{"Disk volumes", diskVolumeCountUsage, diskVolumeCountLimit})
			table.Append([]string{"Disk snapshots", diskSnapshotCountUsage, diskSnapshotCountLimit})
			table.Append([]string{"Public IP addresses", publicIpAddressUsage, publicIpAddressLimit})
			table.Append([]string{"Private subnets", subnetCountUsage, subnetCountLimit})
			table.Append([]string{"Private networks", networkCountUsage, networkCountLimit})
			table.Append([]string{"Security groups", securityGroupUsage, securityGroupLimit})
			table.Append([]string{"Security group rules", securityGroupRuleUsage, securityGroupRuleLimit})
			table.Append([]string{"Number of ports (network connections)", portCountUsage, portCountLimit})
		}
		table.Render()

		if changing {
			_, err := api.QuotaSet(quota)
			if err != nil {
				fmt.Printf("An error occured: ", err)
				return
			}
			fmt.Println("Quota updated for account", quota.Account)
		}
	},
}

func init() {
	RootCmd.AddCommand(quotaCmd)
	if config.Admin() {
		quotaCmd.Flags().StringVarP(&quota.Account, "account", "", "", "The account to update the quota for")
		quotaCmd.Flags().StringVarP(&quota.InstanceCount, "instance-count", "i", "", "The limit to the number of instances available")
		quotaCmd.Flags().StringVarP(&quota.CpuCore, "cpu-core", "c", "", "The limit to the number of CPU cores available")
		quotaCmd.Flags().StringVarP(&quota.RamMB, "ram-mb", "r", "", "The limit to the amount of RAM (in MB) available")
		quotaCmd.Flags().StringVarP(&quota.DiskGB, "disk-gb", "d", "", "The limit to the of disk space (in GB) available")
		quotaCmd.Flags().StringVarP(&quota.DiskVolumeCount, "disk-volume-count", "v", "", "The limit to the number of disk volumes available")
		quotaCmd.Flags().StringVarP(&quota.DiskSnapshotCount, "disk-snapshot-count", "s", "", "The limit to the number of disk snapshots available")
		quotaCmd.Flags().StringVarP(&quota.PublicIPAddress, "public-ip-address", "a", "", "The limit to the number of public IP addresses available")
		quotaCmd.Flags().StringVarP(&quota.SubnetCount, "subnet-count", "u", "", "The limit to the number of subnets available")
		quotaCmd.Flags().StringVarP(&quota.NetworkCount, "network-count", "w", "", "The limit to the number of networks available")
		quotaCmd.Flags().StringVarP(&quota.SecurityGroup, "security-group", "g", "", "The limit to the number of security groups available")
		quotaCmd.Flags().StringVarP(&quota.SecurityGroupRule, "security-group-rule", "l", "", "The limit to the number of security group rules available")
		quotaCmd.Flags().StringVarP(&quota.PortCount, "port-count", "p", "", "The limit to the number of ports (network connections) available")
	}
}
