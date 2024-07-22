package main

import (
	pb "alumniportal.com/shared/grpc/proto"
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
	"strconv"
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
	proxyURL, err := url.Parse("http://AQStkg:4Pzo25@168.81.66.72:8000")
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

	promptHeader := "Grade the " + suffix + " description according to the rules. WRITE ONLY THE GRADE FROM 1 TO 10"
	promptBody := "Description: \n" + description + "\n" +
		"Rules:\n" +
		"1. Submissions must be related to IT, computer science, or technology fields.\n" +
		"2. Projects must have a clear objective, scope, and potential impact.\n" +
		"3. A detailed project description must be provided, including goals, methodology, expected outcomes.\n" +
		"4. Submissions should clearly state the problem the project aims to solve and how it proposes to solve it.\n" +
		"5. Projects requiring specialized equipment or software must detail how these will be procured and used.\n" +
		"6. Projects must comply with all university policies and guidelines.\n" +
		"7. Submissions must adhere to ethical standards, including respect for intellectual property, privacy, and data protection laws.\n" +
		"8. Projects involving human subjects must obtain appropriate ethical approvals.\n"

	requestBodyMap := map[string]interface{}{
		"modelUri": "gpt://b1ge4v0vv3t1uubfd7an/yandexgpt-lite",
		"completionOptions": map[string]interface{}{
			"stream":      false,
			"temperature": 0.1,
			"maxTokens":   "1000",
		},
		"messages": []map[string]interface{}{
			{
				"role": "system",
				"text": promptHeader,
			},
			{
				"role": "user",
				"text": promptBody,
			},
		},
	}

	requestBody, err := json.Marshal(requestBodyMap)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to marshal request body")
		return 1, err
	}

	req, err := http.NewRequest("POST", "https://llm.api.cloud.yandex.net/foundationModels/v1/completion", bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create request")
		return 1, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer t1.9euelZqcio7HjsmbxsmUmsyQzM2Pk-3rnpWak5qanI6Nk8iVipWbkZ6bzZnl9Pd_IAlL-e8cEADN3fT3P08GS_nvHBAAzc3n9euelZqdnInLnouXzceNnovOi8-TlO_8xeuelZqdnInLnouXzceNnovOi8-TlA.pL6IWii8ept17G3HeYVL_vwTxpefufxan8gfdFNGInctrx3abmw0IKKcU6O5Ghahn3GnR4GZLM43QVPcJEn2Cw")

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
		Result struct {
			Alternatives []struct {
				Message struct {
					Role string `json:"role"`
					Text string `json:"text"`
				} `json:"message"`
				Status string `json:"status"`
			} `json:"alternatives"`
			Usage struct {
				InputTextTokens  string `json:"inputTextTokens"`
				CompletionTokens string `json:"completionTokens"`
				TotalTokens      string `json:"totalTokens"`
			} `json:"usage"`
			ModelVersion string `json:"modelVersion"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to decode response body")
		return 1, err
	}

	gradeStr := responseBody.Result.Alternatives[0].Message.Text

	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		log.Fatalf("Failed to convert grade to integer: %v", err)
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
