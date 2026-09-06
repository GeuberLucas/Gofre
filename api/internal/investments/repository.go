package investments

import "database/sql"

type IPortfolioRepository interface {
	Create(model Portfolio) error
	GetAll(userId uint) ([]Portfolio, error)
	GetById(id uint) (Portfolio, error)
	Update(model Portfolio) error
	Delete(id uint, userId uint) error
}

type PortfolioRepository struct {
	db *sql.DB
}

// Create implements IPortfolioRepository.
func (p *PortfolioRepository) Create(model Portfolio) error {
	sqlCommand := `INSERT INTO investments.portfolio(user_id,asset_id,deposit_date,broker,amount,description,is_done)
					values ($1,$2,$3,$4,$5,$6,$7);`

	statement, err := p.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}

	_, err = statement.Exec(model.User_id, model.Asset_id, model.Deposit_date, model.Broker, model.Amount, model.Description, model.IsDone)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

// Delete implements IPortfolioRepository.
func (p *PortfolioRepository) Delete(id uint, userId uint) error {
	sqlCommand := `DELETE from investments.portfolio where id=$1 and user_id=$2;`

	statement, err := p.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id, userId)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

// GetAll implements IPortfolioRepository.
func (p *PortfolioRepository) GetAll(userId uint) ([]Portfolio, error) {
	var sqlCommand string = `SELECT 
	id,
	user_id,
	asset_id,
	deposit_date,
	broker,
	amount,
	description,
	is_done 
	FROM investments.portfolio 
	where user_id=$1`

	rows, err := p.db.Query(sqlCommand, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var portfolioDtos []Portfolio
	for rows.Next() {
		var portfolioModel Portfolio
		err := rows.Scan(&portfolioModel.Id,
			&portfolioModel.User_id, &portfolioModel.Asset_id,
			&portfolioModel.Deposit_date, &portfolioModel.Broker, &portfolioModel.Amount, &portfolioModel.Description, &portfolioModel.IsDone)
		if err != nil {
			return nil, err
		}
		portfolioDtos = append(portfolioDtos, portfolioModel)
	}
	return portfolioDtos, nil
}

// GetById implements IPortfolioRepository.
func (p *PortfolioRepository) GetById(id uint) (Portfolio, error) {
	var portfolioModel Portfolio
	var sqlCommand string = `SELECT 
	id,
	user_id,
	asset_id,
	deposit_date,
	broker,
	amount,
	description,
	is_done 
	FROM investments.portfolio 
	where id=$1;`

	row := p.db.QueryRow(sqlCommand, id)
	err := row.Scan(&portfolioModel.Id,
		&portfolioModel.User_id, &portfolioModel.Asset_id,
		&portfolioModel.Deposit_date, &portfolioModel.Broker, &portfolioModel.Amount, &portfolioModel.Description, &portfolioModel.IsDone)
	if err != nil {
		return portfolioModel, err
	}
	return portfolioModel, nil
}

// Update implements IPortfolioRepository.
func (p *PortfolioRepository) Update(model Portfolio) error {
	var sqlCommand string = `UPDATE investments.portfolio SET asset_id=$1, deposit_date=$2, broker=$3, amount=$4, description=$5, is_done=$6 WHERE id=$7 and user_id=$8`

	statement, err := p.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(model.Asset_id, model.Deposit_date, model.Broker, model.Amount, model.Description, model.IsDone, model.Id, model.User_id)
	if err != nil {

		return err
	}
	return nil
}

func NewPortfolioRepository(db *sql.DB) IPortfolioRepository {
	return &PortfolioRepository{db: db}
}
