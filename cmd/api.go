package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/apicmd"
	"os/exec"
	"strings"
)

// cronCmd represents the cron command
var handlerCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate code based on template files",
	Long: `Such as:
Url: /api/user/info/:id
Method: GET

Command: 'gotools api -m GET -d api/user -n info -p :id'
-m method: Indicates the method of the api
-d dir: corresponds to '/api/user' in the url
-n name: corresponds to '/info' in the url
-p params: corresponds to '/:id' int the url, and ':name/:id' can be used for multiple arguments
 `,
	Run: func(cmd *cobra.Command, args []string) {
		if strings.ToUpper(apicmd.VarStringMethod) == "RESTFUL" {
			err := command()
			if err != nil {
				fmt.Println("command err:", err)
			}
			return
		}
		app := &apicmd.Command{
			Command: cmd.CommandPath(),
		}
		err := app.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			//commandLog(app.Command)
			fmt.Println(app.Command, "Down .")
		}
	},
}

func command() error {
	err := restfulList()
	if err != nil {
		return err
	}
	err = restfulDetail()
	if err != nil {
		return err
	}
	err = restfulUpdate()
	if err != nil {
		return err
	}
	err = restfulCreate()
	if err != nil {
		return err
	}
	return restfulDelete()
}
func restfulCreate() error {
	command := exec.Command("gotools", "api", "-d", apicmd.VarStringDir, "-n", "create", "-m", "POST")
	_, err := command.Output()
	return err
}

func restfulDetail() error {
	command := exec.Command("gotools", "api", "-d", apicmd.VarStringDir, "-n", "detail", "-m", "GET", "-p", ":id")
	_, err := command.Output()
	return err
}
func restfulUpdate() error {
	command := exec.Command("gotools", "api", "-d", apicmd.VarStringDir, "-n", "update", "-m", "PUT", "-p", ":id")
	_, err := command.Output()
	return err
}
func restfulDelete() error {
	command := exec.Command("gotools", "api", "-d", apicmd.VarStringDir, "-n", "delete", "-m", "DELETE", "-p", ":id")
	_, err := command.Output()
	return err
}
func restfulList() error {
	command := exec.Command("gotools", "api", "-d", apicmd.VarStringDir, "-n", "list", "-m", "GET")
	_, err := command.Output()
	return err
}

func init() {
	handlerCmd.Flags().StringVarP(&apicmd.VarStringName, "name", "n", "", "The name of the API")
	handlerCmd.Flags().StringVarP(&apicmd.VarStringDir, "dir", "d", "", "The directory path to the API")
	handlerCmd.Flags().StringVarP(&apicmd.VarStringMethod, "method", "m", "POST", "The method of the API, default POST")
	handlerCmd.Flags().StringVarP(&apicmd.VarStringParams, "params", "p", "", "Route parameters contained in the api")
	rootCmd.AddCommand(handlerCmd)
}
