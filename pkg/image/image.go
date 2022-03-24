package image

import "github.com/containers/storage"

type BuildOptions struct {
}

type ListOptions struct {
}

type ClusterManifest storage.Container

type RegistryService interface {
	Login(domain, username, passwd string) error
	Logout(domain, username string) error
	// platform info
	Pull(image string) error
	PullIfNotExist(image string) error
	Push(image string) error
	// Save a image to local dir witch can mount by a private registry
	Sync(localDir, imageName string) error
}

type ImageService interface {
	//Mount(images ...string) (mountDir string, err error)
	//Unount(dir ...string) error
	Rename(src, dst string) error
	Remove(images ...string) error
	Inspect(image string) (storage.Image, error)
	Build(options BuildOptions, contextDir, imageName string) error
	Prune() error
	ListImages(opt ListOptions) ([]storage.Image, error)
}

type ClusterService interface {
	Create(name, image string) (ClusterManifest, error)
	Delete(name string) error
	Inspect(name string) (ClusterManifest, error)
}
