package booking

import (
	"Accounting/services"
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type bookingRepository struct {
	ctx context.Context
	db  *sql.DB
}

//NewBookingRepository create student repository
func NeBookingRepository(ctx context.Context, db *sql.DB) *bookingRepository {
	return &bookingRepository{
		ctx: ctx,
		db:  db,
	}
}

func (b *bookingRepository) GetBookingFromDatabase(filter map[string]string) ([]BookingModel, error) {
	var rs BookingModel
	var rsl []BookingModel

	qb := new(services.BuilderMapping)
	qb.Select("bk.id", &rs.ID)
	qb.Select("tp.id as typeid", &rs.NotetypeID)
	qb.Select("tp.type as type", &rs.Notetype)
	qb.Select("cg.id as categoryid", &rs.CategoryID)
	qb.Select("cg.category as category", &rs.Category)
	qb.Select("bk.amount", &rs.Amount)
	qb.Select("bk.note", &rs.Note)
	qb.Select("bk.created_at", &rs.CreatedAt)
	qb.Select("bk.updated_at", &rs.UpdatedAt)

	for k, v := range filter {
		if !services.StringEmpty(v) {
			switch k {
			case "dateFrom":
				qb.Where("bk.created_at >= ?", v)
			case "dateTo":
				qb.Where("bk.created_at <= ?", v)
			default:
				qb.Where("bk."+k+" = ?", v)
			}
		}
	}

	qb.Join("inner join category as cg on bk.category = cg.ID")
	qb.Join("inner join type as tp on bk.type = tp.ID")

	query := qb.Get("bookkeeping as bk")
	rows, err := b.db.QueryContext(b.ctx, query.Query, query.Value...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(query.Object...); err != nil {
			return nil, err
		}
		rsl = append(rsl, rs)
	}
	return rsl, nil
}

func (b *bookingRepository) PostBookingToDatabase(req *BookingModel) error {
	qb := new(services.BuilderMapping)
	qb.Insert("type", req.NotetypeID)
	qb.Insert("category", req.CategoryID)
	qb.Insert("amount", req.Amount)
	qb.Insert("note", req.Note)
	qb.Insert("created_at", time.Now())
	qb.Insert("updated_at", time.Now())

	query := qb.Post("bookkeeping")
	ex, err := b.db.PrepareContext(b.ctx, query.Query)

	if err != nil {
		return err
	}

	_, err = ex.ExecContext(b.ctx, query.Value...)

	if err != nil {
		return err
	}

	return nil
}

func (b *bookingRepository) PutBookingOnDatabase(req *BookingModel, filter map[string]string) error {
	qb := new(services.BuilderMapping)

	if !services.IntegerEmpty(req.CategoryID) {
		qb.Update("category = ?", req.CategoryID)
	}
	if !services.IntegerEmpty(req.NotetypeID) {
		qb.Update("type = ?", req.NotetypeID)
	}
	if !services.IntegerEmpty(req.Amount) {
		qb.Update("amount = ?", req.Amount)
	}
	if !services.StringEmpty(req.Note) {
		qb.Update("note = ?", req.Note)
	}

	qb.Update("updated_at = ?", time.Now())
	qb.Where("id = ?", filter["id"])

	query := qb.Put("bookkeeping")
	rsl, err := b.db.PrepareContext(b.ctx, query.Query)
	if err != nil {
		return err
	}

	_, err = rsl.ExecContext(b.ctx, query.Value...)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookingRepository) DeleteBookingOnDatabase(filter map[string]string) error {
	qb := new(services.BuilderMapping)
	qb.Where("id = ?", filter["id"])

	query := qb.Delete("bookkeeping")
	rsl, err := b.db.PrepareContext(b.ctx, query.Query)
	if err != nil {
		return err
	}

	_, err = rsl.ExecContext(b.ctx, query.Value...)
	if err != nil {
		return err
	}
	return nil
}
