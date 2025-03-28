build:
	go build -o bin/FGSB.exe -ldflags="-H windowsgui"
	xcopy themes bin\themes /s /e /d
	xcopy config.json bin
