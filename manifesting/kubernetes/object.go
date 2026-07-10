package kubernetes

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
)

func GetObject(data []byte) (runtime.Object, error) {

	newscheme := runtime.NewScheme()
	err := scheme.AddToScheme(newscheme)
	if err != nil {
		return nil, err
	}

	decoder := serializer.NewCodecFactory(newscheme).UniversalDeserializer()
	object, _, err := decoder.Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}

	return object, nil
}
