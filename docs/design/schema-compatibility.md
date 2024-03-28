# Support for forward and backward compatibility in DCL schemes

Schema changes can cover a wide range of modifications with varying impacts on application compatibility and data integrity. Below are use cases with strategies to manage schema changes and ensure compatibility.

## I. Multiple versions can live in parallel

### 1. Strategy for Compatible Changes

For changes that are backward-compatible, such as adding optional fields or extending enumerations.

#### Option A: Add an optional version field to all DCL schema

**Description:**
Implement an optional version field in all DCL schemas to track the schema version. This approach is simple and quick to execute, suitable primarily for compatible updates.

**Strategy steps:**

- One time actions:
  - Add an optional version field to all DCL schema
- For each update:
  - Update the schema by introducing compatible changes (such as adding a new optional field).
  - Update transactions and queries if needed.
  - DCL doesn't fulfill the Schema version automatically
  - It will be up to the transaction submitter (Vendor) to specify a correct Schema version
  - If Schema Version is not set - then the initial version (version 0 or 1) is assumed
  - It will be up to the client application to process the Schema version

### 2. Strategy for Non-Compatible Changes

For significant changes that directly impact compatibility, such as adding mandatory fields or removing fields, splitting or merging schemas, changing enumerations.

#### Option B: Separate Schemas for Each Version

**Description:**
Each version has its distinct schema, state and its own queries/requests. This strategy eliminates the need for data migration and allows different schema versions to coexist seamlessly.

**Strategy steps:**

- For each update:
  - Create a new version of a Schema and state (a new .proto file)
  - Implement transactions and queries for the new schema version.

#### Option C: Generic Schema Storage (Not Recommended for Production)

**Description:**
Implement a flexible, generic schema structure that can support a wide range of data formats.

While offering a robust solution for handling radical changes, this method requires careful planning and development, which can potentially take a significant amount of time.

**Strategy steps:**

- One time actions:
  - Create a more flexible, generic schema structure to hold a wide range of data formats (Can be used [Any](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/any.proto) as described in [ADR-19](https://docs.cosmos.network/v0.47/build/architecture/adr-019-protobuf-state-encoding#usage-of-any-to-encode-interfaces))
  - Migrate old states to the newer, generic schema.
  - Remove the states associated with the older schema versions.
  - Optioanlly can be implemented queries for requesting schemas with any return type
- For each update:
  - Create a new Schema version (a new .proto file)
  - Implement transactions and queries that can handle data according to its version, including mechanisms for converting generic values into the corresponding schema version.

## II. New version replaces the legacy one (V2 replaces V1)

### 1. Strategy for Compatible or Convertible changes

For changes that are backward-compatible, such as adding optional or mandatory fields or extending enumerations

#### Option D: Not keeping backward compatibility in API

**Description:**
This strategy focuses on updating the schema without ensuring backward compatibility at the API level. Since the schemas are compatible, there will likely be no need for migration.

**Strategy steps:**

- For each update:
  - Update the schema by introducing compatible changes (such as adding a new optional field).
  - Migrate old states to the newer if needed.
  - Update transactions and queries if needed.

#### Option E: Keeping backward compatibility in API

**Description:**
The main idea of this strategy is the dynamically converting newer schemas into older ones. However, this method is only possible if there is compatibility between the newer and legacy schemas, allowing them to be converted to each other. Due to the on-the-fly data conversion, this approach does not support the Light Client in legacy APIs because the converted data is not stored in the state, preventing the generation of proofs.

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

### 2. Strategy for Non-Compatible changes

For significant changes that directly impact compatibility, such as adding mandatory fields or removing fields, splitting or merging schemas, changing enumerations.

#### Optiona F: Not keeping backward compatibility in API

**Description:**
This strategy focuses on updating the schema without ensuring backward compatibility at the API level. Since the schemas are not compatible, migration is carried out manually through a special transaction.

**Strategy steps:**

- For each update:
  - Update the schema by introducing changes.
  - Update transactions and queries if needed.
  - Add a new transaction to fulfill new required fields (essentially this is a manual migration via transactions)

#### Option G: Keeping backward compatibility in API

**Description:**
It's not possible to replace an old version here. [Multiple versions can live in parallel: Strategy for Non-Compatible Changes](#2-strategy-for-non-compatible-changes) options should be used instead.

## Conclusion

To lay the foundation for future compatibility improvements, it's a good idea to start by adding a version field to each schema. For subsequent changes, we will then select the most appropriate strategy based on the nature of these changes.
