package engine

import (
	"math/rand"
	"time"

	"github.com/cloudsonic/sonic-server/domain"
	"github.com/cloudsonic/sonic-server/utils"
)

type ListGenerator interface {
	GetNewest(offset int, size int) (Entries, error)
	GetRecent(offset int, size int) (Entries, error)
	GetFrequent(offset int, size int) (Entries, error)
	GetHighest(offset int, size int) (Entries, error)
	GetRandom(offset int, size int) (Entries, error)
	GetByName(offset int, size int) (Entries, error)
	GetByArtist(offset int, size int) (Entries, error)
	GetStarred(offset int, size int) (Entries, error)
	GetAllStarred() (albums Entries, mediaFiles Entries, err error)
	GetNowPlaying() (Entries, error)
	GetRandomSongs(size int) (Entries, error)
}

func NewListGenerator(alr domain.AlbumRepository, mfr domain.MediaFileRepository, npr NowPlayingRepository) ListGenerator {
	return &listGenerator{alr, mfr, npr}
}

type listGenerator struct {
	albumRepo    domain.AlbumRepository
	mfRepository domain.MediaFileRepository
	npRepo       NowPlayingRepository
}

func (g *listGenerator) query(qo domain.QueryOptions, offset int, size int) (Entries, error) {
	qo.Offset = offset
	qo.Size = size
	albums, err := g.albumRepo.GetAll(qo)

	return FromAlbums(albums), err
}

func (g *listGenerator) GetNewest(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "CreatedAt", Desc: true, Alpha: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetRecent(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "PlayDate", Desc: true, Alpha: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetFrequent(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "PlayCount", Desc: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetHighest(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "Rating", Desc: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetByName(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "Name", Alpha: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetByArtist(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{SortBy: "Artist", Alpha: true}
	return g.query(qo, offset, size)
}

func (g *listGenerator) GetRandom(offset int, size int) (Entries, error) {
	ids, err := g.albumRepo.GetAllIds()
	if err != nil {
		return nil, err
	}
	size = utils.MinInt(size, len(ids))
	perm := rand.Perm(size)
	r := make(Entries, size)

	for i := 0; i < size; i++ {
		v := perm[i]
		al, err := g.albumRepo.Get((ids)[v])
		if err != nil {
			return nil, err
		}
		r[i] = FromAlbum(al)
	}
	return r, nil
}

func (g *listGenerator) GetRandomSongs(size int) (Entries, error) {
	ids, err := g.mfRepository.GetAllIds()
	if err != nil {
		return nil, err
	}
	size = utils.MinInt(size, len(ids))
	perm := rand.Perm(size)
	r := make(Entries, size)

	for i := 0; i < size; i++ {
		v := perm[i]
		mf, err := g.mfRepository.Get(ids[v])
		if err != nil {
			return nil, err
		}
		r[i] = FromMediaFile(mf)
	}
	return r, nil
}

func (g *listGenerator) GetStarred(offset int, size int) (Entries, error) {
	qo := domain.QueryOptions{Offset: offset, Size: size, Desc: true}
	albums, err := g.albumRepo.GetStarred(qo)
	if err != nil {
		return nil, err
	}

	return FromAlbums(albums), nil
}

func (g *listGenerator) GetAllStarred() (Entries, Entries, error) {
	albums, err := g.GetStarred(0, -1)
	if err != nil {
		return nil, nil, err
	}

	mediaFiles, err := g.mfRepository.GetStarred(domain.QueryOptions{Desc: true})

	return albums, FromMediaFiles(mediaFiles), err
}

func (g *listGenerator) GetNowPlaying() (Entries, error) {
	npInfo, err := g.npRepo.GetAll()
	if err != nil {
		return nil, err
	}
	entries := make(Entries, len(npInfo))
	for i, np := range npInfo {
		mf, err := g.mfRepository.Get(np.TrackId)
		if err != nil {
			return nil, err
		}
		entries[i] = FromMediaFile(mf)
		entries[i].UserName = np.Username
		entries[i].MinutesAgo = int(time.Now().Sub(np.Start).Minutes())
		entries[i].PlayerId = np.PlayerId
		entries[i].PlayerName = np.PlayerName

	}
	return entries, nil
}
