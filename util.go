// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package amppackager

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

// CertURLPrefix must start without a slash, for PackagerBase's sake.
const CertURLPrefix = "amppkg/cert"

// CertName returns the basename for the given cert, as served by this
// packager's cert cache. Should be stable and unique (e.g.
// content-addressing). Clients should url.PathEscape this, just in case its
// format changes to need escaping in the future.
func CertName(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.URLEncoding.EncodeToString(sum[:])
}
