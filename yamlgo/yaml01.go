package main

import (
     "fmt"
     "io/ioutil"
     "log"

     "gopkg.in/yaml.v3"
)

func main() {

     // read the contents of, "startrek.yaml"
     yfile, err := ioutil.ReadFile("startrek.yaml")

     if err != nil {

          log.Fatal(err)
     }

     // A map in which we read the data is defined
     data := make(map[any]any)

     // Unmarshal the data into the map
     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {

          log.Fatal(err2)
     }

     fmt.Println(data["flagship"]) // 1701

     for k, v := range data {

          fmt.Printf("%s -> %d\n", k, v)
     }
}

