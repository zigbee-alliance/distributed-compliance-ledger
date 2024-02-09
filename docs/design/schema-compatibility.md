# Support for forward and backward compatibility in DCL schemes

Schema changes can cover a wide range of modifications with varying impacts on application compatibility and data integrity. Below are use cases with strategies to manage schema changes and ensure compatibility.

## Non Breaking Changes

This section provides potential strategies for maintaining backward compatibility for changes that either do not break compatibility or can be converted without loss of information.

### Strategy for Compatible Changes

For changes that are backward-compatible, such as adding optional fields or extending enumerations:

**Strategy steps:**

- One time actions:
  - Add an optional version field to all DCL schema to track the schema version.
- For each update:
  - Update the schema by introducing compatible changes (such as adding a new optional field).
  - Probably need to update transactions and queries.
  - DCL doesn't fulfill the Schema version automatically
  - It will be up to the transaction submitter (Vendor) to specify a correct Schema version
  - If Schema Version is not set - then the initial version (version 0 or 1) is assumed
  - It will be up to the client application to process the Schema version

This strategy is straightforward and quick to implement, but only suitable for compatible changes.

### Strategy for Convertible Changes

For changes that affect compatibility but can be converted, like renaming fields, and changing enumerations:

**Strategy steps:**

- For each update:
  - Create a new version of a Schema and state (a new .proto file)
  - Migrate older states to newer schema version.
  - Remove the states associated with the older schema versions.
  - Implement transactions and queries for the new schema version.
  - Update older transactions and queries to converting data between the latest and older schema version,  ensuring backward compatibility.
  - There will be separated API for each version of the schema, for example::
    - models/vid/pid
    - modelsV2/vid/pid
    - modelsV3/vid/pid

The specific implementation details will vary with each change, each time requires an individual approach.

Support for the Light Client feature will not be extended to legacy APIs due to on-the-fly data migration (as the data is not in the State and proofs can be generated).

## Breaking changes

This section provides potential strategies for maintaining backward compatibility for breaking changes. However, it is unlikely that these options will be implemented in production, as we do not plan to support backward compatibility for breaking changes.

### Strategy for Handling Non-Convertible Changes

For significant changes that directly impact compatibility, such as adding mandatory fields or removing fields:

**Strategy steps:**

- One time actions:
  - Create a more flexible, generic schema structure to hold a wide range of data formats (Can be used [Any](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/any.proto) as described in [ADR-19](https://docs.cosmos.network/v0.47/build/architecture/adr-019-protobuf-state-encoding#usage-of-any-to-encode-interfaces))
  - Migrate old states to the newer, generic schema.
  - Remove the states associated with the older schema versions.
  - Optioanlly can be implemented queries for requesting schemas with any return type
- For each update:
  - Create a new Schema version (a new .proto file)
  - Implement transactions and queries that can handle data according to its version, including mechanisms for converting generic values into the corresponding schema version.

While offering a robust solution for handling radical changes, this method requires careful planning and development, which can potentially take a significant amount of time.

### Strategy for Structural Changes

For major organizational changes to the schema, such as splitting or merging schemas:

**Strategy steps:**

- For each update:
  - Create a new version of a Schema and state (a new .proto file)
  - Migrate older states to newer schema version.
  - Remove the states associated with the older schema versions.
  - Implement transactions and queries for the new schema version.
  - Implement transactions and queries that are backward compatible, specifically designed for the reorganized schema.

These types of changes are relatively rare. Implementing backward compatibility and migration can be complex.

## Conclusion

To lay the foundation for future compatibility improvements, it's a good idea to start by adding a version field to each schema. This step not only provides compatible updates to schemas, but also prepares the schemas for seamless integration with upcoming compatibility enhancements.
