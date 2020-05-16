// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListTables operation.
type ListTablesInput struct {
	_ struct{} `type:"structure"`

	// The first table name that this operation will evaluate. Use the value that
	// was returned for LastEvaluatedTableName in a previous operation, so that
	// you can obtain the next page of results.
	ExclusiveStartTableName *string `min:"3" type:"string"`

	// A maximum number of table names to return. If this parameter is not specified,
	// the limit is 100.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s ListTablesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTablesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTablesInput"}
	if s.ExclusiveStartTableName != nil && len(*s.ExclusiveStartTableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartTableName", 3))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListTables operation.
type ListTablesOutput struct {
	_ struct{} `type:"structure"`

	// The name of the last table in the current page of results. Use this value
	// as the ExclusiveStartTableName in a new request to obtain the next page of
	// results, until all the table names are returned.
	//
	// If you do not receive a LastEvaluatedTableName value in the response, this
	// means that there are no more table names to be retrieved.
	LastEvaluatedTableName *string `min:"3" type:"string"`

	// The names of the tables associated with the current account at the current
	// endpoint. The maximum size of this array is 100.
	//
	// If LastEvaluatedTableName also appears in the output, you can use this value
	// as the ExclusiveStartTableName parameter in a subsequent ListTables request
	// and obtain the next page of results.
	TableNames []string `type:"list"`
}

// String returns the string representation
func (s ListTablesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTables = "ListTables"

// ListTablesRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Returns an array of table names associated with the current account and endpoint.
// The output from ListTables is paginated, with each page returning a maximum
// of 100 table names.
//
//    // Example sending a request using ListTablesRequest.
//    req := client.ListTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListTables
func (c *Client) ListTablesRequest(input *ListTablesInput) ListTablesRequest {
	op := &aws.Operation{
		Name:       opListTables,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"ExclusiveStartTableName"},
			OutputTokens:    []string{"LastEvaluatedTableName"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTablesInput{}
	}

	req := c.newRequest(op, input, &ListTablesOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}
	return ListTablesRequest{Request: req, Input: input, Copy: c.ListTablesRequest}
}

// ListTablesRequest is the request type for the
// ListTables API operation.
type ListTablesRequest struct {
	*aws.Request
	Input *ListTablesInput
	Copy  func(*ListTablesInput) ListTablesRequest
}

// Send marshals and sends the ListTables API request.
func (r ListTablesRequest) Send(ctx context.Context) (*ListTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTablesResponse{
		ListTablesOutput: r.Request.Data.(*ListTablesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTablesRequestPaginator returns a paginator for ListTables.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTablesRequest(input)
//   p := dynamodb.NewListTablesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTablesPaginator(req ListTablesRequest) ListTablesPaginator {
	return ListTablesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTablesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTablesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTablesPaginator struct {
	aws.Pager
}

func (p *ListTablesPaginator) CurrentPage() *ListTablesOutput {
	return p.Pager.CurrentPage().(*ListTablesOutput)
}

// ListTablesResponse is the response type for the
// ListTables API operation.
type ListTablesResponse struct {
	*ListTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTables request.
func (r *ListTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
