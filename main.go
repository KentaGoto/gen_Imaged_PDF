package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func runCommand(paths []string, imgResolution string) string {
	wg := &sync.WaitGroup{}

	flag := 0 // PDFファイルが有るかどうかのフラグ

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			//log.Println(runtime.NumGoroutine()) // goroutineの数
			defer wg.Done()
			ext := strings.LastIndex(path, ".") // 拡張子（.pdf）
			// PDFが存在したら以下の処理をする
			if path[ext:] == ".pdf" {
				flag = 1
				pdfDir := filepath.Dir(path)            // PDFファイルのディレクトリパス
				filename := getFileNameWithoutExt(path) // ファイル名
				saveDir := pdfDir + "/" + filename      // 抽出画像を保存するフォルダ
				err1 := os.Mkdir(saveDir, 0777)         // PDFファイル名のフォルダを作る
				if err1 != nil {
					panic(err1)
				}
				// mutoolを叩いて画像出力
				imageFile := saveDir + "/" + "p%04d.png"
				cmd := exec.Command("mutool.exe", "draw", "-o", imageFile, "-r", imgResolution, path)
				err2 := cmd.Run()
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("\n", path) // 処理後のファイルをフルパスで出力
				// 実行コマンドのステータスを表示
				state := cmd.ProcessState
				fmt.Printf("  %s\n", state.String())               // 終了コードと状態
				fmt.Printf("    Pid: %d\n", state.Pid())           // プロセスID
				fmt.Printf("    System: %v\n", state.SystemTime()) // システム時間（カーネル内で行われた処理の時間）
				fmt.Printf("    User: %v\n", state.UserTime())     // ユーザー時間（プロセス内で消費された時間）
			}
		}(path)
	}
	wg.Wait() // goroutineが終わるまで待つ

	// PDFが無ければ終了
	if flag == 0 {
		fmt.Println("PDF file is missing.")
		os.Exit(1)
	}

	return "\nDone."
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func main() {
	// 引数が2つ以外は終了
	if len(os.Args) != 3 {
		fmt.Println("The number of arguments specified is incorrect. Only one argument is allowed.")
		os.Exit(1)
	}

	dir := os.Args[1]           // 第1引数（処理対象ルートディレクトリ）
	imgResolution := os.Args[2] // 解像度
	paths := dirwalk(dir)       // ルートディレクトリを再帰で読みにいく

	// ファイルが何もなければ終了
	if paths == nil {
		fmt.Println("File is missing.")
		os.Exit(1)
	}

	fmt.Println("Processing...")

	// コマンド起動
	result := runCommand(paths, imgResolution)

	fmt.Println("\n" + result)
}
