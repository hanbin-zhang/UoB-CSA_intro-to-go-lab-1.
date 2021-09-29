package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}
	for n := 0; n < p.imageHeight; n++ {
		for i := 0; i < p.imageWidth; i++ {
			numberLive := 0

			for _, wide := range [3]int{i - 1, i, i + 1} {
				for _, height := range [3]int{n - 1, n, n + 1} {

					if wide < 0 {
						wide = p.imageWidth - 1
					} else if wide >= p.imageWidth {
						wide = 0
					}

					if height < 0 {
						height = p.imageHeight - 1
					} else if height >= p.imageHeight {
						height = 0
					}

					if world[height][wide] == byte(255) {
						numberLive += 1
					}
				}
			}

			if world[n][i] == byte(255) {
				numberLive -= 1
				if numberLive < 2 {
					newWorld[n][i] = byte(0)
				} else if numberLive > 3 {
					newWorld[n][i] = byte(0)
				} else {
					newWorld[n][i] = byte(255)
				}
			} else {
				if numberLive == 3 {
					newWorld[n][i] = byte(255)
				} else {
					newWorld[n][i] = byte(0)
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var list []cell

	for n := 0; n < p.imageHeight; n++ {
		for i := 0; i < p.imageWidth; i++ {
			if world[n][i] == byte(255) {
				list = append(list, cell{i, n})
			}
		}
	}
	return list
}

/*func test_valid(p golParams, coordinateType int, coordinate int) int {
	// wide
	result := 0
	if coordinateType == 0 {
		if coordinate < 0 || coordinate >= p.imageWidth {
			result = -1
		}
	}
	// height
	if coordinateType == 1 {
		if coordinate < 0 || coordinate >= p.imageHeight {
			result = -1
		}
	}
	return result
}*/
