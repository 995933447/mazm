package main

import (
	"fmt"
	"os"
)

//点位置
type point struct {
	i, j int //i纵坐标,j横左边
}

//广度迷宫算法四个方向：上左下右
var dirs = []point{
	{
		-1,0,
	},
	{
		0,-1,
	},
	{
		1, 0,
	},
	{
		0, 1,
	},
}

//迷宫入口
var start = point{0, 0}

//迷宫出口
var end = point{5, 4}

func main() {
	mazm := getMazm()

	fmt.Println("迷宫图(0代表道路,1代表墙):")
	printGrid(mazm)

	steps := getSteps(mazm)
	Q := []point{start}

	var next point
	for len(Q) > 0 {
		if next == end {
			break
		}

		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next = cur.add(dir)
			if val, isExist := next.at(mazm); !isExist || val == 1 {
				continue
			}

			if val, isExist := next.at(steps); !isExist || val != 0 {
				continue
			} else if next == start {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			Q = append(Q, next)
		}
	}

	fmt.Println("迷宫破解走势:")
	printGrid(steps)

	process := getStepProcess(steps, end)
	sort(&process)
	fmt.Println("迷宫最终出路为:")
	fmt.Println(process)
}

//把最终走势由倒序转为正序
func sort(process *[]point)  {
	sort := []point{}
	for i := len(*process) - 1; i >= 0; i-- {
		sort = append(sort, (*process)[i])
	}
	*process = sort
}

//打印方阵
func printGrid(grid [][]int)  {
	for i := range grid {
		for _, v := range grid[i] {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
	fmt.Println()
}

//获得迷宫破解最终路径
func getStepProcess(steps [][]int, end point) []point {
	Q := []point{end}
	process := []point{end}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		for _, dir := range dirs {
			next := cur.add(dir)
			nextVal, isExist := next.at(steps)
			if !isExist {
				continue
			}

			curVal, _ := cur.at(steps)
			if  curVal - nextVal == 1 {
				process = append(process, next)
				Q = append(Q, next)
			}
		}
	}

	return process
}

//获得当前点在图阵的位置
func (cur point) add(dir point) point {
	return point{cur.i + dir.i, cur.j + dir.j}
}

//获得当前所在方阵位置代表的数值
func (cur point) at(grid [][]int) (int, bool) {
	if cur.i < 0 || cur.i > len(grid) - 1 {
		return -1, false
	}

	if cur.j < 0 || cur.j > len(grid[cur.i]) - 1 {
		return -1, false
	}

	return grid[cur.i][cur.j], true
}

//获得迷宫图阵
func getMazm() [][]int {
	file, err := os.Open("mazm.ini")
	if err != nil {
		panic(err)
	}

	var rows, cols int
	fmt.Fscanf(file, "%d %d", &rows, &cols)
	mazm := make([][]int, rows)
	for i := range mazm {
		mazm[i] = make([]int, cols)
		for j := range mazm[i] {
			fmt.Fscanf(file, "%d", &mazm[i][j])
		}
	}

	return mazm
}

//获得迷宫破解走势记录图阵
func getSteps(mazm [][]int) [][]int {
	steps := [][]int{}
	for _ = range mazm {
		steps = append(steps, make([]int, len(mazm[0])))
	}
	return steps
}