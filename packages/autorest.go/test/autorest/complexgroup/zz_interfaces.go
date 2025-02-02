//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

// DotFishClassification provides polymorphic access to related types.
// Call the interface's GetDotFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DotFish, *DotSalmon
type DotFishClassification interface {
	// GetDotFish returns the DotFish content of the underlying type.
	GetDotFish() *DotFish
}

// FishClassification provides polymorphic access to related types.
// Call the interface's GetFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Fish, *Goblinshark, *Salmon, *Sawshark, *Shark, *SmartSalmon
type FishClassification interface {
	// GetFish returns the Fish content of the underlying type.
	GetFish() *Fish
}

// MyBaseTypeClassification provides polymorphic access to related types.
// Call the interface's GetMyBaseType() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MyBaseType, *MyDerivedType
type MyBaseTypeClassification interface {
	// GetMyBaseType returns the MyBaseType content of the underlying type.
	GetMyBaseType() *MyBaseType
}

// SalmonClassification provides polymorphic access to related types.
// Call the interface's GetSalmon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Salmon, *SmartSalmon
type SalmonClassification interface {
	FishClassification
	// GetSalmon returns the Salmon content of the underlying type.
	GetSalmon() *Salmon
}

// SharkClassification provides polymorphic access to related types.
// Call the interface's GetShark() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Goblinshark, *Sawshark, *Shark
type SharkClassification interface {
	FishClassification
	// GetShark returns the Shark content of the underlying type.
	GetShark() *Shark
}
