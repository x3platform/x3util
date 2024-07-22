package osutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceEnvValues(t *testing.T) {
	os.Setenv("MYSQL_HOST", "mysql.cluster.local")
	os.Setenv("MYSQL_PORT", "3306")

	dsn := ReplaceEnvValues("@tcp(${MYSQL_HOST}:${MYSQL_PORT})")

	// if dsn == "@tcp(mysql.cluster.local:3306)" {
	// 	t.Fail()
	// }

	assert.Equal(t, "@tcp(mysql.cluster.local:3306)", dsn)
}
