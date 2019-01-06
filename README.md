# Police Police - Mashed Potatoes

This application will query the open Swedetish police events API, listening in on what happens in the fair
kingdom of Sweden!

## Installation

_(These instructions are written for a Windows environment. If you use Linux then you either know how to adapt to these instructions or you switch back to Windows)_

Clone the repository:

```
git clone https://github.com/Morkalork/police-police-mashed-potatoes
```

Enter the directory and build the project:

```
go build cmd/ppmp/
```

This will generate a main.exe file which you can execute in whatever way you find most suiting
(such as `./main.exe`)

You can also install the software with the following command:

```
go install ./...
```

After this the package will be installed under `$GOPATH/bin/ppmp.exe` and can be executed by just entering `ppmp` in your terminal. Use the --help flag to get help, such as:

```
ppmp --help                         #Show the help
ppmp --hours=48                     #Show the events for the last 48 hours
ppmp --location=Stockholm           #Show the events only for Stockholm
ppmp --hours=12 --crimetype=Brand   #Show only the last 12 hours and only fires (Brand)
``` 
Since the Api will only allow for you to fetch the last 500 events that your search renders, try to be specific in your search.

It may look like this:

![Police Police Mashed Potatoes in action!](/resources/screenshots/ppmp.png)

## Tests

There are tests available which can be executed by running the following command:

```
go test ./... -v
```

To get coverage, execute the following command:

```
go test -coverprofile=c.out ./...
```

## Why?

I'm learning go, that's why. 

## No no dumb snut, why the name?

Oh, because of the swell old Swedish police film ["Polis, polis potatismos"](https://www.imdb.com/title/tt0107846/) with 
Gösta Ekman as DC Beck. Good times!