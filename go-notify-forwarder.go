package main

import (
	"fmt"
	"os"
	"strings"
	"log"
	"github.com/godbus/dbus/v5"
	"github.com/gregdel/pushover"
	"github.com/spf13/viper"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()
	app := pushover.New(viperEnvVariable("PUSHOVER_APP_ID"))
	recipient := pushover.NewRecipient(viperEnvVariable("PUSHOVER_RECIPIENT"))

	rules := []string{
		//"type='signal',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		"type='method_call',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		//"type='method_return',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		//"type='error',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}
	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)
	if call.Err != nil {
		fmt.Fprintln(os.Stderr, "Failed to become monitor:", call.Err)
		os.Exit(1)
	}

	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	fmt.Println("Monitoring notifications")
	for v := range c {
		go forward(v, app, recipient);
	}
}

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
	  fmt.Println("Config file not found, trying enviroment variables")
	}
	value, ok := viper.Get(key).(string)
  
	if !ok {
	  log.Fatalf("Invalid type assertion; ENV var probably not set")
	}
  
	return value
  }

func forward(msg *dbus.Message, app *pushover.Pushover, recipient *pushover.Recipient) {
	if msg.Type == dbus.TypeMethodCall {
		fmt.Println("Forwarding notification", msg.Serial() )
		var name = msg.Body[0];
	    var summary = msg.Body[3];
		var body string = fmt.Sprint(msg.Body[4])
		body  = strings.Replace(body, "\n", "", -1)

		message := &pushover.Message{
			Message:     body,
			Title:       fmt.Sprint(name," : " , summary) ,
			HTML:		 true}
		response, err := app.SendMessage(message, recipient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(response)
	}
}

