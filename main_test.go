package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCSV(t *testing.T) {
	const testfile string = "testdata/problems.csv"
	expected := []problem{
		{ "5+5","10"},
		{"1+1","2"},
		{"what 2+2, sir?", "4"},
	}

	t.Run("happy", func(t *testing.T) {
		problems, err := parseCSV(testfile)
		require.NoError(t, err)
		for i, problem := range problems {
			assert.Equal(t, expected[i], problem)
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		const badFile string = "testdata/doesnotexist.csv"
		_, err := parseCSV(badFile)
		require.Errorf(t, err, "open %s: no such file or directory", badFile)
	})
}

func TestParseProblems(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		expected := []problem{
			{"1+5","6"},
			{"2+2","4"},
		}

		csv := [][]string{
			{"1+5", "6"},
			{"2+2", "4"},
		}

		results, err := parseProblems(csv)
		require.NoError(t, err)
		require.Equal(t, expected, results)
	})

	t.Run("empty", func(t *testing.T) {
		csv := [][]string{}

		_, err := parseProblems(csv)
		require.EqualError(t, err, "no problems in csv file")
	})

	t.Run("invalid format", func(t *testing.T) {
		csv := [][]string{
			[]string{"1+5", "6"},
			[]string{"2+2", "4", "should not be here"},
		}

		_, err := parseProblems(csv)
		require.EqualError(t, err, "invalid csv structure on line 2")
	})
}
