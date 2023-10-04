package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixSubMatrix(t *testing.T) {
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix, _ := matrix.NewFromElements(data)

	// Extract a submatrix
	subMatrix, err := matrix.SubMatrix(1, 1, 2, 2)
	if err != nil {
		t.Errorf("Error extracting submatrix: %v", err)
	}

	// Check the dimensions of the submatrix
	expectedRows := 2
	expectedColumns := 2
	rows := subMatrix.Rows()
	columns := subMatrix.Columns()

	if rows != expectedRows {
		t.Errorf("Expected %d rows in submatrix, but got %d", expectedRows, rows)
	}

	if columns != expectedColumns {
		t.Errorf("Expected %d columns in submatrix, but got %d", expectedColumns, columns)
	}

	// Check the values in the submatrix
	expectedData := [][]int{{5, 6}, {8, 9}}
	for i := 0; i < expectedRows; i++ {
		for j := 0; j < expectedColumns; j++ {
			val, _ := subMatrix.Get(i, j)
			if val != expectedData[i][j] {
				t.Errorf("Expected value %d at (%d, %d) in submatrix, but got %d", expectedData[i][j], i, j, val)
			}
		}
	}
}

func TestMatrixInvalidSubMatrix(t *testing.T) {
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix, _ := matrix.NewFromElements(data)

	// Attempt to extract an invalid submatrix
	_, err := matrix.SubMatrix(1, 1, 3, 3)

	// Check if an error is returned
	if err == nil {
		t.Error("Expected an error for invalid submatrix dimensions, but got nil")
	}

	// Check the error message
	expectedErrorMessage := "submatrix dimensions exceed matrix bounds"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
