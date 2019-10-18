/*
Copyright 2019 cspengl

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Package containerd contains the implementation of an agent connecting
//to the containerd daemon on the machine
package containerd

import(
  "time"

  "github.com/containerd/containerd"
  "github.com/containerd/containerd/containers"
  "github.com/containerd/containerd/diff/apply"
  "github.com/containerd/containerd/mount"
  "github.com/containerd/containerd/rootfs"
  "github.com/containerd/containerd/snapshots"

  "github.com/opencontainers/image-spec/identity"

  "golang.org/x/net/context"
)

const (
	gcRoot           = "containerd.io/gc.root"
	timestampFormat  = "01-02-2006-15:04:05"
	PreviousLabel    = "boss.io/revision.previous"
	ImageLabel       = "boss.io/revision.image"
	ContainerIDLabel = "boss.io/revision.container"
)

type Revision struct {
  Timestamp time.Time
  Key       string
  mounts    []mount.Mount
}

func newRevision(id string) *Revision {
  now := time.Now()
  return &Revision{
    Timestamp:  now,
    Key:        id,
  }
}

func WithCommit(image containerd.Image) containerd.UpdateContainerOpts {
  return func(ctx context.Context, client *containerd.Client, c *containers.Container) error {
    revision, err := save(ctx, client, image, c)

    if err != nil {
      return err
    }

    c.Image = image.Name()
    c.SnapshotKey = revision.Key
    return nil
  }
}

func save(ctx context.Context, client *containerd.Client, updatedImage containerd.Image, c *containers.Container) (*Revision, error) {
  snapshot, err := create(ctx, client, updatedImage, c, c.ID, c.SnapshotKey)
  if err != nil {
    return nil, err
  }

  service := client.SnapshotService(c.Snapshotter)

  diff, err := rootfs.CreateDiff(ctx, c.SnapshotKey, service, client.DiffService())

  if err != nil {
    return nil, err
  }

  applier := apply.NewFileSystemApplier(client.ContentStore())
  if _, err := applier.Apply(ctx, diff, snapshot.mounts); err != nil {
    return nil, err
  }

  return snapshot, nil
}

func create(ctx context.Context, client *containerd.Client, i containerd.Image, c *containers.Container, id, previous string) (*Revision, error) {
  diffIDs, err := i.RootFS(ctx)
  if err != nil {
    return nil, err
  }

  var (
    parent  = identity.ChainID(diffIDs).String()
    r       = newRevision(id)
  )

  labels := map[string]string{
		gcRoot:           r.Timestamp.Format(time.RFC3339),
		ImageLabel:       i.Name(),
		ContainerIDLabel: id,
	}
	if previous != "" {
		labels[PreviousLabel] = previous
	}

  mounts, err := client.SnapshotService(c.Snapshotter).Prepare(ctx, r.Key, parent, snapshots.WithLabels(labels))

  if err != nil {
    panic(err)
    return nil, err
  }

  r.mounts = mounts
  return r, nil
}
