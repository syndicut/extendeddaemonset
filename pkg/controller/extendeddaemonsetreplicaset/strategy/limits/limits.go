// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package limits

// Parameters use to provide the parameters to the Calculation function
type Parameters struct {
	NbNodes            int
	NbPods             int
	NbAvailablesPod    int
	NbOldAvailablesPod int
	NbCreatedPod       int

	MaxPodCreation      int
	MaxUnavailablePod   int
	MaxUnschedulablePod int
}

// CalculatePodToCreateAndDelete from the parameters return:
// * nbCreation: the number of pods to create
// * nbDeletion: the number of pods to delete
func CalculatePodToCreateAndDelete(params Parameters) (nbCreation, nbDeletion int) {
	nbCreation = params.NbNodes - params.NbPods - params.NbOldAvailablesPod
	if nbCreation > params.MaxPodCreation {
		nbCreation = params.MaxPodCreation
	}
	nbCreation -= (params.NbCreatedPod - params.NbAvailablesPod)
	if nbCreation < 0 {
		nbCreation = 0
	}

	nbDeletion = params.MaxUnavailablePod - (params.NbNodes - params.NbAvailablesPod - params.NbOldAvailablesPod)
	if nbDeletion < 0 {
		nbDeletion = 0
	}

	if nbDeletion > params.NbPods+params.NbOldAvailablesPod {
		nbDeletion = params.NbPods + params.NbOldAvailablesPod
	}

	return nbCreation, nbDeletion
}
