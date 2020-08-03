package logic

var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

// GetSize ...
func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]

	diff := newScores - oldScores

	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < 2000.0 {
			return size
		}
		return 2000.0
	} else if diff > -5.0 && diff <= 0.0 {
		return 100.0 + 18.0*diff
	} else {
		return 10.0
	}
}

//SetScore ...
func SetScore(a float64) bool {
	pastFourScores = append(pastFourScores, a)
	pastFourScores = pastFourScores[1:]
	return true
}
