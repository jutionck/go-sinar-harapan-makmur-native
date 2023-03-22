package repository

import (
	"database/sql"
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/utils/common"
)

type VehicleRepository interface {
	BaseRepository[model.Vehicle]
	BaseRepositoryPaging[model.Vehicle]
}

type vehicleRepository struct {
	db *sql.DB
}

func (v *vehicleRepository) Create(newData model.Vehicle) error {
	sql := "INSERT INTO vehicle (id, brand, model, production_year, color, is_automatic, sale_price, stock, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := v.db.Exec(sql, newData.Id, newData.Brand, newData.Model, newData.ProductionYear, newData.Color, newData.IsAutomatic, newData.SalePrice, newData.Stock, newData.Status)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) List() ([]model.Vehicle, error) {
	sql := `SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle`
	rows, err := v.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var vehicle []model.Vehicle
	for rows.Next() {
		var vehilce model.Vehicle
		err := rows.Scan(&vehilce.Id, &vehilce.Brand, &vehilce.Model, &vehilce.ProductionYear, &vehilce.Color, &vehilce.IsAutomatic, &vehilce.SalePrice, &vehilce.Stock, &vehilce.Status)
		if err != nil {
			return nil, err
		}
		vehicle = append(vehicle, vehilce)
	}
	return vehicle, nil
}

func (v *vehicleRepository) Get(id string) (model.Vehicle, error) {
	sql := `SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle WHERE id = $1`
	var vehilce model.Vehicle
	err := v.db.QueryRow(sql, id).Scan(&vehilce.Id, &vehilce.Brand, &vehilce.Model, &vehilce.ProductionYear, &vehilce.Color, &vehilce.IsAutomatic, &vehilce.SalePrice, &vehilce.Stock, &vehilce.Status)
	if err != nil {
		return model.Vehicle{}, err
	}
	return vehilce, nil
}

func (v *vehicleRepository) Update(newData model.Vehicle) error {
	sql := "UPDATE vehicle set brand = $1, model = $2, production_year = $3, color = $4, is_automatic = $5, sale_price = $6, stock = $7, status = $8 WHERE id = $9"
	_, err := v.db.Exec(sql, &newData.Brand, &newData.Model, &newData.ProductionYear, &newData.Color, &newData.IsAutomatic, &newData.SalePrice, &newData.Stock, &newData.Status, newData.Id)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) Delete(id string) error {
	sql := "DELETE FROM vehicle WHERE id = $1"
	_, err := v.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestQueryParams.PaginationParam)
	orderQuery := "ORDER BY id"
	if requestQueryParams.QueryParams.Order != "" && requestQueryParams.QueryParams.Sort != "" {
		sorting := "ASC"
		if requestQueryParams.QueryParams.Sort == "desc" {
			sorting = "DESC"
		}
		orderQuery = fmt.Sprintf("ORDER BY %s %s", requestQueryParams.QueryParams.Order, sorting)
	}
	sql := fmt.Sprintf("SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle %s LIMIT $1 OFFSET $2", orderQuery)
	rows, err := v.db.Query(sql, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}
	}

	var vehicle []model.Vehicle
	for rows.Next() {
		var vehilce model.Vehicle
		err := rows.Scan(&vehilce.Id, &vehilce.Brand, &vehilce.Model, &vehilce.ProductionYear, &vehilce.Color, &vehilce.IsAutomatic, &vehilce.SalePrice, &vehilce.Stock, &vehilce.Status)
		if err != nil {
			return nil, dto.Paging{}
		}
		vehicle = append(vehicle, vehilce)
	}
	totalRows, err := v.getTotalRows()
	if err != nil {
		return nil, dto.Paging{}
	}
	return vehicle, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows)
}

func (v *vehicleRepository) getTotalRows() (int, error) {
	sql := `SELECT COUNT(*) FROM vehicle`
	row := v.db.QueryRow(sql)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewVehicleRepository(db *sql.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}
