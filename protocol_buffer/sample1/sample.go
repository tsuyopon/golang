package main

import (
    "log"

    "./protobuffer"
    "github.com/golang/protobuf/proto"
)

func main() {

    //Serialize
    test := &sample.User{
        Id:   33,
        Name: "laughing_cat",
    }
    buff, err := proto.Marshal(test)
    if err != nil {
        log.Print("marshal error")
        return
    }
    log.Print("serialize success: ", buff)

    // Deserialize
    parsedTest := &sample.User{}
    err = proto.Unmarshal(buff, parsedTest)

    if err != nil {
        log.Print("deserialize error: ", err)
    }
    log.Print("deserialized id :", parsedTest.Id)
    log.Print("deserialized name: ", parsedTest.Name)
}

