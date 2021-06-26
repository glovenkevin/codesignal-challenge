package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hasilnya harus salah, result: ", chessBoardCellColor("A1", "H3"))
	fmt.Println("hasilnya harus benar, result: ", chessBoardCellColor("A1", "C3"))
}

func chessBoardCellColor(cell1 string, cell2 string) bool {

	cell1Color := getChessColor(cell1)
	cell2Color := getChessColor(cell2)

	return cell1Color == cell2Color
}

func getChessColor(cell string) string {
	chessCharacter := strings.Split("ABCDEFGH", "")
	var cellHorizontal, cellVertical int

	cellArray := strings.Split(cell, "")
	cellVertical, _ = strconv.Atoi(cellArray[1])

	for i := 0; i < len(chessCharacter); i++ {
		if cellArray[0] == chessCharacter[i] {
			cellHorizontal = i + 1
		}
	}

	if (cellHorizontal%2 == 1 && cellVertical%2 == 1) || (cellHorizontal%2 == 0 && cellVertical%2 == 0) {
		return "black"
	} else {
		return "white"
	}
}
