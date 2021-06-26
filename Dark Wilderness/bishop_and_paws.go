package main

import "math"

func bishopAndPawn(bishop string, pawn string) bool {
    var bishopPos = []rune(bishop)
    var pawnPos = []rune(pawn)
    
    var selisihNo = math.Abs(float64(int(bishopPos[0]) - int(pawnPos[0])))
    var selisihChar = math.Abs(float64(int(bishopPos[1]) - int(pawnPos[1])))
    
    if (selisihNo == selisihChar) {
        return true
    } else {
        return false
    }
}
