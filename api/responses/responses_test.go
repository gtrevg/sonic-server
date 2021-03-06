package responses_test

import (
	"testing"
	"time"

	. "github.com/cloudsonic/sonic-server/api/responses"
	. "github.com/cloudsonic/sonic-server/tests"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsonicResponses(t *testing.T) {

	response := &Subsonic{Status: "ok", Version: "1.0.0"}

	Convey("Subject: Subsonic Responses", t, func() {
		Convey("EmptyResponse", func() {
			Convey("XML", func() {
				So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"></subsonic-response>`)
			})
			Convey("JSON", func() {
				So(response, ShouldMatchJSON, `{"status":"ok","version":"1.0.0"}`)
			})
		})

		Convey("License", func() {
			response.License = &License{Valid: true}
			Convey("XML", func() {
				So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><license valid="true"></license></subsonic-response>`)
			})
			Convey("JSON", func() {
				So(response, ShouldMatchJSON, `{"license":{"valid":true},"status":"ok","version":"1.0.0"}`)
			})
		})

		Convey("MusicFolders", func() {
			response.MusicFolders = &MusicFolders{}

			Convey("With data", func() {
				folders := make([]MusicFolder, 2)
				folders[0] = MusicFolder{Id: "111", Name: "aaa"}
				folders[1] = MusicFolder{Id: "222", Name: "bbb"}
				response.MusicFolders.Folders = folders

				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><musicFolders><musicFolder id="111" name="aaa"></musicFolder><musicFolder id="222" name="bbb"></musicFolder></musicFolders></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"musicFolders":{"musicFolder":[{"id":"111","name":"aaa"},{"id":"222","name":"bbb"}]},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><musicFolders></musicFolders></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"musicFolders":{},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("Indexes", func() {
			artists := make([]Artist, 1)
			artists[0] = Artist{Id: "111", Name: "aaa"}
			response.Indexes = &Indexes{LastModified: "1", IgnoredArticles: "A"}

			Convey("With data", func() {
				index := make([]Index, 1)
				index[0] = Index{Name: "A", Artists: artists}
				response.Indexes.Index = index
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><indexes lastModified="1" ignoredArticles="A"><index name="A"><artist id="111" name="aaa"></artist></index></indexes></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"indexes":{"ignoredArticles":"A","index":[{"artist":[{"id":"111","name":"aaa"}],"name":"A"}],"lastModified":"1"},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><indexes lastModified="1" ignoredArticles="A"></indexes></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"indexes":{"ignoredArticles":"A","lastModified":"1"},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("Child", func() {
			response.Directory = &Directory{Id: "1", Name: "N"}
			Convey("With all data", func() {
				child := make([]Child, 1)
				t := time.Date(2016, 03, 2, 20, 30, 0, 0, time.UTC)
				child[0] = Child{
					Id: "1", IsDir: true, Title: "title", Album: "album", Artist: "artist", Track: 1,
					Year: 1985, Genre: "Rock", CoverArt: "1", Size: "8421341", ContentType: "audio/flac",
					Suffix: "flac", TranscodedContentType: "audio/mpeg", TranscodedSuffix: "mp3",
					Duration: 146, BitRate: 320, Starred: &t,
				}
				response.Directory.Child = child
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><directory id="1" name="N"><child id="1" isDir="true" title="title" album="album" artist="artist" track="1" year="1985" genre="Rock" coverArt="1" size="8421341" contentType="audio/flac" suffix="flac" starred="2016-03-02T20:30:00Z" transcodedContentType="audio/mpeg" transcodedSuffix="mp3" duration="146" bitRate="320"></child></directory></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"directory":{"child":[{"album":"album","artist":"artist","bitRate":320,"contentType":"audio/flac","coverArt":"1","duration":146,"genre":"Rock","id":"1","isDir":true,"size":"8421341","starred":"2016-03-02T20:30:00Z","suffix":"flac","title":"title","track":1,"transcodedContentType":"audio/mpeg","transcodedSuffix":"mp3","year":1985}],"id":"1","name":"N"},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("Directory", func() {
			response.Directory = &Directory{Id: "1", Name: "N"}
			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><directory id="1" name="N"></directory></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"directory":{"id":"1","name":"N"},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("With just required data", func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.Directory.Child = child
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><directory id="1" name="N"><child id="1" isDir="false" title="title"></child></directory></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"directory":{"child":[{"id":"1","isDir":false,"title":"title"}],"id":"1","name":"N"},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("AlbumList", func() {
			response.AlbumList = &AlbumList{}
			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><albumList></albumList></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"albumList":{},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("With just required data", func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.AlbumList.Album = child
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><albumList><album id="1" isDir="false" title="title"></album></albumList></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"albumList":{"album":[{"id":"1","isDir":false,"title":"title"}]},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("User", func() {
			response.User = &User{Username: "deluan"}
			Convey("Without optional fields", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><user username="deluan" scrobblingEnabled="false" adminRole="false" settingsRole="false" downloadRole="false" uploadRole="false" playlistRole="false" coverArtRole="false" commentRole="false" podcastRole="false" streamRole="false" jukeboxRole="false" shareRole="false" videoConversionRole="false"></user></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"status":"ok","user":{"adminRole":false,"commentRole":false,"coverArtRole":false,"downloadRole":false,"jukeboxRole":false,"playlistRole":false,"podcastRole":false,"scrobblingEnabled":false,"settingsRole":false,"shareRole":false,"streamRole":false,"uploadRole":false,"username":"deluan","videoConversionRole":false},"version":"1.0.0"}`)
				})
			})
			Convey("With optional fields", func() {
				response.User.Email = "cloudsonic@deluan.com"
				response.User.Folder = []int{1}
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><user username="deluan" email="cloudsonic@deluan.com" scrobblingEnabled="false" adminRole="false" settingsRole="false" downloadRole="false" uploadRole="false" playlistRole="false" coverArtRole="false" commentRole="false" podcastRole="false" streamRole="false" jukeboxRole="false" shareRole="false" videoConversionRole="false"><folder>1</folder></user></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"status":"ok","user":{"adminRole":false,"commentRole":false,"coverArtRole":false,"downloadRole":false,"email":"cloudsonic@deluan.com","folder":[1],"jukeboxRole":false,"playlistRole":false,"podcastRole":false,"scrobblingEnabled":false,"settingsRole":false,"shareRole":false,"streamRole":false,"uploadRole":false,"username":"deluan","videoConversionRole":false},"version":"1.0.0"}`)
				})
			})
		})
		Convey("Playlists", func() {
			response.Playlists = &Playlists{}

			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><playlists></playlists></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"playlists":{},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("With data", func() {
				pls := make([]Playlist, 2)
				pls[0] = Playlist{Id: "111", Name: "aaa"}
				pls[1] = Playlist{Id: "222", Name: "bbb"}
				response.Playlists.Playlist = pls

				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><playlists><playlist id="111" name="aaa"></playlist><playlist id="222" name="bbb"></playlist></playlists></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"playlists":{"playlist":[{"id":"111","name":"aaa"},{"id":"222","name":"bbb"}]},"status":"ok","version":"1.0.0"}`)
				})
			})
		})

		Convey("Playlist", func() {
			response.Playlist = &PlaylistWithSongs{}
			response.Playlist.Id = "1"
			response.Playlist.Name = "My Playlist"
			Convey("Without data", func() {
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><playlist id="1" name="My Playlist"></playlist></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"playlist":{"id":"1","name":"My Playlist"},"status":"ok","version":"1.0.0"}`)
				})
			})
			Convey("With just required data", func() {
				entry := make([]Child, 1)
				entry[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.Playlist.Entry = entry
				Convey("XML", func() {
					So(response, ShouldMatchXML, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.0.0"><playlist id="1" name="My Playlist"><entry id="1" isDir="false" title="title"></entry></playlist></subsonic-response>`)
				})
				Convey("JSON", func() {
					So(response, ShouldMatchJSON, `{"playlist":{"entry":[{"id":"1","isDir":false,"title":"title"}],"id":"1","name":"My Playlist"},"status":"ok","version":"1.0.0"}`)
				})
			})
		})
		Reset(func() {
			response = &Subsonic{Status: "ok", Version: "1.0.0"}
		})
	})

}
