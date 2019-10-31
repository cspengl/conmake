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
    "bytes"
  	"encoding/json"
    "context"

  	"github.com/containerd/containerd/content"
  	"github.com/containerd/containerd/images"
  	digest "github.com/opencontainers/go-digest"
  	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

)

func (a *ContainerdAgent) loadManifest(imageRef string) (map[string]json.RawMessage, error) {

  //getting image
  image, err := a.client.GetImage(a.ctx, imageRef)

  if err != nil {
    return nil, err
  }

  //loading manifest from image
  rawManifest, err := content.ReadBlob(a.ctx, a.client.ContentStore(), image.Target())

  if err != nil {
    return nil, err
  }

  //unmarshaling manifest
  manifest := map[string]json.RawMessage{}
  if err := json.Unmarshal(rawManifest, &manifest); err != nil {
    return nil, err
  }

  return manifest, nil
}

func (a *ContainerdAgent) addLayerToManifest(manifestDesc, layer, diff ocispec.Descriptor) (ocispec.Descriptor, error) {

  //load json from blob store
  manifest, err := loadJSONFromDescriptor(a.ctx, a.client.ContentStore(), manifestDesc)

  if err != nil {
    return ocispec.Descriptor{}, err
  }

  // -------------------------------------------------

  //patching the image config

  //load jsonobject of the descriptor from json
  configJSON, err := manifest["config"].MarshalJSON()

  if err != nil {
    return ocispec.Descriptor{}, err
  }

  //unmarshal descriptor from json object
  var config = ocispec.Descriptor{}
  err = json.Unmarshal(configJSON, &config)

  if err != nil {
    return ocispec.Descriptor{}, err
  }

  //patching
  config, err = a.patchImageConfig(config, diff.Digest)

  if err != nil {
    return ocispec.Descriptor{}, err
  }

  //---------------------------------------------------

  //updating layers

  //assuming working with docker image
  layer.MediaType = images.MediaTypeDockerSchema2LayerGzip

  //getting layers from manifest
  layers := []ocispec.Descriptor{}
  layersJSON, err := manifest["layers"].MarshalJSON()

  if err != nil {
    return ocispec.Descriptor{}, err
  }

  if err = json.Unmarshal(layersJSON, &layers); err != nil {
    return ocispec.Descriptor{}, err
  }

  //append new layer
  layers = append(layers, layer)


  //writing layers back to json object
  layersJSON, err = json.Marshal(layers)

  if err != nil {
    return ocispec.Descriptor{}, err
  }


  //save new layers in manifest json object
  manifest["layers"] = layersJSON

  // -------------------------------------------------------

  // labels := map[string]string{
  //   "containerd.io/gc.ref.content.0": config.Digest.String(),
  // }
  //
  // for i, layer := range layers {
  //   labels[fmt.Sprintf("containerd.io/gc.reg.content.%d", i+1)] = layer.Digest.String()
  // }


  //constructing new descriptor
  manifestBytes, err := json.Marshal(manifest)

  newDesc := manifestDesc

  newDesc.Digest = digest.FromBytes(manifestBytes)
  newDesc.Size = int64(len(manifestBytes))

  //writing new descriptor to blob store
  if err := content.WriteBlob(
    a.ctx,
    a.client.ContentStore(),
    "custom-ref",
    bytes.NewReader(manifestBytes),
    newDesc); err != nil {
      return ocispec.Descriptor{}, err
    }

  return newDesc, nil

}

func (a *ContainerdAgent) patchImageConfig(imageConfig ocispec.Descriptor, newLayerDiff digest.Digest) (ocispec.Descriptor, error) {

  result := imageConfig

  //Loading JSON from config descriptor
  config, err := loadJSONFromDescriptor(a.ctx, a.client.ContentStore(), imageConfig)

  //getting rootfs from config
  var rootFS ocispec.RootFS
  rootFSJSON, err := config["rootfs"].MarshalJSON()
  if err != nil {
    return result, err
  }

  if err = json.Unmarshal(rootFSJSON, &rootFS); err != nil {
    return result, err
  }

  //appending diff to diffids
  rootFS.DiffIDs = append(rootFS.DiffIDs, newLayerDiff)

  //save RootFS to back to json object of config
  rootFSJSON, err = json.Marshal(rootFS)
  if err != nil {
    return ocispec.Descriptor{}, err
  }

  config["rootfs"] = rootFSJSON

  //convert json object to bytes
  rawConfig, err := json.Marshal(config)
  if err != nil {
    return ocispec.Descriptor{}, err
  }

  //saving new bytes to blob store
  result.Digest = digest.FromBytes(rawConfig)
  result.Size = int64(len(rawConfig))
  err = content.WriteBlob(
    a.ctx,
    a.client.ContentStore(),
    "custom-ref",
    bytes.NewReader(rawConfig),
    result,
  )

  //returning patched config as descriptor
  return result, err

}

func loadJSONFromDescriptor(context context.Context, store content.Store, descriptor ocispec.Descriptor) (map[string]json.RawMessage, error) {

  //rawjson
  rawJSON, err := content.ReadBlob(context, store, descriptor)

  if err != nil {
    return nil, err
  }

  //marshalling
  jsonObj := map[string]json.RawMessage{}
  err = json.Unmarshal(rawJSON, &jsonObj)

  return jsonObj, err
}
