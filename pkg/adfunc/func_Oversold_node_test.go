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

func TestAdd(t *testing.T) {
	//tests := []struct {
	//	a        resource.Quantity
	//	b        resource.Quantity
	//	expected resource.Quantity
	//}{
	//	{decQuantity(10, 0, resource.DecimalSI), decQuantity(1, 1, resource.DecimalSI), decQuantity(20, 0, resource.DecimalSI)},
	//	{decQuantity(10, 0, resource.DecimalSI), decQuantity(1, 0, resource.BinarySI), decQuantity(11, 0, resource.DecimalSI)},
	//	{decQuantity(10, 0, resource.BinarySI), decQuantity(1, 0, resource.DecimalSI), decQuantity(11, 0, resource.BinarySI)},
	//	{resource.Quantity{Format: resource.DecimalSI}, decQuantity(50, 0, resource.DecimalSI), decQuantity(50, 0, resource.DecimalSI)},
	//	{decQuantity(50, 0, resource.DecimalSI), resource.Quantity{Format: resource.DecimalSI}, decQuantity(50, 0, resource.DecimalSI)},
	//	{resource.Quantity{Format: resource.DecimalSI}, resource.Quantity{Format: resource.DecimalSI}, decQuantity(0, 0, resource.DecimalSI)},
	//}
	//
	//for i, test := range tests {
	//	test.a.Add(test.b)
	//	if test.a.Cmp(test.expected) != 0 {
	//		t.Errorf("[%d] Expected %q, got %q", i, test.expected.String(), test.a.String())
	//	}
	//}
}
