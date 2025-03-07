/*
Copyright 2019 The Jetstack cert-manager contributors.

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

package issuers

import (
	"context"
	"fmt"
	"reflect"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/errors"

	apiutil "github.com/jetstack/cert-manager/pkg/api/util"
	"github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"github.com/jetstack/cert-manager/pkg/internal/apis/certmanager/validation"
	logf "github.com/jetstack/cert-manager/pkg/logs"
	"github.com/jetstack/cert-manager/pkg/metrics"
)

const (
	errorInitIssuer = "ErrInitIssuer"
	errorConfig     = "ConfigError"

	messageErrorInitIssuer = "Error initializing issuer: "
)

func (c *controller) Sync(ctx context.Context, iss *v1alpha2.Issuer) (err error) {
	metrics.Default.IncrementSyncCallCount(ControllerName)

	log := logf.FromContext(ctx)

	issuerCopy := iss.DeepCopy()
	defer func() {
		if _, saveErr := c.updateIssuerStatus(iss, issuerCopy); saveErr != nil {
			err = errors.NewAggregate([]error{saveErr, err})
		}
	}()

	el := validation.ValidateIssuer(issuerCopy)
	if len(el) > 0 {
		msg := fmt.Sprintf("Resource validation failed: %v", el.ToAggregate())
		apiutil.SetIssuerCondition(issuerCopy, v1alpha2.IssuerConditionReady, cmmeta.ConditionFalse, errorConfig, msg)
		return
	}

	// Remove existing ErrorConfig condition if it exists
	for i, c := range issuerCopy.Status.Conditions {
		if c.Type == v1alpha2.IssuerConditionReady {
			if c.Reason == errorConfig && c.Status == cmmeta.ConditionFalse {
				issuerCopy.Status.Conditions = append(issuerCopy.Status.Conditions[:i], issuerCopy.Status.Conditions[i+1:]...)
				break
			}
		}
	}

	i, err := c.issuerFactory.IssuerFor(issuerCopy)

	if err != nil {
		return err
	}

	err = i.Setup(ctx)
	if err != nil {
		s := messageErrorInitIssuer + err.Error()
		log.Info(s)
		c.recorder.Event(issuerCopy, v1.EventTypeWarning, errorInitIssuer, s)
		return err
	}

	return nil
}

func (c *controller) updateIssuerStatus(old, new *v1alpha2.Issuer) (*v1alpha2.Issuer, error) {
	if reflect.DeepEqual(old.Status, new.Status) {
		return nil, nil
	}
	// TODO: replace Update call with UpdateStatus. This requires a custom API
	// server with the /status subresource enabled and/or subresource support
	// for CRDs (https://github.com/kubernetes/kubernetes/issues/38113)
	return c.cmClient.CertmanagerV1alpha2().Issuers(new.Namespace).Update(new)
}
