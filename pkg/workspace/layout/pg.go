package layout

func (gl GoLayout) PgMigrationDir(databaseName string) WorkspaceFolder {
	return gl.path.append("migrations", databaseName)
}
