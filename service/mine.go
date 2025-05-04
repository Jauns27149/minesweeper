package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"minesweeper/constant"
	"minesweeper/util"
	"strconv"
	"time"
)

type Mine struct {
	row      int
	col      int
	mines    int
	safeArea int

	Cells []cell
}

type cell struct {
	value  int
	button *widget.Button
}

var mines = map[string]*Mine{
	constant.Simple: {row: 9, col: 9, mines: 10},
	constant.Common: {row: 16, col: 16, mines: 40},
}

func (m *Mine) Run() {
	n := m.row * m.col
	m.Cells = make([]cell, n)
	for i := range n {
		m.Cells[i] = cell{
			button: widget.NewButton("", func() {
				c := m.Cells[i]
				if c.value == -1 {
					m.GameOver("踩雷了，是否重新开始")
				} else {
					m.safeArea--
					if m.safeArea == 0 {
						m.GameOver("恭喜，扫雷成功，是否重新开始")
					}
				}
				c.button.SetText(strconv.Itoa(c.value))
				c.button.Disable()
				if c.value == 0 {
					m.findSafe(i)
				}
			}),
		}
	}
	m.initCellValue()
}

func (m *Mine) GameOver(message string) {
	dialog.ShowConfirm("", message, func(confirmed bool) {
		if confirmed {
			m.initCellValue()
		} else {
			fyne.CurrentApp().Quit()
		}
	}, util.GetTopWindow())
}

/*
系统会从当前0格出发，向周围8个方向（上、下、左、右、4个对角）检查相邻格子。
如果相邻格子是未揭开的空白格（0），则自动揭开它，并继续以该格子为中心递归展开。
如果相邻格子是数字格（1~8），则揭开该格子但停止递归（数字格是展开的边界）。
如果相邻格子是地雷或已标记的旗子/问号，则跳过不揭开。
*/
func (m *Mine) findSafe(i int) {
	scope := m.getScope(i)
	for _, v := range scope {
		c := m.Cells[v]
		if !c.button.Disabled() && c.value >= 0 {
			c.button.SetText(strconv.Itoa(c.value))
			c.button.Disable()
			if c.value == 0 {
				m.findSafe(v)
			}

			m.safeArea--
			if m.safeArea == 0 {
				m.GameOver("恭喜，扫雷成功，是否重新开始")
			}
		}
	}
}

func (m *Mine) initCellValue() {
	m.safeArea = len(m.Cells) - m.mines
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(m.Cells))
	for i, v := range perm {
		if i < m.mines {
			m.Cells[v].value = -1
		} else {
			m.Cells[v].value = 0
		}
	}

	for i, c := range m.Cells {
		if c.value == 0 {
			m.Cells[i].value = m.Compute(i)
		}

		c.button.Enable()
		c.button.SetText("")
	}
}

func (m *Mine) getScope(i int) []int {
	// 九宫格下标
	temp := []int{
		i - m.col - 1, i - m.col, i - m.col + 1,
		i - 1, i, i + 1,
		i + m.col - 1, i + m.col, i + m.col + 1,
	}

	result := make([]int, 0, 9)
	for ii, v := range temp {
		if v < 0 || v >= len(m.Cells) {
			continue
		}
		if (ii+1)%3 == 1 && (v+1)%m.col == m.col {
			continue
		}
		if (ii+1)%3 == 0 && (v+1)%m.col == 1 {
			continue
		}
		result = append(result, v)
	}

	return result
}

func (m *Mine) Compute(i int) int {
	scope := m.getScope(i)
	var result int
	for _, v := range scope {
		if m.Cells[v].value == -1 {
			result++
		}
	}
	return result
}
