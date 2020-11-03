# Governance module

This module implements voting procedure for decisions that should be approved by a majority of partners.

## Usage

This module provides only the abstract type of proposal.
Other modules should implement concrete proposal types and register them in the `gov` module.

General usage algorithm:

- Implement concrete type of the proposal (type must implement `Coontent` interface).
- Register this type in Amino codec of the `gov` module (?).
- Register handler for the proposal. This handler will be triggered when voting procedure succeed.
- Implement transaction that will add proposal to the `gov` handler.

## Transactions

### Vote

__Description:__

Adds a positive or negative vote to the proposal.

__Parameters:__

- Proposal ID
- Author ID
- Vote type [approve | decline]

__Signature:__

Should be signed by a partner.

TODO: Allow other roles to send proposals?

## Queries

### Proposal

__Description:__

Returns one proposal by ID.

__Parameters:__

- Proposal ID

__Response:__

Single proposal model.

### Active proposals

__Description:__

Returns active proposals.

__Parameters:__

- Pagination parameters (?)

__Response:__

List of proposal models.

### All proposals

__Description:__

Returns all proposals.

__Parameters:__

- Pagination parameters (?)

__Response:__

List of proposal models.

## Models

### Proposal

Fields:

- Content - Proposal content interface (implementation of proposal)
- ProposalID - ID of the proposal
- Status - Status of the Proposal {Active, Passed, Rejected}
- Votes - List of `Vote`'s

### Vote

Fields:

- Author - User id
- Action - (Approve | Reject)

## State schema

Items are stored with the following keys and values:

- List of all proposals:
    - 0x00<proposalID_Bytes>: Proposal
- Index for active proposals;
    - 0x01<proposalID_Bytes>: activeProposalID
- Key for the next id value:
    - 0x02: nextProposalID

## Content interface

Content interface is the interface that concrete proposals must implement.

See `types/content.go` for more information. 

## Triggers

Update of the state of a proposal should happen only on one event:

- Partner adds new voice

TODO: Possible improvement here is to check for status update in `end blocker` handler.
It will solve issues with partner revocations.

If proposal status changes to `passes`, the handler for particular proposal type should be executed. See `Routing`.

## Routing

A handler for each type of proposal should be registered during application initialization (in `app.go`).

## Implementation notes

Firstly I tried to adapt `gov` module from the Cosmos SDK. Now it's in `gov_old` folder and should be deleted.
It had a lot of unnecessary functionality, so it was decided to write a new module.

Key features:

- No expiration. A proposal is available until it's rejected or approved.
- No queues. Partners can vote for proposals in any order.
- No staking.
