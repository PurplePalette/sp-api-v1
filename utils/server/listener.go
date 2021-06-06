package server

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Listener struct {
	client *firestore.Client
	cache  *potato.CacheService
}

func NewListener(client *firestore.Client, cache *potato.CacheService) *Listener {
	return &Listener{client: client, cache: cache}
}

func (l *Listener) ListenFirestoreUpdate(collectionName string) {
	ctx := context.Background()
	snapIter := l.client.Collection(collectionName).Snapshots(ctx)
	defer snapIter.Stop()

	for {
		snap, err := snapIter.Next()
		if err != nil {
			log.Fatalln(err)
		}
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			log.Fatalln(err)
		}
		for _, change := range snap.Changes {
			switch change.Kind {
			case firestore.DocumentAdded:
				switch collectionName {
				case "backgrounds":
					var data potato.Background
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new background: ", change.Doc.Ref.ID)
					}
				case "effects":
					var data potato.Effect
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new effect: ", change.Doc.Ref.ID)
					}
				case "engines":
					var data potato.Engine
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new engine: ", change.Doc.Ref.ID)
					}
				case "levels":
					var data potato.Level
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new level: ", change.Doc.Ref.ID)
					}
				case "particles":
					var data potato.Particle
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new particle: ", change.Doc.Ref.ID)
					}
				case "skins":
					var data potato.Skin
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new skin: ", change.Doc.Ref.ID)
					}
				case "users":
					var data potato.User
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					if err := l.cache.Add(change.Doc.Ref.ID, data); err != nil {
						log.Print("Added new user: ", change.Doc.Ref.ID)
					}
				}
			case firestore.DocumentModified:
				switch collectionName {
				case "backgrounds":
					var data potato.Background
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified background: ", change.Doc.Ref.ID)
				case "effects":
					var data potato.Effect
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified effect: ", change.Doc.Ref.ID)
				case "engines":
					var data potato.Engine
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified engine: ", change.Doc.Ref.ID)
				case "levels":
					var data potato.Level
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified level: ", change.Doc.Ref.ID)
				case "particles":
					var data potato.Particle
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified particle: ", change.Doc.Ref.ID)
				case "skins":
					var data potato.Skin
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified skin: ", change.Doc.Ref.ID)
				case "users":
					var data potato.User
					if err := change.Doc.DataTo(&data); err != nil {
						log.Fatal(err)
					}
					l.cache.Set(change.Doc.Ref.ID, data)
					log.Print("Modified user: ", change.Doc.Ref.ID)
				}
			}
		}
	}
}
