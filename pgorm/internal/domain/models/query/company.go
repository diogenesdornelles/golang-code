package query

// import (
//     "github.com/diogenesdornelles/pgorm/internal/domain/models"
// )

// type CompanyQuery[T any] interface {
//     // SELECT * FROM @@table WHERE id=@id
//     GetByID(id int) (T, error)

//     // SELECT * FROM @@table
//     // {{where}}
//     //   {{if company.Name != ""}} name LIKE CONCAT('%', @company.Name, '%') {{end}}
//     // {{end}}
//     // ORDER BY id DESC
//     // LIMIT @limit OFFSET @offset
//     SearchCompanies(company models.Company, limit, offset int) ([]T, error)

//     // INSERT INTO @@table (name) VALUES (@company.Name)
//     CreateCompany(company models.Company) error

//     // UPDATE @@table
//     // {{set}}
//     //   {{if company.Name != ""}} name=@company.Name, {{end}}
//     // {{end}}
//     // WHERE id=@id
//     UpdateCompany(company models.Company, id int) error

//     // DELETE FROM @@table WHERE id=@id
//     DeleteCompany(id int) error

//     // SELECT c.*, u.* FROM companies c
//     // LEFT JOIN users u ON u.company_id = c.id
//     // WHERE c.id=@id
//     GetCompanyWithUsers(id int) (T, error)
// }