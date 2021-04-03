// entity
// entitybase, id_wisudawan, ip, time

// Usecase AddViews

// Repository AddViews

// request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record

package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"fmt"
    "net"
    "net/http"
    "strings"
)

type ViewsUsecase struct {
	viewsrepo entity.ViewsRepository
}
func AddViewsUsecase(uc ViewsUsecase, r *http.Request) entity.ViewsUsecase {
	ip, err := getIP(r)
	if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("No valid ip"))
    }
	w.WriteHeader(200)
    w.Write([]byte(ip))
	//bingung kalo gaada serializer
	
}

func getIP(r *http.Request) (string, error) {
    ip := r.Header.Get("X-REAL-IP")
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }

    ips := r.Header.Get("X-FORWARDED-FOR")
    splitIps := strings.Split(ips, ",")
    for _, ip := range splitIps {
        netIP := net.ParseIP(ip)
        if netIP != nil {
            return ip, nil
        }
    }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP = net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }
    return "", fmt.Errorf("No valid ip found")
