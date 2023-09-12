<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2023-09-12 19:06:10
 * @LastEditTime: 2023-09-12 19:24:03
 * @LastEditors: FunctionSir
 * @Description: README.md
 * @FilePath: /BatchPeakNormise/README.md
-->
# BPN

Batch Peak Normalization using FFmpeg.  

## How to use

### 1. Install FFmpeg

#### Linux

Use your package manager.  
For Arch Linux, it is:  

```shell
sudo pacman -S ffmpeg
```

#### Windows

1. Get FFmpeg from ffmpeg.org.  
2. Copy ffmpeg.exe to System32. (You may need admin permissions to do that)  

### 2. Use it

Open a terminal, run:

```shell
bpn xxx yyy zzz
```

P.S. xxx/yyy/zzz is dirs, musics are in these dirs.  
P.S. Child dirs will be ignored.  

### 3. Get processed files

If xxx.ogg is the original file, xxx.new.ogg will be the name of processed file.  

## How to build

### On linux

Build for Linux:  

```shell
go build main.go
```

Build for Windows:  

```shell
GOOS=windows go build main.go
```
