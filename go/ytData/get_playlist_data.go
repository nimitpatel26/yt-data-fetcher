// Sample Go code for user authorization

package ytData

import (
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

// Request that is received from the client
// Contains the ID of the playlist whose meta data the client is requesting
type PlayListId struct {
	Id string `json:"id"`
}

// Stores the metadata info for a particular video
type VideoData struct {
	Title        string
	VideoId      string
	ThumbnailUrl string
}

// Response that will be sent back to the client
type Response struct {
	Data []VideoData `json:"data"`
}

/* Gets metadata for a playlist by its ID on a particular page
 * Also takes in youtube.Service that will be used to get data
 */
func getPlaylistDataOnPage(service *youtube.Service, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
	call := service.PlaylistItems.List([]string{"contentDetails", "id", "snippet", "status"}).PlaylistId(playlistId).MaxResults(50)
	if pageToken != "" {
		call.PageToken(pageToken)
	}
	resp, _ := call.Do()
	return resp
}

/* Extracts data returned from YouTube API into the VideoData struct
 * Returns an array of VideoData struct that has metadata for all videos on a page
 */
func extractData(resp *youtube.PlaylistItemListResponse) []VideoData {

	vDataList := []VideoData{}

	// Go through metadata returned from YouTube API for multiple videos
	for _, vData := range resp.Items {

		// Extract data needed to store in VideoData struct
		snippet := vData.Snippet
		tnUrl := "N/A"

		if snippet.Thumbnails.Maxres != nil {
			tnUrl = snippet.Thumbnails.Maxres.Url
		} else if snippet.Thumbnails.High != nil {
			tnUrl = snippet.Thumbnails.High.Url
		}

		vId := snippet.ResourceId.VideoId
		vTitle := snippet.Title

		// Append a new VideoData struct to the array
		vDataList = append(vDataList, VideoData{Title: vTitle, VideoId: vId, ThumbnailUrl: tnUrl})

	}
	return vDataList
}

/* Returns an response with metadata about a YouTube playlist
 * Takes in ID of the playlist
 */
func GetPlayListData(playlistId string) Response {
	ctx := context.Background()

	// For large playlists YouTube API returns data in pages
	// Following code will call YouTube API for all the pages and store the result in an array of VideoData
	service, _ := youtube.NewService(ctx, option.WithCredentialsJSON([]byte(os.Getenv("YT_API_KEY"))))
	resp := getPlaylistDataOnPage(service, playlistId, "")

	nextPageToken := resp.NextPageToken
	vDataList := extractData(resp)

	// Keep calling YouTube API until the last page
	for nextPageToken != "" {
		resp := getPlaylistDataOnPage(service, playlistId, nextPageToken)
		nextPageToken = resp.NextPageToken
		vDataList = append(vDataList, extractData(resp)...)
	}

	// Response returned to the client
	return Response{
		Data: vDataList,
	}
}
