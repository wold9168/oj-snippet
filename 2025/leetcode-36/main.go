package main

func judge(slice []byte) bool {
    var mp [10]bool
    for _, num := range slice {
        if num != '.' {
            if mp[num-'0'] == false {
                mp[num-'0'] = true
            } else {
                return false
            }
        }
    }
    return true
}

func checkCurCol(board [][]byte, colno int) bool {
    d := make([]byte, 0)
    for i := 0; i < 9; i++ {
        d = append(d, board[i][colno])
    }
    return judge(d)
}

func checkCurRow(board [][]byte, rowno int) bool {
    d := board[rowno][:]
    return judge(d)
}

func getInd(blockInd int) []int {
    const blockLength int = 3
    res := make([]int, 0)
    for i := 0; i < blockLength; i++ {
        res = append(res, blockInd*blockLength+i)
    }
    return res
}

func checkBlock(board [][]byte, colno int, rowno int) bool {
    d := make([]byte, 0)
    for _, c := range getInd(colno) {
        for _, r := range getInd(rowno) {
            d = append(d, board[r][c])
        }
    }
    return judge(d)
}

func isValidSudoku(board [][]byte) bool {
    var res bool = true
    for i := 0; i < 9; i++ {
        if checkCurCol(board, i) == false || checkCurRow(board, i) == false {
            res = false
            break
        }
    }

    if res == false {
        return res
    }

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if checkBlock(board, i, j) == false {
                res = false
                break
            }
        }
        if res == false {
            break
        }
    }
    return res
}
