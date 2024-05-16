package statesyncstore

import (
	"context"
	"time"

	serrors "project-v/internal/errors"

	"github.com/go-redis/redis/v8"
)

// stateSyncStoreImpl ...
type stateSyncStoreImpl struct {
	client     *redis.Client
	defaultTTL int
	namespace  string
}

// New ...
func New(redisClient *redis.Client, namespace string) *stateSyncStoreImpl {
	return &stateSyncStoreImpl{
		client:     redisClient,
		namespace:  namespace,
		defaultTTL: 300, // 5 minutes
	}
}

const prefix = "statesync:"

func (o *stateSyncStoreImpl) resolveKey(key string) string {
	return prefix + o.namespace + ":" + key
}

/*
Lock try to lock the workspace with a default TTL of 5 minutes.

Try 5 times to lock the workspace with a 100ms delay between each attempt.
If can lock the workspace, return nil.
If can't lock the workspace, return error "workspace is syncing".
*/
func (o *stateSyncStoreImpl) Lock(ctx context.Context, workspaceID string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	expiration := time.Duration(o.defaultTTL) * time.Second

	// Try to set the key 3 times with a 100ms delay between each attempt.
	for i := 0; i < 5; i++ {
		// SetNX: SET if Not eXists
		// SetNX returns a boolean indicating if the key was set.
		ok, err := o.client.SetNX(ctx, o.resolveKey(workspaceID), true, expiration).Result()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return serrors.ErrWorkspaceIsSyncing
}

// UnLock unlock the workspace.
func (o *stateSyncStoreImpl) UnLock(ctx context.Context, workspaceID string) error {
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	return o.client.Del(ctx, o.resolveKey(workspaceID)).Err()
}
