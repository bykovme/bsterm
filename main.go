package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/bykovme/bslib"
	"github.com/bykovme/gotrans"
	"github.com/rivo/tview"
)

func main() {
	errLocales := gotrans.InitLocales("langs")
	if errLocales != nil {
		log.Fatal(errLocales)
	}
	errLocales = gotrans.SetDefaultLocale("ru")
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	appFolder := usr.HomeDir + "/" + cAppFolder
	dbFile := appFolder + "/" + cDbFile
	if _, err := os.Stat(appFolder); os.IsNotExist(err) {
		err := os.MkdirAll(appFolder, os.ModePerm)
		if err != nil {
			log.Fatal("Initiation error: " + err.Error())
			return
		}
	}
	bsInstance := bslib.GetInstance()
	log.Println("BSApp is initiating ")
	err = bsInstance.Open(dbFile)
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}
	defer func() {
		log.Println("BSApp stopped, closing storage")
		err := bsInstance.Close()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Storage is closed")
	}()

	fmt.Println(cAppName + " " + cTermAppVersion)

	app := tview.NewApplication()
	form := tview.NewForm().
		AddPasswordField("Master Password", "", 10, '*', nil).
		AddButton("Login", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle(cAppName).SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
