/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package cluster

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ecs"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ecs/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeClustersInput returns input for read
// operation.
func GenerateDescribeClustersInput(cr *svcapitypes.Cluster) *svcsdk.DescribeClustersInput {
	res := &svcsdk.DescribeClustersInput{}

	return res
}

// GenerateCluster returns the current state in the form of *svcapitypes.Cluster.
func GenerateCluster(resp *svcsdk.DescribeClustersOutput) *svcapitypes.Cluster {
	cr := &svcapitypes.Cluster{}

	found := false
	for _, elem := range resp.Clusters {
		if elem.ActiveServicesCount != nil {
			cr.Status.AtProvider.ActiveServicesCount = elem.ActiveServicesCount
		} else {
			cr.Status.AtProvider.ActiveServicesCount = nil
		}
		if elem.Attachments != nil {
			f1 := []*svcapitypes.Attachment{}
			for _, f1iter := range elem.Attachments {
				f1elem := &svcapitypes.Attachment{}
				if f1iter.Details != nil {
					f1elemf0 := []*svcapitypes.KeyValuePair{}
					for _, f1elemf0iter := range f1iter.Details {
						f1elemf0elem := &svcapitypes.KeyValuePair{}
						if f1elemf0iter.Name != nil {
							f1elemf0elem.Name = f1elemf0iter.Name
						}
						if f1elemf0iter.Value != nil {
							f1elemf0elem.Value = f1elemf0iter.Value
						}
						f1elemf0 = append(f1elemf0, f1elemf0elem)
					}
					f1elem.Details = f1elemf0
				}
				if f1iter.Id != nil {
					f1elem.ID = f1iter.Id
				}
				if f1iter.Status != nil {
					f1elem.Status = f1iter.Status
				}
				if f1iter.Type != nil {
					f1elem.Type = f1iter.Type
				}
				f1 = append(f1, f1elem)
			}
			cr.Status.AtProvider.Attachments = f1
		} else {
			cr.Status.AtProvider.Attachments = nil
		}
		if elem.AttachmentsStatus != nil {
			cr.Status.AtProvider.AttachmentsStatus = elem.AttachmentsStatus
		} else {
			cr.Status.AtProvider.AttachmentsStatus = nil
		}
		if elem.CapacityProviders != nil {
			f3 := []*string{}
			for _, f3iter := range elem.CapacityProviders {
				var f3elem string
				f3elem = *f3iter
				f3 = append(f3, &f3elem)
			}
			cr.Spec.ForProvider.CapacityProviders = f3
		} else {
			cr.Spec.ForProvider.CapacityProviders = nil
		}
		if elem.ClusterArn != nil {
			cr.Status.AtProvider.ClusterARN = elem.ClusterArn
		} else {
			cr.Status.AtProvider.ClusterARN = nil
		}
		if elem.ClusterName != nil {
			cr.Spec.ForProvider.ClusterName = elem.ClusterName
		} else {
			cr.Spec.ForProvider.ClusterName = nil
		}
		if elem.Configuration != nil {
			f6 := &svcapitypes.ClusterConfiguration{}
			if elem.Configuration.ExecuteCommandConfiguration != nil {
				f6f0 := &svcapitypes.ExecuteCommandConfiguration{}
				if elem.Configuration.ExecuteCommandConfiguration.KmsKeyId != nil {
					f6f0.KMSKeyID = elem.Configuration.ExecuteCommandConfiguration.KmsKeyId
				}
				if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil {
					f6f0f1 := &svcapitypes.ExecuteCommandLogConfiguration{}
					if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil {
						f6f0f1.CloudWatchEncryptionEnabled = elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled
					}
					if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil {
						f6f0f1.CloudWatchLogGroupName = elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName
					}
					if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil {
						f6f0f1.S3BucketName = elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName
					}
					if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil {
						f6f0f1.S3EncryptionEnabled = elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled
					}
					if elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil {
						f6f0f1.S3KeyPrefix = elem.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix
					}
					f6f0.LogConfiguration = f6f0f1
				}
				if elem.Configuration.ExecuteCommandConfiguration.Logging != nil {
					f6f0.Logging = elem.Configuration.ExecuteCommandConfiguration.Logging
				}
				f6.ExecuteCommandConfiguration = f6f0
			}
			cr.Spec.ForProvider.Configuration = f6
		} else {
			cr.Spec.ForProvider.Configuration = nil
		}
		if elem.DefaultCapacityProviderStrategy != nil {
			f7 := []*svcapitypes.CapacityProviderStrategyItem{}
			for _, f7iter := range elem.DefaultCapacityProviderStrategy {
				f7elem := &svcapitypes.CapacityProviderStrategyItem{}
				if f7iter.Base != nil {
					f7elem.Base = f7iter.Base
				}
				if f7iter.CapacityProvider != nil {
					f7elem.CapacityProvider = f7iter.CapacityProvider
				}
				if f7iter.Weight != nil {
					f7elem.Weight = f7iter.Weight
				}
				f7 = append(f7, f7elem)
			}
			cr.Spec.ForProvider.DefaultCapacityProviderStrategy = f7
		} else {
			cr.Spec.ForProvider.DefaultCapacityProviderStrategy = nil
		}
		if elem.PendingTasksCount != nil {
			cr.Status.AtProvider.PendingTasksCount = elem.PendingTasksCount
		} else {
			cr.Status.AtProvider.PendingTasksCount = nil
		}
		if elem.RegisteredContainerInstancesCount != nil {
			cr.Status.AtProvider.RegisteredContainerInstancesCount = elem.RegisteredContainerInstancesCount
		} else {
			cr.Status.AtProvider.RegisteredContainerInstancesCount = nil
		}
		if elem.RunningTasksCount != nil {
			cr.Status.AtProvider.RunningTasksCount = elem.RunningTasksCount
		} else {
			cr.Status.AtProvider.RunningTasksCount = nil
		}
		if elem.Settings != nil {
			f11 := []*svcapitypes.ClusterSetting{}
			for _, f11iter := range elem.Settings {
				f11elem := &svcapitypes.ClusterSetting{}
				if f11iter.Name != nil {
					f11elem.Name = f11iter.Name
				}
				if f11iter.Value != nil {
					f11elem.Value = f11iter.Value
				}
				f11 = append(f11, f11elem)
			}
			cr.Spec.ForProvider.Settings = f11
		} else {
			cr.Spec.ForProvider.Settings = nil
		}
		if elem.Statistics != nil {
			f12 := []*svcapitypes.KeyValuePair{}
			for _, f12iter := range elem.Statistics {
				f12elem := &svcapitypes.KeyValuePair{}
				if f12iter.Name != nil {
					f12elem.Name = f12iter.Name
				}
				if f12iter.Value != nil {
					f12elem.Value = f12iter.Value
				}
				f12 = append(f12, f12elem)
			}
			cr.Status.AtProvider.Statistics = f12
		} else {
			cr.Status.AtProvider.Statistics = nil
		}
		if elem.Status != nil {
			cr.Status.AtProvider.Status = elem.Status
		} else {
			cr.Status.AtProvider.Status = nil
		}
		if elem.Tags != nil {
			f14 := []*svcapitypes.Tag{}
			for _, f14iter := range elem.Tags {
				f14elem := &svcapitypes.Tag{}
				if f14iter.Key != nil {
					f14elem.Key = f14iter.Key
				}
				if f14iter.Value != nil {
					f14elem.Value = f14iter.Value
				}
				f14 = append(f14, f14elem)
			}
			cr.Spec.ForProvider.Tags = f14
		} else {
			cr.Spec.ForProvider.Tags = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateClusterInput returns a create input.
func GenerateCreateClusterInput(cr *svcapitypes.Cluster) *svcsdk.CreateClusterInput {
	res := &svcsdk.CreateClusterInput{}

	if cr.Spec.ForProvider.CapacityProviders != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.CapacityProviders {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetCapacityProviders(f0)
	}
	if cr.Spec.ForProvider.ClusterName != nil {
		res.SetClusterName(*cr.Spec.ForProvider.ClusterName)
	}
	if cr.Spec.ForProvider.Configuration != nil {
		f2 := &svcsdk.ClusterConfiguration{}
		if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration != nil {
			f2f0 := &svcsdk.ExecuteCommandConfiguration{}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.KMSKeyID != nil {
				f2f0.SetKmsKeyId(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.KMSKeyID)
			}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil {
				f2f0f1 := &svcsdk.ExecuteCommandLogConfiguration{}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil {
					f2f0f1.SetCloudWatchEncryptionEnabled(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil {
					f2f0f1.SetCloudWatchLogGroupName(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil {
					f2f0f1.SetS3BucketName(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil {
					f2f0f1.SetS3EncryptionEnabled(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil {
					f2f0f1.SetS3KeyPrefix(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix)
				}
				f2f0.SetLogConfiguration(f2f0f1)
			}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.Logging != nil {
				f2f0.SetLogging(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.Logging)
			}
			f2.SetExecuteCommandConfiguration(f2f0)
		}
		res.SetConfiguration(f2)
	}
	if cr.Spec.ForProvider.DefaultCapacityProviderStrategy != nil {
		f3 := []*svcsdk.CapacityProviderStrategyItem{}
		for _, f3iter := range cr.Spec.ForProvider.DefaultCapacityProviderStrategy {
			f3elem := &svcsdk.CapacityProviderStrategyItem{}
			if f3iter.Base != nil {
				f3elem.SetBase(*f3iter.Base)
			}
			if f3iter.CapacityProvider != nil {
				f3elem.SetCapacityProvider(*f3iter.CapacityProvider)
			}
			if f3iter.Weight != nil {
				f3elem.SetWeight(*f3iter.Weight)
			}
			f3 = append(f3, f3elem)
		}
		res.SetDefaultCapacityProviderStrategy(f3)
	}
	if cr.Spec.ForProvider.Settings != nil {
		f4 := []*svcsdk.ClusterSetting{}
		for _, f4iter := range cr.Spec.ForProvider.Settings {
			f4elem := &svcsdk.ClusterSetting{}
			if f4iter.Name != nil {
				f4elem.SetName(*f4iter.Name)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetSettings(f4)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f5 := []*svcsdk.Tag{}
		for _, f5iter := range cr.Spec.ForProvider.Tags {
			f5elem := &svcsdk.Tag{}
			if f5iter.Key != nil {
				f5elem.SetKey(*f5iter.Key)
			}
			if f5iter.Value != nil {
				f5elem.SetValue(*f5iter.Value)
			}
			f5 = append(f5, f5elem)
		}
		res.SetTags(f5)
	}

	return res
}

// GenerateUpdateClusterInput returns an update input.
func GenerateUpdateClusterInput(cr *svcapitypes.Cluster) *svcsdk.UpdateClusterInput {
	res := &svcsdk.UpdateClusterInput{}

	if cr.Spec.ForProvider.Configuration != nil {
		f1 := &svcsdk.ClusterConfiguration{}
		if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration != nil {
			f1f0 := &svcsdk.ExecuteCommandConfiguration{}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.KMSKeyID != nil {
				f1f0.SetKmsKeyId(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.KMSKeyID)
			}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil {
				f1f0f1 := &svcsdk.ExecuteCommandLogConfiguration{}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil {
					f1f0f1.SetCloudWatchEncryptionEnabled(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil {
					f1f0f1.SetCloudWatchLogGroupName(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil {
					f1f0f1.SetS3BucketName(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil {
					f1f0f1.SetS3EncryptionEnabled(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled)
				}
				if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil {
					f1f0f1.SetS3KeyPrefix(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix)
				}
				f1f0.SetLogConfiguration(f1f0f1)
			}
			if cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.Logging != nil {
				f1f0.SetLogging(*cr.Spec.ForProvider.Configuration.ExecuteCommandConfiguration.Logging)
			}
			f1.SetExecuteCommandConfiguration(f1f0)
		}
		res.SetConfiguration(f1)
	}
	if cr.Spec.ForProvider.Settings != nil {
		f2 := []*svcsdk.ClusterSetting{}
		for _, f2iter := range cr.Spec.ForProvider.Settings {
			f2elem := &svcsdk.ClusterSetting{}
			if f2iter.Name != nil {
				f2elem.SetName(*f2iter.Name)
			}
			if f2iter.Value != nil {
				f2elem.SetValue(*f2iter.Value)
			}
			f2 = append(f2, f2elem)
		}
		res.SetSettings(f2)
	}

	return res
}

// GenerateDeleteClusterInput returns a deletion input.
func GenerateDeleteClusterInput(cr *svcapitypes.Cluster) *svcsdk.DeleteClusterInput {
	res := &svcsdk.DeleteClusterInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}