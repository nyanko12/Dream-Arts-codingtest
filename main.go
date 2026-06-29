package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 始点ID，終点ID，距離を格納するための構造体
type path struct {
	start_id int
	end_id   int
	distance float64
}

// 行き先のIDと距離を格納するための構造体
type edge struct {
	next int
	cost float64
}

var visited = map[int]bool{} // 訪れた頂点を記録するための変数
var edges = map[int][]edge{} // 各ID毎の次に進むIDをまとめた配列
var cur_path []int           // IDの進行状態を管理するための配列
var goal_path []int          // 最長経路を通るIDを保存する配列
var max_dis = 0.0            // 最長経路を記録するための変数

// DFS関数
func dfs(cur int, cur_dis float64) {
	visited[cur] = true
	cur_path = append(cur_path, cur)

	if cur_dis > max_dis {
		max_dis = cur_dis

		// 配列の場合は単純な代入だとポインタがコピーされてしまいcur_pathの変更に合わせてgoal_pathも変化してしまう
		goal_path = make([]int, len(cur_path)) // 新しく配列を作ることで独立させる
		copy(goal_path, cur_path)
	}

	// 入力IDをもとに次のIDに進む
	for _, e := range edges[cur] {
		// 訪問済みでない場合のみさらに深い階層へ潜る
		if !visited[e.next] {
			dfs(e.next, cur_dis+e.cost)
		}
	}

	// 元の階層に戻るため進めた部分を戻す
	visited[cur] = false
	cur_path = cur_path[:len(cur_path)-1]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	starts := map[int]bool{} // 始点のIDの集合を記録する変数，重複箇所に関してはDFSで対応できるためまとめている

	// fmt.Println("始点のID（整数型），終点のID（整数型），距離（浮動小数点数）を入力してください")
	for sc.Scan() {
		parts := strings.Split(sc.Text(), ",")

		// 空行がある場合はスキップする
		if len(parts) < 3 {
			continue
		}

		str, _ := strconv.Atoi(strings.TrimSpace(parts[0]))           // 始点のIDを取得
		end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))           // 終点のIDを取得
		dis, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64) // 距離を取得

		starts[str] = true

		// 各IDの次に進むIDとそのときの距離を記録
		edges[str] = append(edges[str], edge{end, dis})
	}

	// dfsを呼び出して全パスのルートの距離をチェックする
	for s := range starts {
		dfs(s, 0)
	}

	for i := 0; i < len(goal_path); i++ {
		fmt.Println(goal_path[i])
	}
}
