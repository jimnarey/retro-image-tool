package main


/*
#include<stdio.h>
#include "src/bchunk.c"
int fortytwo() {
    return 42;
}
*/
import "C"


func main() {

	var _bin *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.bin")
	var _cue *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.cue")
	var _target *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.iso")

	var arguments [3]*C.char
	arguments[0] = _bin
	arguments[1] = _cue
	arguments[2] = _target


	var argLength C.int
	argLength = C.int(len(arguments))
	firstValue := &(arguments[0])
	C._bchunk(argLength, firstValue)

}