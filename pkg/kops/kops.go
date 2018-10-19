package kops

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/weaveworks/eksctl/pkg/eks/api"

	"k8s.io/kops/pkg/resources/aws"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

type KopsWrapper struct {
	clusterName string
	cloud       awsup.AWSCloud
	spec        *api.ClusterConfig
}

func NewKopsWrapper(cfg *api.ClusterConfig, kopsClusterName string) (*KopsWrapper, error) {
	cloud, err := awsup.NewAWSCloud(cfg.Region, nil)
	if err != nil {
		return nil, err
	}

	return &KopsWrapper{kopsClusterName, cloud, cfg}, nil
}

func (k *KopsWrapper) isOwned(t *ec2.Tag) bool {
	return *t.Key == "kubernetes.io/cluster/"+k.clusterName && *t.Value == "owned"
}

func (k *KopsWrapper) UseVPC() error {
	allVPCs, err := aws.ListVPCs(k.cloud, k.clusterName)
	if err != nil {
		return err
	}

	allSubnets, err := aws.ListSubnets(k.cloud, k.clusterName)
	if err != nil {
		return err
	}

	vpcs := []string{}
	for _, vpc := range allVPCs {
		vpc := vpc.Obj.(*ec2.Vpc)
		for _, tag := range vpc.Tags {
			if k.isOwned(tag) {
				vpcs = append(vpcs, *vpc.VpcId)
			}
		}
	}
	logger.Debug("vpcs = %#v", vpcs)
	if len(vpcs) > 1 {
		return fmt.Errorf("more then one VPC found for kops cluster %q", k.clusterName)
	}
	k.spec.VPC = vpcs[0]

	for _, subnet := range allSubnets {
		subnet := subnet.Obj.(*ec2.Subnet)
		for _, tag := range subnet.Tags {
			if k.isOwned(tag) && *subnet.VpcId == vpcs[0] {
				k.spec.Subnets = append(k.spec.Subnets, *subnet.SubnetId)
				k.spec.AvailabilityZones = append(k.spec.AvailabilityZones, *subnet.AvailabilityZone)
			}
		}
	}
	logger.Debug("subnets = %#v", k.spec.Subnets)
	if len(k.spec.Subnets) < 3 {
		return fmt.Errorf("cannot use VPC from kops cluster less then 3 subnets")
	}

	return nil
}
