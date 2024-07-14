package quiz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	bolt "go.etcd.io/bbolt"
)

var quizBucket string = "QuizBucket"
var fileName = "/tmp/quiz-users.db"

type CLI struct {
	baseURL string
	db      *bolt.DB
}

func NewCLI(baseURL string) (*CLI, error) {
	db, err := bolt.Open(fileName, 0666, nil)
	if err != nil {
		return nil, err
	}
	// we don't close here because we need the db open to perform operations

	err = db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		_, err := tx.CreateBucketIfNotExists([]byte(quizBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &CLI{baseURL: baseURL, db: db}, nil
}

func (c *CLI) Register(userId string) error {
	url := fmt.Sprintf("%s/quiz/user/register", c.baseURL)

	userReq := UserRequest{
		Id: userId,
	}

	data, err := json.Marshal(userReq)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", string(respBody))

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return errors.New(response.Error)
	}

	c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(quizBucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		err := b.Put([]byte(userId), []byte("1"))
		return err
	})

	fmt.Println("Response written to file successfully")

	return c.db.Close()
}

func (c *CLI) Questions() error {
	url := fmt.Sprintf("%s/quiz/questions", c.baseURL)

	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(response.Error)
	}

	result, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(result))

	return c.db.Close()
}

func (c *CLI) Start(userId string) error {
	var position []byte
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(quizBucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		position = b.Get([]byte(userId))

		if position == nil {
			return errors.New("user not found")
		}
		return nil
	})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/quiz/%s", c.baseURL, string(position))

	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(response.Error)
	}

	result, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(result))

	return c.db.Close()
}
