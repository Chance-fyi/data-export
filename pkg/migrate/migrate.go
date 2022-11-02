package migrate

import (
	"data-export/pkg/console"
	"data-export/pkg/g"
	"gorm.io/gorm"
)

// migrationFunc define the types of up and down callback methods
type migrationFunc func(db *gorm.DB) error

// migrationFiles array of all migrated files, group by data connect
var migrationFiles = make(map[string][]MigrationFile, 2)

// MigrationFile single migration file
type MigrationFile struct {
	FileName string
	Up       migrationFunc
	Down     migrationFunc
}

// Migration record table structure
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// Add a migration file
func Add(mf MigrationFile, connect string) {
	migrationFiles[connect] = append(migrationFiles[connect], mf)
}

// Run migration method
func Run(function func(mf []MigrationFile, db *gorm.DB) error) {
	for connect, files := range migrationFiles {
		db := g.DB(connect)
		createMigrationsTable(db)

		err := db.Transaction(func(tx *gorm.DB) error {
			return function(files, tx)
		})
		console.ExitIf(err)
	}
}

// create migration table
func createMigrationsTable(db *gorm.DB) {
	migration := Migration{}

	if !db.Migrator().HasTable(&migration) {
		err := db.Migrator().CreateTable(&migration)
		console.ExitIf(err)
	}
}

// Up execute migration file
func Up(files []MigrationFile, db *gorm.DB) error {
	batch := getBatch(db)
	// get migrated data
	var migrations []Migration
	db.Find(&migrations)

	run := false
	for _, mf := range files {
		if mf.isNotMigrated(migrations) {
			if err := mf.runUpMigration(db, batch); err != nil {
				return err
			}
			run = true
		}
	}
	if !run {
		console.Successln("database is up to date.")
	}
	return nil
}

// get the value of the current batch
func getBatch(db *gorm.DB) int {
	batch := 1

	lastMigration := Migration{}
	db.Order("id DESC").First(&lastMigration)

	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// not migrated
func (mf *MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mf.FileName {
			return false
		}
	}
	return true
}

// up method for migration
func (mf *MigrationFile) runUpMigration(db *gorm.DB, batch int) error {
	if mf.Up != nil {
		console.Logln("migrating " + mf.FileName)
		if err := mf.Up(db); err != nil {
			return err
		}
		console.Logln("migrated " + mf.FileName)
	}
	// adding migration records to the database
	return db.Create(&Migration{
		Migration: mf.FileName,
		Batch:     batch,
	}).Error
}

// Rollback the last migration operation
func Rollback(files []MigrationFile, db *gorm.DB) error {
	lastMigration := Migration{}
	db.Order("id DESC").First(&lastMigration)

	var migrations []Migration
	db.Where("batch = ?", lastMigration.Batch).Find(&migrations)

	return runDownMigrations(migrations, files, db)
}

// Get the migration object by the name of the migration file
func getMigrationFile(name string, files []MigrationFile) MigrationFile {
	for _, mf := range files {
		if name == mf.FileName {
			return mf
		}
	}
	return MigrationFile{}
}

// down method for migration
func runDownMigrations(migrations []Migration, files []MigrationFile, db *gorm.DB) error {
	if len(migrations) == 0 {
		console.Successln("[migrations] table is empty, nothing to rollback.")
		return nil
	}

	for _, migration := range migrations {
		mf := getMigrationFile(migration.Migration, files)
		if mf.Down != nil {
			console.Logln("rollback " + mf.FileName)
			if err := mf.Down(db); err != nil {
				return err
			}
			if err := db.Delete(&migration).Error; err != nil {
				return err
			}
			console.Logln("finish " + mf.FileName)
		}
	}

	return nil
}

// Reset rollback all migrations that have been run
func Reset(files []MigrationFile, db *gorm.DB) error {
	var migrations []Migration
	db.Order("id DESC").Find(&migrations)

	return runDownMigrations(migrations, files, db)
}
