/*
 * @Author: FunctionSir (2023214592@sdtbu.edu.cn)
 * @License: AGPLv3
 * @Date: 2023-09-11 22:30:49
 * @LastEditTime: 2023-09-12 19:20:46
 * @LastEditors: FunctionSir
 * @Description: To do peak normise batchly.
 * @Thanks: ASCII art by figlet, FFmpeg cmd line from https://zhuanlan.zhihu.com/p/473952525.
 * @FilePath: /BatchPeakNormise/bpn.go
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	NAME          string = "BATCH PEAK NORMISE"
	VER           string = "0.1-alpha"
	CODENAME      string = "KuroNeko"
	LICENSE       string = "AGPLv3"
	AUTHOR        string = "FunctionSir (2023214592@sdtbu.edu.cn)"
	ASCII_ART     string = " ____  ____  _   _ \n| __ )|  _ \\| \\ | |\n|  _ \\| |_) |  \\| |\n| |_) |  __/| |\\  |\n|____/|_|   |_| \\_|"
	LAW_INFO      string = "This is a FOSS, and comes with ABSOLUTELY NO WARRANTY to the extent permitted by applicable law."
	SPLIT_LINE    string = "--------------------------------"
	FFMPEG_CHK_TS int    = 64
)

// Is it windows?
func is_windows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

// Find str in a []string.
func find_str(source []string, target string) int {
	for i := 0; i < len(source); i++ {
		if source[i] == target {
			return i
		}
	}
	return -1
}

// Let path unified.
func unify_path(p string, f bool) string {
	check_EOP := func(path string, flag bool) bool {
		if ((!strings.HasSuffix(path, "/")) && flag) || (strings.HasSuffix(path, "/") && (!flag)) {
			return false
		}
		return true
	}
	if is_windows() {
		p = strings.ReplaceAll(p, "\\", "/")
	}
	if check_EOP(p, false) && f {
		p = p + "/"
	}
	for check_EOP(p, true) && is_windows() && (!f) {
		p = p[:len(p)-2]
	}
	return p
}

// Do that first!
func initial() {
	fmt.Println(ASCII_ART)
	fmt.Println(SPLIT_LINE)
	fmt.Println(NAME)
	fmt.Println("Version: " + VER + " " + CODENAME)
	fmt.Println("Author: " + AUTHOR + " License: " + LICENSE)
	fmt.Println(LAW_INFO)
}

// Press Enter to continue...
func pause() {
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
}

// Check ffmpeg (very simple check).
func check_ffmpeg(ffmpeg string) bool {
	ff := exec.Command(ffmpeg)
	ffOut, _ := ff.CombinedOutput()
	if len(ffOut) < FFMPEG_CHK_TS {
		return false
	} else {
		return true
	}
}

// The main part.
func main() {
	initial()
	if len(os.Args) <= 1 {
		fmt.Println("[ERROR] NO ANY DIR PROVIDED.")
		pause()
		os.Exit(1)
	}
	ffmpeg := "ffmpeg"
	if !check_ffmpeg(ffmpeg) {
		fmt.Println("[ERROR] NO VAILD FFMPEG.")
		pause()
		os.Exit(1)
	}
	pause()
	for i := 1; i < len(os.Args); i++ {
		dir, err := os.ReadDir(os.Args[i])
		if err != nil {
			fmt.Println("[ERROR] " + err.Error())
		}
		for _, file := range dir {
			if !file.IsDir() {
				filePath := unify_path(os.Args[i], true) + file.Name()
				fmt.Println("Processing: " + filePath + "...")
				var ff *exec.Cmd
				if is_windows() {
					ff = exec.Command("cmd", "/c", ffmpeg, "-i", filePath, "-af", "volumedetect", "-vn", "-sn", "-dn", "-f", "null", "null")
				} else {
					ff = exec.Command(ffmpeg, "-i", filePath, "-af", "volumedetect", "-vn", "-sn", "-dn", "-f", "null", "null")
				}
				ffOut, _ := ff.CombinedOutput()
				sFFOut := string(ffOut)
				fmt.Println("[Begin FFmpeg Output]")
				fmt.Println(sFFOut)
				fmt.Println("[End FFmpeg Output]")
				splitedOut := strings.Split(sFFOut, " ")
				toGain := "-0.0"
				if find_str(splitedOut, "max_volume:") != -1 {
					toGain = splitedOut[find_str(splitedOut, "max_volume:")+1]
				}
				splitedFP := strings.Split(filePath, ".")
				splitedFP[len(splitedFP)-1] = "new." + splitedFP[len(splitedFP)-1]
				tmp := splitedFP[0]
				for i := 1; i < len(splitedFP); i++ {
					tmp = tmp + "." + splitedFP[i]
				}
				if is_windows() {
					ff = exec.Command("cmd", "/c", ffmpeg, "-y", "-i", filePath, "-af", "volume="+toGain[1:]+"dB", tmp)
				} else {
					ff = exec.Command(ffmpeg, "-y", "-i", filePath, "-af", "volume="+toGain[1:]+"dB", tmp)
				}
				ffOut, _ = ff.CombinedOutput()
				fmt.Println("[Begin FFmpeg Output]")
				fmt.Println(string(ffOut))
				fmt.Println("[End FFmpeg Output]")
			}
		}
	}
	fmt.Println("Done!")
	pause()
}
