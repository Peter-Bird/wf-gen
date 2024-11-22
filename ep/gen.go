// wf-gen/ep/gen.go
package ep

import (
	"encoding/json"
	"log"

	"github.com/Peter-Bird/models"
	"github.com/Peter-Bird/ws"
)

// GenWorkFlowService struct implementing the Service interface
type GenWorkFlowService struct{}

func init() {
	ws.RegisterService("workflow", func() ws.Service { return &GenWorkFlowService{} })
}

func (g GenWorkFlowService) Process(input map[string]interface{}) (map[string]interface{}, error) {
	log.Printf("Generate Workflow invoked")

	wf := GenerateWorkflow()

	return map[string]interface{}{"work-flow": wf}, nil
}

func GenerateWorkflow() models.Workflow {
	// Define the nested workflow for the "parameters" field
	nestedWorkflow := models.Workflow{
		Id:   "WF2",
		Name: "WF2",
		Steps: []models.Step{
			{
				Endpoint:     "http://localhost:8083/workflows",
				Method:       "GET",
				Parameters:   nil,
				Dependencies: []string{},
			},
		},
	}

	// Convert `nestedWorkflow` to `map[string]interface{}`
	var parameters map[string]interface{}
	wfJSON, err := json.Marshal(nestedWorkflow)
	if err != nil {
		log.Fatalf("Failed to marshal nestedWorkflow: %v", err)
	}
	if err := json.Unmarshal(wfJSON, &parameters); err != nil {
		log.Fatalf("Failed to unmarshal to map[string]interface{}: %v", err)
	}

	return models.Workflow{
		Id:   "WF1",
		Name: "Initial Workflow",
		Steps: []models.Step{
			{
				Endpoint:     "http://localhost:8083/workflows",
				Method:       "GET",
				Parameters:   nil,
				Dependencies: nil,
			},
			{
				Endpoint:     "http://localhost:8083/workflows/submit",
				Method:       "POST",
				Parameters:   parameters, // `parameters` is now of type map[string]interface{}
				Dependencies: []string{"step1"},
			},
			{
				Endpoint:     "http://localhost:8083/workflows/get/WF1",
				Method:       "GET",
				Parameters:   nil,
				Dependencies: []string{"step2"},
			},
		},
	}
}
