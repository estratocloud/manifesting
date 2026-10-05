package kubernetes

import (
	"testing"

	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

var fromEnv = []config.EnvFrom{
	{SecretRef: "the-secret-name", SecretProviderClass: "cloud-secret-provider"},
}

func assertSecretVolumeSetup(t *testing.T, template *corev1.PodSpec) {
	volumes := template.Volumes
	assert.Len(t, volumes, 1)
	assert.Equal(t, corev1.Volume{
		Name: "the-secret-name",
		VolumeSource: corev1.VolumeSource{
			CSI: &corev1.CSIVolumeSource{
				Driver:   "secrets-store.csi.k8s.io",
				ReadOnly: ptr.To(true),
				VolumeAttributes: map[string]string{
					"secretProviderClass": "cloud-secret-provider",
				},
			},
		},
	}, volumes[0])

	containers := template.Containers
	assert.NotEmpty(t, containers)
	for _, container := range containers {
		assert.Equal(t, []corev1.EnvFromSource{
			{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "the-secret-name",
					},
				},
			},
		}, container.EnvFrom)
		assert.Equal(t, []corev1.VolumeMount{{
			Name:      "the-secret-name",
			ReadOnly:  true,
			MountPath: "/mnt/the-secret-name",
		}}, container.VolumeMounts)
	}

	containers = template.InitContainers
	assert.NotEmpty(t, containers)
	for _, container := range containers {
		assert.Equal(t, []corev1.EnvFromSource{
			{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "the-secret-name",
					},
				},
			},
		}, container.EnvFrom)
		assert.Equal(t, []corev1.VolumeMount{{
			Name:      "the-secret-name",
			ReadOnly:  true,
			MountPath: "/mnt/the-secret-name",
		}}, container.VolumeMounts)
	}
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can get handle a Pod
func Test_SetEnvironmentVariablesFromSecretsStore1(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers:     []corev1.Container{{}},
			InitContainers: []corev1.Container{{}},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a Deployment
func Test_SetEnvironmentVariablesFromSecretsStore2(t *testing.T) {
	object := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a DaemonSet
func Test_SetEnvironmentVariablesFromSecretsStore3(t *testing.T) {
	object := &appsv1.DaemonSet{
		Spec: appsv1.DaemonSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a ReplicaSet
func Test_SetEnvironmentVariablesFromSecretsStore4(t *testing.T) {
	object := &appsv1.ReplicaSet{
		Spec: appsv1.ReplicaSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a StatefulSet
func Test_SetEnvironmentVariablesFromSecretsStore5(t *testing.T) {
	object := &appsv1.StatefulSet{
		Spec: appsv1.StatefulSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a Job
func Test_SetEnvironmentVariablesFromSecretsStore6(t *testing.T) {
	object := &batchv1.Job{
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a CronJob
func Test_SetEnvironmentVariablesFromSecretsStore7(t *testing.T) {
	object := &batchv1.CronJob{
		Spec: batchv1.CronJobSpec{
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers:     []corev1.Container{{}},
							InitContainers: []corev1.Container{{}},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.JobTemplate.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure we can handle a ReplicationController
func Test_SetEnvironmentVariablesFromSecretsStore8(t *testing.T) {
	object := &corev1.ReplicationController{
		Spec: corev1.ReplicationControllerSpec{
			Template: &corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:     []corev1.Container{{}},
					InitContainers: []corev1.Container{{}},
				},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assertSecretVolumeSetup(t, &object.Spec.Template.Spec)
}

// SetEnvironmentVariablesFromSecretsStore Ensure all containers are updated
func Test_SetEnvironmentVariablesFromSecretsStore9(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{},
				{},
				{},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assert.Equal(t, []corev1.EnvFromSource{{
		SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "the-secret-name"}},
	}}, object.Spec.Containers[0].EnvFrom)
	assert.Equal(t, []corev1.EnvFromSource{{
		SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "the-secret-name"}},
	}}, object.Spec.Containers[1].EnvFrom)
	assert.Equal(t, []corev1.EnvFromSource{{
		SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "the-secret-name"}},
	}}, object.Spec.Containers[2].EnvFrom)
}

// SetEnvironmentVariablesFromSecretsStore Ensure empty definitions don't panic
func Test_SetEnvironmentVariablesFromSecretsStore10(t *testing.T) {
	object := &corev1.ReplicationController{}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assert.True(t, true)
}

// SetEnvironmentVariablesFromSecretsStore Ensure any existing envFrom with the same secretRef.name is not overwritten
func Test_SetEnvironmentVariablesFromSecretsStore11(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				EnvFrom: []corev1.EnvFromSource{{
					Prefix:       "prefix_",
					ConfigMapRef: &corev1.ConfigMapEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "ExistingConfig"}},
					SecretRef:    &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "the-secret-name"}},
				}},
			}},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	got := object.Spec.Containers[0].EnvFrom
	assert.Equal(t, []corev1.EnvFromSource{{
		Prefix:       "prefix_",
		ConfigMapRef: &corev1.ConfigMapEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "ExistingConfig"}},
		SecretRef:    &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "the-secret-name"}},
	}}, got)
}

// SetEnvironmentVariablesFromSecretsStore Ensure existing volumes are not overwritten
func Test_SetEnvironmentVariablesFromSecretsStore12(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{{
				Name: "the-secret-name",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/existing/stuff",
					},
				}},
			},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assert.Equal(t, []corev1.Volume{{
		Name: "the-secret-name",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/existing/stuff",
			},
		}},
	}, object.Spec.Volumes)
}

// SetEnvironmentVariablesFromSecretsStore Ensure existing volume mounts are not overwritten
func Test_SetEnvironmentVariablesFromSecretsStore13(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				VolumeMounts: []corev1.VolumeMount{{
					Name:      "the-secret-name",
					MountPath: "/mnt/existing-container",
					ReadOnly:  false,
				}},
			}},
			InitContainers: []corev1.Container{{
				VolumeMounts: []corev1.VolumeMount{{
					Name:      "the-secret-name",
					MountPath: "/mnt/existing-init-container",
					ReadOnly:  false,
				}},
			}},
		},
	}

	SetEnvironmentVariablesFromSecretsStore(object, fromEnv)
	assert.Equal(t, []corev1.VolumeMount{{
		Name:      "the-secret-name",
		MountPath: "/mnt/existing-container",
		ReadOnly:  false,
	}}, object.Spec.Containers[0].VolumeMounts)
	assert.Equal(t, []corev1.VolumeMount{{
		Name:      "the-secret-name",
		MountPath: "/mnt/existing-init-container",
		ReadOnly:  false,
	}}, object.Spec.InitContainers[0].VolumeMounts)
}
