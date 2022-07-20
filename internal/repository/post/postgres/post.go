package postgres

import (
	"database/sql"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	postRepo "postservice/internal/repository/post"
	"postservice/pkg/logger"
	"strconv"
)

type postRepository struct {
	lg      logger.ILogger
	adapter *sql.DB
}

func (p *postRepository) Load(url string) (postRepo.IResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		p.lg.Error("PostgresPostRepository.Load.Get.Error", zap.Error(err))
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.lg.Error("PostgresPostRepository.Load.ReadBody.Error", zap.Error(err))
		return nil, err
	}

	var response postRepo.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		p.lg.Error("PostgresPostRepository.Load.Get.Error", zap.Error(err))
		return nil, err
	}

	return &response, err
}

func (p *postRepository) SaveBatch(posts []postRepo.Post) (int64, error) {
	vals := []interface{}{}

	sqlStr := "insert into " + p.TableName() + "(id, user_id, title, body) values"
	for idx, post := range posts {
		sqlStr += "($" + strconv.Itoa(idx*4+1) + ", $" + strconv.Itoa(idx*4+2) + ", $" + strconv.Itoa(idx*4+3) + ", $" + strconv.Itoa(idx*4+4) + "),"
		vals = append(vals, post.GetId(), post.GetUserId(), post.GetTitle(), post.GetBody())
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	sqlStr += " on conflict (id) do update set user_id = EXCLUDED.user_id, title = EXCLUDED.title, body = EXCLUDED.body;"

	stmt, err := p.adapter.Prepare(sqlStr)
	if err != nil {
		p.lg.Error("PostgresPostRepository.SaveBatch.Prepare.Error", zap.Error(err))
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(vals...)
	if err != nil {
		p.lg.Error("PostgresPostRepository.SaveBatch.stmtExec.Error", zap.Error(err))
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		p.lg.Error("PostgresPostRepository.SaveBatch.RowsAffected.Error", zap.Error(err))
		return 0, err
	}

	return rowsAffected, nil
}

func (p *postRepository) GetAll() ([]postRepo.Post, error) {
	rows, err := p.adapter.Query("select * from " + p.TableName())
	if err != nil {
		p.lg.Error("PostgresPostRepository.GetAll.GetQueryRowsError", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var (
		posts []postRepo.Post
		post  postRepo.Post
	)
	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body); err != nil {
			p.lg.Error("PostgresPostRepository.GetAll.ScanFromRowsError", zap.Error(err))
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *postRepository) Get(id int64) (post postRepo.Post, err error) {
	row := p.adapter.QueryRow("select * from "+p.TableName()+" where id=$1;", id)

	err = row.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		p.lg.Error("PostgresPostRepository.Get.QueryRowScanError", zap.Error(err))
		return
	}

	return
}

func (p *postRepository) Update(post postRepo.IPost) (int64, error) {
	result, err := p.adapter.Exec("update "+p.TableName()+" set user_id = $2, title = $3, body = $4 where id=$1", post.GetId(), post.GetUserId(), post.GetTitle(), post.GetBody())
	if err != nil {
		p.lg.Error("PostgresPostRepository.Update.QueryExecError", zap.Error(err))
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		p.lg.Error("PostgresPostRepository.Update.GetRowsAffectedError", zap.Error(err))
		return 0, err
	}

	return rowsAffected, nil
}

func (p *postRepository) Delete(id int64) (int64, error) {
	result, err := p.adapter.Exec("delete from "+p.TableName()+" where id=$1", id)
	if err != nil {
		p.lg.Error("PostgresPostRepository.Delete.Error", zap.Error(err))
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		p.lg.Error("PostgresPostRepository.Delete.GetRowsAffectedError", zap.Error(err))
		return 0, err
	}

	return rowsAffected, nil
}

func (p *postRepository) TableName() string {
	return "posts"
}

func NewPostRepository(lg logger.ILogger, adapter *sql.DB) postRepo.IPostRepository {
	return &postRepository{
		lg:      lg,
		adapter: adapter,
	}
}
