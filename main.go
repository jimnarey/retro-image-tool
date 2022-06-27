package main


/*
#define DEBUG
#include<stdio.h>
#include "src/bchunk.c"
#include "src/iso2opl.h"
#include "src/iso2opl.c"
#include "src/isofs.c"

int fortytwo() {
    return 42;
}
*/
import "C"
import (
	"fmt"
	"os"
)

func main() {

	var _bchunk_exe *C.char = C.CString("bchunk")
	var _raw *C.char = C.CString("-r")
	var _verbose *C.char = C.CString("-v")
	var _bin *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.bin")
	var _cue *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.cue")
	var _target *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/cd_bin_cue/Unreal Tournament/Unreal Tournament.iso")

	var bc_arguments [6]*C.char
	bc_arguments[0] = _bchunk_exe
	bc_arguments[1] = _raw
	bc_arguments[2] = _verbose
	bc_arguments[3] = _bin
	bc_arguments[4] = _cue
	bc_arguments[5] = _target


	var argLength C.int
	argLength = C.int(len(bc_arguments))
	firstValue := &(bc_arguments[0])
	//C._bchunk(argLength, firstValue)

	var _iso2opl_exe *C.char = C.CString("iso2opl")
	var _iso *C.char = C.CString("/media/jimnarey/HDD_Data_B/Retro/Workdir/PS2/iso_dvd/Killzone (PAL).iso")
	var _media *C.char = C.CString("/media/jimnarey/7C06-C755/")
	var _title *C.char = C.CString("Black")
	var _disc_type *C.char = C.CString("DVD")

	var i2_arguments [5]*C.char
	i2_arguments[0] = _iso2opl_exe
	i2_arguments[1] = _iso
	i2_arguments[2] = _media
	i2_arguments[3] = _title
	i2_arguments[4] = _disc_type

	env_vars := os.Environ()
	//num_vars:= len(env_vars)

	c_env_vars := make([]*C.char, len(env_vars))

	//var c_env_vars []*C.char

	for i:= 0; i < len(env_vars); i++ {
		fmt.Println(env_vars[i])
		c_env_vars[i] = C.CString(env_vars[i])
	}

	argLength = C.int(len(i2_arguments))
	firstValue = &(i2_arguments[0])
	env := &(c_env_vars[0])

	C._iso2opl(argLength, firstValue, env)

}