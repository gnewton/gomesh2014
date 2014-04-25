
package main
 
import (
	"fmt"
	"runtime"
	"github.com/gnewton/jianGoMeSHi"
)


//////
func main() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	
	if false {
		_, err := jianGoMeSHi.SupplementalMapFromFile("/home/newtong/2014/mesh/supp2014.xml.bz2")
		if err != nil{
			fmt.Println("error: ", err)
			return
		}

		/////////////
		_, err = jianGoMeSHi.QualifierMapFromFile("/home/newtong/2014/mesh/qual2014.xml.bz2")
		if err != nil{
			fmt.Println("error: ", err)
			return
		}
	}
	
	descMap, err := jianGoMeSHi.DescriptorMapFromFile("/home/newtong/2014/mesh/desc2014.xml.bz2")
	//descMap, err := DescriptorMapFromFile("/home/newtong/gocode/jiangomeshi/testData/desc2014_29records.xml.bz2")
	if err != nil{
		fmt.Println(err)
		return
	}


	jianGoMeSHi.SelfLinkDescriptor(descMap)

	
	/*
	jsonOutFile,err  := os.Create("./out.json2.gz")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	jsonOutWriter := bufio.NewWriter(jsonOutFile)
	jsonOut := gzip.NewWriter(jsonOutWriter)


			b, err := json.Marshal(record)
			if err != nil {
				fmt.Println("error:", err)
			}

			m[record.SupplementalRecordUI] = record
			//fmt.Println(record.SupplementalRecordUI)
			if len(m)%10000==0{
				fmt.Println(len(m))
			}
			//jsonOut.Write(b)
			//fmt.Printf("%+v\n", record)
		}
	}
	jsonOut.Flush()
	jsonOut.Close()

	jsonOutWriter.Flush()

	jsonOutFile.Sync()
	jsonOutFile.Close()
}
	 */
}
