import React, {useState} from 'react';
import ReactDOM from "react-dom";
import PlaylistView from "./PlaylistView";

import {Input} from "antd";
const {Search } = Input;
function App() {
    return (
        <div className="app appCenter">
            <Header/>
            <Body/>
        </div>
    );
}

// Title
function Header(){
    return <h1>YouTube Data Fetcher</h1>
}

// Search Box
function Body(){
    const [searching, setSearching] = useState(false);

    return (
        <>
            <Search className="search" placeholder="Enter Playlist URL Here!" enterButton="Search" size="large"  loading={searching} onSearch={async (value) => {
                setSearching(true);
                await searchPlaylist(value, () => setSearching(false));

            }

            }/>
        </>
    );
}


// Get data for a particular playlist ID and change the view
async function searchPlaylist(value){
    let params = new URLSearchParams(value);
    let playlistId = params.get("list");

    if (playlistId === null){
        playlistId = params.get('https://www.youtube.com/playlist?list');
    }


    try {
        var api = "https://yt.nimitpatel.me/.netlify/functions/playlist_data";
        var url = api + "?id=" + playlistId;

        var videos = await fetch(url).then(response => response.json()).then(jsonData => {
            return jsonData.data.map((video) => {
                return (
                    <div className="item">
                        <a href={"https://www.youtube.com/watch?v=" + video.VideoId}>
                            <img src={video.ThumbnailUrl} alt={video.Title}/>
                        </a>
                    </div>

                );
            });
        });

        SetPlaylistView(videos);
    }catch (e) {
        console.log(e);

    }
}

function SetPlaylistView(videos){
    ReactDOM.render(
        <React.StrictMode>
            <PlaylistView playlistVideos = {videos}/>
        </React.StrictMode>,
        document.getElementById('root')
    );
}


export default App;
