package main

import (
	"testing"
	"time"
)

func TestFormattingCharts(t *testing.T) {
	type test struct {
		name         string
		transactions []*Transaction
		intervalType IntervalType
		expected     []*Transaction
		err          error
	}

	tests := []test{
		{
			name:         "Test empty slice",
			transactions: []*Transaction{},
			intervalType: Hour,
			expected:     []*Transaction{},
		},
		{
			name:         "Test unkown interval",
			transactions: []*Transaction{},
			intervalType: 5,
			expected:     []*Transaction{},
			err:          UnknownIntervalError{},
		},
		{
			name: "Test hour interval",
			transactions: []*Transaction{
				{
					Value:     4000,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5000,
					Timestamp: time.Date(2024, 1, 1, 2, 1, 1, 0, time.Local),
				},
				{
					Value:     6000,
					Timestamp: time.Date(2024, 1, 1, 3, 1, 1, 0, time.Local),
				},
				{
					Value:     3999,
					Timestamp: time.Date(2024, 1, 1, 1, 2, 1, 0, time.Local),
				},
				{
					Value:     4999,
					Timestamp: time.Date(2024, 1, 1, 2, 2, 1, 0, time.Local),
				},
				{
					Value:     5999,
					Timestamp: time.Date(2024, 1, 1, 3, 2, 1, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 1, 1, 3, 1, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 1, 1, 2, 3, 1, 0, time.Local),
				},
				{
					Value:     5998,
					Timestamp: time.Date(2024, 1, 1, 3, 3, 1, 0, time.Local),
				},
			},
			intervalType: Hour,
			expected: []*Transaction{
				{
					Value:     5998,
					Timestamp: time.Date(2024, 1, 1, 3, 0, 0, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 1, 1, 2, 0, 0, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 1, 1, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "Test day interval",
			transactions: []*Transaction{
				{
					Value:     4456,
					Timestamp: time.Unix(1616026248, 0),
				},
				{
					Value:     4231,
					Timestamp: time.Unix(1616022648, 0),
				},
				{
					Value:     5212,
					Timestamp: time.Unix(1616019048, 0),
				},
				{
					Value:     4321,
					Timestamp: time.Unix(1615889448, 0),
				},
				{
					Value:     4567,
					Timestamp: time.Unix(1615871448, 0),
				},
			},
			intervalType: Day,
			expected: []*Transaction{
				{
					Value:     4456,
					Timestamp: time.Unix(1616025600, 0),
				},
				{
					Value:     4231,
					Timestamp: time.Unix(1615939200, 0),
				},
				{
					Value:     4321,
					Timestamp: time.Unix(1615852800, 0),
				},
			},
		},
		{
			name: "Test week interval",
			transactions: []*Transaction{
				{
					Value:     4000,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5000,
					Timestamp: time.Date(2024, 1, 8, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     6000,
					Timestamp: time.Date(2024, 1, 15, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     3999,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5999,
					Timestamp: time.Date(2024, 1, 15, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     4999,
					Timestamp: time.Date(2024, 1, 8, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5998,
					Timestamp: time.Date(2024, 1, 15, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 1, 8, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
			},
			intervalType: Week,
			expected: []*Transaction{
				{
					Value:     5998,
					Timestamp: time.Date(2024, 1, 15, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 1, 8, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
			},
		},
		{
			name: "Test month interval",
			transactions: []*Transaction{
				{
					Value:     4000,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5000,
					Timestamp: time.Date(2024, 2, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     6000,
					Timestamp: time.Date(2024, 3, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     4999,
					Timestamp: time.Date(2024, 2, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5999,
					Timestamp: time.Date(2024, 3, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     3999,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     5998,
					Timestamp: time.Date(2024, 3, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 2, 1, 1, 1, 1, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 1, 1, 1, 1, 0, time.Local),
				},
			},
			intervalType: Month,
			expected: []*Transaction{
				{
					Value:     5998,
					Timestamp: time.Date(2024, 3, 0, 0, 0, 0, 0, time.Local),
				},
				{
					Value:     4998,
					Timestamp: time.Date(2024, 2, 0, 0, 0, 0, 0, time.Local),
				},
				{
					Value:     3998,
					Timestamp: time.Date(2024, 1, 0, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := FormattingCharts(tt.transactions, tt.intervalType)
			if err != tt.err {
				t.Logf("expected error: %v", err)
				t.FailNow()
			}
			if len(actual) != len(tt.expected) {
				t.Log("slices has different length")
				t.FailNow()
			}
			for i, expectedT := range tt.expected {
				actualT := actual[i]
				if expectedT.Value != actualT.Value && expectedT.Timestamp != actualT.Timestamp {
					t.Logf("expected: %v; actual: %v as position: %v", expectedT, actualT, i)
					t.Fail()
				}
			}
		})
	}
}
