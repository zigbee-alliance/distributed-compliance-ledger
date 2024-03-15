package common

type SchemaVersion interface {
	GetSchemaVersion() uint32
}

func GetSchemaVersionOrDefault(schemaVersion SchemaVersion) uint32 {
	currentVersion := schemaVersion.GetSchemaVersion()
	if currentVersion == 0 {
		return 1
	}

	return currentVersion
}
