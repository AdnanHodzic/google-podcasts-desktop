# google-podcasts-desktop

google-podcats-desktop app allows you to paste podcast link, after which it'll automatically convert it to desktop (web) friendly URL and open it in your default web browser. 

### How to run google-podcats-desktop

* Download binary for your OS from [releases page](https://github.com/AdnanHodzic/google-podcasts-web-url/releases). Supported platforms are: Linux, MacOS, Windows.
* Run google-podcasts-desktop from terminal, i.e: 
`./google-podcast-desktop-linux`
* Paste [Google Podcasts URL](https://www.google.com/podcasts?feed=aHR0cDovL2pvZXJvZ2FuZXhwLmpvZXJvZ2FuLmxpYnN5bnByby5jb20vcnNz&episode=N2U0ZTEzZDUyZjE4NDNlYzkxNDhkZDhhZTgzYTI0ODY) (copied from Android app)

![Alt Text](https://foolcontrol.org/wp-content/uploads/2019/05/google-podcasts-desktop-screencast.gif)

Linux/MacOS users should consider moving binary to `/usr/bin` directory for easy access. 
I.e: `sudo mv google-podcasts-desktop-linux /usr/bin/google-podcasts-desktop`.

After which app can be run by executing `google-podcasts-desktop` anywhere in terminal.

### Current state of affair with Google Podcats on desktop

Currently [Google Podcasts](https://podcasts.google.com/about) is only available as an [Android app](https://play.google.com/store/apps/details?id=com.google.android.apps.podcasts) and on [Google Home](https://store.google.com/gb/product/google_home) devices.

If you were listening to Google Podcasts on your phone and you wanted to continue on desktop, things would quickly get tricky. Only way this would be possible is by copying podcast URL from Google Podcasts Android App, manually [modifying it](https://9to5google.com/2019/03/20/google-podcasts-desktop-web-app/) each time and opening modified link in a web browser.  

### Technical

google-podcasts-desktop is a simple Go (golang) CLI app which will automatically make necessary changes to link copied from Google Podcasts Android app to get a web friendly version of same URL. After which this URL will be opened in a new tab of default web browser.

If you have Golang installed and setup, you can run google-podcasts-desktop from source by running:
`go run google-podcasts-desktop.go`

### Discussion

[google-podcasts-desktop app - Listen to Google Podcasts on your desktop!](https://foolcontrol.org/?p=3095)
