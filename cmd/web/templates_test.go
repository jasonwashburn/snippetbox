package main

import (
	"testing"
	"time"

	"github.com/jasonwashburn/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2023 at 10:15",
		},

		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2023 at 09:15",
		},
		{
			name: "LeapYear",
			tm:   time.Date(2024, 2, 29, 10, 15, 0, 0, time.UTC),
			want: "29 Feb 2024 at 10:15",
		},
		{
			name: "PST",
			tm:   time.Date(2023, 3, 17, 10, 15, 0, 0, time.FixedZone("PST", -8*60*60)),
			want: "17 Mar 2023 at 18:15",
		},
		{
			name: "EndOfMonth",
			tm:   time.Date(2023, 3, 31, 10, 15, 0, 0, time.UTC),
			want: "31 Mar 2023 at 10:15",
		},
		{
			name: "StartOfYear",
			tm:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			want: "01 Jan 2023 at 00:00",
		},
		{
			name: "EndOfYear",
			tm:   time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
			want: "31 Dec 2023 at 23:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
