package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User representa um modelo simples de usuário.
type User struct {
	ID   uint   // ID do usuário (chave primária)
	Name string // Nome do usuário
}

// DatabaseTestSuite é a suíte de testes.
type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB // Instância do banco de dados
}

// SetupSuite é chamado uma vez antes da execução da suíte de testes.
func (suite *DatabaseTestSuite) SetupSuite() {
	// Configura uma base de dados PostgreSQL para os testes
	dsn := "user=testuser password=testpassword dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	suite.Require().NoError(err, "Erro ao conectar à base de dados de teste")

	// Habilita o log do Gorm durante os testes
	suite.db = db.Debug()

	// Realiza a automigração das tabelas
	err = suite.db.AutoMigrate(&User{})
	suite.Require().NoError(err, "Erro ao realizar a automigração das tabelas do banco de dados")
}

// TestUserInsertion testa a inserção de um registro de usuário.
func (suite *DatabaseTestSuite) TestUserInsertion() {
	// Cria um usuário
	user := User{Name: "John Doe"}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err, "Erro ao criar o registro do usuário")

	// Recupera o usuário inserido
	var retrievedUser User
	err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
	suite.Require().NoError(err, "Erro ao recuperar o registro do usuário")

	// Verifica se o usuário recuperado corresponde ao usuário inserido
	suite.Equal(user.Name, retrievedUser.Name, "Os nomes devem coincidir")
}

// TearDownSuite é chamado uma vez após a execução da suíte de testes.
func (suite *DatabaseTestSuite) TearDownSuite() {
	// Limpeza: Fecha a conexão com o banco de dados
	err := suite.db.Exec("DROP TABLE users;").Error
	suite.Require().NoError(err, "Erro ao dropar a tabela de teste")

	err = suite.db.Close()
	suite.Require().NoError(err, "Erro ao fechar a base de dados de teste")
}

// TestSuite executa a suíte de testes.
func TestSuite(t *testing.T) {
	// Pula os testes se os detalhes da conexão PostgreSQL não forem fornecidos
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("Pulando os testes de PostgreSQL; forneça a variável de ambiente POSTGRES_DSN.")
	}

	suite.Run(t, new(DatabaseTestSuite))
}
