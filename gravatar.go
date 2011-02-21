package gravatar

import (
	"strings"
	"crypto/md5"
	"os"
)

type G struct {
	size       int
	defaultImg string
	rating     string
}

var (
	errBadParm       = os.NewError("bad parameter")
	errInvalidRating = os.NewError("invalid rating")
)

func md5sum(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	sum := h.Sum()
	ret := make([]byte, len(sum)*2)
	for i := 0; i < len(sum); i++ {
		a := sum[i] / 0x10
		b := sum[i] % 0x10
		if a <= 9 {
			a += '0'
		} else {
			a += 'a' - 10
		}
		if b <= 9 {
			b += '0'
		} else {
			b += 'a' - 10
		}
		ret[i*2] = a
		ret[i*2+1] = b
	}
	return string(ret)
}

func New() *G {
	return &G{}
}

// set the rating
func (g *G) Rating(rating string) (err os.Error) {
	switch rating {
	default:
		return errInvalidRating
	case "g", "pg", "r", "x":
		g.rating = rating
	}
	return nil
}

// configure the Default Image,
// valid parameters are: "404", "mm", "identicon", "monsterid", "wavatar", "retro"
// alternatively a URL pointing to the image
func (g *G) Default(url string) (err os.Error) {
	lowercase := strings.ToLower(url)
	switch lowercase {
	default:
		if (len(url) > 4) && ("http" == lowercase[0:4]) {
			g.defaultImg = url
		} else {
			return errBadParm
		}
	case "", "404", "mm", "identicon", "monsterid", "wavatar", "retro":
		g.defaultImg = lowercase
	}
	return
}

// return the URL to the non-encrypted Gravatar
func (g *G) URL(email string) string {
	url := "http://www.gravatar.com/avatar/" + Hash(email)
	return url
}

// return the SSL URL for Gravatar
func (g *G) URLs(email string) string {
	return "https://secure.gravatar.com/avatar/" + Hash(email)
}

func URL(email string) string {
	return "http://www.gravatar.com/avatar/" + Hash(email)
}

func Hash(email string) string {
	return md5sum(strings.ToLower(strings.TrimSpace(email)))
}
