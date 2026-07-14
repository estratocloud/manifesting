package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

var defaultEnvVars = map[string]corev1.EnvVar{
	"KEY_NAME1": {Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	"KEY_NAME2": {Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	"KEY_NAME3": {Name: "KEY_NAME3", Value: "DEFAULT_VALUE3"},
}

func assertPodSpec(t *testing.T, template *corev1.PodTemplateSpec, expected []corev1.EnvVar, expectedInit []corev1.EnvVar) {
	containers := template.Spec.Containers
	assert.NotEmpty(t, containers)
	for _, container := range containers {
		assert.Equal(t, expected, container.Env)
	}

	containers = template.Spec.InitContainers
	assert.NotEmpty(t, containers)
	for _, container := range containers {
		assert.Equal(t, expectedInit, container.Env)
	}
}

// SetEnvironmentVariableDefaults Ensure we can get handle a Pod
func Test_SetEnvironmentVariableDefaults1(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Env: []corev1.EnvVar{
						{Name: "KEY_NAME1"},
					},
				},
			},
			InitContainers: []corev1.Container{
				{
					Env: []corev1.EnvVar{
						{Name: "KEY_NAME2"},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	got := object.Spec.Containers[0].Env
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, got)
	got = object.Spec.InitContainers[0].Env
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	}, got)
}

// SetEnvironmentVariableDefaults Ensure we can handle a Deployment
func Test_SetEnvironmentVariableDefaults2(t *testing.T) {
	object := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a DaemonSet
func Test_SetEnvironmentVariableDefaults3(t *testing.T) {
	object := &appsv1.DaemonSet{
		Spec: appsv1.DaemonSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a ReplicaSet
func Test_SetEnvironmentVariableDefaults4(t *testing.T) {
	object := &appsv1.ReplicaSet{
		Spec: appsv1.ReplicaSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a StatefulSet
func Test_SetEnvironmentVariableDefaults5(t *testing.T) {
	object := &appsv1.StatefulSet{
		Spec: appsv1.StatefulSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a Job
func Test_SetEnvironmentVariableDefaults6(t *testing.T) {
	object := &batchv1.Job{
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a CronJob
func Test_SetEnvironmentVariableDefaults7(t *testing.T) {
	object := &batchv1.CronJob{
		Spec: batchv1.CronJobSpec{
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Env: []corev1.EnvVar{
										{Name: "KEY_NAME1"},
									},
								},
							},
							InitContainers: []corev1.Container{
								{
									Env: []corev1.EnvVar{
										{Name: "KEY_NAME2"},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, &object.Spec.JobTemplate.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure we can handle a ReplicationController
func Test_SetEnvironmentVariableDefaults8(t *testing.T) {
	object := &corev1.ReplicationController{
		Spec: corev1.ReplicationControllerSpec{
			Template: &corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME1"},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "KEY_NAME2"},
							},
						},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assertPodSpec(t, object.Spec.Template, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	})
}

// SetEnvironmentVariableDefaults Ensure envvars that already have a value aren't overwritten
func Test_SetEnvironmentVariableDefaults9(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Env: []corev1.EnvVar{
						{Name: "KEY_NAME1", Value: "CUSTOM_VALUE1"},
						{Name: "KEY_NAME2"},
						{Name: "KEY_NAME3", ValueFrom: &corev1.EnvVarSource{}},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	got := object.Spec.Containers[0].Env
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "CUSTOM_VALUE1"},
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
		{Name: "KEY_NAME3", ValueFrom: &corev1.EnvVarSource{}},
	}, got)
}

// SetEnvironmentVariableDefaults Ensure empty definitions don't panic
func Test_SetEnvironmentVariableDefaults10(t *testing.T) {
	object := &corev1.ReplicationController{}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assert.True(t, true)
}

// SetEnvironmentVariableDefaults Ensure all containers are updated
func Test_SetEnvironmentVariableDefaults11(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Env: []corev1.EnvVar{{Name: "KEY_NAME1"}}},
				{Env: []corev1.EnvVar{{Name: "KEY_NAME2"}}},
				{Env: []corev1.EnvVar{{Name: "KEY_NAME3"}}},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME1", Value: "DEFAULT_VALUE1"},
	}, object.Spec.Containers[0].Env)
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME2", Value: "DEFAULT_VALUE2"},
	}, object.Spec.Containers[1].Env)
	assert.Equal(t, []corev1.EnvVar{
		{Name: "KEY_NAME3", Value: "DEFAULT_VALUE3"},
	}, object.Spec.Containers[2].Env)
}

// SetEnvironmentVariableDefaults Ensure an undefined envvar doesn't cause a panic
func Test_SetEnvironmentVariableDefaults12(t *testing.T) {
	object := &corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Env: []corev1.EnvVar{
						{Name: "NO_SUCH_KEY"},
					},
				},
			},
		},
	}

	SetEnvironmentVariableDefaults(object, defaultEnvVars)
	got := object.Spec.Containers[0].Env
	assert.Equal(t, []corev1.EnvVar{
		{Name: "NO_SUCH_KEY"},
	}, got)
}
