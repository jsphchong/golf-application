package main

import "time"

type Golfer struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
	Birthday  time.Time
	StartDate time.Time
	Stats     *Statistics
	Rounds    []Round
}

type Statistics struct {
	Eagles            int
	Birdies           int
	Pars              int
	Bogeys            int
	DoubleBogeys      int
	TripleBogeys      int
	FourOverAndHigher int
}

type Round struct {
	Id                    string
	CourseName            string
	StrokeNumPerHole      []int
	TotalStrokes          int
	StrokesOverParPerHole []int
	StrokesOverPar        int
	StatisticsFlag        bool
}
