// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"regexp"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"

	"github.com/spf13/cobra"
)

//var cfgFile string
var registerName, registerPassword string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new User",
	Long: `register:register a new User

	For example:
	register a new user,with name:User1,password:12345678
	agenda register -n=User1 -p=12345678 
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		register(registerName, registerPassword)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&registerName, "name", "n", "", "user's name")
	registerCmd.Flags().StringVarP(&registerPassword, "password", "p", "", "user's password")

}

func register(name string, password string) {

	if isValidName(name) && isValidPassword(password) {

		var email, phone string
		fmt.Println("please input your email:")
		fmt.Scanln(&email)
		fmt.Println("please input your phone:")
		fmt.Scanln(&phone)

		if isValidEmail(email) && isValidPhone(phone) {
			users := datarw.GetUsers()
			if !hasName(name, users) {
				newuser := entity.User{Name: name, Password: password, Email: email, Phone: phone}
				users = append(users, newuser)
				datarw.SaveUsers(users)
				fmt.Println("Registration complete")
				return

			}

		}

	}

	fmt.Println("Register fail")

}

//Judge username exists
func hasName(name string, users []entity.User) bool {

	for _, user := range users {
		if user.Name == name {
			fmt.Println("The Username has been registered")
			return true
		}
	}

	return false
}
func isValidName(n string) bool {
	b := []byte(n)
	val, _ := regexp.Match(".+", b)
	if !val {
		fmt.Println("flag -n ,name is invaild")
	}
	return val
}
func isValidPassword(p string) bool {
	b := []byte(p)
	val, _ := regexp.Match(".+", b)
	if len(p) < 8 {
		fmt.Println("The password must be longer than 8 digits")
		val = false
	}
	if !val {
		fmt.Println("flag -p ,password is invaild")
	}
	return val
}
func isValidEmail(e string) bool {
	b := []byte(e)
	val, _ := regexp.Match("\\w*@\\w*\\.w*", b)

	if !val {
		fmt.Println("email is invaild")
	}
	return val
}
func isValidPhone(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match("[0-9]+", b)

	if !val {
		fmt.Println("phone is invaild")
	}
	return val
}
