# google-podcasts-desktop

***Please note:** in meantime [Google has rolled out changes to their backend which make this tool redundant](https://github.com/AdnanHodzic/google-podcasts-desktop/issues/1). Now, Google Podcasts URL’s shared via its Android App are automatically redirected to their web counterpart when run on desktop.*

google-podcasts-desktop app allows you to paste podcast link, after which it'll automatically convert it to desktop (web) friendly URL and open it in your default web browser. 

### How to run google-podcasts-desktop

* Download binary for your OS from [releases page](https://github.com/AdnanHodzic/google-podcasts-web-url/releases). Supported platforms are: [Linux](https://github.com/AdnanHodzic/google-podcasts-desktop/releases/download/0.2/google-podcasts-desktop-linux), [MacOS](https://github.com/AdnanHodzic/google-podcasts-desktop/releases/download/0.2/google-podcasts-desktop-mac), [Windows](https://github.com/AdnanHodzic/google-podcasts-desktop/releases/download/0.2/google-podcasts-desktop.exe).
* Run google-podcasts-desktop from terminal, i.e: 
`./google-podcast-desktop-linux`
* Paste [Google Podcasts URL](https://www.google.com/podcasts?feed=aHR0cDovL2pvZXJvZ2FuZXhwLmpvZXJvZ2FuLmxpYnN5bnByby5jb20vcnNz&episode=N2U0ZTEzZDUyZjE4NDNlYzkxNDhkZDhhZTgzYTI0ODY) (copied from Android app)
* Enjoy :)

![Alt Text align="center"](https://github.com/user-attachments/assets/f5018dae-fe1c-4ce5-833a-3320dc505f6e)

Linux/MacOS users should consider moving binary to `/usr/bin` directory for easy access. 
I.e: `sudo mv google-podcasts-desktop-linux /usr/bin/google-podcasts-desktop`.

After which app can be run by executing `google-podcasts-desktop` anywhere in terminal.

### Current state of affair with Google Podcats on desktop

Currently [Google Podcasts](https://podcasts.google.com/about) is only available as an [Android app](https://play.google.com/store/apps/details?id=com.google.android.apps.podcasts) and on [Google Home](https://store.google.com/gb/product/google_home) devices.

If you were listening to Google Podcasts on your phone and you wanted to continue on desktop, things would quickly get tricky. Only way this would be possible is by copying podcast URL from Google Podcasts Android App to Notes app of your choice, manually [modifying it](https://9to5google.com/2019/03/20/google-podcasts-desktop-web-app/) each time and opening modified link in a web browser.

### Technical

google-podcasts-desktop is a [simple Go (golang) CLI app](https://github.com/AdnanHodzic/google-podcasts-desktop/blob/master/google-podcasts-desktop.go) which will automatically make necessary changes to link copied from Google Podcasts Android app to get a web friendly version of same URL and openening it in a new tab of default web browser.

If you have Golang installed and setup, you can run google-podcasts-desktop from source by running:
`go run google-podcasts-desktop.go`

Code is released under GPL3 open source license and any changes you may have are most welcome.

### Discussion

[google-podcasts-desktop app - Listen to Google Podcasts on your desktop!](https://foolcontrol.org/?p=3095)
