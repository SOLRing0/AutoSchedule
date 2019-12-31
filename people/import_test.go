package people

import (
	"github.com/stretchr/testify/require"
	"github.com/tealeg/xlsx"
	"testing"
)
func TestSimpleImport(t *testing.T) {

	testFileName := "test/people-import.xlsx"
	xlFile, err := xlsx.OpenFile(testFileName)
	require.NoError(t, err)

	require.True(t, len(xlFile.Sheets) > 0)
}
