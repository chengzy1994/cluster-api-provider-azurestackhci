/*
Copyright 2020 The Kubernetes Authors.
Portions Copyright © Microsoft Corporation.

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

package controllers

import (
	"encoding/json"

	infrav1 "github.com/microsoft/cluster-api-provider-azurestackhci/api/v1alpha4"
)

// updateMachineAnnotationJSON updates the `annotation` on `machine` with
// `content`. `content` in this case should be a `map[string]interface{}`
// suitable for turning into JSON. This `content` map will be marshalled into a
// JSON string before being set as the given `annotation`.
func (r *AzureStackHCIMachineReconciler) updateMachineAnnotationJSON(machine *infrav1.AzureStackHCIMachine, annotation string, content map[string]interface{}) error {
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}

	r.updateMachineAnnotation(machine, annotation, string(b))
	return nil
}

// updateMachineAnnotation updates the `annotation` on the given `machine` with
// `content`.
func (r *AzureStackHCIMachineReconciler) updateMachineAnnotation(machine *infrav1.AzureStackHCIMachine, annotation string, content string) {
	// Get the annotations
	annotations := machine.GetAnnotations()

	// Set our annotation to the given content.
	annotations[annotation] = content

	// Update the machine object with these annotations
	machine.SetAnnotations(annotations)
}

// Returns a map[string]interface from a JSON annotation.
// This method gets the given `annotation` from the `machine` and unmarshalls it
// from a JSON string into a `map[string]interface{}`.
func (r *AzureStackHCIMachineReconciler) machineAnnotationJSON(machine *infrav1.AzureStackHCIMachine, annotation string) (map[string]interface{}, error) {
	out := map[string]interface{}{}

	jsonAnnotation := r.machineAnnotation(machine, annotation)
	if len(jsonAnnotation) == 0 {
		return out, nil
	}

	err := json.Unmarshal([]byte(jsonAnnotation), &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// Fetches the specific machine annotation.
func (r *AzureStackHCIMachineReconciler) machineAnnotation(machine *infrav1.AzureStackHCIMachine, annotation string) string {
	return machine.GetAnnotations()[annotation]
}
