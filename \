/* RZFeeser | Alta3 Research
   Without Pointers - Why we need pointers  */    

package main

import (
 "fmt"
)

type User struct {
 Name string
 Pets []string
}

func (u *User) newPet() {
 u.Pets = append(u.Pets, "Fido")                    // Simple function should ensure "Fido" is added to User
}

func main() {
 u := User{Name: "Anna", Pets: []string{"Bailey"}}
 u.newPet()                                         // {Anna [Bailey Fido]} -- This *should* add "Fido" to "u"
 fmt.Println(u)                                     // We **DO NOT** see Fido -- what happened?
}

