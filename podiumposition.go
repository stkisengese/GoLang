package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := range podium {
		for j := 0; j < len(podium); j++ {
			for k := j + 1; k < len(podium[i]); k++ {
				if podium[i][j][0] < podium[i][k][0] {
					swap := podium[i][j]
					podium[i][j] = podium[i][k]
					podium[i][k] = swap
				}
			}
		}
	}
	return podium
}
