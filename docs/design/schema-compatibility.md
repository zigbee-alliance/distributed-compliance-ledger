# Support for forward and backward compatibility in DCL schemes

Schema changes can cover a wide range of modifications with varying impacts on application compatibility and data integrity. Below are use cases with strategies to manage schema changes and ensure compatibility.

## Strategy for Compatible Changes

For changes that are backward-compatible, such as adding optional fields or extending enumerations:

**Strategy steps:**

- One time actions:
  - Add a version field to DCL schema to track the schema version.
- For each update:
  - Update the schema.
  - Probably need to update transactions and queries.

This strategy is straightforward and quick to implement, but only suitable for compatible changes.

## Strategy for Convertible Changes

For changes that affect compatibility but can be converted, like field type changes, renaming fields, and changing enumerations:

**Strategy steps:**

- For each update:
  - Create a new schema version and state
  - Migrate older states to newer schema version.
  - Implement transactions and queries capable of converting data between the latest and older schema version.

The specific implementation details will vary with each change, each time requires an individual approach.

## Strategy for Handling Non-Convertible Changes

For significant changes that directly impact compatibility, such as adding mandatory fields or removing fields:

**Strategy steps:**

- One time actions:
  - Create a more flexible, generic schema structure to hold a wide range of data formats (Can be used [Any](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/any.proto) as described in [ADR-19](https://docs.cosmos.network/v0.47/build/architecture/adr-019-protobuf-state-encoding#usage-of-any-to-encode-interfaces))
  - Migrate old states to the newer, generic schema.
  - Implement transactions and queries that can handle data according to its version, including mechanisms for converting generic values into the corresponding schema version.
- For each update:
  - Create a new schema version
  - Probably need to update transactions and queries

While offering a robust solution for handling radical changes, this method requires careful planning and development, which can potentially take a significant amount of time.

## Strategy for Structural Changes

For major organizational changes to the schema, such as splitting or merging schemas:

**Strategy steps:**

- For each update:
  - Create a new schema version and state
  - Migrate older states to newer schema version.
  - Implement transactions and queries specifically designed for the reorganized schema.

These types of changes are relatively rare. Implementing backward compatibility and migration can be complex.

## Conclusion

To lay the foundation for future compatibility improvements, it's a good idea to start by adding a version field to each schema. This step not only provides compatible updates to schemas, but also prepares the schemas for seamless integration with upcoming compatibility enhancements.
