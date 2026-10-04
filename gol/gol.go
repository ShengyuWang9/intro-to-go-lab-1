package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	nextWorld := make([][]byte, len(world))

	for y, row := range world {
		nextWorld[y] = make([]byte, len(row))
	}

	for y, row := range world {
		for x, current := range row {
			count := 0

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					neighbourX := (x + dx + len(row)) % len(row)
					neighbourY := (y + dy + len(world)) % len(world)

					neighbour := world[neighbourY][neighbourX]

					if neighbour == byte(0xFF) {
						count++
					}
				}
			}

			nextState := byte(0x00)

			if (current == byte(0xFF) && (count == 2 || count == 3)) ||
				(current == byte(0x00) && count == 3) {
				nextState = byte(0xFF)
			}
			nextWorld[y][x] = nextState

		}

	}

	return nextWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	activeCells := []cell{}

	for y, row := range world {
		for x, value := range row {
			if value == byte(0xFF) {
				activeCells = append(activeCells, cell{x: x, y: y})
			}
		}

	}
	return activeCells
}
