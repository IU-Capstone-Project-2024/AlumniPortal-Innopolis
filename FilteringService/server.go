package FilteringService

import (
	"context"
	"github.com/ayush6624/go-chatgpt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"

	pb "FilteringService/grpc/proto"
)

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
	client, err := chatgpt.NewClient(os.Getenv("OPENAI_KEY"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed create GPT New Client")
		return 1, err
	}

	ctx := context.Background()

	var suffix string

	if isProject {
		suffix = "project"
	} else {
		suffix = "event"
	}

	res, err := client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role: chatgpt.ChatGPTModelRoleSystem,
				Content: "Grade the " + suffix + " description according to the rules.\n" +
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
					"WRITE ONLY THE GRADE FROM 1 TO 10",
			},
		},
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to use OpenAI API")
		return 1, err
	}

	grade, err := strconv.Atoi(res.Choices[0].Message.Content)
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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve filtering service: %v", err)
	} else {
		logrus.Info("Filtering Service is running on port 50051")
	}
}
