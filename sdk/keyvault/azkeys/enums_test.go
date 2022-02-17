//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeys

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPtrMethods(t *testing.T) {
	d := DeletionRecoveryLevelValues.CustomizedRecoverable()
	require.Equal(t, d.ToPtr(), &d)

	j := JSONWebKeyCurveNameValues.P256()
	require.Equal(t, j.ToPtr(), &j)

	o := JSONWebKeyOperationValues.Decrypt()
	require.Equal(t, o.ToPtr(), &o)

	a := KeyEncryptionAlgorithmValues.RSAAESKEYWRAP256()
	require.Equal(t, a.ToPtr(), &a)
}

//nolint
func TestToGeneratedMethods(t *testing.T) {
	// If done incorrectly, this will have a nil pointer reference
	var l *ListKeyVersionsOptions = nil
	_ = l.toGenerated()
}
