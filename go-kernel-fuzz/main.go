package main

import (
	"fmt"
    "os"
	"syscall"

)

type READMSG struct  {
	Address uint64  
	Size uint32
}

type READRSP struct  {
	Data uint64  
}

var SwapAddr uint64

func main() {

	
	testFileName := "\\\\.\\vulnerable"
	
	var data uint64 = 0
	//os.Remove(testFileName)
	f, err := os.Create(testFileName)
	if err != nil {
		fmt.Println(err)
	}
	//defer f.Close()

	
	fuzzIOctl(f)
	fmt.Println(fmt.Sprintf("response to fuzzing : %x\n",  data))
	
}

func fuzzIOctl(f *os.File) bool {	

	// start with 48 bytes buffers, adjust size as you progress with reversing the driver
	bufin := make([]byte, 48)
	bufout := make([]byte, 48)
	var dwBytesReturned uint32
	
	// get cyclic pattern generator from pwntools
	// from pwn import *
	// a = cyclic_gen()
	// generate 48 bytes for bufin
	// a.get(48) 
	// And 48 more bytes to initialise bufout
	// Don't trust the driver to use bufout only as output
	
	copy(bufin,[]byte("aaaabaaacaaadaaaeaaafaaagaaahaaaiaaajaaakaaalaaa"))
	copy(bufout, []byte("maaanaaaoaaapaaaqaaaraaasaaataaauaaavaaawaaaxaaa"))
	
	
	//return false
	err := syscall.DeviceIoControl(syscall.Handle(f.Fd()), 0x222808, &bufin[0], 48, &bufout[0], 48, &dwBytesReturned, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if (dwBytesReturned != 0) {
		fmt.Println("kernel sent something")
		fmt.Println(bufout)
		return true
	}


	return false
	
	

}

