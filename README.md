# YouTube Data Fetcher UI
## Introduction
This is a front-end application made using React and ant.design. It takes in a YouTube playlist url and shows the videos in a grid format. Clicking on the video will link to the video. It calls an AWS lambda function that returns the meta-data about the playlist. This website can be viewed live at [yt.nimitpatel.me](https://yt.nimitpatel.me).

## File Info
### **MainPage.js**
* Contains the header and search bar
* Playlist ID is parsed
* Data is retrived from backend API

### **index.less**
* CSS for the application

### **PlaylistView.js**
* New view that is rendered after a search is executed

## Technology Used
- React.js
- HTML
- CSS
- JavaScript

## Steps To Reproduce
* use "npm install" in the folder to install packages
* use "npm start" to run the application
* use "npm run-script build" to build a static HTML website
* Backend API will need to be modified in MainPage.js

## Links
* [Backend Lambda Function](https://github.com/nimitpatel26/YouTube-Data-Fetcher)
* [Project Demo](https://youtu.be/VKYjhC71Sxg)
