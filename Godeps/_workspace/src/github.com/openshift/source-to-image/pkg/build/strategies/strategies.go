package strategies

import (
	dockerclient "github.com/fsouza/go-dockerclient"
	"github.com/golang/glog"
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/build"
	"github.com/openshift/source-to-image/pkg/build/strategies/onbuild"
	"github.com/openshift/source-to-image/pkg/build/strategies/sti"
	"github.com/openshift/source-to-image/pkg/docker"
)

// GetStrategy decides what build strategy will be used for the STI build.
func GetStrategy(config *api.Config) (build.Builder, error) {
	image, err := GetBuilderImage(config)
	if err != nil {
		return nil, err
	}

	if image.OnBuild {
		return onbuild.New(config)
	}

	return sti.New(config)
}

// GetBuilderImage processes the config and performs operations necessary to make
// the Docker image specified as BuilderImage available locally.
// It returns information about the base image, containing metadata necessary
// for choosing the right STI build strategy.
func GetBuilderImage(config *api.Config) (*docker.PullResult, error) {
	d, err := docker.New(config.DockerConfig, config.PullAuthentication)
	result := docker.PullResult{}
	if err != nil {
		return nil, err
	}

	var image *dockerclient.Image
	if config.ForcePull {
		image, err = d.PullImage(config.BuilderImage)
		if err != nil {
			glog.Warningf("An error occurred when pulling %s: %v. Attempting to use local image.", config.BuilderImage, err)
			image, err = d.CheckImage(config.BuilderImage)
		}
	} else {
		image, err = d.CheckAndPullImage(config.BuilderImage)
	}

	if err != nil {
		return nil, err
	}
	result.Image = image
	result.OnBuild = d.IsImageOnBuild(config.BuilderImage)
	return &result, nil
}
