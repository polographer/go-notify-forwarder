# go-notify-forwarder
A freedesktop notification forwarder written in go

## what is this?
The idea for this software was born when I started using a rapsberry pi with an iPad pro. I wanted a way to see the notifications of the raspberry on iOS.

## how it works?
The way it works is very simple it set up's itself as a Monitor of the freedesktop notification bus, and when a message arrives it formats it and sends it to iOS via an app called "pushover"

## what do you need
A raspberry, an ipad or ios, a paid pushover account (it's just 5 bucks)

## does it work
yes

## whats next
- make proper releases
- enable multiple outputs/notification systems
- make the code prettier, this is my first go program bear with me 
- write tests (not in order of importance)
- create init scripts
- (maybe) implement cobra

## where is the config ?
Environment variables, you need to set ;

`PUSHOVER_API_KEY -> your app key`
`PUSHOVER_USER_KEY -> the key of your device`

you can also create a.env file
