package query

// import (
//     "github.com/diogenesdornelles/pgorm/internal/domain/models"
// )


// type UserQuery[T any] interface {
//     // SELECT * FROM @@table WHERE id=@id
//     GetByID(id int) (T, error)

//     // SELECT * FROM @@table WHERE @@column=@value
//     FilterWithColumn(column string, value string) ([]T, error)

//     // SELECT * FROM @@table
//     // {{where}}
//     //   {{if user.Name != ""}} nome=@user.Name {{end}}
//     //   {{if user.Age > 0}} AND idade=@user.Age {{end}}
//     //   {{if user.Email.Valid}} AND email=@user.Email.String {{end}}
//     //   {{if user.CompanyID > 0}} AND company_id=@user.CompanyID {{end}}
//     // {{end}}
//     // ORDER BY id DESC
//     // LIMIT @limit OFFSET @offset
//     SearchUsers(user models.User, limit, offset int) ([]T, error)

//     // INSERT INTO @@table
//     //   (nome, idade, email, company_id, score, aniversario, numero_membro, ativado_em)
//     // VALUES
//     //   (@user.Name, @user.Age, @user.Email, @user.CompanyID, @user.Score,
//     //    @user.Birthday, @user.MemberNumber, @user.ActivatedAt)
//     CreateUser(user models.User) error

//     // UPDATE @@table
//     // {{set}}
//     //   {{if user.Name != ""}} nome=@user.Name, {{end}}
//     //   {{if user.Age > 0}} idade=@user.Age, {{end}}
//     //   {{if user.Email.Valid}} email=@user.Email.String, {{end}}
//     //   {{if user.CompanyID > 0}} company_id=@user.CompanyID, {{end}}
//     //   {{if user.Score > 0}} score=@user.Score, {{end}}
//     // {{end}}
//     // WHERE id=@id
//     UpdateUser(user models.User, id int) error

//     // UPDATE @@table SET deleted_at=NOW() WHERE id=@id AND deleted_at IS NULL
//     SoftDeleteUser(id int) error

//     // DELETE FROM @@table WHERE id=@id
//     HardDeleteUser(id int) error
// }