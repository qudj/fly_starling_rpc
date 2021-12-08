package middlewares

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
	"time"
)

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()
	pos := strings.LastIndex(info.FullMethod, "/") + 1
	method := info.FullMethod[pos:]

	in, _ := json.Marshal(req)
	inStr := string(in)

	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		log.Println(" ip:", remoteAddr, " method:", method, " in:", inStr, " out:", outStr, " err:", err, " duration/ms:", duration)
		if err != nil {
			log.Fatal()
		}
	}()
	resp, err = handler(ctx, req)
	return
}
