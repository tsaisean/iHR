package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ = Describe("", func() {
	var (
		mockDb          *sql.DB
		mockSQL         sqlmock.Sqlmock
		mockRedisClient *redis.Client
		mockRedis       redismock.ClientMock
		empRepo         EmployeeRepository
	)
	BeforeEach(func() {
		mockDb, mockSQL, _ = sqlmock.New()
		mockSQL.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))
		mockRedisClient, mockRedis = redismock.NewClientMock()
		dialector := mysql.New(mysql.Config{
			Conn:       mockDb,
			DriverName: "mysql",
		})
		db, _ := gorm.Open(dialector, &gorm.Config{})
		empRepo = NewEmployeeRepo(db, mockRedisClient)
	})

	Context("When the request is valid", func() {
		It("should delete an employee and cache with no error", func() {
			// Arrange
			employeeID := uint(1)
			limit := 1

			mockSQL.ExpectQuery("SELECT \\* FROM \\`employees\\` WHERE \\`employees\\`\\.\\`id\\` = \\? ORDER BY \\`employees\\`\\.\\`id\\` LIMIT \\?").
				WithArgs(employeeID, limit).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(employeeID))
			mockSQL.ExpectBegin()
			mockSQL.ExpectExec("DELETE FROM `employees` WHERE `employees`.`id` = ?").
				WithArgs(employeeID).
				WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected
			mockSQL.ExpectCommit()

			// MiddleAuth will cache the employee result
			mockRedis.ExpectGet(GetEmployeeCacheKey(1)).RedisNil()
			mockRedis.ExpectEval(LuaScript, []string{}, "employees:*").SetVal(2)

			// Act
			err := empRepo.DeleteEmployee(nil, employeeID)

			// Assert
			gomega.Expect(err).ToNot(gomega.HaveOccurred())                     // Ensure no error
			gomega.Expect(mockSQL.ExpectationsWereMet()).To(gomega.Succeed())   // Ensure SQL expectations met
			gomega.Expect(mockRedis.ExpectationsWereMet()).To(gomega.Succeed()) // Ensure Redis expectations met
		})
	})
})
