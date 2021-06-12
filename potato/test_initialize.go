package potato

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// DeleteCollection deletes entire data stored in specified ref
func DeleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

// InsertCollection inserts entire data of json migration file to specified ref
func InsertCollection(ctx context.Context, firestore *firestore.Client, collectionName string) error {
	col := firestore.Collection(collectionName)
	bytes, err := ioutil.ReadFile("./migration/" + collectionName + ".json")
	if err != nil {
		panic(err)
	}
	switch collectionName {
	case "backgrounds":
		var bgs []Background
		if err = json.Unmarshal(bytes, &bgs); err != nil {
			panic(err)
		}
		for _, bg := range bgs {
			if _, err := col.Doc(bg.Name).Set(ctx, bg); err != nil {
				return err
			}
		}
	case "effects":
		var efs []Effect
		if err = json.Unmarshal(bytes, &efs); err != nil {
			panic(err)
		}
		for _, ef := range efs {
			if _, err := col.Doc(ef.Name).Set(ctx, ef); err != nil {
				return err
			}
		}
	case "engines":
		var efs []Engine
		if err = json.Unmarshal(bytes, &efs); err != nil {
			panic(err)
		}
		for _, ef := range efs {
			if _, err := col.Doc(ef.Name).Set(ctx, ef); err != nil {
				return err
			}
		}
	case "levels":
		var lvs []Level
		if err = json.Unmarshal(bytes, &lvs); err != nil {
			panic(err)
		}
		for _, lv := range lvs {
			if _, err := col.Doc(lv.Name).Set(ctx, lv); err != nil {
				return err
			}
		}
	case "particles":
		var pts []Particle
		if err = json.Unmarshal(bytes, &pts); err != nil {
			panic(err)
		}
		for _, pt := range pts {
			if _, err := col.Doc(pt.Name).Set(ctx, pt); err != nil {
				return err
			}
		}
	case "skins":
		var sks []Skin
		if err = json.Unmarshal(bytes, &sks); err != nil {
			panic(err)
		}
		for _, sk := range sks {
			if _, err := col.Doc(sk.Name).Set(ctx, sk); err != nil {
				return err
			}
		}
	case "users":
		var uss []User
		if err = json.Unmarshal(bytes, &uss); err != nil {
			panic(err)
		}
		for _, us := range uss {
			if _, err := col.Doc(us.UserID).Set(ctx, us); err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported collection type")
	}
	return nil
}

// ReGenerateDatabase drop and insert data to firestore
func ReGenerateDatabase(firestore *firestore.Client) error {
	// Drop and insert database
	cols := []string{"backgrounds", "effects", "engines", "levels", "particles", "skins", "users"}
	for _, col := range cols {
		ref := firestore.Collection(col)
		ctx := context.Background()
		if err := DeleteCollection(ctx, firestore, ref, 100); err != nil {
			return err
		}
		if err := InsertCollection(ctx, firestore, col); err != nil {
			return err
		}
	}
	return nil
}
