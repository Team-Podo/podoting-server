package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func TestSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sqlmock Suite")
}

var _ = Describe("Product", func() {
	var repo ProductRepository
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).NotTo(HaveOccurred())

		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{})
		Expect(err).NotTo(HaveOccurred())

		repo = ProductRepository{DB: gormDB}
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet() // 모든 기대가 충족되었는지 확인
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("GetWithQueryMap", func() {
		const defaultSelectQuery = "SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL"
		var toInsertRows = []string{"id", "title", "file_id", "created_at", "updated_at", "deleted_at"}

		It("Success", func() {
			mock.ExpectQuery(regexp.QuoteMeta(defaultSelectQuery)).
				WillReturnRows(sqlmock.NewRows(toInsertRows).
					AddRow(1, "웃는 남자", nil, time.Now(), time.Now(), nil).
					AddRow(2, "안웃는 남자", nil, time.Now(), time.Now(), nil))
			products, err := repo.GetWithQueryMap(map[string]any{})

			Expect(err).NotTo(HaveOccurred())
			Expect(products).To(HaveLen(2))
		})

		It("Reversed", func() {
			const sqlQuery = defaultSelectQuery + " ORDER BY id desc"
			mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
				WillReturnRows(sqlmock.NewRows(toInsertRows).
					AddRow(1, "웃는 남자", nil, time.Now(), time.Now(), nil).
					AddRow(2, "안웃는 남자", nil, time.Now(), time.Now(), nil))
			products, err := repo.GetWithQueryMap(map[string]any{
				"reversed": true,
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(products).To(HaveLen(2))
		})

		It("Limit", func() {
			const sqlQuery = defaultSelectQuery + " LIMIT 1"
			mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
				WillReturnRows(sqlmock.NewRows(toInsertRows).
					AddRow(1, "웃는 남자", nil, time.Now(), time.Now(), nil))
			products, err := repo.GetWithQueryMap(map[string]any{
				"limit": 1,
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(products).To(HaveLen(1))
		})

		It("Offset", func() {
			const sqlQuery = defaultSelectQuery + " OFFSET 1"
			mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
				WillReturnRows(sqlmock.NewRows(toInsertRows).
					AddRow(2, "안웃는 남자", nil, time.Now(), time.Now(), nil))
			products, err := repo.GetWithQueryMap(map[string]any{
				"offset": 1,
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(products).To(HaveLen(1))
		})
	})

	Describe("FindByID", func() {
		It("Success", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL AND `products`.`id` = ? ORDER BY `products`.`id` LIMIT 1")).
				WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "title", "file_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "웃는 남자", nil, time.Now(), time.Now(), nil))
			product, err := repo.FindByID(1)

			Expect(err).NotTo(HaveOccurred())
			Expect(product.ID).To(Equal(uint(1)))
		})
	})
})
