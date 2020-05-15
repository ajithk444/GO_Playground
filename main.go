package main

import (
	"flag"
	"fmt"

	"github.com/mpetavy/go-dicom"
	"github.com/mpetavy/go-dicom/dicomtag"
)

var (
	extract = flag.Bool("x", false, "Extract PixelData")
)

func main() {

	//data, err := dicom.ReadDataSetFromFile("files/I_000010.dcm", //10200904.dcm",
	data, err := dicom.ReadDataSetFromFile("files/10200904.dcm",
    //data, err := dicom.ReadDataSetFromFile("files/CT-MONO2-16-ort.dcm", //I_000010.dcm", //10200904.dcm",
		dicom.ReadOptions{DropPixelData: !*extract})
	if err != nil {
		panic(err)
	}


	for _, elem := range data.Elements {
		
		if elem.Tag == dicomtag.PatientSex {
			fmt.Println("PatientSex data : ", elem.Value)
		}

		if elem.Tag == dicomtag.Manufacturer {
			fmt.Println("Manufacturer data : ", elem.Value)
		}

		if elem.Tag == dicomtag.ProtocolName {
			fmt.Println("ProtocolName data : ", elem.Value)
		}

	}

}
