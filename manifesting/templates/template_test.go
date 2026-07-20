package templates

import (
	"testing"

	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewTemplate Ensure we get an error if there's no template
func Test_NewTemplate1(t *testing.T) {
	resource := &config.Resource{Name: "myapp", Template: "no-such-file"}
	got, err := NewTemplate(resource, nil, newWorkingDirectory(t))
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to read template for resource 'myapp': file '/app/tests/samples/templates/no-such-file.yaml.gotmpl' does not exist")
}

// NewTemplate Ensure we get an error if the template isn't parsable
func Test_NewTemplate2(t *testing.T) {
	resource := &config.Resource{Template: "bad-syntax"}
	got, err := NewTemplate(resource, nil, newWorkingDirectory(t))
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to parse header for template 'bad-syntax': error converting YAML to JSON: yaml: line 3: found unexpected end of stream")
}

// NewTemplate Ensure we get an error for invalid template syntax
func Test_NewTemplate3(t *testing.T) {
	resource := &config.Resource{Template: "bad-template"}
	got, err := NewTemplate(resource, nil, newWorkingDirectory(t))
	assert.Nil(t, got)
	assert.EqualError(t, err, `unable to parse the template 'bad-template': template: template:1: function "unknown" not defined`)
}

// Render Ensure we can render a basic template
func Test_Render1(t *testing.T) {
	wd := newWorkingDirectory(t)
	resource := &config.Resource{Template: "no-header-doc"}
	template, err := NewTemplate(resource, nil, wd)
	require.NoError(t, err)

	got, err := template.Render(map[string]any{}, nil, wd)
	require.NoError(t, err)
	assert.Equal(t, &batchv1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1",
		},
	}, got)
}

// Render Ensure we can render a template that extends another
func Test_Render2(t *testing.T) {
	wd := newWorkingDirectory(t)
	resource := &config.Resource{Template: "extends"}
	template, err := NewTemplate(resource, nil, wd)
	require.NoError(t, err)

	got, err := template.Render(map[string]any{}, nil, wd)
	require.NoError(t, err)
	assert.Equal(t, &batchv1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "base",
			Name:      "override",
		},
	}, got)
}
