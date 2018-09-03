package oci_registry_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	registry "code.cloudfoundry.org/eirini/oci_registry"
	"code.cloudfoundry.org/eirini/oci_registry/oci_registryfakes"
)

var _ = Describe("Registry", func() {
	var (
		fakeServer   *httptest.Server
		handler      http.Handler
		url          string
		imageManager *oci_registryfakes.FakeImageManager
		res          *http.Response
		err          error
	)

	BeforeEach(func() {
		imageManager = new(oci_registryfakes.FakeImageManager)
		handler = registry.NewHandler(imageManager)
		fakeServer = httptest.NewServer(handler)
	})

	Context("when requesting a manifest", func() {
		Context("for an image name and tag", func() {

			BeforeEach(func() {
				url = "/v2/image-name/manifest/image-tag"
			})

			JustBeforeEach(func() {
				url = fmt.Sprintf("%s%s", fakeServer.URL, url)
				res, err = http.Get(url)
			})

			It("should not fail to request the endpoint", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should serve the GET image manifest endpoint ", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should ask for the manifest of desired image and tag", func() {
				name, tag := imageManager.GetManifestArgsForCall(0)
				Expect(name).To(Equal("image-name"))
				Expect(tag).To(Equal("image-tag"))
			})

			Context("When requesting an image manifest fails", func() {
				BeforeEach(func() {
					imageManager.GetManifestReturns(nil, errors.New("retrieving manifest failed"))
				})

				It("should return an error response", func() {
					response, _ := ioutil.ReadAll(res.Body)
					Expect(string(response)).To(ContainSubstring("could not"))
					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				})
			})
		})

		Context("for image names have multiple paths or special chars", func() {
			It("it should support / in the name path parameter", func() {
				url := fmt.Sprintf("%s%s", fakeServer.URL, "/v2/image/name/manifest/image-tag")
				res, err := http.Get(url)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("it should support mulitple / in the name path parameter", func() {
				url := fmt.Sprintf("%s%s", fakeServer.URL, "/v2/image/tag/v/22/name/manifest/image-tag")
				res, err := http.Get(url)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("it NOT should support special characters in the name path parameter", func() {
				url := fmt.Sprintf("%s%s", fakeServer.URL, "/v2/image/tag@/v/!22/name/manifest/image-tag")
				res, err := http.Get(url)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("When requesting a layer", func() {
		BeforeEach(func() {
			url = "/v2/image-name/blobs/my-droplet-digest"
			buf := bytes.NewBuffer([]byte("a-tar-file"))
			imageManager.GetLayerReturns(buf, nil)
			imageManager.HasReturns(true)
		})

		JustBeforeEach(func() {
			url = fmt.Sprintf("%s%s", fakeServer.URL, url)
			res, err = http.Get(url)
		})

		It("should not fail to request the endpoint", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should serve the image layer GET endpoint", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should call the ImageHandler for the given image name and digest", func() {
			name, digest := imageManager.GetLayerArgsForCall(0)
			Expect(name).To(Equal("image-name"))
			Expect(digest).To(Equal("my-droplet-digest"))
		})

		Context("and the request fails", func() {
			BeforeEach(func() {
				imageManager.GetLayerReturns(nil, errors.New("something bad happend"))
			})

			It("should fail with InternalServerError", func() {
				response, readErr := ioutil.ReadAll(res.Body)
				Expect(readErr).ToNot(HaveOccurred())
				Expect(string(response)).To(ContainSubstring("could not receive layer"))
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("and the layer does not exist", func() {
			BeforeEach(func() {
				imageManager.HasReturns(false)
			})

			It("should fail with a NotFound response", func() {
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
