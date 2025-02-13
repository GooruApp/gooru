package migrations

import "embed"

//go:embed postgres/*
var PGMigrations embed.FS

//go:embed sqlite/*
var SQLiteMigrations embed.FS
