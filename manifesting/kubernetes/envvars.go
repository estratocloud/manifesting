package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func SetEnvironmentVariableDefaults(object runtime.Object, defaults map[string]corev1.EnvVar) {

	switch o := object.(type) {
	case *corev1.Pod:
		processContainers(o.Spec.Containers, defaults)
		processContainers(o.Spec.InitContainers, defaults)

	case *appsv1.Deployment:
		processPodTemplate(&o.Spec.Template, defaults)

	case *appsv1.DaemonSet:
		processPodTemplate(&o.Spec.Template, defaults)

	case *appsv1.ReplicaSet:
		processPodTemplate(&o.Spec.Template, defaults)

	case *appsv1.StatefulSet:
		processPodTemplate(&o.Spec.Template, defaults)

	case *batchv1.Job:
		processPodTemplate(&o.Spec.Template, defaults)

	case *batchv1.CronJob:
		processPodTemplate(&o.Spec.JobTemplate.Spec.Template, defaults)

	case *corev1.ReplicationController:
		processPodTemplate(o.Spec.Template, defaults)

	}
}

func processPodTemplate(template *corev1.PodTemplateSpec, defaults map[string]corev1.EnvVar) {
	if template == nil {
		return
	}
	processContainers(template.Spec.Containers, defaults)
	processContainers(template.Spec.InitContainers, defaults)
}

func processContainers(containers []corev1.Container, defaults map[string]corev1.EnvVar) {
	for _, container := range containers {
		processContainer(container, defaults)
	}
}

func processContainer(container corev1.Container, defaults map[string]corev1.EnvVar) {
	for key := range container.Env {
		setEnvironmentVariableDefault(&container.Env[key], defaults)
	}
}

func setEnvironmentVariableDefault(envvar *corev1.EnvVar, defaults map[string]corev1.EnvVar) {

	// If the current already has a value then don't do anything
	if envvar.Value != "" || envvar.ValueFrom != nil {
		return
	}

	// Get the default for this env var
	defaultEnv, found := defaults[envvar.Name]

	// If there is no default then don't do anything
	if !found {
		return
	}

	// Provide the default value/valueFrom
	envvar.Value = defaultEnv.Value
	envvar.ValueFrom = defaultEnv.ValueFrom
}
