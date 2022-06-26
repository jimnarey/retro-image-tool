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

	var arg0 *C.char = C.CString("bchunk")
	var arg1 *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.bin")
	var arg2 *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.cue")
	var arg3 *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.iso")

	var arguments [4]*C.char
	arguments[0] = arg0
	arguments[1] = arg1
	arguments[2] = arg2
	arguments[3] = arg3


	var argLength C.int
	argLength = C.int(len(arguments))
	firstValue := &(arguments[0])
	C._bchunk(argLength, firstValue)

}