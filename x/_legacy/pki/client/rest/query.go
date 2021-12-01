// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

func getAllX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		performPkiQuery(restCtx,
			fmt.Sprintf("custom/%s/all_x509_root_certs", storeName), "", "")
	}
}

func getAllProposedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		performPkiQuery(restCtx,
			fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", storeName), "", "")
	}
}

func getAllX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		performPkiQuery(restCtx, fmt.Sprintf("custom/%s/all_x509_certs", storeName), rootSubject, rootSubjectKeyID)
	}
}

func getAllSubjectX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		vars := restCtx.Variables()
		subject := vars[subject]
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		performPkiQuery(restCtx, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s",
			storeName, subject), rootSubject, rootSubjectKeyID)
	}
}

func getProposedX509RootCertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetProposedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID).Error())

			return
		}

		var proposedCertificate types.ProposedCertificate

		restCtx.Codec().MustUnmarshalBinaryBare(res, &proposedCertificate)

		restCtx.EncodeAndRespondWithHeight(proposedCertificate, height)
	}
}

func getX509CertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrCertificateDoesNotExist(subject, subjectKeyID).Error())

			return
		}

		var certificates types.Certificates

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certificates)

		restCtx.EncodeAndRespondWithHeight(certificates, height)
	}
}

func getX509CertChainHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		chain := types.NewCertificates([]types.Certificate{})

		height, err := chainCertificates(restCtx, storeName, subject, subjectKeyID, &chain)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrCertificateDoesNotExist(subject, subjectKeyID).Error())

			return
		}

		restCtx.EncodeAndRespondWithHeight(chain, height)
	}
}

func getAllProposedX509RootCertsToRevokeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		performPkiQuery(restCtx,
			fmt.Sprintf("custom/%s/all_proposed_x509_root_cert_revocations", storeName), "", "")
	}
}

func getProposedX509RootCertToRevokeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetProposedCertificateRevocationKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrProposedCertificateRevocationDoesNotExist(subject, subjectKeyID).Error())

			return
		}

		var revocation types.ProposedCertificateRevocation

		restCtx.Codec().MustUnmarshalBinaryBare(res, &revocation)

		restCtx.EncodeAndRespondWithHeight(revocation, height)
	}
}

func getAllRevokedX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		performPkiQuery(restCtx, fmt.Sprintf("custom/%s/all_revoked_x509_certs", storeName),
			rootSubject, rootSubjectKeyID)
	}
}

func getAllRevokedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		performPkiQuery(restCtx, fmt.Sprintf("custom/%s/all_revoked_x509_root_certs", storeName), "", "")
	}
}

func getRevokedX509CertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetRevokedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrRevokedCertificateDoesNotExist(subject, subjectKeyID).Error())

			return
		}

		var certificates types.Certificates

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certificates)

		restCtx.EncodeAndRespondWithHeight(certificates, height)
	}
}

func chainCertificates(restCtx rest.RestContext, storeName string,
	subject string, subjectKeyID string, chain *types.Certificates) (int64, sdk.Error) {
	res, height, err := restCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), storeName)
	if err != nil || res == nil {
		return height, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	var certificates types.Certificates

	restCtx.Codec().MustUnmarshalBinaryBare(res, &certificates)

	certificate := certificates.Items[len(certificates.Items)-1]
	chain.Items = append(chain.Items, certificate)

	if !certificate.IsRoot {
		return chainCertificates(restCtx, storeName, certificate.Issuer, certificate.AuthorityKeyID, chain)
	}

	return height, nil
}

func performPkiQuery(restCtx rest.RestContext, path string, rootSubject string, rootSubjectKeyID string) {
	paginationParams, err := restCtx.ParsePaginationParams()
	if err != nil {
		return
	}

	params := types.NewPkiQueryParams(paginationParams, rootSubject, rootSubjectKeyID)
	restCtx.QueryList(path, params)
}
