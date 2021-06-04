# YouTube Data Fetcher
## Introduction
This is a front-end application made using React and ant.design. It takes in a YouTube playlist url and retrieves its metadata using YouTube API. Then, the metadata is shown in a grid format. Clicking on the video will link to the video. It's hosted on Netlify. This website can be viewed live at [yt.nimitpatel.me](https://yt.nimitpatel.me). Demo for the application can be viewed [here](https://youtu.be/l-i_7hmStG4)

## File/Folder Info
### **go**
* Contains Go lambda functions that use the YouTube API

### **MainPage.js**
* Contains the header and search bar
* Playlist ID is parsed
* Data is retrived from backend API

### **index.less**
* CSS for the application

### **PlaylistView.js**
* New view that is rendered after a search is executed
