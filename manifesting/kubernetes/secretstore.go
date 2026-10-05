package kubernetes

import (
	"fmt"

	"github.com/estratocloud/manifesting/manifesting/config"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
)

func SetEnvironmentVariablesFromSecretsStore(object runtime.Object, froms []config.EnvFrom) {

	for _, from := range froms {

		switch o := object.(type) {
		case *corev1.Pod:
			addSecretStoreVolume(&o.Spec, from)
			addEnvFromVolumeMountToContainers(o.Spec.Containers, from)
			addEnvFromVolumeMountToContainers(o.Spec.InitContainers, from)

		case *appsv1.Deployment:
			addSecretStoreVolume(&o.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.Template, from)

		case *appsv1.DaemonSet:
			addSecretStoreVolume(&o.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.Template, from)

		case *appsv1.ReplicaSet:
			addSecretStoreVolume(&o.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.Template, from)

		case *appsv1.StatefulSet:
			addSecretStoreVolume(&o.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.Template, from)

		case *batchv1.Job:
			addSecretStoreVolume(&o.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.Template, from)

		case *batchv1.CronJob:
			addSecretStoreVolume(&o.Spec.JobTemplate.Spec.Template.Spec, from)
			addEnvFromVolumeMountToPodTemplate(&o.Spec.JobTemplate.Spec.Template, from)

		case *corev1.ReplicationController:
			template := o.Spec.Template
			if template != nil {
				addSecretStoreVolume(&template.Spec, from)
				addEnvFromVolumeMountToPodTemplate(template, from)
			}
		}
	}
}

func addSecretStoreVolume(template *corev1.PodSpec, from config.EnvFrom) {
	for _, volume := range template.Volumes {
		if volume.Name == from.SecretRef {
			return
		}
	}

	template.Volumes = append(template.Volumes, corev1.Volume{
		Name: from.SecretRef,
		VolumeSource: corev1.VolumeSource{
			CSI: &corev1.CSIVolumeSource{
				Driver:   "secrets-store.csi.k8s.io",
				ReadOnly: ptr.To(true),
				VolumeAttributes: map[string]string{
					"secretProviderClass": from.SecretProviderClass,
				},
			},
		},
	})
}

func addEnvFromVolumeMountToPodTemplate(template *corev1.PodTemplateSpec, from config.EnvFrom) {
	addEnvFromVolumeMountToContainers(template.Spec.Containers, from)
	addEnvFromVolumeMountToContainers(template.Spec.InitContainers, from)
}

func addEnvFromVolumeMountToContainers(containers []corev1.Container, from config.EnvFrom) {
	for key := range containers {
		addEnvFromToContainer(&containers[key], from)
		addVolumeMountToContainer(&containers[key], from)
	}
}

func addEnvFromToContainer(container *corev1.Container, from config.EnvFrom) {
	for _, existing := range container.EnvFrom {
		if existing.SecretRef != nil {
			if existing.SecretRef.Name == from.SecretRef {
				return
			}
		}
	}

	container.EnvFrom = append(container.EnvFrom, corev1.EnvFromSource{
		SecretRef: &corev1.SecretEnvSource{
			LocalObjectReference: corev1.LocalObjectReference{
				Name: from.SecretRef,
			},
		},
	})
}

func addVolumeMountToContainer(container *corev1.Container, from config.EnvFrom) {
	for _, existing := range container.VolumeMounts {
		if existing.Name == from.SecretRef {
			return
		}
	}

	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
		Name:      from.SecretRef,
		ReadOnly:  true,
		MountPath: fmt.Sprintf("/mnt/%s", from.SecretRef),
	})
}
