package main
 
import (
  "encoding/json"
  "strconv"
 
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
  ID int `json:"id"`
  Name string `json:"name"`
}

var movies = []struct {
  ID int `json:"id"`
  Name string `json:"name"`
}{
    {
      ID: 1,
      Name: "Avengers",
    },
    {
      ID: 2,
      Name: "Ant-Man",
    },
    {
      ID: 3,
      Name: "Thor",
    },
    {
      ID: 4,
      Name: "Hulk",
    }, {
      ID: 5,
      Name: "Doctor Strange",
    },
}
 
func findAll() (events.APIGatewayProxyResponse, error) {
  response, err := json.Marshal(movies)
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }
 
  return events.APIGatewayProxyResponse{
    StatusCode: 200,
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
    Body: string(response),
  }, nil
}

func findOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
 
  id, err := strconv.Atoi(req.PathParameters["id"])
 
  if err != nil {
 
    return events.APIGatewayProxyResponse{
 
      StatusCode: 500,
 
      Body:       "ID must be a number",
 
    }, nil
 
  }
 
  response, err := json.Marshal(movies[id-1])
 
  if err != nil {
 
    return events.APIGatewayProxyResponse{
 
      StatusCode: 500,
 
      Body:       err.Error(),
 
    }, nil
 
  }
 
  return events.APIGatewayProxyResponse{
 
    StatusCode: 200,
 
    Headers: map[string]string{
 
      "Content-Type": "application/json",
 
    },
 
    Body: string(response),
 
  }, nil
 
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  var movie Movie
  err := json.Unmarshal([]byte(req.Body), &movie)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: 400,
      Body: "Invalid payload",
    }, nil
  }
 
  movies = append(movies, movie)
 
  response, err := json.Marshal(movies)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: 500,
      Body: err.Error(),
    }, nil
  }
 
  return events.APIGatewayProxyResponse{
    StatusCode: 200,
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
    Body: string(response),
  }, nil
}

func main() {
  lambda.Start(findAll)
}