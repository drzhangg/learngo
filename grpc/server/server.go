package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	pb "learngo/grpc/pb"
	"math"
	"sync"
)

type routeGuideServer struct {
	pb.UnimplementedRouteServer
	savedFeatures []*pb.Feature

	mu         sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

// 普通grpc请求
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{
		Name:     "",
		Location: point,
	}, nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.Route_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if inRange {

		}
	}
}

func (s *routeGuideServer) RecordRoute(stream pb.Route_RecordRouteServer) error {

}

func (s *routeGuideServer) RouteChat(stream pb.Route_RouteChatServer) error {

}
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
}

func main() {

}
