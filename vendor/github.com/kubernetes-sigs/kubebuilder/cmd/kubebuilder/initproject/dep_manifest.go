/*
Copyright 2017 The Kubernetes Authors.

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

package initproject

const depManifestHeader = `
# Users add deps lines here

[prune]
  go-tests = true
  #unused-packages = true

# Note: Stanzas below are generated by Kubebuilder and may be rewritten when
# upgrading kubebuilder versions.
`

// depManifestKBMarker acts as a separater between the user managed dependencies
// and KB generated dependencies. Content above this marker is user managed and
// needs to be preserved across 'vendor update' operations. Content below this
// marker is generated by KB and will be updated by KB during 'vendor update'.
const depManifestKBMarker = `# DO NOT MODIFY BELOW THIS LINE.`

// template for dep's manifest file (Gopkg.toml). This is generated using
// scripts/generate_dep_manifest.sh scripts.
const depManifestOverride = `
[[override]]
name="cloud.google.com/go"
version="v0.21.0"

[[override]]
name="github.com/PuerkitoBio/purell"
version="v1.1.0"

[[override]]
name="github.com/PuerkitoBio/urlesc"
revision="de5bf2ad457846296e2031421a34e2568e304e35"

[[override]]
name="github.com/davecgh/go-spew"
version="v1.1.0"

[[override]]
name="github.com/emicklei/go-restful"
version="v2.7.0"

[[override]]
name="github.com/ghodss/yaml"
version="v1.0.0"

[[override]]
name="github.com/go-openapi/jsonpointer"
revision="3a0015ad55fa9873f41605d3e8f28cd279c32ab2"

[[override]]
name="github.com/go-openapi/jsonreference"
revision="3fb327e6747da3043567ee86abd02bb6376b6be2"

[[override]]
name="github.com/go-openapi/spec"
revision="bcff419492eeeb01f76e77d2ebc714dc97b607f5"

[[override]]
name="github.com/go-openapi/swag"
revision="811b1089cde9dad18d4d0c2d09fbdbf28dbd27a5"

[[override]]
name="github.com/gogo/protobuf"
version="v1.0.0"

[[override]]
name="github.com/golang/glog"
revision="23def4e6c14b4da8ac2ed8007337bc5eb5007998"

[[override]]
name="github.com/golang/groupcache"
revision="66deaeb636dff1ac7d938ce666d090556056a4b0"

[[override]]
name="github.com/golang/protobuf"
version="v1.1.0"

[[override]]
name="github.com/google/gofuzz"
revision="24818f796faf91cd76ec7bddd72458fbced7a6c1"

[[override]]
name="github.com/googleapis/gnostic"
version="v0.1.0"

[[override]]
name="github.com/hashicorp/golang-lru"
revision="0fb14efe8c47ae851c0034ed7a448854d3d34cf3"

[[override]]
name="github.com/howeyc/gopass"
revision="bf9dde6d0d2c004a008c27aaee91170c786f6db8"

[[override]]
name="github.com/imdario/mergo"
version="v0.3.4"

[[override]]
name="github.com/json-iterator/go"
version="1.1.3"

[[override]]
name="github.com/mailru/easyjson"
revision="8b799c424f57fa123fc63a99d6383bc6e4c02578"

[[override]]
name="github.com/modern-go/concurrent"
version="1.0.3"

[[override]]
name="github.com/modern-go/reflect2"
version="1.0.0"

[[override]]
name="github.com/onsi/ginkgo"
version="v1.4.0"

[[override]]
name="github.com/onsi/gomega"
version="v1.3.0"

[[override]]
name="github.com/spf13/pflag"
version="v1.0.1"

[[override]]
name="golang.org/x/crypto"
revision="4ec37c66abab2c7e02ae775328b2ff001c3f025a"

[[override]]
name="golang.org/x/net"
revision="640f4622ab692b87c2f3a94265e6f579fe38263d"

[[override]]
name="golang.org/x/oauth2"
revision="cdc340f7c179dbbfa4afd43b7614e8fcadde4269"

[[override]]
name="golang.org/x/sys"
revision="7db1c3b1a98089d0071c84f646ff5c96aad43682"

[[override]]
name="golang.org/x/text"
version="v0.3.0"

[[override]]
name="golang.org/x/time"
revision="fbb02b2291d28baffd63558aa44b4b56f178d650"

[[override]]
name="google.golang.org/appengine"
version="v1.0.0"

[[override]]
name="gopkg.in/inf.v0"
version="v0.9.1"

[[override]]
name="gopkg.in/yaml.v2"
version="v2.2.1"

[[override]]
name="k8s.io/api"
version="kubernetes-1.10.0"

[[override]]
name="k8s.io/apiextensions-apiserver"
version="kubernetes-1.10.1"

[[override]]
name="k8s.io/apimachinery"
version="kubernetes-1.10.0"

[[override]]
name="k8s.io/client-go"
version="kubernetes-1.10.1"

[[override]]
name="k8s.io/kube-aggregator"
version="kubernetes-1.10.1"

[[override]]
name="k8s.io/kube-openapi"
revision="f08db293d3ef80052d6513ece19792642a289fea"

[[override]]
name="sigs.k8s.io/testing_frameworks"
revision="f53464b8b84b4507805a0b033a8377b225163fea"

[[override]]
name = "github.com/kubernetes-sigs/kubebuilder"
{{ if eq .Version "unknown" -}}
branch="master"
{{ else if eq .Version "master" -}}
branch="master"
{{ else -}}
version="{{.Version}}"
{{ end }}

`
