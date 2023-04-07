# Publish CRL distibution point for an Intermediate or Leaf certificate

## API

Please check [transactions.md](TBD LINK).

## Model

Model described in [propose_publish_crl](TBD LINK) and [approve_publish_crl](TBD LINK) files.

### Propose publish CRL message structure

described in [transactions.md](TBD LINK).

### Approve publish CRL message structure

described in [transactions.md](TBD LINK).

## State

described in [transactions.md](TBD LINK).

## Auth rules

described in [transactions.md](TBD LINK).

## Questions

- Is `subject` field needed?
  `subject` field is needed because we use it for querying certificates
- Do we need to validate `vid` somehow?
  We need to validate if `Vendor` with provided `vid` exists and it owns the certificate
- Are existing models need to be adjusted?
  New functionality does not requre some changes in existing models
- Are migration of data is needed?
  No migration of data needed since existing models were not adjusted
- How this new functionality should affect existing certificate "revocation" logic?
  If certificate removed it should be removed from CRL
