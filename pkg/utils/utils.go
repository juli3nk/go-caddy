package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func GetURL(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RunCommand(command []string, dir string) error {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = env
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
