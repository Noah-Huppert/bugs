package main

import (
	"fmt"

	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

func main() {
	fmt.Printf("%#v\n", v1alpha1.SchemeBuilder) // Random print just to use pkg
}
