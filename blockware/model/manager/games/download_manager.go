package games

/*


 <RequestDownload>  <--  [peer joins]
				|                     |
	[block found] - no - <DeferredRequest>
				|
			 yes
				|
[download block]
				|
<DownloadProgress>


*/
type DownloadManager struct {
	/*
		Download manager threads will send requests down this channel
		to prompt a peer listener to attempt to download the block
	*/
	RequestDownload chan DownloadRequest

	/*

	 */
	DeferredRequests chan DownloadRequest

	/*
		Once a block has been downloaded a message will be sent down
		this channel.
		This can be used to signal to the UI that a download has been
		completed
	*/
	DownloadProgress chan DownloadRequest
}

func NewDownloadManager() *DownloadManager {
	return &DownloadManager{
		RequestDownload:  make(chan DownloadRequest),
		DeferredRequests: make(chan DownloadRequest),
		DownloadProgress: make(chan DownloadRequest),
	}
}

func (d *DownloadManager) Close() {
	close(d.DeferredRequests)
	close(d.DownloadProgress)
	close(d.RequestDownload)
}
