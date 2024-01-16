package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewProductMysql creates a new instance of the product repository
func NewProductMysql(db *sql.DB) *ProductMysql {
	return &ProductMysql{db}
}

// ProductMysql is the mysql implementation of the product repository
type ProductMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all products from the database
func (r *ProductMysql) FindAll() (products []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `p.id`, `p.product_code`, `p.description`, `p.height`, `p.length`, `p.width`, `p.weight`, `p.expiration_rate`, `p.freezing_rate`, `p.recom_freez_temp`, `p.product_type_id`, `p.seller_id` FROM `products` AS `p`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new product
		var product internal.Product
		err = rows.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID)
		if err != nil {
			return
		}

		// append the product to the slice
		products = append(products, product)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a product from the database by its id
func (r *ProductMysql) FindByID(id int) (product internal.Product, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `p.id`, `p.product_code`, `p.description`, `p.height`, `p.length`, `p.width`, `p.weight`, `p.expiration_rate`, `p.freezing_rate`, `p.recom_freez_temp`, `p.product_type_id`, `p.seller_id` FROM `products` AS `p` WHERE `p.id` = ?", id)

	// scan the row into the product
	err = row.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrProductRepositoryNotFound
		}
		return
	}

	return
}

// Save saves a product into the database
func (r *ProductMysql) Save(product *internal.Product) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `products` (`product_code`, `description`, `height`, `length`, `width`, `weight`, `expiration_rate`, `freezing_rate`, `recom_freez_temp`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		(*product).ProductCode, (*product).Description, (*product).Height, (*product).Length, (*product).Width, (*product).Weight, (*product).ExpirationRate, (*product).FreezingRate, (*product).RecomFreezTemp, (*product).ProductTypeID, (*product).SellerID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrProductRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// get the id of the inserted product
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the product
	(*product).ID = int(id)

	return
}

// Update updates a product in the database
func (r *ProductMysql) Update(product *internal.Product) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `products` SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recom_freez_temp` = ?, `product_type_id` = ?, `seller_id` = ? WHERE `id` = ?",
		(*product).ProductCode, (*product).Description, (*product).Height, (*product).Length, (*product).Width, (*product).Weight, (*product).ExpirationRate, (*product).FreezingRate, (*product).RecomFreezTemp, (*product).ProductTypeID, (*product).SellerID, (*product).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrProductRepositoryDuplicated
			default:
				// ...
			}
			return
		}
	}

	return
}

// Delete deletes a product from the database by its id
func (r *ProductMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `products` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}