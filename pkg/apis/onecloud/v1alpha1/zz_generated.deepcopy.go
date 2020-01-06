// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalConfig) DeepCopyInto(out *BaremetalConfig) {
	*out = *in
	out.ServiceCommonOptions = in.ServiceCommonOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalConfig.
func (in *BaremetalConfig) DeepCopy() *BaremetalConfig {
	if in == nil {
		return nil
	}
	out := new(BaremetalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudUser) DeepCopyInto(out *CloudUser) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudUser.
func (in *CloudUser) DeepCopy() *CloudUser {
	if in == nil {
		return nil
	}
	out := new(CloudUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceRequirement)
		**out = **in
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceRequirement)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobSpec) DeepCopyInto(out *CronJobSpec) {
	*out = *in
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobSpec.
func (in *CronJobSpec) DeepCopy() *CronJobSpec {
	if in == nil {
		return nil
	}
	out := new(CronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBConfig) DeepCopyInto(out *DBConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBConfig.
func (in *DBConfig) DeepCopy() *DBConfig {
	if in == nil {
		return nil
	}
	out := new(DBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetSpec) DeepCopyInto(out *DaemonSetSpec) {
	*out = *in
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetSpec.
func (in *DaemonSetSpec) DeepCopy() *DaemonSetSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(appsv1.DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageStatus != nil {
		in, out := &in.ImageStatus, &out.ImageStatus
		*out = new(ImageStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlanceConfig) DeepCopyInto(out *GlanceConfig) {
	*out = *in
	out.ServiceDBCommonOptions = in.ServiceDBCommonOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlanceConfig.
func (in *GlanceConfig) DeepCopy() *GlanceConfig {
	if in == nil {
		return nil
	}
	out := new(GlanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlanceStatus) DeepCopyInto(out *GlanceStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlanceStatus.
func (in *GlanceStatus) DeepCopy() *GlanceStatus {
	if in == nil {
		return nil
	}
	out := new(GlanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostConfig) DeepCopyInto(out *HostConfig) {
	*out = *in
	out.ServiceCommonOptions = in.ServiceCommonOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostConfig.
func (in *HostConfig) DeepCopy() *HostConfig {
	if in == nil {
		return nil
	}
	out := new(HostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneConfig) DeepCopyInto(out *KeystoneConfig) {
	*out = *in
	out.ServiceBaseConfig = in.ServiceBaseConfig
	out.DB = in.DB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneConfig.
func (in *KeystoneConfig) DeepCopy() *KeystoneConfig {
	if in == nil {
		return nil
	}
	out := new(KeystoneConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneSpec) DeepCopyInto(out *KeystoneSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneSpec.
func (in *KeystoneSpec) DeepCopy() *KeystoneSpec {
	if in == nil {
		return nil
	}
	out := new(KeystoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneStatus) DeepCopyInto(out *KeystoneStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneStatus.
func (in *KeystoneStatus) DeepCopy() *KeystoneStatus {
	if in == nil {
		return nil
	}
	out := new(KeystoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterConfig) DeepCopyInto(out *MeterConfig) {
	*out = *in
	out.ServiceDBCommonOptions = in.ServiceDBCommonOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterConfig.
func (in *MeterConfig) DeepCopy() *MeterConfig {
	if in == nil {
		return nil
	}
	out := new(MeterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterStatus) DeepCopyInto(out *MeterStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterStatus.
func (in *MeterStatus) DeepCopy() *MeterStatus {
	if in == nil {
		return nil
	}
	out := new(MeterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mysql) DeepCopyInto(out *Mysql) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mysql.
func (in *Mysql) DeepCopy() *Mysql {
	if in == nil {
		return nil
	}
	out := new(Mysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnecloudCluster) DeepCopyInto(out *OnecloudCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnecloudCluster.
func (in *OnecloudCluster) DeepCopy() *OnecloudCluster {
	if in == nil {
		return nil
	}
	out := new(OnecloudCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnecloudCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnecloudClusterConfig) DeepCopyInto(out *OnecloudClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Keystone = in.Keystone
	out.RegionServer = in.RegionServer
	out.Glance = in.Glance
	out.Webconsole = in.Webconsole
	out.Logger = in.Logger
	out.Yunionconf = in.Yunionconf
	out.Yunionagent = in.Yunionagent
	out.KubeServer = in.KubeServer
	out.AnsibleServer = in.AnsibleServer
	out.Cloudnet = in.Cloudnet
	out.Cloudevent = in.Cloudevent
	out.APIGateway = in.APIGateway
	out.Notify = in.Notify
	out.HostAgent = in.HostAgent
	out.BaremetalAgent = in.BaremetalAgent
	out.S3gateway = in.S3gateway
	out.Devtool = in.Devtool
	out.Meter = in.Meter
	out.AutoUpdate = in.AutoUpdate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnecloudClusterConfig.
func (in *OnecloudClusterConfig) DeepCopy() *OnecloudClusterConfig {
	if in == nil {
		return nil
	}
	out := new(OnecloudClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnecloudClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnecloudClusterList) DeepCopyInto(out *OnecloudClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OnecloudCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnecloudClusterList.
func (in *OnecloudClusterList) DeepCopy() *OnecloudClusterList {
	if in == nil {
		return nil
	}
	out := new(OnecloudClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnecloudClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnecloudClusterSpec) DeepCopyInto(out *OnecloudClusterSpec) {
	*out = *in
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Mysql = in.Mysql
	if in.CertSANs != nil {
		in, out := &in.CertSANs, &out.CertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		copy(*out, *in)
	}
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.RegionServer.DeepCopyInto(&out.RegionServer)
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	in.Glance.DeepCopyInto(&out.Glance)
	in.Climc.DeepCopyInto(&out.Climc)
	in.Webconsole.DeepCopyInto(&out.Webconsole)
	in.Logger.DeepCopyInto(&out.Logger)
	in.Yunionconf.DeepCopyInto(&out.Yunionconf)
	in.Yunionagent.DeepCopyInto(&out.Yunionagent)
	in.Influxdb.DeepCopyInto(&out.Influxdb)
	in.Kapacitor.DeepCopyInto(&out.Kapacitor)
	in.APIGateway.DeepCopyInto(&out.APIGateway)
	in.Web.DeepCopyInto(&out.Web)
	in.KubeServer.DeepCopyInto(&out.KubeServer)
	in.AnsibleServer.DeepCopyInto(&out.AnsibleServer)
	in.Cloudnet.DeepCopyInto(&out.Cloudnet)
	in.Cloudevent.DeepCopyInto(&out.Cloudevent)
	in.Notify.DeepCopyInto(&out.Notify)
	in.HostAgent.DeepCopyInto(&out.HostAgent)
	in.HostDeployer.DeepCopyInto(&out.HostDeployer)
	in.BaremetalAgent.DeepCopyInto(&out.BaremetalAgent)
	in.S3gateway.DeepCopyInto(&out.S3gateway)
	in.Devtool.DeepCopyInto(&out.Devtool)
	in.Meter.DeepCopyInto(&out.Meter)
	in.AutoUpdate.DeepCopyInto(&out.AutoUpdate)
	in.CloudmonPing.DeepCopyInto(&out.CloudmonPing)
	in.CloudmonReportUsage.DeepCopyInto(&out.CloudmonReportUsage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnecloudClusterSpec.
func (in *OnecloudClusterSpec) DeepCopy() *OnecloudClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OnecloudClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnecloudClusterStatus) DeepCopyInto(out *OnecloudClusterStatus) {
	*out = *in
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.RegionServer.DeepCopyInto(&out.RegionServer)
	in.Glance.DeepCopyInto(&out.Glance)
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	in.Webconsole.DeepCopyInto(&out.Webconsole)
	in.Influxdb.DeepCopyInto(&out.Influxdb)
	in.Kapacitor.DeepCopyInto(&out.Kapacitor)
	in.Logger.DeepCopyInto(&out.Logger)
	in.APIGateway.DeepCopyInto(&out.APIGateway)
	in.Web.DeepCopyInto(&out.Web)
	in.Yunionconf.DeepCopyInto(&out.Yunionconf)
	in.Yunionagent.DeepCopyInto(&out.Yunionagent)
	in.KubeServer.DeepCopyInto(&out.KubeServer)
	in.AnsibleServer.DeepCopyInto(&out.AnsibleServer)
	in.Cloudnet.DeepCopyInto(&out.Cloudnet)
	in.Cloudevent.DeepCopyInto(&out.Cloudevent)
	in.Notify.DeepCopyInto(&out.Notify)
	in.BaremetalAgent.DeepCopyInto(&out.BaremetalAgent)
	in.S3gateway.DeepCopyInto(&out.S3gateway)
	in.Devtool.DeepCopyInto(&out.Devtool)
	in.Meter.DeepCopyInto(&out.Meter)
	in.AutoUpdate.DeepCopyInto(&out.AutoUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnecloudClusterStatus.
func (in *OnecloudClusterStatus) DeepCopy() *OnecloudClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OnecloudClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionConfig) DeepCopyInto(out *RegionConfig) {
	*out = *in
	out.ServiceDBCommonOptions = in.ServiceDBCommonOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionConfig.
func (in *RegionConfig) DeepCopy() *RegionConfig {
	if in == nil {
		return nil
	}
	out := new(RegionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionSpec) DeepCopyInto(out *RegionSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionSpec.
func (in *RegionSpec) DeepCopy() *RegionSpec {
	if in == nil {
		return nil
	}
	out := new(RegionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionStatus) DeepCopyInto(out *RegionStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionStatus.
func (in *RegionStatus) DeepCopy() *RegionStatus {
	if in == nil {
		return nil
	}
	out := new(RegionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirement) DeepCopyInto(out *ResourceRequirement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirement.
func (in *ResourceRequirement) DeepCopy() *ResourceRequirement {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBaseConfig) DeepCopyInto(out *ServiceBaseConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBaseConfig.
func (in *ServiceBaseConfig) DeepCopy() *ServiceBaseConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceBaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCommonOptions) DeepCopyInto(out *ServiceCommonOptions) {
	*out = *in
	out.ServiceBaseConfig = in.ServiceBaseConfig
	out.CloudUser = in.CloudUser
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCommonOptions.
func (in *ServiceCommonOptions) DeepCopy() *ServiceCommonOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceCommonOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDBCommonOptions) DeepCopyInto(out *ServiceDBCommonOptions) {
	*out = *in
	out.ServiceCommonOptions = in.ServiceCommonOptions
	out.DB = in.DB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDBCommonOptions.
func (in *ServiceDBCommonOptions) DeepCopy() *ServiceDBCommonOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceDBCommonOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulDeploymentSpec) DeepCopyInto(out *StatefulDeploymentSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulDeploymentSpec.
func (in *StatefulDeploymentSpec) DeepCopy() *StatefulDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebconsoleStatus) DeepCopyInto(out *WebconsoleStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebconsoleStatus.
func (in *WebconsoleStatus) DeepCopy() *WebconsoleStatus {
	if in == nil {
		return nil
	}
	out := new(WebconsoleStatus)
	in.DeepCopyInto(out)
	return out
}
