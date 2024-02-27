package chains

import (
	"context"
	"os"
	"testing"

	"github.com/robermar23/langchaingo/llms/openai"
	"github.com/robermar23/langchaingo/tools/sqldatabase"
	"github.com/robermar23/langchaingo/tools/sqldatabase/mysql"
	"github.com/stretchr/testify/require"
)

func TestSQLDatabaseChain_Call(t *testing.T) {
	t.Parallel()
	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
		t.Skip("OPENAI_API_KEY not set")
	}

	llm, err := openai.New()
	require.NoError(t, err)

	// export LANGCHAINGO_TEST_MYSQL=user:p@ssw0rd@tcp(localhost:3306)/test
	mysqlURI := os.Getenv("LANGCHAINGO_TEST_MYSQL")
	if mysqlURI == "" {
		t.Skip("LANGCHAINGO_TEST_MYSQL not set")
	}
	engine, err := mysql.NewMySQL(mysqlURI)
	require.NoError(t, err)

	db, err := sqldatabase.NewSQLDatabase(engine, nil)
	require.NoError(t, err)

	chain := NewSQLDatabaseChain(llm, 5, db)
	input := map[string]interface{}{
		"query":              "How many cards are there?",
		"table_names_to_use": []string{"AllianceAuthority", "AllianceGift", "Card"},
	}
	result, err := chain.Call(context.Background(), input)
	require.NoError(t, err)

	ret, ok := result["result"].(string)
	require.True(t, ok)
	require.NotEmpty(t, ret)

	t.Log(ret)
}
