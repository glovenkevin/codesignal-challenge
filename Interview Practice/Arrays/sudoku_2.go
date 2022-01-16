package main

import "fmt"

func main() {
	param := [][]string{
		{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}

	fmt.Println(sudoku2(param)) // Expected: true

	param = [][]string{
		{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}

	fmt.Println(sudoku2(param)) // Expected: false
}

func sudoku2(grid [][]string) bool {

	for pos := 0; pos < 9; pos++ {
		flagRow := make(map[string]bool)
		flagCol := make(map[string]bool)
		flagGrid := make(map[string]bool)
		for idx := 0; idx < 9; idx++ {
			if grid[pos][idx] != "." && flagRow[grid[pos][idx]] {
				return false
			}
			flagRow[grid[pos][idx]] = true

			if grid[idx][pos] != "." && flagCol[grid[idx][pos]] {
				return false
			}
			flagCol[grid[idx][pos]] = true

			if grid[pos-pos%3+idx/3][pos%3*3+idx%3] != "." && flagGrid[grid[pos-pos%3+idx/3][pos%3*3+idx%3]] {
				return false
			}
			flagGrid[grid[pos-pos%3+idx/3][pos%3*3+idx%3]] = true
		}

		if pos == 0 || pos == 3 || pos == 6 {
			flagGrid = make(map[string]bool)
		}
	}

	return true
}

/*

In Java by Henry11 - CodeSignal
Interesting idea with HashSet

boolean sudoku2(char[][] grid) {
    int n = grid.length;

    Set<String> set = new HashSet<String>();

    for (int row = 0; row < n ; row++) {
        for (int col = 0; col < n; col++) {
            if (grid[row][col] != '.' && !set.add(grid[row][col] + " in row " + row))
                return false;
            if (grid[row][col] != '.' && !set.add(grid[row][col] + " in col " + col))
                return false;
            if (grid[row][col] != '.' && !set.add(grid[row][col] + " in square " + row/3 + " " + col/3))
                return false;
        }
    }

    return true;
}
*/
