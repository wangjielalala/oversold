package adfunc

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"testing"
)

func TestFunc(t *testing.T) {
	Allcalalbe := &v1.ResourceList{
		v1.ResourceCPU:    resource.MustParse("2"),
		v1.ResourceMemory: resource.MustParse("1Mi"),
	}

	klog.Info(overcpu(Allcalalbe.Cpu(), "2"))
}
