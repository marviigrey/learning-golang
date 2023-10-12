package main

import (
	//"fmt"
	log "github.com/sirupsen/logrus"
	// "os"
)

func main() {
	/*	file, err := os.OpenFile("/home/ec2-user/environment/learning-golang/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) //created a file for the output of our log.
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	log.SetLevel(log.TraceLevel)
	log.Trace("hello world")
	//	log.SetOutput(file)

}
