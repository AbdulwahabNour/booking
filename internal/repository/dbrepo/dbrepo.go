package dbrepo

import (
	"context"
	"time"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/repository"
	"github.com/jmoiron/sqlx"
)



type postgresDBRepo struct{
    App *config.AppConfig
    DB *sqlx.DB
}

func NewPostgressRepo(conn *sqlx.DB, a *config.AppConfig) repository.DatabaseRepo{

    return &postgresDBRepo{
         App: a,
         DB: conn,
    }
}

func(p *postgresDBRepo)InsertReservation(ctx context.Context, res *models.Reservation)(int ,error){
   var id int
   stmt := `INSERT INTO reservations
            (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
            VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
       `
   row := p.DB.QueryRowContext(ctx, stmt,
                        res.FirstName,
                        res.LastName,
                        res.Email,
                        res.Phone, 
                        res.StartDate,
                        res.EndDate,
                        res.RoomId,
                        time.Now(),
                        time.Now())
    err := row.Scan(&id)

    if err !=nil {
        p.App.ErrorLog.Println(err)
        return id, err
    }
    return id, nil
}

func(p *postgresDBRepo)InsertRoomRestrictions(ctx context.Context, res *models.RoomRestriction)error{
    
    stmt := `INSERT INTO room_restrictions(start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
                                    VALUES($1, $2, $3, $4, $5, $6, $7)`
    
    _, err := p.DB.ExecContext(ctx, stmt, res.StartDate,
                                          res.EndDate,
                                          res.RoomID,
                                          res.ReservationID,
                                          res.RestrictionID,
                                          time.Now(),
                                          time.Now())   
                                          
    if err != nil{
        p.App.ErrorLog.Println(err)
        return err
    }
 
    return nil
}
func(p *postgresDBRepo) CheckAvailabilityByDateAndRoom(ctx context.Context,roomId int, start, end time.Time) (bool, error){
    var rowCount int
    stmt := `SELECT count(id) FROM room_restrictions WHERE room_id= $1 and $2 <= end_date and $3 >= start_date;`

    row := p.DB.QueryRowContext(ctx, stmt,roomId, start, end)

    err := row.Scan(&rowCount)
    
    if err != nil{
        p.App.ErrorLog.Println(err)
        return false, err
    }

    return rowCount > 0, nil
}

func (p *postgresDBRepo) SearchAvailabilityForRooms(ctx context.Context, pageSize int, offset int, start, end time.Time) ([]models.Room, error){

    stmt := `
            SELECT room.id, room.room_name, room.image
            FROM
                rooms room
            WHERE room.id NOT IN(SELECT room_id FROM room_restrictions roomRes WHERE  $1 <= roomRes.end_date and $2 >= roomRes.start_date)
            ORDER BY id DESC
            LIMIT $3 OFFSET $4
    `
    rows, err := p.DB.QueryContext(ctx, stmt, start, end, pageSize,  offset)

    if err != nil{
        return nil, err
    }


    var rooms []models.Room

    for rows.Next(){

        var room models.Room
        if err := rows.Scan(&room.ID,
                            &room.RoomName,
                            &room.RoomImage); err != nil{
                return nil, err
        }
      rooms = append(rooms, room)
    }

    if err := rows.Close(); err != nil{
        p.App.ErrorLog.Println(err)
        return nil, err
    }
    
	if err := rows.Err(); err != nil {
        p.App.ErrorLog.Println(err)
		return nil, err
	}

    return rooms, nil
}
func (p *postgresDBRepo) GetRoomById(ctx context.Context, id int) (*models.Room, error){
    
    var room models.Room

    stmt := "SELECT  *  FROM rooms WHERE id = $1"

    err := p.DB.QueryRowContext(ctx, stmt, id).Scan(&room.ID, &room.RoomImage, &room.RoomName, &room.CreatedeAt, &room.UpdatedAt)
    if  err != nil{
        p.App.ErrorLog.Println(err)
        return nil, err
    }

    return &room, nil
}

func (p *postgresDBRepo) CountRooms(ctx context.Context) (int, error){
    var totalRoomsNumber int
    stmt := "SELECT COUNT(id) FROM rooms"

    err := p.DB.QueryRowContext(ctx, stmt).Scan(&totalRoomsNumber)
    if  err != nil{
        p.App.ErrorLog.Println(err)
        return 0, err
    }

    return totalRoomsNumber, nil
}

func (p *postgresDBRepo) GetReservationById(ctx context.Context, id int) (*models.Reservation, error){
    
    var res models.Reservation

    stmt := "SELECT  *  FROM reservations WHERE id = $1"

    err := p.DB.QueryRowContext(ctx, stmt, id).Scan(&res.ID,
                                                     &res.FirstName,
                                                    &res.LastName,
                                                    &res.Email, &res.Phone,
                                                    &res.StartDate, 
                                                    &res.EndDate, 
                                                    &res.RoomId, 
                                                    &res.CreatedeAt,  
                                                    &res.UpdatedAt)
    if  err != nil{
        p.App.ErrorLog.Println(err)
        return nil, err
    }

    return &res, nil
}

func (p *postgresDBRepo)GetRoomsByOffset(ctx context.Context, pageSize int, offset int) ([]models.Room, error){

 stmt := `SELECT id, room_name, image FROM rooms ORDER BY id DESC LIMIT $1 OFFSET $2`
 rows, err := p.DB.QueryContext(ctx, stmt, pageSize, offset)
 if  err !=nil {
    return nil, err
 }

 var rooms []models.Room

 for rows.Next(){
        var room models.Room
        if err := rows.Scan(&room.ID,
                            &room.RoomName,
                            &room.RoomImage); err != nil{
                return nil, err
        }
        rooms = append(rooms, room)
    }

    if  err := rows.Close(); err != nil{
        p.App.ErrorLog.Println(err)
        return nil, err
    }
	if err := rows.Err(); err != nil {
        p.App.ErrorLog.Println(err)
		return nil, err
	}
    return rooms, nil
}