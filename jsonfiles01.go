/* RZFeeser | Alta3 Research
Reading JSON - with structs  */

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

// Users struct which contains
// an array of users
type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users struct
    var november Users

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'november' which we defined above
    json.Unmarshal(byteValue, &november)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(november.Users); i++ {
        fmt.Println("User Type: " + november.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(november.Users[i].Age))
        fmt.Println("User Name: " + november.Users[i].Name)
        fmt.Println("Facebook Url: " + november.Users[i].Social.Facebook)
    }

}

