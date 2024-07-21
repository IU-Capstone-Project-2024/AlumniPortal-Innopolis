package main

import (
	"alumniportal.com/shared/initializers"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"

	pb "alumniportal.com/shared/grpc/proto"
)

func init() {
	initializers.LoadEnvVariables()
}

type server struct {
	pb.UnimplementedFilteringServiceServer
}

func (s *server) GradeDescription(ctx context.Context, req *pb.GradeRequest) (*pb.GradeResponse, error) {
	grade, err := Filter(req.Description, req.IsProject)
	if err != nil {
		return &pb.GradeResponse{Grade: int32(grade), Error: err.Error()}, nil
	}
	return &pb.GradeResponse{Grade: int32(grade), Error: ""}, nil
}

func Filter(description string, isProject bool) (int, error) {
	proxyURL, err := url.Parse("http://207.180.234.234:3128")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to parse proxy URL")
		return 1, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	var suffix string
	if isProject {
		suffix = "project"
	} else {
		suffix = "event"
	}

	prompt := "Grade the " + suffix + " description according to the rules.\n" +
		"Description: \n" + description + "\n" +
		"Rules:\n" +
		"1. Submissions must be related to IT, computer science, or technology fields.\n" +
		"2. Projects must have a clear objective, scope, and potential impact.\n" +
		"3. A detailed project description must be provided, including goals, methodology, expected outcomes.\n" +
		"4. Submissions should clearly state the problem the project aims to solve and how it proposes to solve it.\n" +
		"5. Projects requiring specialized equipment or software must detail how these will be procured and used.\n" +
		"6. Projects must comply with all university policies and guidelines.\n" +
		"7. Submissions must adhere to ethical standards, including respect for intellectual property, privacy, and data protection laws.\n" +
		"8. Projects involving human subjects must obtain appropriate ethical approvals.\n\n" +
		"WRITE ONLY THE GRADE FROM 1 TO 10"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-4",
		"messages": []map[string]string{{"role": "system", "content": prompt}},
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to marshal request body")
		return 1, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/gpt-4/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create request")
		return 1, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_KEY"))

	resp, err := httpClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to send request")
		return 1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
		}).Error("Failed to get a successful response")
		return 1, errors.New("failed to get a successful response")
	}

	var responseBody struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to decode response body")
		return 1, err
	}

	// Extract and convert the grade
	grade, err := strconv.Atoi(responseBody.Choices[0].Message.Content)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to convert grade to integer")
		return 1, err
	}

	return grade, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to listen: " + err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterFilteringServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve filtering service: %v", err)
	} else {
		logrus.Info("Filtering Service is running on port 50051")
	}
}
