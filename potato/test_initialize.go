package potato

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func deleteCollection(ctx context.Context, client *firestore.Client,
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

func ReGenerateDatabase(firestore *firestore.Client) error {
	// Drop database
	drops := []string{"backgrounds", "effects", "engines", "levels", "particles", "skins", "users"}
	for _, d := range drops {
		ref := firestore.Collection(d)
		ctx := context.Background()
		err := deleteCollection(ctx, firestore, ref, 100)
		if err != nil {
			return err
		}
	}
	if err := initBackgroundsDatabase(firestore); err != nil {
		return err
	}
	if err := initEffectsDatabase(firestore); err != nil {
		return err
	}
	if err := initEnginesDatabase(firestore); err != nil {
		return err
	}
	if err := initLevelsDatabase(firestore); err != nil {
		return err
	}
	if err := initParticlesDatabase(firestore); err != nil {
		return err
	}
	if err := initSkinsDatabase(firestore); err != nil {
		return err
	}
	if err := initUsersDatabase(firestore); err != nil {
		return err
	}
	return nil
}
