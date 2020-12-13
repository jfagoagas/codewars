package speedcontrol

/*
In John's car the GPS records every s seconds the distance travelled from an origin
(distances are measured in an arbitrary but consistent unit).

For example, below is part of a record with s = 15:

x = [0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25]

The sections are:

0.0-0.19, 0.19-0.5, 0.5-0.75, 0.75-1.0, 1.0-1.25, 1.25-1.50, 1.5-1.75, 1.75-2.0, 2.0-2.25

We can calculate John's average hourly speed on every section and we get:

[45.6, 74.4, 60.0, 60.0, 60.0, 60.0, 60.0, 60.0, 60.0]

Given s and x the task is to return as an integer the *floor* of the maximum average speed per hour obtained
on the sections of x. If x length is less than or equal to 1 return 0 since the car didn't move.

*/
func gps(s int, x []float64) (maxSpeed int) {
	var sections []float64
	maxSpeed = 0
	for i := 0; i < len(x)-1; i++ {
		// delta between distances
		element := x[i+1] - x[i]
		// hourly speed
		speed := ((float64(3600) * element) / float64(s))
		sections = append(sections, speed)
	}
	// find max element in slice and convert float64 to int
	for _, v := range sections {
		if int(v) > maxSpeed {
			maxSpeed = int(v)
		}
	}
	return
}
