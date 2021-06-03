
import React from 'react';
import './index.less';


// Playlist grid page
function PlaylistView(props){
    return (
        <div className="app">
            <h1>YouTube Data Fetcher</h1>
            <PlaylistContainer {...props}/>
        </div>
    );
}

// Playlist grid
function PlaylistContainer(props){
    return (
        <div className="playlistContainer">
            {props.playlistVideos}

        </div>
    );
}

export default PlaylistView
